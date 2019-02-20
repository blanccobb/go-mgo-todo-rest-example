package app 

import (
	"net/http"
	"log"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/globalsign/mgo"
	"github.com/blanccobb/go-mgo-todo-rest-example/config"
	"github.com/blanccobb/go-mgo-todo-rest-example/app/handler"
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
	defer session.Close()
	
	session.SetMode(mgo.Monotonic, true)
	fmt.Printf("Connected to %v!\n", session.LiveServers())
	
	
	app.DB.Session = session.New()
	app.Router = mux.NewRouter()
	app.setRouters()	
	
}

func (app *App) setRouters() {
	app.Get("/todo", app.GetAllTodo)
	app.Post("/todo", app.CreateTodo)
	app.Get("/todo/{id}", app.GetTodo)
	app.Put("/todo/{id}", app.UpdateTodo)
	app.Delete("/todo/{id}", app.DeleteTodo)
	app.Put("/todo/{id}/archeive", app.ArcheiveTodo)
	app.Delete("/todo/{id}/archeive", app.RestoreTodo)
	
	app.Get("/todo/{id}/tasks", app.GetAllTasks)
	app.Post("/todo/{id}/tasks", app.CreateTask)
	app.Get("/todo/{id}/tasks/{title}", app.GetTask)
	app.Put("/todo/{id}/tasks/{title}", app.UpdateTask)
	app.Delete("/todo/{id}/tasks/{title}", app.DeleteTask)
	app.Put("/todo/{id}/tasks/{title}/complete", app.CompleteTask)
	app.Delete("/todo/{id}/tasks/{title}/complete", app.UndoTask)
}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("Post")
}

func (app *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("Put")
}

func (app *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("Delete")
}


// Todo Handler

func (app *App) GetAllTodo(w http.ResponseWriter, r *http.Request) {
	handler.GetAllTodo(app.DB, w, r)	
}

func (app *App) CreateTodo(w http.ResponseWriter, r *http.Request) {
	handler.CreateTodo(app.DB, w, r)	
}

func (app *App) GetTodo(w http.ResponseWriter, r *http.Request) {
	handler.GetTodo(app.DB, w, r)
}

func (app *App) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	handler.UpdateTodo(app.DB, w, r)
}

func (app *App) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	handler.DeleteTodo(app.DB, w, r)
}

func (app *App) ArcheiveTodo(w http.ResponseWriter, r *http.Request) {
	handler.ArchiveTodo(app.DB, w, r)
}

func (app *App) RestoreTodo(w http.ResponseWriter, r *http.Request) {
	handler.RestoreTodo(app.DB, w, r)
}

// Task Handler

func (app *App) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	handler.GetAllTasks(app.DB, w, r)
}

func (app *App) CreateTask(w http.ResponseWriter, r *http.Request) {
	handler.CreateTasks(app.DB, w, r)
}

func (app *App) GetTask(w http.ResponseWriter, r *http.Request) {
	handler.GetTasks(app.DB, w, r)
}

func (app *App) UpdateTask(w http.ResponseWriter, r *http.Request) {
	handler.UpdateTasks(app.DB, w, r)
}

func (app *App) DeleteTask(w http.ResponseWriter, r *http.Request) {
	handler.DeleteTasks(app.DB, w, r)
}

func (app *App) CompleteTask(w http.ResponseWriter, r *http.Request) {
	handler.CompleteTasks(app.DB, w, r)
}

func (app *App) UndoTask(w http.ResponseWriter, r *http.Request) {
	handler.UndoTasks(app.DB, w, r)
}

func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}


