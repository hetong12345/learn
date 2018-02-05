package hello

import (
	//"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello()
	expect := "hello func in package hello."
	if got != expect {
		t.Errorf("got [%s] expected [%s]", got, expect)
	}
}
