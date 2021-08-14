package gyms

import (
	"pokemontrainer/business/gyms"
	"time"

	"gorm.io/gorm"
)

// Gym database table structure
type Gym struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// ToDomain convert database structure to domain structure
func ToDomain(gymData *Gym) gyms.Domain {
	return gyms.Domain{
		ID:        gymData.ID,
		Name:      gymData.Name,
		Address:   gymData.Address,
		CreatedAt: gymData.CreatedAt,
		UpdatedAt: gymData.UpdatedAt,
		DeletedAt: gymData.DeletedAt,
	}
}

// FromDomain convert domain structure to database structure
func FromDomain(gymDomain *gyms.Domain) Gym {
	return Gym{
		ID:        gymDomain.ID,
		Name:      gymDomain.Name,
		Address:   gymDomain.Address,
		CreatedAt: gymDomain.CreatedAt,
		UpdatedAt: gymDomain.UpdatedAt,
		DeletedAt: gymDomain.DeletedAt,
	}
}

// ToSliceDomain convert slice Gym to Slice Domain Gym
func ToSliceDomain(sliceGym []Gym) []gyms.Domain {
	var convertedGyms []gyms.Domain
	for _, v := range sliceGym {
		convertedGyms = append(convertedGyms, ToDomain(&v))
	}
	return convertedGyms
}
