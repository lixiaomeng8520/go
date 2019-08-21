package main

import (
	"fmt"
	"log"
	"net"
	"io"
	"time"
)

func main() {
	// l : *TCPListener
	l, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatal(err)
	}
	
	for {
		fmt.Println("---------------------------------------------------")
		// cLeft : *TCPConn
		cLeft, err := l.Accept()
		if err != nil {
			fmt.Println("err:", err)
			continue
		}

		go func() {
			defer cLeft.Close()

			// cRight : *TCPConn
			cRight, err := net.Dial("tcp", "125.46.11.188:80")
			if err != nil {
				fmt.Println("err:", err)
				return
			}
			defer cRight.Close()

			// io.Copy会阻塞，直到检测到EOF(一方断开连接)或者错误
			// SetDeadline是指时间内没有读写，则判定超时

			fmt.Println("start to transfer")

			ch := make(chan bool)
			
			go func() {
				_, err := io.Copy(cRight, cLeft)
				fmt.Println("left -> right", err)
				cRight.SetDeadline(time.Now())
				ch <- true
			}()
			

			go func() {
				_, err := io.Copy(cLeft, cRight)
				fmt.Println("right -> left", err)
				cLeft.SetDeadline(time.Now())
			}()
			
			<-ch

			fmt.Println("over")
		}()

	}
}

func relay(left, right net.Conn) (int64, int64, error) {
	type res struct {
		N   int64
		Err error
	}
	ch := make(chan res)

	go func() {
		n, err := io.Copy(right, left)
		fmt.Println("left->right", n, err)
		//right.SetDeadline(time.Now()) // wake up the other goroutine blocking on right
		//left.SetDeadline(time.Now())  // wake up the other goroutine blocking on left
		ch <- res{n, err}
	}()
	n, err := io.Copy(left, right)
	fmt.Println("right->left", n, err)
	//right.SetDeadline(time.Now()) // wake up the other goroutine blocking on right
	left.SetDeadline(time.Now())  // wake up the other goroutine blocking on left
	rs := <-ch

	if err == nil {
		err = rs.Err
	}
	return n, rs.N, err
}