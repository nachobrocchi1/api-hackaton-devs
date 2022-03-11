package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type RepositoryError struct {
	ErrorMsg string
}

func (re *RepositoryError) Error() string {
	return fmt.Sprintf("Repository Error: %s", re.ErrorMsg)
}

type Repository interface {
	InsertHackaton(h *Hackaton) error
	GetHackatonWithBestDevs() (*[]Hackaton, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) InsertHackaton(h *Hackaton) error {
	if err := r.db.Create(h).Error; err != nil {
		return &RepositoryError{
			ErrorMsg: "Error creating Hackaton",
		}
	}
	return nil
}

func (r *repository) GetHackatonWithBestDevs() (*[]Hackaton, error) {
	hackatons := []Hackaton{}
	if err := r.db.Preload("Devs", "position IN ?", []int{1, 2, 3}).Find(&hackatons).Error; err != nil {
		return nil, &RepositoryError{
			ErrorMsg: "Error trying to retrieve hackatons",
		}
	}
	return &hackatons, nil
}
