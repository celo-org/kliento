// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

use bls_crypto::{DirectHasher, PrivateKey, TryAndIncrement};

use algebra::bytes::ToBytes;

use clap::{App, Arg};
use rand::thread_rng;
use std::{fs::File, io::Write};

fn main() {
    let matches = App::new("BLS Proof of Possession test vectors")
        .about("Generates many proof of posession for random keys")
        .arg(
            Arg::with_name("num")
                .short("n")
                .value_name("NUM")
                .help("Sets the number of test vectors")
                .required(true),
        )
        .arg(
            Arg::with_name("out")
                .short("o")
                .value_name("OUT")
                .help("Sets the output file path")
                .required(true),
        )
        .get_matches();

    let num: i32 = matches.value_of("num").unwrap().parse().unwrap();
    let out = matches.value_of("out").unwrap();

    let direct_hasher = DirectHasher::new().unwrap();
    let try_and_increment = TryAndIncrement::new(&direct_hasher);
    let rng = &mut thread_rng();
    let mut file = File::create(out).unwrap();
    for _ in 0..num {
        let sk = PrivateKey::generate(rng);
        let pk = sk.to_public();
        let address_bytes = hex::decode("60515f8c59451e04ab4b22b3fc9a196b2ad354e6").unwrap();
        let mut pk_bytes = vec![];
        pk.write(&mut pk_bytes).unwrap();
        let pop = sk.sign_pop(&address_bytes, &try_and_increment).unwrap();
        let mut pop_bytes = vec![];
        pop.write(&mut pop_bytes).unwrap();

        let mut sk_bytes = vec![];
        sk.write(&mut sk_bytes).unwrap();

        pk.verify_pop(&address_bytes, &pop, &try_and_increment)
            .unwrap();

        file.write_all(
            format!(
                "{},{},{}\n",
                hex::encode(sk_bytes),
                hex::encode(pk_bytes),
                hex::encode(pop_bytes)
            )
            .as_bytes(),
        )
        .unwrap();
    }
}
