package test

import (
	"fmt"
	"github.com/hetong12345/learn/golang-demo-testing2/helper"
	"testing"
)

func TestValidateMobile(t *testing.T) {
	result := helper.ValidateMobile("130")
	if result != false {
		t.Error("the result is wrong")
	}
	fmt.Println("done")
}
