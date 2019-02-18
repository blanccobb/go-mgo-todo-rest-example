package app 

import (
	"net/http"
	"log"

	"github.com/gorilla/mux"
	"github.com/globalsign/mgo"
	"github.com/blanccobb/go-mgo-todo-rest-example/config"
)

type App struct {
	Router 	*mux.Router
	DB		*mgo.Database
}

func (app *App) Init() {
	
	session, err := mgo.DialWithInfo(config.GetConfig().DB)
	
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	
	app.DB.Session = session
	app.Router = mux.NewRouter()
	app.setRouters()	
}

func (app *App) setRouters() {
	app.Get("/todo", app.GetAllTodo)
	app.Post("/todo", app.CreateTodo)
	app.Get("/todo/{title}", app.GetTodo)
	app.Put("/todo/{title}", app.UpdateTodo)
	app.Delete("/todo/{title}", app.DeleteTodo)
	app.Put("/todo/{title}/archeive", app.ArcheiveTodo)
	app.Delete("/todo/{title}/archeive", app.RestoreTodo)
	
	app.Get("/todo/{title}/tasks", app.GetAllTasks)
	app.Post("/todo/{title}/tasks", app.CreateTask)
	app.Get("/todo/{title}/tasks/{id:[0-9]+}", app.GetTask)
	app.Put("/todo/{title}/tasks/{id:[0-9]+}", app.UpdateTask)
	app.Delete("/todo/{title}/tasks/{id:[0-9]+}", app.DeleteTask)
	app.Put("/todo/{title}/tasks/{id:[0-9]+}/complete", app.CompleteTask)
	app.Delete("/todo/{title}/tasks/{id:[0-9]+}/complete", app.UndoTask)
}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Response)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Response)) {
	app.Router.HandleFunc(path, f).Methods("Post")
}

func (app *App) Put(path string, f func(w http.ResponseWriter, r *http.Response)) {
	app.Router.HandleFunc(path, f).Methods("Put")
}

func (app *App) Delete(path string, f func(w http.ResponseWriter, r *http.Response)) {
	app.Router.HandleFunc(path, f).Methods("Delete")
}


// Todo Handler

func (app *App) GetAllTodo(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) CreateTodo(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) GetTodo(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) UpdateTodo(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) DeleteTodo(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) ArcheiveTodo(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) RestoreTodo(w http.ResponseWriter, r *http.Response) {
	
}

// Task Handler

func (app *App) GetAllTasks(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) CreateTask(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) GetTask(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) UpdateTask(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) DeleteTask(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) CompleteTask(w http.ResponseWriter, r *http.Response) {
	
}

func (app *App) UndoTask(w http.ResponseWriter, r *http.Response) {
	
}



