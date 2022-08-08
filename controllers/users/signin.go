package users

import (
	"api/database"
	"log"

	"api/model/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Signin struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func HandleSignin(c *gin.Context) {
	var signin Signin
	db := database.ConnectDatabase()

	if c.ShouldBind(&signin) == nil {
		if signin.Email == "" || signin.Password == "" {
			c.String(400, "incorrect form submission")
			return
		}

		var user user.User
		result := db.Where("email = ?", signin.Email).First(&user)
		if result.Error != nil {
			c.Status(500)
			return
		}

		unmatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signin.Password))
		log.Println("unmatch", unmatch)
		if unmatch == nil {
			//c.Status(403)
			// generate jwt
			// return jwt
		} else {
			c.String(403, "Wrong password")
			return
		}
	}

	c.String(200, "success")
}
