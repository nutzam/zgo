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

func Test_SBuilder(t *testing.T) {
	sb := z.StringBuilder()
	sb.Append("abc")
	if sb.String() != "abc" {
		t.Errorf("sb has error, should be %s, but be %s", "abc", sb.String())
	}
	sb.Append('d')
	if sb.String() != "abcd" {
		t.Errorf("sb has error, should be %s, but be %s", "abcd", sb.String())
	}
	if sb.Len() != 4 {
		t.Errorf("sb length should be 4, but be %d", sb.Len())
	}
	sb2 := z.StringBuilder()
	sb2.Append([]string{"acb", "111", "gdgdg"})
	if sb2.String() != "[acb 111 gdgdg]" {
		t.Errorf("sb2 has error, should be %s, but be %s", "[acb 111 gdgdg]", sb2.String())
	}
}
