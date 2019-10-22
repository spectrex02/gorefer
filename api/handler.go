package api

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//handler function is here.
//func responseHandler(c echo.Context) error {
//
//}

func getPackage(c echo.Context) error {
	//get package name analyzed
	name := c.Param("package")
	log.Printf("Return result of analysing (package name: %s) to client\n", name)
	return c.String(http.StatusOK, name)
}

func getFunction(c echo.Context) error {
	name := c.Param("package")
	log.Printf("Return function list package(%s) to client\n", name)
	data := serveJson(name, "function-list")
	return c.JSONBlob(http.StatusOK, data)
}

func getStruct(c echo.Context) error {
	name := c.Param("package")
	log.Printf("Return structure list package(%s) to client\n", name)
	data := serveJson(name, "struct-list")
	return c.JSONBlob(http.StatusOK, data)
}

func getInterface(c echo.Context) error {
	name := c.Param("package")
	log.Printf("Return interface list package(%s) to client\n", name)
	data := serveJson(name, "interface-list")
	return c.JSONBlob(http.StatusOK, data)
}

func getVar(c echo.Context) error {
	name := c.Param("package")
	log.Printf("Return variable list package(%s) to client\n", name)
	data := serveJson(name, "var-list")
	return c.JSONBlob(http.StatusOK, data)
}

func getNetwork(c echo.Context) error {
	name := c.Param("package")
	log.Printf("Return function network package(%s) to client\n", name)
	return c.String(http.StatusOK, "todo")
}
func serveJson(packageName string, filename string) []byte {
	f, err := os.Open(path(packageName, filename))
	if err != nil {
		panic(err)
	}
	//data := make([]byte, 1024)
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return data
}


//make filepath to result/${pakcagename}.json
func path(pakcageName string, filename string) string {
	path := "result/" + pakcageName + "/"
	return path + filename + ".json"
}
