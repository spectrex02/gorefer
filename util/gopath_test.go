package util

import (
	"fmt"
	"testing"
)

func TestGetGOPATH(t *testing.T) {
	wanted := "/Users/spectre/go/"
	result, err := GetGOPATH()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	if wanted != result {
		fmt.Errorf("hoge")
	}
}
