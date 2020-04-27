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

use algebra::bytes::{FromBytes, ToBytes};

use clap::{App, Arg};

fn main() {
    let matches = App::new("BLS Proof of Possession")
        .about("Generates a proof of posession for the given private key")
        .arg(
            Arg::with_name("key")
                .short("k")
                .value_name("KEY")
                .help("Sets the BLS private key")
                .required(true),
        )
        .get_matches();

    let key = matches.value_of("key").unwrap();

    let key_bytes = hex::decode(key).unwrap();

    let direct_hasher = DirectHasher::new().unwrap();
    let try_and_increment = TryAndIncrement::new(&direct_hasher);
    let sk = PrivateKey::read(key_bytes.as_slice()).unwrap();
    let pk = sk.to_public();
    let mut pk_bytes = vec![];
    pk.write(&mut pk_bytes).unwrap();
    let pop = sk.sign_pop(&pk_bytes, &try_and_increment).unwrap();
    let mut pop_bytes = vec![];
    pop.write(&mut pop_bytes).unwrap();

    /*
    let mut pk_bytes = vec![];
    pk.write(&mut pk_bytes).unwrap();
    println!("{}", hex::encode(&pk_bytes));
    */

    pk.verify_pop(&pk_bytes, &pop, &try_and_increment).unwrap();

    let pop_hex = hex::encode(&pop_bytes);
    println!("{}", pop_hex);
}
