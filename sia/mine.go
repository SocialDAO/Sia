package sia

import (
	"bytes"
	"crypto/rand"
)

// Hashcash brute-forces a nonce that produces a hash less than target.
func Hashcash(target Hash) (nonce []byte, i int) {
	nonce = make([]byte, 8)
	for {
		i++
		rand.Read(nonce)
		h := HashBytes(nonce)
		if bytes.Compare(h[:], target[:]) < 0 {
			return
		}
	}
}

// Creates a new block.  This function creates a new block given a previous
// block, isn't happy with being interrupted.  Need a different thread that can
// be updated by listening on channels or something.
func (s *State) GenerateBlock() (b *Block) {
	b = &Block{
		Version:     1,
		ParentBlock: s.CurrentBlock,
		// Timestamp
		// Miner Address
		// Merkle Root
		// List of Transactions
	}

	// Perform work until the block matches the desired header value.

	return
}