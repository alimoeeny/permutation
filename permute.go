package permutation

// Naive permutation implementation on bytes

import (
	"bytes"
	"sort"
)

type byteSlice []byte

func (p byteSlice) Len() int           { return len(p) }
func (p byteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p byteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func nextBytePermutation(seq []byte) bool {
	for j := len(seq) - 1; j > 0; j-- {
		if v := seq[j-1]; v < seq[j] {
			m := len(seq) - 1
			for v > seq[m] {
				m--
			}
			seq[j-1], seq[m] = seq[m], seq[j-1]
			reverse(seq[j:])
			return true
		}
	}
	return false
}

// reverse a sequence of bytes in place. Thanks Russ Cox!
func reverse(seq []byte) {
	for i, j := 0, len(seq)-1; i < j; i, j = i+1, j-1 {
		seq[i], seq[j] = seq[j], seq[i]
	}
}

// permutation on bytes
type BytePermutation struct {
	seq []byte
}


// Initializes a byte permutation and returns it
func NewBytePermutation(seq []byte) *BytePermutation {
	if !sort.IsSorted(byteSlice(seq)) {
		current := make([]byte, len(seq), len(seq))
		copy(current, seq)
		sort.Sort(byteSlice(current))
		return &BytePermutation{seq: current}
	}
	return &BytePermutation{seq: seq}
}

// Returns the current sequence of bytes
func (p *BytePermutation) Current() []byte {
	return p.seq
}

// Generates the next permutation of bytes and returns, whether it is different
// Use Current() to get the next one
func (p *BytePermutation) Next() bool {
	next, valid := nextPerm(p.seq)
	p.seq = next
	return valid
}

// How many possible permutations of this sequence exist?
func (p *BytePermutation) Len() int {
	result := 1
	for i := 2; i <= len(p.seq); i++ {
		result *= i
	}
	return result
}

func nextPerm(seq []byte) (next []byte, valid bool) {
	current := make([]byte, len(seq), len(seq))
	copy(current, seq)

	for i := 1; i < len(seq); i++ {
		if ok := nextBytePermutation(seq); !ok {
			return seq, false
		}
		if ok := bytes.Compare(current, seq) != 0; ok {
			return seq, true
		}
	}

	return seq, false
}
