package repository

import (
	"errors"
	"fmt"

	"github.com/vkhoa145/go-training/app/models"
)

func (r UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	result := r.DB.Table(models.User{}.TableName()).Where("email = ?", email).First(&user)
	fmt.Println("email of user", email)
	fmt.Println("result of get user by email", result)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
