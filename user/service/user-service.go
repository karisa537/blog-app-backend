package service

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/karisa537/blog-app/user/model"
	"github.com/karisa537/blog-app/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *model.User) error
	Login(username, password string) (string, error)
	GetByUsername(username string) (*model.User, error)
}

type userService struct {
	repo      repository.UserRepository
	jwtSecret string
}

// GetByUsername implements UserService.
func (s *userService) GetByUsername(username string) (*model.User, error) {
	return s.repo.GetByUsername(username)
}

// Login implements UserService.
func (s *userService) Login(username string, password string) (string, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil || user.Password != password {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}

// Register implements UserService.
func (s *userService) Register(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.Create(user)
}

func NewUserService(repo repository.UserRepository, jwtSecret string) UserService {
	return &userService{repo: repo, jwtSecret: jwtSecret}
}
