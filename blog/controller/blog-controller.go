package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/karisa537/blog-app/blog/model"
	"github.com/karisa537/blog-app/blog/service"
	"net/http"
	"strconv"
)

type BlogController struct {
	Service service.BlogService
}

func NewBlogController(service service.BlogService) *BlogController {
	return &BlogController{Service: service}
}

func (bc *BlogController) GetBlogs(c *gin.Context){
	blogs := bc.Service.GetBlogs()
	c.JSON(http.StatusOK, blogs)
}

func (bc *BlogController) GetBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	blog, err := bc.Service.GetBlog(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}

func(bc *BlogController) CreateBlog(c *gin.Context) {
	var blog model.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	bc.Service.CreateBlog(&blog)
	c.JSON(http.StatusCreated, blog)
}

func (bc *BlogController) UpdateBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var blog model.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := bc.Service.UpdateBlog(uint(id), &blog)
	if err != nil {
		c.JSON(http.StatusFound, gin.H{"error" : "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}

func (bc *BlogController) DeleteBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := bc.Service.DeleteBlog(uint(id))
	
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not forund"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted"})
}