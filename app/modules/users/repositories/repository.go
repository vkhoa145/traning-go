package repository

import (
	"github.com/vkhoa145/go-training/app/models"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CheckEmailExisting(email string) bool
	CreateUser(data *models.SignUpInput) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id float64) (*models.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}
