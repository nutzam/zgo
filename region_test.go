package z_test

import (
	z "github.com/nutzam/zgo"
	"testing"
	"time"
)

func Test_Region(t *testing.T) {

	z.DebugOn()

	r1 := new(z.Region)
	r1.Left = 18
	r1.Right = 55
	if r1.LeftInt() != 18 {
		t.Errorf("left should be %d", 18)
	}
	if r1.RightInt() != 55 {
		t.Errorf("right should be %d", 55)
	}

	t1 := time.Now()
	t2 := time.Date(2013, 9, 10, 11, 32, 50, 200, time.Local)
	r1.Left = time.Unix(t1.Unix(), 0)
	r1.Right = time.Unix(0, t2.UnixNano())
	if r1.LeftDate().Unix() != t1.Unix() {
		t.Errorf("left should be %d, but be %d", t1.Unix(), r1.LeftDate().Unix())
	}
	if r1.RightDate().UnixNano() != t2.UnixNano() {
		t.Errorf("right should be %d, but be %d", t2.UnixNano(), r1.RightDate().UnixNano())
	}

	r2 := z.MakeRegion("[12355, 33535)")
	if r2.LeftOpen == true {
		t.Error("left should be close")
	}
	if r2.RightOpen == false {
		t.Error("right should be open")
	}
	if r2.Left != int(12355) {
		t.Error("left should be int 12355")
	}

	r3 := z.MakeRegion("[12.333, -0.555]")
	if r3.Left != float32(12.333) {
		t.Error("left should be float 12.333")
	}
	if r3.Right != float32(-0.555) {
		t.Error("right should be float -0.555, but be ", r3.Right)
	}

	r5 := z.MakeRegion("[2013-05-22 11:22:33, 2013-11-12 15:45:00]")
	if r5.LeftDate().Unix() != time.Date(2013, 5, 22, 11, 22, 33, 0, time.Local).Unix() {
		t.Error("left should be 2013-05-22 11:22:33")
	}

	r4 := z.MakeRegion("[2013-05-22, 2013-11-12]")
	if r4.LeftDate().Unix() != time.Date(2013, 5, 22, 0, 0, 0, 0, time.Local).Unix() {
		t.Error("left should be 2013-05-22")
	}
}
