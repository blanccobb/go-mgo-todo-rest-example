package handler

import (
	"encoding/json"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/blanccobb/go-mgo-todo-rest-example/app/model"
	"github.com/blanccobb/go-mgo-todo-rest-example/config"
)


func GetAllTodo(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	todo := []model.Todo{}
	if err := session.DB(config.AuthDatabase).C(config.COLLECTION).Find(bson.M{}).All(&todo); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, todo)
}

func CreateTodo(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	todo := model.Todo{}
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return 
	}
	
	defer r.Body.Close()
	
	if err := session.DB(config.AuthDatabase).C(config.COLLECTION).Insert(&todo); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, todo)
}

func GetTodo(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	todo := getTodoOr404(session, id, w, r)
	if todo == nil {
		return
	}
	
	respondJSON(w, http.StatusOK, todo)
}

func UpdateTodo(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	todo := getTodoOr404(session, id, w, r)
	if todo == nil {
		return 
	}
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return 
	}
	defer r.Body.Close()
	
	if err := session.DB(config.AuthDatabase).C(config.COLLECTION).Insert(&todo); err != nil {
		respondJSON(w, http.StatusInternalServerError, err.Error())
		return 
	}
	
	respondJSON(w, http.StatusOK, todo)
}

func DeleteTodo(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	todo := getTodoOr404(session, id, w, r)
	if todo == nil {
		return
	}
	if err := session.DB(config.AuthDatabase).C(config.COLLECTION).Remove(&todo); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}

func ArchiveTodo(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	todo := getTodoOr404(session, id, w, r)
	if todo == nil {
		return 
	}
	
	todo.Achive()
	if err := session.DB(config.AuthDatabase).C(config.COLLECTION).Insert(&todo); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, todo)
}

func RestoreTodo(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	todo := getTodoOr404(session, id, w, r)
	if todo == nil {
		return 
	}
	
	todo.Save()
	if err := session.DB(config.AuthDatabase).C(config.COLLECTION).Insert(&todo); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, todo)	
}

// getTotoOr404 gets a Todo instance if exists, or respond the 404 error otherwise
func getTodoOr404(session *mgo.Session, id string, w http.ResponseWriter, r *http.Request) *model.Todo {
	todo := model.Todo{}
	
	if err := session.DB(config.AuthDatabase).C(config.COLLECTION).FindId(bson.ObjectIdHex(id)).One(&todo); err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	
	return &todo
}