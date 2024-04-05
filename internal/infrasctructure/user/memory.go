package user

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/banggibima/go-fiber-redis/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type UserMemory struct {
	rdb   *redis.Client
	mutex *sync.RWMutex
}

func NewUserMemory(rdb *redis.Client) *UserMemory {
	return &UserMemory{
		rdb:   rdb,
		mutex: &sync.RWMutex{},
	}
}

func (m *UserMemory) GetAll() ([]*entities.User, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	ctx := context.Background()
	keys := m.rdb.Keys(ctx, "user:*").Val()

	var users []*entities.User
	for _, key := range keys {
		userBytes, err := m.rdb.Get(ctx, key).Bytes()
		if err != nil {
			return nil, err
		}

		var user *entities.User
		if err := json.Unmarshal(userBytes, &user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (m *UserMemory) GetByID(id uuid.UUID) (*entities.User, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	ctx := context.Background()
	userBytes, err := m.rdb.Get(ctx, "user:"+id.String()).Bytes()
	if err != nil {
		return nil, err
	}

	var user *entities.User
	if err := json.Unmarshal(userBytes, &user); err != nil {
		return nil, err
	}

	return user, nil
}

func (m *UserMemory) Create(user *entities.User) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	ctx := context.Background()

	user.ID = uuid.New()

	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err := m.rdb.Set(ctx, "user:"+user.ID.String(), userBytes, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (m *UserMemory) Update(id uuid.UUID, user *entities.User) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	ctx := context.Background()

	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err := m.rdb.Set(ctx, "user:"+id.String(), userBytes, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (m *UserMemory) Delete(id uuid.UUID) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	ctx := context.Background()

	if err := m.rdb.Del(ctx, "user:"+id.String()).Err(); err != nil {
		return err
	}

	return nil
}
