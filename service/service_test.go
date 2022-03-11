package service_test

import (
	"api-hackaton-devs/entity"
	"api-hackaton-devs/repository"
	mockRepo "api-hackaton-devs/repository/mocks"
	"api-hackaton-devs/service"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

var (
	ctx = context.TODO()
)

func TestService(t *testing.T) {
	t.Run("Execute service successfully", func(t *testing.T) {
		repo := new(mockRepo.Repository)
		repo.On("GetHackatonWithBestDevs").Return(aModel, nil)
		svc := service.NewService(repo)
		response, err := svc.GetHackatonWithBestDevs(ctx, &entity.Request{})
		assert.Nil(t, err)
		assert.NotNil(t, response)
		hackatons := *response.Hackatons
		assert.Equal(t, hackatons[0].ID, anEntity[0].ID)
		assert.Equal(t, hackatons[0].Name, anEntity[0].Name)
		assert.Equal(t, hackatons[0].Devs[0].ID, anEntity[0].Devs[0].ID)
		assert.Equal(t, hackatons[0].Devs[0].Name, anEntity[0].Devs[0].Name)
		assert.Equal(t, hackatons[0].Devs[0].LastName, anEntity[0].Devs[0].LastName)
		assert.Equal(t, hackatons[0].Devs[0].Position, anEntity[0].Devs[0].Position)
	})

	t.Run("Execute service and returns error", func(t *testing.T) {
		repo := new(mockRepo.Repository)
		repo.On("GetHackatonWithBestDevs").Return(nil, &repository.RepositoryError{ErrorMsg: "Error trying to retrieve hackatons"})
		svc := service.NewService(repo)
		response, err := svc.GetHackatonWithBestDevs(ctx, &entity.Request{})
		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.EqualValues(t, "Repository Error: Error trying to retrieve hackatons", err.Error())
	})
}

var (
	aModel = &[]repository.Hackaton{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "Hackaton 1",
			Devs: []repository.Developer{
				{
					Model: gorm.Model{
						ID: 1,
					},
					HackatonID: 1,
					Position:   1,
					Name:       "name1",
					LastName:   "lastname1",
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					HackatonID: 1,
					Position:   2,
					Name:       "name2",
					LastName:   "lastname2",
				},
			},
		},
	}

	anEntity = []entity.HackatonResponse{
		{
			ID:   1,
			Name: "Hackaton 1",
			Devs: []entity.DeveloperResponse{
				{
					ID:       1,
					Position: 1,
					Name:     "name1",
					LastName: "lastname1",
				},
				{
					ID:       1,
					Position: 2,
					Name:     "name2",
					LastName: "lastname2",
				},
			},
		},
	}
)
