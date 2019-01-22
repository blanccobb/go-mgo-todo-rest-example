package config

import (
	"time"
	"github.com/globalsign/mgo"
)

const (
	MongoDBHosts = "localhost:27017"
	AuthDatabase = "todoDB"
	AuthUserName = "todoOwner"
	AuthPassword = "todomanager"
)

type Config struct {
	DB	*mgo.DialInfo
}

func getConfig() *Config {
	return &Config {
		DB: &mgo.DialInfo {
			Addrs:		[]string{MongoDBHosts},
			Timeout:	60 * time.Second,
			Database:	AuthDatabase,
			Username:	AuthUserName,
			Password:	AuthPassword,
		},
	}
}

