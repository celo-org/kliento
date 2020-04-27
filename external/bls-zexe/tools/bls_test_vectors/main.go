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

package main

import (
  "fmt"
  "golang.org/x/crypto/blake2s"
  "math/rand"
  "encoding/hex"
)

func main() {
  rand.Seed(1000)

  hashLength := 96
  xof, _ := blake2s.NewXOF(uint16(hashLength), nil)
  msg := make([]byte, 140)
  rand.Read(msg)
  fmt.Printf("msg: %s\n", hex.EncodeToString(msg))

  xof.Write(msg)

  hash := make([]byte, hashLength)
  xof.Read(hash)
  fmt.Printf("hash: %s\n", hex.EncodeToString(hash))
}
