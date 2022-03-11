package mapper

import (
	"api-hackaton-devs/entity"
	"api-hackaton-devs/repository"
)

func MapModelResponseToEntityResponse(hackatons *[]repository.Hackaton) *entity.Response {
	entityHackatons := make([]entity.HackatonResponse, len(*hackatons))
	for i, h := range *hackatons {
		entityHackatons[i] = hackatonToEntity(h)
	}
	return &entity.Response{
		Hackatons: &entityHackatons,
	}
}

func hackatonToEntity(h repository.Hackaton) entity.HackatonResponse {
	devs := make([]entity.DeveloperResponse, len(h.Devs))
	for i, d := range h.Devs {
		devs[i] = developerToEntity(d)
	}
	return entity.HackatonResponse{
		ID:   h.ID,
		Name: h.Name,
		Devs: devs,
	}
}

func developerToEntity(d repository.Developer) entity.DeveloperResponse {
	return entity.DeveloperResponse{
		ID:       d.ID,
		Position: d.Position,
		Name:     d.Name,
		LastName: d.LastName,
	}
}
