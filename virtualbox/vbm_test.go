package virtualbox

import (
	"testing"
)

func init() {
	verbose = true
}

func TestVBMOut(t *testing.T) {
	b, err := vbmOut("list", "vms")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", b)
}
