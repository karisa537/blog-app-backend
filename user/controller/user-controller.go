package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karisa537/blog-app/user/model"
	"github.com/karisa537/blog-app/user/service"
)

// type UserController struct {
// 	Service service.UserService
// }

// func NewUserController(service service.UserService) *UserController {
//     return &UserController{Service: service}
// }

// func (uc *UserController) Register(c *gin.Context) {
// 	var user model.User
// 	if err := c.ShouldBindJSON(&user); err != nil{
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := uc.Service.Register(&user); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
// }

// func (uc *UserController) Login(c *gin.Context) {
//     var credentials struct {
//         Username string `json:"username"`
//         Password string `json:"password"`
//     }
//     if err := c.ShouldBindJSON(&credentials); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }
//     token, err := uc.Service.Login(credentials.Username, credentials.Password)
//     if err != nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
//         return
//     }
//     c.JSON(http.StatusOK, gin.H{"token": token})
// }



type UserController struct {
    Service service.UserService
}

func NewUserController(service service.UserService) *UserController {
    return &UserController{Service: service}
}

func (uc *UserController) Register(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := uc.Service.Register(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    token, err := uc.Service.Login(credentials.Username, credentials.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}