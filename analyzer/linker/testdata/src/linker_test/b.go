package linker_test

import "fmt"

type B interface {
	hoge() BB
}
type BB struct {
	name string
}

type BBB struct {
	name string
}
func NewBBB() *BBB {
	return &BBB{name:"hoge"}
}

func (bbb *BBB) hoge() BB {
	return BB{name:"hogehogehoge"}
}

func (bb *BB) show() {
	fmt.Println(bb.name)
}

func complexType() {
	NewBBB().hoge()
	b := NewBBB().hoge()
	fmt.Println(b)

	if true {
		b.show()
	}
}
