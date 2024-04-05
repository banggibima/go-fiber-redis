package services

import (
	"github.com/banggibima/go-fiber-redis/internal/domain/entities"
	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]*entities.User, error)
	GetUserByID(id uuid.UUID) (*entities.User, error)
	CreateUser(user *entities.User) error
	UpdateUser(id uuid.UUID, user *entities.User) error
	DeleteUser(id uuid.UUID) error
}
