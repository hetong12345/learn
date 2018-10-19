package main

import (
	"github.com/hetong12345/learn/single"
	"testing"
)

func TestNew(t *testing.T) {
	company := &single.BaseCompany{}
	developingCompany := &single.DevelopingCompany{NewCompany: company}
	developingCompany.Showing()
	bigCompany := &single.NewBigCompany{NewCompany: developingCompany}
	bigCompany.Showing()
}
