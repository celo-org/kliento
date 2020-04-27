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

use super::{BLSCurve, CPCurve, CPField, CPFrParams};
use crate::encoding::EncodingError;
use crate::epoch_block::{hash_first_last_epoch_block, EpochBlock};
use crate::gadgets::pack;
use groth16::{prepare_verifying_key, verify_proof, Proof, VerifyingKey};
use r1cs_core::SynthesisError;
use thiserror::Error;
use tracing::info;

#[derive(Debug, Error)]
pub enum VerificationError {
    #[error("Verification failed")]
    VerificationFailed,
    #[error("Synthesis Error: {0}")]
    ZexeSynthesisError(#[from] SynthesisError),
    #[error("Encoding Error: {0}")]
    EpochEncodingError(#[from] EncodingError),
}

/// The verifier only needs to check that the edges were calculated correctly
pub fn verify(
    vk: &VerifyingKey<CPCurve>,
    first_epoch: &EpochBlock,
    last_epoch: &EpochBlock,
    proof: &Proof<CPCurve>,
) -> Result<(), VerificationError> {
    info!("Verifying proof");
    // Hash the first-last block together
    let hash = hash_first_last_epoch_block(first_epoch, last_epoch)?;
    // packs them
    let public_inputs = pack::<CPField, CPFrParams>(&hash);
    // verifies the BLS proof by using the First/Last epoch as public inputs over CP
    if verify_proof(&prepare_verifying_key(vk), proof, &public_inputs)? {
        Ok(())
    } else {
        Err(VerificationError::VerificationFailed)
    }
}
