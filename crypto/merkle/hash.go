package merkle

import (
	"hash"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

const (
	leafPrefix  = 0
	innerPrefix = 1
)

type merkleHasher struct {
	state hash.Hash
}

func newMerkleHasher() merkleHasher {
	return merkleHasher{state: tmhash.New()}
}

func (mh *merkleHasher) emptyHash() []byte {
	mh.state.Reset()
	return mh.state.Sum(nil)
}

func (mh *merkleHasher) leafHash(leaf []byte) []byte {
	mh.state.Reset()
	mh.state.Write([]byte{leafPrefix})
	mh.state.Write(leaf)
	return mh.state.Sum(nil)
}

func (mh *merkleHasher) innerHash(left []byte, right []byte) []byte {
	mh.state.Reset()
	mh.state.Write([]byte{innerPrefix})
	mh.state.Write(left)
	mh.state.Write(right)
	return mh.state.Sum(nil)
}
