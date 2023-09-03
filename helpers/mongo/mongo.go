package mongo

import (
	"fmt"

	"github.com/shivamsouravjha/influenza/config"
	"gopkg.in/mgo.v2"
)

var (
	DbSession *mgo.Session
	Db        *mgo.Database
)

// package main

// import (
// 	"fmt"

// 	"gopkg.in/mgo.v2"
// )

// func InitDb() (*mgo.Database, error) {
// 	mongoDBName := "ds"
// 	mongoURI := "mongodb://admin:adminpassword@localhost:27017/Influenza"
// 	fmt.Println("Connecting to MongoDB:", mongoURI)

// 	session, err := mgo.Dial(mongoURI)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
// 	}
// 	defer session.Close()

// 	// Test the connection by performing a simple query
// 	if err := session.Ping(); err != nil {
// 		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
// 	}

// 	return session.DB(mongoDBName), nil
// }

// func main() {
// 	db, err := InitDb()
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}
// 	defer db.Session.Close()

// 	fmt.Println("Connected to MongoDB")
// 	// Your application logic here
// }

func InitDb() (*mgo.Database, error) {
	// Construct the MongoDB URI
	mongoURI := "mongodb://" + config.Get().MongoUsername + ":" + config.Get().MongoPassword + "@127.0.0.1:27017"
	fmt.Println(mongoURI)
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		fmt.Println("ssasaasda", err.Error())

		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	if config.Get().DBUserName != "" && config.Get().DBPassword != "" {
		err = session.DB(config.Get().DBName).Login(config.Get().DBUserName, config.Get().DBPassword)
		if err != nil {
			return nil, fmt.Errorf("failed to authenticate to MongoDB: %v", err)
		}
	}

	db := session.DB(config.Get().DBName)
	fmt.Println("Connected to MongoDB")
	return db, nil
}

func Client() *mgo.Database {
	return Db
}
