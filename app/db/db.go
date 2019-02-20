package dao

import (
	"log"
	
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/blanccobb/go-mgo-todo-rest-example/config"
)


var (
	Session		*mgo.Session
	
	Todos		*mgo.Collection
)

func Init() {
	
	Session, err := mgo.DialWithInfo(config.GetConfig().DB)
	if err != nil {
		log.Fatal("CreateSession: %s\n", err)
	}
	
	// COLLECTION
	Todos = Session.DB(config.AuthDatabase).C(config.COLLECTION)
	
}

func Close() {
	Session.Close()
}

// Insert...
func Insert(collection *mgo.Collection, i interface{}) {
	collection.Insert(i)
}

// Like GET...
func Get(collection *mgo.Collection, id string, i interface{}) {
	collection.FindId(bson.ObjectIdHex(id)).One(i)
}

func GetListByQ(collection *mgo.Collection, q interface{}, i interface{}) {
	collection.Find(q).All(i)
}

// Update...
func Update(collection *mgo.Collection, q interface{}, i interface{}) error {
	return collection.Update(q, i)
}

// Delete...
func Delete(collection *mgo.Collection, i interface{}) error {
	return collection.Remove(i)
}


//func Err(err error) bool {
//	if err != nil {
//		fmt.Println(err)
//		if err.Error() == "not found" {
//			return true
//		}
//		return false
//	}
//	return true
//}

