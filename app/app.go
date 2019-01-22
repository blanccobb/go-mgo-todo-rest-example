package app 

import (


	"github.com/gorilla/mux"
	"github.com/globalsign/mgo"
	"github.com/blanccob//go-mgo-todo-rest-example/config"
)

type App struct {
	Router 	*mux.Router
	DB		*mgo.Database
}


func (app *App) init(config *config.Config) {
	
}


