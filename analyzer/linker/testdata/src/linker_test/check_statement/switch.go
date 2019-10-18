package check_statement

import (
	"fmt"
	"log"
)

func a() int {
	return 2
}
func checkSwitch() {
	switch i := a(); i {
	case 1:
		fmt.Println("hogehoge")
	case 2:
		log.Println("gottttttt")
	default:
		fmt.Println("fugagufa")
	}
}
