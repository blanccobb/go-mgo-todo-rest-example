package model

import (
	"time"
	
	"github.com/jinzhu/gorm"
	"github.com/globalsign/mgo/bson"
)

type Todo struct {
	ID		bson.ObjectId		`bson:"_id" json:"id"`
	Title	string				`json:"title"`
	Achived	bool				`json:"archived"`
	Task	[]Task				`json:"tasks"`
}

func (todo *Todo) Achive() {
	todo.Achived = true
}

func (todo *Todo) Save() {
	todo.Achived = false
}

type Task struct {
	gorm.Model
	Title		string			`json:"title"`
	Priority	string			`gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
	Deadline	*time.Time		`gorm:"default:null" json:"deadline"`
	Done		bool			`json:"done"`
}

func (task *Task) Complete() {
	task.Done = true
}

func (task *Task) Undo() {
	task.Done = false
}
