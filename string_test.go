package z_test

import (
	z "github.com/nutzam/zgo"
	"testing"
)

func Test_Is_Space(t *testing.T) {
	if !z.IsSpace(' ') {
		t.Error("' ' is space")
	}
	if !z.IsSpace('\t') {
		t.Error("'\t' is space")
	}
	if !z.IsSpace('\n') {
		t.Error("'\n' is space")
	}
	if !z.IsSpace('\r') {
		t.Error("'\r' is space")
	}
}

func Test_Trim_Extra_Space(t *testing.T) {
	s1 := " a b  c    d\te \n"
	s2 := z.TrimExtraSpace(s1)
	if s2 != "a b c d e" {
		t.Errorf("can't trim [%s]", s1)
	}
}

func Test_Dup_Char(t *testing.T) {
	if z.DupChar('a', 5) != "aaaaa" {
		t.Error("DupChar 'a' 5 times != 'aaaaa'")
	}
}

func Test_Dup(t *testing.T) {
	if z.Dup("abc", 3) != "abcabcabc" {
		t.Error("Dup 'abc' 3 times != 'abcabcabc'")
	}
}
