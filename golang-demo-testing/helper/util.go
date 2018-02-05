package helper

import (
	"crypto/sha1"
	"fmt"
	"io"
	"regexp"
)

const (
	regex = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\\d{8}$"
)

func CryptoSha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("x%", t.Sum(nil))
}

func ValidateMobile(mobileNum string) bool {
	reg := regexp.MustCompile(regex)
	return reg.MatchString(mobileNum)
}
