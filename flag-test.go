package main

import (
	"flag"
	"fmt"
)

func main() {
	var ip = flag.Int("flagname", 1234, "help message for flagname")
	flag.Parse()
	fmt.Println("ip has value ", *ip)
	// fmt.Println("flagvar has value ", flagvar)
}

// var flagvar int

// func init() {
// 	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
// }
