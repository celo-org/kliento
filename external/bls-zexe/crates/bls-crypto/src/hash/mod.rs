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

pub mod composite;
pub mod direct;
use std::error::Error;

pub trait XOF {
    fn crh(
        &self,
        domain: &[u8],
        message: &[u8],
        xof_digest_length: usize,
    ) -> Result<Vec<u8>, Box<dyn Error>>;
    fn xof(
        &self,
        domain: &[u8],
        hashed_message: &[u8],
        output_size_in_bytes: usize,
    ) -> Result<Vec<u8>, Box<dyn Error>>;
    fn hash(
        &self,
        domain: &[u8],
        message: &[u8],
        output_size_in_bytes: usize,
    ) -> Result<Vec<u8>, Box<dyn Error>>;
}
