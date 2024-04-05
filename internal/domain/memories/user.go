package memories

import (
	"github.com/banggibima/go-fiber-redis/internal/domain/entities"
	"github.com/google/uuid"
)

type UserMemory interface {
	GetAll() ([]*entities.User, error)
	GetByID(id uuid.UUID) (*entities.User, error)
	Create(user *entities.User) error
	Update(id uuid.UUID, user *entities.User) error
	Delete(id uuid.UUID) error
}
