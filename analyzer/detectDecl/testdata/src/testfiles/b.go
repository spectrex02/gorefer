package main


type B struct {
	name string
}
func (b *B) hoge() string {
	return b.name
}

const testVar  = "string"
var bvar int