package permutation

// Naive permutation implementation

import (
	"sort"
)

// Sequence is the interface you need to implement, to allow permutations of it.
// It must consist of independent elements.
type Sequence interface {
	sort.Interface
	// Return whether seqence q matches the receiver
	Equal(q Sequence) bool
	// Create a copy of the sequence. Only the indexes need to be independent.
	Copy() Sequence
}

func nextPermutation(seq Sequence) bool {
	for j := seq.Len() - 1; j > 0; j-- {
		if seq.Less(j-1, j) {
			m := seq.Len() - 1
			for seq.Less(m, j-1) {
				m--
			}
			seq.Swap(j-1, m)
			reverse(seq, j)
			return true
		}
	}
	return false
}

// reverse a sequence of in place.
func reverse(seq Sequence, start int) {
	for i, j := start, seq.Len()-1; i < j; i, j = i+1, j-1 {
		seq.Swap(i, j)
	}
}

// Generic permutation type
type Permutation struct {
	seq Sequence
}

// Initializes a permutation from seqence seq and returns it
// TODO: Always copy it?
func New(seq Sequence) *Permutation {
	if !sort.IsSorted(seq) {
		sort.Sort(seq)
	}
	return &Permutation{seq: seq}
}

// Returns the current sequence
func (p *Permutation) Current() Sequence {
	return p.seq
}

// Generates the next permutation in place and returns, whether it is different.
// Use Current() to get newly generated sequence
func (p *Permutation) Next() bool {
	next, valid := nextPerm(p.seq)
	p.seq = next
	return valid
}

// How many possible permutations of this sequence exist?
func (p *Permutation) Len() int {
	result := 1
	for i := 2; i <= p.seq.Len(); i++ {
		result *= i
	}
	return result
}

func nextPerm(seq Sequence) (next Sequence, valid bool) {
	current := seq.Copy()

	for i := 1; i < seq.Len(); i++ {
		if ok := nextPermutation(seq); !ok {
			return seq, false
		}
		if ok := !seq.Equal(current); ok {
			return seq, true
		}
	}

	return seq, false
}
