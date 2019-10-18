package check_statement

import "log"

func num() []int {
	return []int{0,1,2,3,4,5,6,7,8,9}
}

func show(i int) {
	log.Fatalln(i)
}
func checkForRange() {
	for i := range num() {
		show(i)
	}
}