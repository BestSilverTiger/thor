// Copyright (c) 2022 The Dexio developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package params

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/BestSilverTiger/thor/muxdb"
	"github.com/BestSilverTiger/thor/state"
	"github.com/BestSilverTiger/thor/thor"
)

func TestParamsGetSet(t *testing.T) {
	db := muxdb.NewMem()
	st := state.New(db, thor.Bytes32{})
	setv := big.NewInt(10)
	key := thor.BytesToBytes32([]byte("key"))
	p := New(thor.BytesToAddress([]byte("par")), st)
	p.Set(key, setv)

	getv, err := p.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, setv, getv)
}
