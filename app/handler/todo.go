package handler

import (
	"encoding/json"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/globalsign/mgo/bson"
	"github.com/blanccobb/go-mgo-todo-rest-example/app/db"
	"github.com/blanccobb/go-mgo-todo-rest-example/app/model"
)

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	todo := []model.Todo{}
	db.GetListByQ(db.Todos, bson.M{}, &todo)
//	if err := db.C(config.COLLECTION).Find(bson.M{}).All(&todo); err != nil {
//		respondError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
	respondJSON(w, http.StatusOK, todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := model.Todo{}
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return 
	}
	
	defer r.Body.Close()
	
//	if err := db.C(config.COLLECTION).Insert(&todo); err != nil {
	if err := db.Insert(db.Todos, &todo); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, todo)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	todo := getTodoOr404(id, w, r)
	if todo == nil {
		return
	}
	
	respondJSON(w, http.StatusOK, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	todo := getTodoOr404(id, w, r)
	if todo == nil {
		return 
	}
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return 
	}
	defer r.Body.Close()
	
//	if err := db.C(config.COLLECTION).Insert(&todo); err != nil {
	if err := db.Insert(db.Todos, &todo); err != nil {
		respondJSON(w, http.StatusInternalServerError, err.Error())
		return 
	}
	
	respondJSON(w, http.StatusOK, todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	todo := getTodoOr404(id, w, r)
	if todo == nil {
		return
	}
	
//	if err := db.C(config.COLLECTION).Remove(&todo); err != nil {
	if err := db.Delete(db.Todos, &todo); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}

func ArchiveTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	todo := getTodoOr404(id, w, r)
	if todo == nil {
		return 
	}
	
	todo.Achive()
//	if err := db.C(config.COLLECTION).Insert(&todo); err != nil {
	if err := db.Insert(db.Todos, &todo); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, todo)
}

func RestoreTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	todo := getTodoOr404(id, w, r)
	if todo == nil {
		return 
	}
	
	todo.Save()
//	if err := db.C(config.COLLECTION).Insert(&todo); err != nil {
	if err := db.Insert(db.Todos, &todo); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, todo)	
}

// getTotoOr404 gets a Todo instance if exists, or respond the 404 error otherwise
func getTodoOr404(id string, w http.ResponseWriter, r *http.Request) *model.Todo {
	todo := model.Todo{}
	
	db.Get(db.Todos, id, &todo)
//	if err := db.C(config.COLLECTION).FindId(bson.ObjectIdHex(id)).One(&todo); err != nil {
//		respondError(w, http.StatusNotFound, err.Error())
//		return nil
//	}
	
	return &todo
}
