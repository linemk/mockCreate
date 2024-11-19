package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"mockProject/mocks"
	"mockProject/workOnMock"
	"testing"
)

func TestRealUserRepository_GetUserByID(t *testing.T) {

	t.Run("testGetWithId", func(t *testing.T) {
		mocker := new(mocks.UserRepository)

		mocker.On("Get", uint64(123)).Return(workOnMock.User{ID: 123, Name: "Максим"}, nil)

		user, err := mocker.Get(123)
		assert.NoError(t, err)
		assert.Equal(t, "Максим", user.Name)
		mocker.AssertCalled(t, "Get", uint64(123))
	})
	t.Run("testGetWithError", func(t *testing.T) {
		mocker := new(mocks.UserRepository)

		mocker.On("Get", uint64(456)).Return(workOnMock.User{}, errors.New("test error"))

		_, err := mocker.Get(456)
		assert.Error(t, err)
		assert.Equal(t, "test error", err.Error())
	})

	t.Run("testAddWithoutError", func(t *testing.T) {
		mocker := new(mocks.UserRepository)
		mocker.On("Add", workOnMock.User{ID: 123, Name: "Святозар"}).Return(workOnMock.User{ID: 123, Name: "Святозар"}, nil)

		user := workOnMock.User{ID: 123, Name: "Святозар"}

		user, err := mocker.Add(user)

		assert.NoError(t, err)
		assert.Equal(t, "Святозар", user.Name)
		assert.Equal(t, uint64(123), user.ID)
	})

	t.Run("testAddWithError", func(t *testing.T) {
		mocker := new(mocks.UserRepository)

		mocker.On("Add", workOnMock.User{}).Return(workOnMock.User{}, errors.New("user not added"))

		_, err := mocker.Add(workOnMock.User{})
		assert.Error(t, err)
		assert.Equal(t, "user not added", err.Error())
	})

}
