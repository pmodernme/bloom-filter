package filter

import (
	"hash"
	"hash/fnv"

	"github.com/zentures/cityhash"
)

// Filter is a bloom filter
type Filter struct {
	bits      []bool        //TODO: bool is 8 bits wide, need single bit implementation
	n         uint          //Number of elements
	hashFuncs []hash.Hash64 // hash functions
}

// New Filter with size s
func New(s uint) *Filter {
	return &Filter{
		bits:      make([]bool, s),
		n:         uint(0),
		hashFuncs: []hash.Hash64{cityhash.New64(), fnv.New64(), fnv.New64a()},
	}
}

// Add an item as []byte to the filter
func (f *Filter) Add(item []byte) {
	hashes := f.hashes(item)

	for _, h := range hashes {
		p := h % uint64(len(f.bits))
		f.bits[p] = true
	}

	f.n++
}

// Lookup if an item as []byte matches the contents of the filter
// (can result in false positives)
func (f *Filter) Lookup(item []byte) bool {
	hashes := f.hashes(item)

	for _, h := range hashes {
		p := h % uint64(len(f.bits))
		if !f.bits[p] {
			return false
		}
	}

	return true
}

func (f *Filter) hashes(item []byte) []uint64 {
	var h []uint64

	for _, hf := range f.hashFuncs {
		hf.Write(item)
		h = append(h, hf.Sum64())
		hf.Reset()
	}

	return h
}
