package api

//server to serve result of analysis as json format

import (
"github.com/labstack/echo"
"log"
"net/http"
)

func Serve() {
	server := echo.New()
	server.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//get package name and serve json file
	server.GET("/result/:package", getPackage)
	server.GET("/result/:package/struct", getStruct)
	server.GET("/result/:package/interface", getInterface)
	server.GET("/result/:package/function", getFunction)
	server.GET("/result/:package/var", getVar)
	server.GET("/resutl/:package/func-network", getNetwork)

	log.Printf("Server start at http://localhost:5731")
	err := server.Start(":5731")
	if err != nil {
		panic(err)
	}

}