package main

import "fmt"

type A struct {}

type I interface {
	hoge() string
}

func (a *A) hoge() string { return "hogehoge"}



func main() {
	fmt.Println("hogehoge")	// want
	fmt.Printf("fugafuga\n") // want
}

