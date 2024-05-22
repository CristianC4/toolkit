package toolkit

import "testing"

func TestTools_RamdonsString(t *testing.T) {
	var tools Tools
	s := tools.RamdonsString(10)
	if len(s) != 10 {
		t.Error("Expected random string of length 10, but got ", len(s))
	}
}
