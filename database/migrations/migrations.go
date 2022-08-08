package migrations

import (
	"api/database"
	"api/model/delivery"
	"api/model/user"
)

//use gorm migrator instead

func Migrate() {

	db := database.ConnectDatabase()
	db.AutoMigrate(&delivery.Delivery{}, &user.User{})

	//db.Create(&user.User{Name: "Test User", Email: "test@email.com"})
}
