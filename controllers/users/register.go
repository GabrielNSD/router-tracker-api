package users

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"

	"api/database"
	"api/model/user"
)

type Register struct {
	Email    string `form:"email"`
	Name     string `form:"name"`
	Password string `form:"password"`
}

func HandleRegister(c *gin.Context) {
	var register Register

	if c.ShouldBind(&register) == nil {
		log.Println(register.Email)
		log.Println(register.Name)
		log.Println(register.Password)
		if register.Email == "" || register.Name == "" || register.Password == "" {
			c.String(400, "incorrect form submisison")
			return
		}

		password := []byte(register.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			c.Status(500)
			panic(err)
		}
		convertedHash := string(hashedPassword[:])
		log.Println(convertedHash)

		user := user.User{Name: register.Name, Email: register.Email, Password: convertedHash}

		db := database.ConnectDatabase()
		result := db.Create(&user)

		log.Println("id", user.ID)
		log.Println("error", result.Error)
		log.Println("affected rows", result.RowsAffected)
		if result.Error != nil {
			c.Status(500)
			return
		}
	}

	c.String(200, "success")
}
