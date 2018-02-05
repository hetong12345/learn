/**
*定位到Helper目录下执行 go test 测试函数

执行 go test -bench=. 测试执行效率

*/

package helper

import (
	"fmt"
	"testing"
)

func TestCryptoSha1(t *testing.T) {
	a := "123"
	str := CryptoSha1(a)
	fmt.Println(a)
	fmt.Println(str)
}

func TestValidateMobile(t *testing.T) {
	result := ValidateMobile("130")
	if result != false {
		t.Error("the result is wrong")
	}
}

func BenchmarkValidateMobile(B *testing.B) {
	for i := 0; i < B.N; i++ {
		ValidateMobile("130")
	}
}
