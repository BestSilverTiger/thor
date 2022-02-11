// Copyright (c) 2022 The Dexio developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package chain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/BestSilverTiger/thor/block"
	"github.com/BestSilverTiger/thor/chain"
)

func TestBlockReader(t *testing.T) {
	repo := newTestRepo()
	b0 := repo.GenesisBlock()

	b1 := newBlock(b0, 10)
	repo.AddBlock(b1, nil)

	b2 := newBlock(b1, 20)
	repo.AddBlock(b2, nil)

	b3 := newBlock(b2, 30)
	repo.AddBlock(b3, nil)

	b4 := newBlock(b3, 40)
	repo.AddBlock(b4, nil)

	repo.SetBestBlockID(b4.Header().ID())

	br := repo.NewBlockReader(b2.Header().ID())

	var blks []*chain.ExtendedBlock

	for {
		r, err := br.Read()
		if err != nil {
			panic(err)
		}
		if len(r) == 0 {
			break
		}
		blks = append(blks, r...)

	}

	assert.Equal(t, []*chain.ExtendedBlock{
		{block.Compose(b3.Header(), b3.Transactions()), false},
		{block.Compose(b4.Header(), b4.Transactions()), false}},
		blks)
}

func TestBlockReaderFork(t *testing.T) {
	repo := newTestRepo()
	b0 := repo.GenesisBlock()

	b1 := newBlock(b0, 10)
	repo.AddBlock(b1, nil)

	b2 := newBlock(b1, 20)
	repo.AddBlock(b2, nil)

	b2x := newBlock(b1, 20)
	repo.AddBlock(b2x, nil)

	b3 := newBlock(b2, 30)
	repo.AddBlock(b3, nil)

	b4 := newBlock(b3, 40)
	repo.AddBlock(b4, nil)

	repo.SetBestBlockID(b4.Header().ID())

	br := repo.NewBlockReader(b2x.Header().ID())

	var blks []*chain.ExtendedBlock

	for {
		r, err := br.Read()
		if err != nil {
			panic(err)
		}
		if len(r) == 0 {
			break
		}

		blks = append(blks, r...)
	}

	assert.Equal(t, []*chain.ExtendedBlock{
		{block.Compose(b2x.Header(), b2x.Transactions()), true},
		{block.Compose(b2.Header(), b2.Transactions()), false},
		{block.Compose(b3.Header(), b3.Transactions()), false},
		{block.Compose(b4.Header(), b4.Transactions()), false}},
		blks)
}
