package permutation

import (
	"bytes"
	"testing"
)

type permTest struct {
	seq   []byte
	perms [][]byte
}

type byteSlice []byte

func (p byteSlice) Len() int               { return len(p) }
func (p byteSlice) Less(i, j int) bool     { return p[i] < p[j] }
func (p byteSlice) Swap(i, j int)          { p[i], p[j] = p[j], p[i] }
func (p byteSlice) Equal(q Interface) bool { return bytes.Compare(p, q.(byteSlice)) == 0 }
func (p byteSlice) Copy() Interface        { return p[:] }

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
		p := New(byteSlice(a.seq))
		i := 0
		for ok, seq := true, p.Current().(byteSlice); ok; ok, seq = p.Next(), p.Current().(byteSlice) {
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
		p := New(byteSlice(a.seq))
		if len(a.perms) != p.Len() {
			t.Errorf("%d: seq=%s, want=%v, got=%v", test, a.seq, len(a.perms), p.Len())
		} else {
			t.Logf("%d: seq=%s, want=%v, got=%v", test, a.seq, len(a.perms), p.Len())
		}
	}
}
