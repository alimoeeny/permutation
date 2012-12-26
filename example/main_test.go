package main

import (
	"fmt"

	"github.com/nightlyone/permutation"
)

func ExampleBytes() {
	// Permutation of simple slice
	perm := permutation.New(byteSlice{3, 5, 1})

	// How you iterate over all permutations
	for ok, seq := true, perm.Current(); ok; ok, seq = perm.Next(), perm.Current() {
		fmt.Println("Permutation := ", seq)
	}
	// Output:
	// Permutation :=  [1 3 5]
	// Permutation :=  [1 5 3]
	// Permutation :=  [3 1 5]
	// Permutation :=  [3 5 1]
	// Permutation :=  [5 1 3]
	// Permutation :=  [5 3 1]
}

func ExampleCharacters() {
	// Permutation of character slice
	toad := permutation.New(stringSlice{'t', 'o', 'a', 'd'})

	for ok, seq := true, toad.Current(); ok; ok, seq = toad.Next(), toad.Current() {
		fmt.Println("Permutation := ", seq)
	}
	// Output:
	// Permutation :=  adot
	// Permutation :=  adto
	// Permutation :=  aodt
	// Permutation :=  aotd
	// Permutation :=  atdo
	// Permutation :=  atod
	// Permutation :=  daot
	// Permutation :=  dato
	// Permutation :=  doat
	// Permutation :=  dota
	// Permutation :=  dtao
	// Permutation :=  dtoa
	// Permutation :=  oadt
	// Permutation :=  oatd
	// Permutation :=  odat
	// Permutation :=  odta
	// Permutation :=  otad
	// Permutation :=  otda
	// Permutation :=  tado
	// Permutation :=  taod
	// Permutation :=  tdao
	// Permutation :=  tdoa
	// Permutation :=  toad
	// Permutation :=  toda
}
