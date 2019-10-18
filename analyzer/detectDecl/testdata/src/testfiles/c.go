package main

import "fmt"

type C struct {}

type I interface {
	hoge() string
}

func (c *C) hoge() string { return "hogehoge"}



func main() {
	fmt.Println("hogehoge")	// want
	fmt.Printf("fugafuga\n") // want
}

