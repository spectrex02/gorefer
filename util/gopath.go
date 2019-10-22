package util

import (
	"errors"
	"os"
)

func GetGOPATH() (string, error) {
	src := "/src/"
	gosrc := "/go/src/"
	gopath, isExist := os.LookupEnv("GOPATH")
	if !isExist {
		home, isExist := os.LookupEnv("HOME")
		if !isExist {
			return "", errors.New("Not found GOPATH, Please input full path.")
		}
		return home + gosrc, nil
	}
	return gopath + src, nil
}
