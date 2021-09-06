package main

import (
	"fmt"
	"net/http"

	"github.com/IsaaCurvelo/go-rest-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	router := httprouter.New()
	consoleController := controllers.NewConsoleController(getSession())

	router.GET("/consoles/:id", consoleController.GetConsole)
	router.POST("/consoles", consoleController.CreateConsole)
	router.DELETE("/consoles/:id", consoleController.DeleteConsole)

	defer http.ListenAndServe(":8080", router)
	fmt.Println("o servidor http subiu...")
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://mongodb:27017")
	if err != nil {
		fmt.Println("porra, valeu, cara!!")
		panic(err)
	}

	return session
}
