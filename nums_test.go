package z_test

import (
	z "github.com/nutzam/zgo"
	"testing"
)

func Test_Index_Of_Bytes(t *testing.T) {
	bs := []byte{'a', 'b', 'c', 'e'}
	if z.IndexOfBytes(bs, 'b') != 1 {
		t.Error("b should be 1")
	}
	if z.IndexOfBytes(bs, 'e') != 3 {
		t.Error("e should be 3")
	}
	if z.IndexOfBytes(bs, 'z') != -1 {
		t.Error("e should be -1")
	}
}

func Test_Index_Of_Strings(t *testing.T) {
	bs := []string{"abc", "bcd", "cde", "efg"}
	if z.IndexOfStrings(bs, "bcd") != 1 {
		t.Error("bcd should be 1")
	}
	if z.IndexOfStrings(bs, "z123") != -1 {
		t.Error("z123 should be -1")
	}
}

func Test_Is_In_Strings(t *testing.T) {
	bs := []string{"abc", "bcd", "cde", "efg"}
	if !z.IsInStrings(bs, "bcd") {
		t.Error("bcd should be 1")
	}
	if z.IsInStrings(bs, "z123") {
		t.Error("z123 should be -1")
	}
}
