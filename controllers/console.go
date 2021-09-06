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

func (consoleController ConsoleController) GetAllConsoles(resWriter http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var allConsoles []models.Console

	consoleController.session.DB("consoles").C("consoles").Find(nil).All(&allConsoles)

	allConsolesJson, err := json.Marshal(allConsoles)

	if err != nil {
		fmt.Println(err)
	}

	resWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(resWriter, "%s\n", allConsolesJson)
}

func (consoleController ConsoleController) GetConsole(resWriter http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		resWriter.WriteHeader(http.StatusNotFound)
	}

	objectId := bson.ObjectIdHex(id)
	console := models.Console{}

	if err := consoleController.session.DB("consoles").C("consoles").FindId(objectId).One(&console); err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		return
	}

	consoleJson, err := json.Marshal(console)
	if err != nil {
		fmt.Println(err)
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(resWriter, "%s\n", consoleJson)
}

func (consoleController ConsoleController) CreateConsole(resWriter http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	newConsole := models.Console{}

	err := json.NewDecoder(req.Body).Decode(&newConsole)

	if err != nil {
		fmt.Println(err)
		resWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(resWriter, "the informed json is invalid!\n")
		return
	}

	newConsole.Id = bson.NewObjectId()

	consoleController.session.DB("consoles").C("consoles").Insert(newConsole)

	newConsoleJson, err := json.Marshal(newConsole)

	if err != nil {
		fmt.Println(err)
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusCreated)
	fmt.Fprintf(resWriter, "%s\n", newConsoleJson)
}

func (consoleController ConsoleController) DeleteConsole(resWriter http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		resWriter.WriteHeader(http.StatusNotFound)
	}

	objectId := bson.ObjectIdHex(id)

	if err := consoleController.session.DB("consoles").C("consoles").RemoveId(objectId); err != nil {
		resWriter.WriteHeader(http.StatusNotFound)
		return
	}

	resWriter.WriteHeader(http.StatusOK)
	fmt.Fprint(resWriter, "Deleted user ", objectId, "\n")
}
