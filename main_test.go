package main

import (
	"fmt"
	"github.com/hetong12345/learn/single"
	"testing"
)

func TestNew(t *testing.T) {
	leshi := &single.LeshiGlasses{FirstCost: 100}
	var glassesbuilder single.Builder
	glassesbuilder = &single.ShanWeiBuilder{}
	glasses := leshi.GetGlasses(glassesbuilder)
	fmt.Println("glasses's price is: ", glasses.Price, " glasses from :", glasses.From)
	glassesbuilder = &single.ShenZhenBuilder{}
	glasses = leshi.GetGlasses(glassesbuilder)
	fmt.Println("glasses's price is: ", glasses.Price, " glasses from :", glasses.From)

	return
}
