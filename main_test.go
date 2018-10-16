package main

import (
	"github.com/hetong12345/learn/single"
	"testing"
)

func TestNew(t *testing.T) {
	instance := single.GetInstance()
	instance2 := single.GetInstance()
	if instance != instance2 {
		t.Errorf("Get %p, Expect %p", instance2, instance)
	}
}
