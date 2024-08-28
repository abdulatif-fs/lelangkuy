package controller

import "github.com/gin-gonic/gin"

type User struct {
	Id         string `json:"user_id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Roles      string `json:"roles"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func CreateUser(ctx *gin.Context) {

}
