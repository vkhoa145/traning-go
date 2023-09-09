package repository

import (
	"errors"
	"fmt"

	"github.com/vkhoa145/go-training/app/models"
)

func (r UserRepo) GetUserById(id float64) (*models.User, error) {
	var user *models.User
	result := r.DB.Table(models.User{}.TableName()).Where("id = ?", id).First(&user)
	fmt.Println("email of user", id)
	fmt.Println("result of get user by email", result)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
