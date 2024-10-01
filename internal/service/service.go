package service

import (
	"gox/internal/models"
	"gox/internal/repository"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user models.User) error
	GetUserById(id uuid.UUID) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}


func NewUserService(firstName string, lastName string, email string) (models.User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil

}

func (s *userService) CreateUser(user models.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *userService) GetUserById(id uuid.UUID) (models.User, error) {
	return s.userRepository.GetUserById(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAllUsers()
}
