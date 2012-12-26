package main

import (
	"bytes"
	"fmt"

	"github.com/nightlyone/permutation"
)

// define custom type
type byteSlice []byte

// Implement the three functions from sort.Interface (part of permutation.Sequence interface)
func (p byteSlice) Len() int           { return len(p) }
func (p byteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p byteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Implement the remaining portions of permutation.Sequence interface
func (p byteSlice) Equal(q permutation.Sequence) bool { return bytes.Equal(p, q.(byteSlice)) }
func (p byteSlice) Copy() permutation.Sequence {
	q := make(byteSlice, len(p), len(p))
	copy(q, p)
	return q
}

// and again

// define custom type
type stringSlice []byte

// Implement the three functions from sort.Interface (part of permutation.Sequence interface)
func (p stringSlice) Len() int           { return len(p) }
func (p stringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p stringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Implement the remaining portions of permutation.Sequence interface
func (p stringSlice) Equal(q permutation.Sequence) bool { return bytes.Equal(p, q.(stringSlice)) }
func (p stringSlice) Copy() permutation.Sequence {
	q := make(stringSlice, len(p), len(p))
	copy(q, p)
	return q
}

// and make it printable as string
func (p stringSlice) String() string { return string(p) }

// Now try it out
func main() {

	// Permutation of simple slice
	perm := permutation.New(byteSlice{3, 5, 1})

	// How you iterate over all permutations
	for ok, seq := true, perm.Current(); ok; ok, seq = perm.Next(), perm.Current() {
		fmt.Println("Permutation := ", seq)
	}

	// Permutation of character slice
	toad := permutation.New(stringSlice{'t', 'o', 'a', 'd'})

	for ok, seq := true, toad.Current(); ok; ok, seq = toad.Next(), toad.Current() {
		fmt.Println("Permutation := ", seq)
	}
}
