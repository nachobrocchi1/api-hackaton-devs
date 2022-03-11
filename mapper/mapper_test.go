package mapper

import (
	"api-hackaton-devs/entity"
	"api-hackaton-devs/repository"
	"testing"

	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

func TestMapper(t *testing.T) {
	t.Run("Maps successfully a model into an entity", func(t *testing.T) {
		response := MapModelResponseToEntityResponse(&aModel)
		assert.NotNil(t, response)
		hackatons := *response.Hackatons
		assert.Equal(t, hackatons[0].ID, anEntity[0].ID)
		assert.Equal(t, hackatons[0].Name, anEntity[0].Name)
		assert.Equal(t, hackatons[0].Devs[0].ID, anEntity[0].Devs[0].ID)
		assert.Equal(t, hackatons[0].Devs[0].Name, anEntity[0].Devs[0].Name)
		assert.Equal(t, hackatons[0].Devs[0].LastName, anEntity[0].Devs[0].LastName)
		assert.Equal(t, hackatons[0].Devs[0].Position, anEntity[0].Devs[0].Position)
	})
}

var (
	aModel = []repository.Hackaton{
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
