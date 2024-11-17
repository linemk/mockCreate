package workOnMock

import "errors"

type User struct {
	ID   uint64
	Name string
}

type UserRepository interface {
	Add(user User) (User, error)
	Get(id uint64) (User, error)
}

type RealUserRepository struct {
	users map[uint64]User
}

func (r *RealUserRepository) GetUserByID(id uint64) (User, error) {
	user, exists := r.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *RealUserRepository) AddUser(user User) error {
	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = user
	return nil
}
