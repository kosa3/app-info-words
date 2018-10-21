package main

import (
	"./database"
	"./web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, err := database.DbInit()
	if err != nil {
		panic(err)
	}
	defer database.DbClose()

	web.Run()
}

