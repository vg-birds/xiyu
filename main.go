package main

import (
	db "test/database"
)

func main() {
	defer db.SqlDB.Close()

	router := initRouter()
	router.Run(":8000")
}
