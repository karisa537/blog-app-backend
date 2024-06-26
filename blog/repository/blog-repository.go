package repository

import (
	"github.com/karisa537/blog-app/blog/model"
	// "golang.org/x/tools/blog"
	"gorm.io/gorm"
)

type BlogRespository interface {
	GetAll() []model.Blog
	GetByID(id uint) (*model.Blog, error)
	Create(blog *model.Blog)
	Update(blog *model.Blog) error
	Delete(blog *model.Blog) error
}

type blogRepository struct {
	db *gorm.DB
}

// Create implements BlogRespository.
func (r *blogRepository) Create(blog *model.Blog) {
	r.db.Create(blog)
}

// Delete implements BlogRespository.
func (r *blogRepository) Delete(blog *model.Blog) error{
	if err := r.db.Delete(blog).Error; err != nil {
		return err
	}
	return nil
}

// GetAll implements BlogRespository.
func (r *blogRepository) GetAll() []model.Blog {
	var blogs []model.Blog
	r.db.Find(&blogs)
	return blogs
}

// GetByID implements BlogRespository.
func (r *blogRepository) GetByID(id uint) (*model.Blog, error) {
	var blog model.Blog
	if err := r.db.Where("id = ?", id).First(&blog).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

// Update implements BlogRespository.
func (r *blogRepository) Update(blog *model.Blog) error{
	if err := r.db.Save(blog).Error; err != nil{
		return err
	}

	return nil
}

func NewBlogRepository(db *gorm.DB) BlogRespository {
	return &blogRepository{db: db}
}
