package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/karisa537/blog-app/user/controller"
	"github.com/karisa537/blog-app/user/model"
	"github.com/karisa537/blog-app/user/repository"
	"github.com/karisa537/blog-app/user/service"
)

func main() {
	db, err := gorm.Open("host=localhost port=5432 user=postgres dbname=user_DB password=yourpassword sslmode=disable")
	if err != nil {
		panic("Failed to connect to db")
	}

	defer db.Close()

	db.AutoMigrate(&model.User{})
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, "your_jwt_secret_key")
	userController := controller.NewUserController


	r := gin.Default()
	r.Use(CorsMiddleware())

	r.POST("/api/register", userController.Register)
	r.POST("/api/login", userController.Login)

	r.Run(":8081")

}

func CorsMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS"{
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}