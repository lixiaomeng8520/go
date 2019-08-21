package main

import (
	"fmt"
	"log"
	"net"
	// "reflect"
)

func main() {
	// l : *TCPListener
	l, err := net.Listen("tcp", "0.0.0.0:10000")
	if err != nil {
		log.Fatal(err)
	}
	
	for {
		// c : *TCPConn
		cFrom, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer cFrom.Close()
		
		cTo, err := net.Dial("tcp", "http://s.dahe.cn")
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer cTo.Close()

		

	}
}