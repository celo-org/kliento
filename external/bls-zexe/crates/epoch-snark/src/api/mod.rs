#![allow(unused)]
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

pub mod prover;
/// API for setting up, generating and verifying proofs for the
/// SNARK instantiated over BLS12-377 / SW6 which proves language
/// ...
pub mod setup;
pub mod verifier;

pub mod ffi;

use algebra::{bls12_377, sw6};

// Instantiate certain types to avoid confusion
pub type Parameters = setup::Parameters<CPCurve, BLSCurve>;

pub type BLSField = bls12_377::Fr;
pub type BLSCurve = bls12_377::Bls12_377;
pub type BLSFrParams = bls12_377::FrParameters;

pub type CPField = sw6::Fr;
pub type CPCurve = sw6::SW6;
pub type CPFrParams = sw6::FrParameters;
