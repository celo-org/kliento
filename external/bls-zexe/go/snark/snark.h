/*
 * Copyright 2020 Celo Org
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

/**
 * Data structure received from consumers of the FFI interface describing
 * an epoch block.
 */
typedef struct {
  /**
   * The epoch's index
   */
  uint16_t index;
  /**
   * Pointer to the public keys array
   */
  const uint8_t *pubkeys;
  /**
   * The number of public keys to be read from the pointer
   */
  uintptr_t pubkeys_num;
  /**
   * Maximum number of non signers for that epoch
   */
  uint32_t maximum_non_signers;
} EpochBlockFFI;

/**
 * Verifies a Groth16 proof about the validity of the epoch transitions
 * between the provided `first_epoch` and `last_epoch` blocks.
 *
 * All elements are assumed to be sent as serialized byte arrays
 * of **compressed elements**. There are no assumptions made about
 * the length of the verifying key or the proof, so that must be
 * provided by the caller.
 *
 * # Safety
 * 1. VK and Proof must be valid pointers
 * 1. The vector of pubkeys inside EpochBlockFFI must point to valid memory
 */
bool verify(const uint8_t *vk,
            uint32_t vk_len,
            const uint8_t *proof,
            uint32_t proof_len,
            EpochBlockFFI first_epoch,
            EpochBlockFFI last_epoch);
