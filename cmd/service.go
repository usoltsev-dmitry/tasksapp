package main

import (
	"fmt"
	"log"
	"tasksapp/pkg/storage"
	"tasksapp/pkg/storage/postgres"
)

var db storage.Interface

func main() {
	var err error

	connstr := "host=localhost port=5433 user=postgres password=pg1234 dbname=postgres sslmode=disable"
	db, err = postgres.New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	db = &postgres.Storage{}
	tasks, err := db.GetTasks(1, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
}
