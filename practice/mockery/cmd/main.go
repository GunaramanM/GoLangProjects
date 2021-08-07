package main

import (
	"awesomeProject/practice/mockery/pkg/model"
	db "awesomeProject/practice/mockery/pkg/repository"
	. "awesomeProject/practice/mockery/router"
)

func main() {
	dbHost := "localhost:27017"
	db.Init(&model.Database{
		Driver:   "mongodb",
		Endpoint: dbHost})
	defer db.Exit()

	Router()

}
