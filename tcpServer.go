package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	l, _ := net.Listen("tcp", "127.0.0.1:8888")

	conn, _ := l.Accept()
	//msg := make([]byte, 4096)
	br := bufio.NewReader(conn)
	i := 0
	for {
		i++
		fmt.Println(i)

		line, err := br.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s\n", line)

	}

	resp := "HTTP/1.1 200 OK\nDate: Mon, 01 Apr 2019 08:50:45 GMT;\nContent-Length: 5\nContent-Type: text/plain; charset=utf-8\n\nhello"
	conn.Write([]byte(resp))

}
