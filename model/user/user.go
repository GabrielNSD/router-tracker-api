package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

func HelloUser() {
	fmt.Println("hello from user")
}

func CreateUser(c *gin.Context) {

}
