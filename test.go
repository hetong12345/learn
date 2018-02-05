package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	target := "World"
	if len(os.Args) > 1 { /* os */
		target = os.Args[2]
	}
	fmt.Println("Hello", target)
	Root, _ = os.Getwd()
	fmt.Println(Root)
	InitLocalIp()
	fmt.Println(LocalIp)

	fmt.Println(os.Hostname())
}

var Root string
var LocalIp string

func InitLocalIp() {
	if true {
		conn, err := net.DialTimeout("tcp", "47.95.253.231:6030", time.Second*10)
		if err != nil {
			log.Println("get local addr failed !")
		} else {
			fmt.Println(conn.RemoteAddr(), conn.LocalAddr())
			LocalIp = strings.Split(conn.LocalAddr().String(), ":")[0]
			conn.Close()
		}
	} else {
		log.Println("hearbeat is not enabled, can't get localip")
	}
}
