package permutation

import (
	"bytes"
	"testing"
)

type byteSlice []byte

func (p byteSlice) Len() int              { return len(p) }
func (p byteSlice) Less(i, j int) bool    { return p[i] < p[j] }
func (p byteSlice) Swap(i, j int)         { p[i], p[j] = p[j], p[i] }
func (p byteSlice) Equal(q Sequence) bool { return bytes.Equal(p, q.(byteSlice)) }
func (p byteSlice) Copy() Sequence        { q := make(byteSlice, len(p), len(p)); copy(q, p); return q }

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

func TestIteration(t *testing.T) {
	for test, a := range permsPositive {
		p := New(byteSlice(a.seq))
		for i, ok, seq := 0, true, p.Current().(byteSlice); ok; i, ok, seq = i+1, p.Next(), p.Current().(byteSlice) {
			if string(seq) != string(a.perms[i]) {
				t.Errorf("%d: seq=%s, want=%s, got=%s", test, a.seq, a.perms[i], seq)
			} else {
				t.Logf("%d: seq=%s, want=%s, got=%s", test, a.seq, a.perms[i], seq)
			}
		}
	}
}

func TestCompleteness(t *testing.T) {
	for test, a := range permsPositive {
		seen := map[string]bool{}
		for i := 0; i < len(a.perms); i++ {
			seen[string(a.perms[i])] = false
		}
		p := New(byteSlice(a.seq))
		for ok, seq := true, p.Current().(byteSlice); ok; ok, seq = p.Next(), p.Current().(byteSlice) {
			seen[string(seq)] = true
		}
		for seq, ok := range seen {
			if ok {
				t.Logf("%d: seen=%s", test, seq)
			} else {
				t.Errorf("%d: missing=%s", test, seq)
			}
		}
	}
}

func TestLen(t *testing.T) {
	for test, a := range permsPositive {
		p := New(byteSlice(a.seq))
		if len(a.perms) != p.Len() {
			t.Errorf("%d: seq=%s, want=%v, got=%v", test, a.seq, len(a.perms), p.Len())
		} else {
			t.Logf("%d: seq=%s, want=%v, got=%v", test, a.seq, len(a.perms), p.Len())
		}
	}
}
