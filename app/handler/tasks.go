package handler

import (
	"encoding/json"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/globalsign/mgo/bson"
	"github.com/blanccobb/go-mgo-todo-rest-example/app/db"
	"github.com/blanccobb/go-mgo-todo-rest-example/app/model"
)


func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todo := getTodoOr404(todoId, w, r)
	
	if todo == nil {
		return 
	}
	
	tasks := []model.Task{}
	tasks = todo.Task
	
	
	respondJSON(w, http.StatusOK, tasks)
}

func CreateTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todo := getTodoOr404(todoId, w, r)
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
	
//	if err := db.C(config.COLLECTION).Update(where, pushArray); err != nil {
	if err := db.Update(db.Todos, where, pushArray); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todoTitle := vars["title"]
	task := getTaskOr404(todoTitle, todoId, w, r)
	if task == nil {
		return 
	}
	
	respondJSON(w, http.StatusOK, task)
}

func UpdateTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todoTitle := vars["title"]
	
	task := getTaskOr404(todoTitle, todoId, w, r)
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
	
	if err := db.Update(db.Todos, where, updateArray); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, task)
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todoTitle := vars["title"]
	
	task := getTaskOr404(todoTitle, todoId, w, r)
	if task == nil {
		return 
	}
	
	where := bson.M{"_id": todoId, "tasks.title": todoTitle}
	deleteArray := bson.M{"$pull": bson.M{"tasks": task}}
	
	if err := db.Update(db.Todos, where, deleteArray); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, nil)
}

//추가할 것 update와 delete는 같은 update메소드로 작동하기 때문에 문자열만 바꿔서 해볼것.

func CompleteTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todoTitle := vars["title"]
	
	task := getTaskOr404(todoTitle, todoId, w, r)
	if task == nil {
		return 
	}

	task.Complete()
	
	where := bson.M{"_id": todoId, "tasks.title": todoTitle}
	updateArray := bson.M{"$set": bson.M{"tasks": task}}
	
	if err := db.Update(db.Todos, where, updateArray); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, task)
}

func UndoTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	todoId := vars["id"]
	todoTitle := vars["title"]
	
	task := getTaskOr404(todoTitle, todoId, w, r)
	if task == nil {
		return 
	}

	task.Undo()
	
	where := bson.M{"_id": todoId, "tasks.title": todoTitle}
	updateArray := bson.M{"$set": bson.M{"tasks": task}}
	
	if err := db.Update(db.Todos, where, updateArray); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, task)
}

func getTaskOr404(title string, id string,  w http.ResponseWriter, r *http.Request) *model.Task {
	todo := model.Todo{}
	
	//find()안에 조건 다시 확인
//	if err := db.C(config.COLLECTION).Find(bson.M{"_id": id, "tasks.title": title}).One(&todo); err != nil {
//			respondError(w, http.StatusNotFound, err.Error())
//			return nil
//	}
	db.GetByQ(db.Todos, bson.M{"_id": id, "tasks.title": title}, &todo)
	
	//배열이 아닌 하나의 Task를 반환
	
	return &todo.Task[0]
}
