package check_statement

import "fmt"

func close() {
	fmt.Println("hogehoge")
}

func checkDefer() {
	defer close()
	for i := 0; i < 10; i++ {
		fmt.Println("llll")
	}
}