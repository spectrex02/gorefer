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

func getPackageName(c echo.Context) error {
	//get package name analyzed
	name := c.Param("package")
	log.Printf("Return result of analysing (package name: %s) to client\n", name)
	data := serveJson(name)
	return c.JSONBlob(http.StatusOK, data)

}

func serveJson(filename string) []byte {
	f, err := os.Open(path(filename))
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
func path(filename string) string {
	path := "result/"
	return path + filename + ".json"
}