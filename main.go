package main

import (
	"fmt"

	"api/database/migrations"
	"api/model/user"
	"api/mypackage"
	"api/router"
)

func main() {
	fmt.Println("Hello, modules!")
	mypackage.PrintHello()
	user.HelloUser()
	migrations.Migrate()
	r := router.SetupRouter()
	r.Run(":8080")
}
