package z_test

import (
	z "github.com/nutzam/zgo"
	"testing"
)

func Test_Untar(t *testing.T) {
	if z.Exists("/root/test.tar.gz") {
		err := z.Untar("/root/test.tar.gz", "/root/test")
		if err != nil {
			t.Error(err.Error())
		}
	}
}
