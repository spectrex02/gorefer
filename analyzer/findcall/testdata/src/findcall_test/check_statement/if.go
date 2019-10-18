package check_statement

import (
	"fmt"
	"log"
)

type C struct {
	name string
	value int
}

func (c *C) show() {
	fmt.Printf("name: %v, value: %v\n", c.name, c.value)
}

func (c *C) Value() int {
	return c.value
}

func (c *C) Name() string {
	return c.name
}

func newC(name string, value int) *C {
	return &C{
		name:  name,
		value: value,
	}
}

func Cmain() {
	c := newC("testc", 1)
	cc := newC("testcc", 10)
	if v := cc.Value(); v > c.value {
		log.Printf("hoge")
	} else {
		log.Println(c.name)
	}
}