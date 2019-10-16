package main

import "fmt"

type A struct {}

func (a *A) hoge() string { return "hogehoge"}

func main() {
	fmt.Println("hogehoge")	// want
	fmt.Printf("fugafuga\n") // want
}
