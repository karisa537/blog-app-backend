package main

import (
	"github.com/gin-gonic/gin"
	"github.com/karisa537/blog-app/blog/controller"
	"github.com/karisa537/blog-app/blog/model"
	"github.com/karisa537/blog-app/blog/repository"
	"github.com/karisa537/blog-app/blog/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
)

func main() {
	dsn := "host=localhost port=5432 user=root dbname=blogs_DB password=root sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}


	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to connect to db handle")
	}
	defer sqlDB.Close()

	// Automatically migrate your schema
	db.AutoMigrate(&model.Blog{})

	// Initialize repository, service, and controller
	blogRepo := repository.NewBlogRepository(db)
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)
	
    // Set up Gin router with CORS middleware
	r := gin.Default()
	r.Use(corsMiddleware())

	r.GET("/api/blogs", blogController.GetBlog)
	r.GET("/api/blogs/:id", blogController.GetBlog)
	r.POST("/api/blogs", blogController.CreateBlog)
	r.PUT("/api/blogs/:id", blogController.UpdateBlog)
	r.DELETE("/api/blogs/:id", blogController.DeleteBlog)

	r.Run(":8080")


}

func corsMiddleware() gin.HandlerFunc{
	return func (c *gin.Context)  {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-headers", "Content-Type, Content-Length, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	}
}