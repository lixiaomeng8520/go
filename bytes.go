package main

import (
	"bytes"
	"fmt"
	"os"
)

func bytes1() {
	f, err := os.Open("1")
	fmt.Println(f, err)
	defer f.Close()

	b := []byte{0}

	n, err := f.Read(b)
	fmt.Println(n, err, b)

	n, err = f.Read(b)
	fmt.Println(n, err, b)

	n, err = f.Read(b)
	fmt.Println(n, err, b)

	n, err = f.Read(b)
	fmt.Println(n, err, b)
}

func bytes2() {
	f, _ := os.Open("1")
	defer f.Close()

	b := bytes.Buffer{}
	fmt.Println(b)

	n, err := b.ReadFrom(f)
	fmt.Println(n, err, b.Bytes(), b.Len(), b.Cap())
}
