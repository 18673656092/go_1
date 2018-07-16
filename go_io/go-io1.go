package main

import (
	"net"
	"fmt"
)

func main() {
	conn,err := net.Dial("ip4:icmp","www.baidu.com")
	fmt.Println(conn,err)
}
