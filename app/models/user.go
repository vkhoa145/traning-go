package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(255)" json:"first_name"`
	LastName  string `gorm:"type:varchar(255)" json:"last_name"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
	Password  string `gorm:"type:varchar(255)" json:"password"`
	Phone     string `gorm:"type:varchar(15)" json:"phone"`
}

func (User) TableName() string {
	return "users"
}

type SignUpInput struct {
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Phone           string `json:"phone" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	FirstName string    `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName  string    `json:"last_name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone     string    `json:"phone" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterUserRecord(user *User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
