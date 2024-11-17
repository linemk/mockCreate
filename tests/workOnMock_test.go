package tests

import (
	"github.com/stretchr/testify/assert"
	"mockProject/mocks"
	"mockProject/workOnMock"
	"testing"
)

func TestRealUserRepository_GetUserByID(t *testing.T) {
	mocker := new(mocks.UserRepository)

	mocker.On("Get", uint64(123)).Return(workOnMock.User{ID: 123, Name: "Максим"}, nil)

	user, err := mocker.Get(123)
	assert.NoError(t, err)
	assert.Equal(t, "Максим", user.Name)
	mocker.AssertCalled(t, "Get", uint64(123))
}
