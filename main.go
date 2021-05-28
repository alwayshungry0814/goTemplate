// main.go

package main

import (
	"log"
	"os"
	"template/controller"
)

func main() {
	a := controller.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	log.Println("Web server is available on port 8080")
	a.Run("localhost:8080")
}
