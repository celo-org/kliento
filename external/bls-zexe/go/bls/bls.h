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

#include <stdbool.h>
#include <stdint.h>

struct PrivateKey;
typedef struct PrivateKey PrivateKey;

struct PublicKey;
typedef struct PublicKey PublicKey;

struct Signature;
typedef struct Signature Signature;

void init();
bool generate_private_key(PrivateKey**);
bool deserialize_private_key(const unsigned char*, int32_t, PrivateKey**);
bool serialize_private_key(const PrivateKey*, unsigned char**, int32_t*);
bool private_key_to_public_key(const PrivateKey*, PublicKey**);
bool sign_message(const PrivateKey*, const unsigned char*, int32_t, const unsigned char*, int32_t, bool, Signature**);
bool sign_pop(const PrivateKey*, const unsigned char*, int32_t, Signature**);
void destroy_private_key(PrivateKey*);
void free_vec(unsigned char*, int32_t);
bool deserialize_public_key(const unsigned char*, int32_t, PublicKey**);
bool serialize_public_key(const PublicKey*, unsigned char**, int32_t*);
void destroy_public_key(PublicKey*);
bool deserialize_signature(const unsigned char*, int32_t, Signature**);
bool serialize_signature(const Signature*, unsigned char**, int32_t*);
void destroy_signature(Signature*);
bool verify_signature(const PublicKey*, const unsigned char*, int32_t, const unsigned char*, int32_t, const Signature*, bool, bool*);
bool verify_pop(const PublicKey*, const unsigned char*, int32_t, const Signature*, bool*);
bool aggregate_public_keys(const PublicKey**, int32_t, PublicKey**);
bool aggregate_public_keys_subtract(const PublicKey*, const PublicKey**, int32_t, PublicKey**);
bool aggregate_signatures(const Signature**, int32_t, Signature**);

bool encode_epoch_block_to_bytes(uint16_t, uint32_t, const PublicKey**, int32_t, bool, unsigned char**, int32_t*);
bool hash_direct(const unsigned char*, int32_t, unsigned char**, int32_t*, bool);
bool hash_composite(const unsigned char*, int32_t, const unsigned char*, int32_t, unsigned char**, int32_t*); 
bool compress_signature(const unsigned char*, int32_t, unsigned char**, int32_t*);
bool compress_pubkey(const unsigned char*, int32_t, unsigned char**, int32_t*);
