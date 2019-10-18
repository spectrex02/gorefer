package main

import "fmt"

type A struct {
	Name string
	Id int
}

func hoge() {
	h := "hogehoge"
	fmt.Println(h)
}

func (a *A) ShowName() string {
	return a.Name
}

func (a *A) Err() {
	panic(a.Name)
}

func New() *A {
	return &A{
		Name: "spectre",
		Id: 1,
	}
}

func HHHH() {
	test := New()
	test.Err()
	str := test.ShowName()
	fmt.Println(str)
}