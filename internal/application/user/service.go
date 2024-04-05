package user

import (
	"github.com/banggibima/go-fiber-redis/internal/domain/entities"
	"github.com/banggibima/go-fiber-redis/internal/domain/memories"
	"github.com/google/uuid"
)

type UserService struct {
	userMemory memories.UserMemory
}

func NewUserService(userMemory memories.UserMemory) *UserService {
	return &UserService{
		userMemory: userMemory,
	}
}

func (uc *UserService) GetAllUsers() ([]*entities.User, error) {
	users, err := uc.userMemory.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UserService) GetUserByID(id uuid.UUID) (*entities.User, error) {
	user, err := uc.userMemory.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserService) CreateUser(user *entities.User) error {
	err := uc.userMemory.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserService) UpdateUser(id uuid.UUID, user *entities.User) error {
	err := uc.userMemory.Update(id, user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserService) DeleteUser(id uuid.UUID) error {
	err := uc.userMemory.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
