// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided ‘as is’ and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package crypto

import (
	"github.com/golang/glog"
	"golang.org/x/crypto/blake2b"

	"github.com/iotexproject/iotex-core/common"
)

// Merkle tree struct
type Merkle struct {
	root common.Hash32B
	leaf []common.Hash32B
	size int
}

// NewMerkleTree creates a merkle tree given hashed leaves
func NewMerkleTree(leaves []common.Hash32B) *Merkle {
	size := len(leaves)
	if size == 0 {
		glog.Warning("Try to create merkle tree with empty leaf list!")
		return nil
	}

	mk := &Merkle{
		leaf: make([]common.Hash32B, (size+1)>>1<<1),
		size: size,
	}

	for i, leaf := range leaves {
		mk.leaf[i] = leaf
	}

	if size == 1 {
		mk.root = mk.leaf[0]
		return mk
	}

	// copy the last hash if original size is odd number
	if size != len(mk.leaf) {
		mk.leaf[size] = mk.leaf[size-1]
		mk.size = len(mk.leaf)
	}

	return mk
}

// HashTree calculates the root hash of a merkle tree
func (mk *Merkle) HashTree() common.Hash32B {
	if mk.root != common.ZeroHash32B {
		return mk.root
	}

	length := mk.size >> 1
	merkle := make([]common.Hash32B, length)

	// first round, compute hash from original leaf
	for i := 0; i < length; i++ {
		h := mk.leaf[i<<1][:]
		h = append(h, mk.leaf[i<<1+1][:]...)
		merkle[i] = blake2b.Sum256(h)
	}

	for length > 1 {
		if length&1 != 0 {
			merkle = append(merkle, merkle[length-1])
			length++
		}

		length >>= 1
		for i := 0; i < length; i++ {
			h := merkle[i<<1][:]
			h = append(h, merkle[i<<1+1][:]...)
			merkle[i] = blake2b.Sum256(h)
		}
		merkle = merkle[0:length]
	}

	mk.root = merkle[0]
	return mk.root
}
