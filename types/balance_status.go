// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"fmt"

	"github.com/thunderheados/go-substrate-rpc-client/v4/scale"
)

type BalanceStatus byte

const (
	// Funds are free, as corresponding to `free` item in Balances.
	Free BalanceStatus = 0
	// Funds are reserved, as corresponding to `reserved` item in Balances.
	Reserved BalanceStatus = 1
)

func (bs *BalanceStatus) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	vb := BalanceStatus(b)
	switch vb {
	case Free, Reserved:
		*bs = vb
	default:
		return fmt.Errorf("unknown BalanceStatus enum: %v", vb)
	}
	return err
}

func (bs BalanceStatus) Encode(encoder scale.Encoder) error {
	return encoder.PushByte(byte(bs))
}
