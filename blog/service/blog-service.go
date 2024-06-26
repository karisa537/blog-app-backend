package service

import (
	"github.com/karisa537/blog-app/blog/model"
	"github.com/karisa537/blog-app/blog/repository"
	// "golang.org/x/tools/blog"
)

type BlogService interface {
	GetBlogs() []model.Blog
	GetBlog(id uint) (*model.Blog, error)
	CreateBlog(blog *model.Blog)
	UpdateBlog(id uint, blog *model.Blog) error
	DeleteBlog(id uint) error
}

type blogService struct {
	repo repository.BlogRespository
}

// CreateBlog implements BlogService.
func (s *blogService) CreateBlog(blog *model.Blog) {
	s.repo.Create(blog)
}

// DeleteBlog implements BlogService.
func (s *blogService) DeleteBlog(id uint) error {
	blog, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(blog)
}

// GetBlog implements BlogService.
func (s *blogService) GetBlog(id uint) (*model.Blog, error) {
	return s.repo.GetByID(id)
}

// GetBlogs implements BlogService.
func (s *blogService) GetBlogs() []model.Blog {
	return s.repo.GetAll()
}

// UpdateBlog implements BlogService.
func (s *blogService) UpdateBlog(id uint, blog *model.Blog) error {
	existingBlog, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	existingBlog.Title = blog.Title
	existingBlog.Summary = blog.Summary
	existingBlog.Content = blog.Author
	existingBlog.Author = blog.Author
	return s.repo.Update(existingBlog)
}

func NewBlogService(repo repository.BlogRespository) BlogService {
	return &blogService{repo: repo}
}
