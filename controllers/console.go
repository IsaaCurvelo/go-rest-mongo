package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IsaaCurvelo/go-rest-mongo/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ConsoleController struct {
	session *mgo.Session
}

func NewConsoleController(session *mgo.Session) *ConsoleController {
	return &ConsoleController{session}
}

func (consoleController ConsoleController) GetConsole(rw http.ResponseWriter, rqst *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		rw.WriteHeader(http.StatusNotFound)
	}

	objectId := bson.ObjectIdHex(id)
	console := models.Console{}

	if err := consoleController.session.DB("consoles").C("consoles").FindId(objectId).One(&console); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	jsonConsole, err := json.Marshal(console)
	if err != nil {
		fmt.Println(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "%s\n", jsonConsole)
}

func (consoleController ConsoleController) CreateConsole(rw http.ResponseWriter, rqst *http.Request, p httprouter.Params) {

}

func (consoleController ConsoleController) DeleteConsole(rw http.ResponseWriter, rqst *http.Request, p httprouter.Params) {

}
