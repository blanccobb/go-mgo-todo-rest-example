package handler

import (
	"encoding/json"
	"net/http"
	
	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"github.com/globalsign/mgo/bson"
	"github.com/blanccobb/go-mgo-todo-rest-example/app/model"
	"github.com/blanccobb/go-mgo-todo-rest-example/config"
)

func GetAllTasks(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todo := getTodoOr404(db, todoId, w, r)
	
	if todo == nil {
		return 
	}
	
	tasks := []model.Task{}
	tasks = todo.Task
	
	
	respondJSON(w, http.StatusOK, tasks)
}

func CreateTasks(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todo := getTodoOr404(db, todoId, w, r)
	if todo == nil {
		return 
	}
	
	task := model.Task{}
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	
	
	where := bson.M{"_id": todoId}
	pushArray := bson.M{"$push": bson.M{"tasks": task}}
	
	if err := db.C(config.COLLECTION).Update(where, pushArray); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, task)
}

func GetTasks(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todoTitle := vars["title"]
	task := getTaskOr404(db, todoTitle, todoId, w, r)
	if task == nil {
		return 
	}
	
	respondJSON(w, http.StatusOK, task)
}

func UpdateTasks(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todoTitle := vars["title"]
	
	task := getTaskOr404(db, todoTitle, todoId, w, r)
	if task == nil {
		return 
	}
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	
	where := bson.M{"_id": todoId, "tasks.title": todoTitle}
	updateArray := bson.M{"$set": bson.M{"tasks": task}}
	
	if err := db.C(config.COLLECTION).Update(where, updateArray); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, task)
}

func DeleteTasks(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todoTitle := vars["title"]
	
	task := getTaskOr404(db, todoTitle, todoId, w, r)
	if task == nil {
		return 
	}
	
	where := bson.M{"_id": todoId, "tasks.title": todoTitle}
	deleteArray := bson.M{"$pull": bson.M{"tasks": task}}
	
	if err := db.C(config.COLLECTION).Update(where, deleteArray); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, nil)
}

//추가할 것 update와 delete는 같은 update메소드로 작동하기 때문에 문자열만 바꿔서 해볼것.

func CompleteTasks(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todoTitle := vars["title"]
	
	task := getTaskOr404(db, todoTitle, todoId, w, r)
	if task == nil {
		return 
	}

	task.Complete()
	
	where := bson.M{"_id": todoId, "tasks.title": todoTitle}
	updateArray := bson.M{"$set": bson.M{"tasks": task}}
	
	if err := db.C(config.COLLECTION).Update(where, updateArray); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, task)
}

func UndoTasks(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todoTitle := vars["title"]
	
	task := getTaskOr404(db, todoTitle, todoId, w, r)
	if task == nil {
		return 
	}

	task.Undo()
	
	where := bson.M{"_id": todoId, "tasks.title": todoTitle}
	updateArray := bson.M{"$set": bson.M{"tasks": task}}
	
	if err := db.C(config.COLLECTION).Update(where, updateArray); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, task)
}

func getTaskOr404(db *mgo.Database, title string, id string,  w http.ResponseWriter, r *http.Request) *model.Task {
	todo := model.Todo{}
	
	if err := db.C(config.COLLECTION).Find(bson.M{"tasks":{"title": title}}).One(&todo); err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	
	return &todo.Task
	
}

