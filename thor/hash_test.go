// Copyright (c) 2022 The Dexio developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package thor_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/BestSilverTiger/thor/blake2b"
)

func BenchmarkKeccak(b *testing.B) {
	data := []byte("hello world")
	for i := 0; i < b.N; i++ {
		hash := sha3.NewKeccak256()
		hash.Write(data)
		hash.Sum(nil)
	}
}

func BenchmarkBlake2b(b *testing.B) {
	data := []byte("hello world")
	for i := 0; i < b.N; i++ {
		hash, _ := blake2b.New256(nil)
		hash.Write(data)
		hash.Sum(nil)
	}
}
