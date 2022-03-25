package usecase_test

import (
	"errors"
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/internal/user/mocks"
	"lottery-web-scrapping/internal/user/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {

	mockRequest := &models.User{
		Username:  "tester@gmail.com",
		Password:  "#SAK$)LSDKA:SKD@#)!@#!23asdasd'as;d",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo := new(mocks.IUserRepository)

		mockUserRepo.On("CreateOne", mockRequest).Return(nil).Once()

		mockUserUsecase := usecase.NewUserService(mockUserRepo)

		err := mockUserUsecase.CreateUser(mockRequest)

		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepo := new(mocks.IUserRepository)

		mockUserRepo.On("CreateOne", mockRequest).Return(errors.New("unexpect error occur")).Once()

		mockUserUsecase := usecase.NewUserService(mockUserRepo)

		err := mockUserUsecase.CreateUser(mockRequest)

		assert.Error(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestFindName(t *testing.T) {

	mockReponse := &models.User{
		ID:        primitive.NewObjectID(),
		Username:  "jojodevilman@gmail.com",
		Password:  "#SAK$)LSDKA:SKD@#)!@#!23asdasd'as;d",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo := new(mocks.IUserRepository)

		mockUserRepo.On("FindUser", mockReponse.Username).Return(mockReponse, nil)
		mockUserRepo.On("FindByUser", mockReponse.Username).Return(mockReponse, nil)

		mockUserUsecase := usecase.NewUserService(mockUserRepo)

		user, err := mockUserUsecase.FindByUser(mockReponse.Username)

		assert.Nil(t, err)

		assert.NotNil(t, user)

		assert.Equal(t, mockReponse, user)

		mockUserRepo.AssertExpectations(t)

	})

	t.Run("error", func(t *testing.T) {
		mockUserRepo := new(mocks.IUserRepository)

		mockUserRepo.On("FindUser", mock.AnythingOfType("string")).Return(nil, errors.New("Unexpected Error")).Once()
		mockUserRepo.On("FindByUser", mockReponse.Username).Return(nil, errors.New("Unexpected Error")).Once()

		mockUserUsecase := usecase.NewUserService(mockUserRepo)

		_, err := mockUserUsecase.FindByUser(mockReponse.Username)

		assert.Error(t, err)

		mockUserRepo.AssertExpectations(t)

	})
}

func TestFindId(t *testing.T) {
	mockReponse := &models.User{
		ID:        primitive.NewObjectID(),
		Username:  "jojodevilman@gmail.com",
		Password:  "#SAK$)LSDKA:SKD@#)!@#!23asdasd'as;d",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo := new(mocks.IUserRepository)
		mockUserRepo.On("FindId", mock.AnythingOfType("string")).Return(mockReponse, nil).Once()

		mockUserRepo.On("FindById", mock.AnythingOfType("string")).Return(mockReponse, nil).Once()
		mockUserUsecase := usecase.NewUserService(mockUserRepo)

		user, err := mockUserUsecase.FindById(mockReponse.ID.String())

		assert.Error(t, err)

		assert.Equal(t, mockReponse, user)

		mockUserRepo.AssertExpectations(t)

	})
}
