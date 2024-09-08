package usecase

import (
	"project/internal/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByID(id int64) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*model.User, error) {
	args := m.Called(username)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := &UserUseCase{UserRepo: mockRepo}

	user := &model.User{Username: "testuser"}

	mockRepo.On("CreateUser", user).Return(nil)

	err := userUseCase.CreateUser(user)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
