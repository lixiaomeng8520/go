package main

import (
	"fmt"
	"log"
	"net"
	"io/ioutil"
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
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		b, err := ioutil.ReadAll(c)
		fmt.Println(string(b), err)
	}
}