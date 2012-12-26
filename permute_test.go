package permutation

import (
	"testing"
)

type permTest struct {
	seq   []byte
	perms [][]byte
}

var permsPositive = [...]permTest{
	{
		seq: []byte("bac"),
		perms: [][]byte{
			[]byte("abc"),
			[]byte("acb"),
			[]byte("bac"),
			[]byte("bca"),
			[]byte("cab"),
			[]byte("cba"),
		},
	},
}

func TestNext(t *testing.T) {
	for test, a := range permsPositive {
		p := NewBytePermutation(a.seq)
		i := 0
		for ok, seq := true, p.Current(); ok; ok, seq = p.Next(), p.Current() {
			if string(seq) != string(a.perms[i]) {
				t.Errorf("%d: seq=%s, want=%s, got=%s", test, a.seq, a.perms[i], seq)
			} else {
				t.Logf("%d: seq=%s, want=%s, got=%s", test, a.seq, a.perms[i], seq)
			}
			i++
		}
	}
}

func TestLen(t *testing.T) {
	for test, a := range permsPositive {
		p := NewBytePermutation(a.seq)
		if len(a.perms) != p.Len() {
			t.Errorf("%d: seq=%s, want=%v, got=%v", test, a.seq, len(a.perms), p.Len())
		} else {
			t.Logf("%d: seq=%s, want=%v, got=%v", test, a.seq, len(a.perms), p.Len())
		}
	}
}
