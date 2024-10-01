package repository

import (
	"gox/internal/models"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserById(id uuid.UUID) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userRepository struct {
	users []models.User
}

var users = []models.User{
	{
		ID:        uuid.MustParse("d9b5a4b1-d1d1-4d92-a14b-441a5e5a5ae5"),
		FirstName: "Olivia",
		LastName:  "Ponton",
		Email:     "olivia.ponton@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        uuid.MustParse("d9b5a4b1-d1d1-4d92-a14b-441a5e5a5ae6"),
		FirstName: "Faith",
		LastName:  "Ordway",
		Email:     "faith.ordway@godess.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func CreateUser(user models.User) error {
	users = append(users, user)
	return nil
}

func GetAllUsers() []models.User {
	return users
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return &models.User{}, nil
}
