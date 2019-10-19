package check_statement

import "fmt"

func rungo() {
	fmt.Println("hogohoge")
}
func checkGo() {
	go rungo()
}
