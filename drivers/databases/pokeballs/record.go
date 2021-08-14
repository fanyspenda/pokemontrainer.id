package pokeballs

import (
	"pokemontrainer/business/pokeballs"
	"time"

	"gorm.io/gorm"
)

// Pokeball table structure
type Pokeball struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	SuccessRate float32        `json:"success_rate"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

// ToDomain convert Pokeball to its Domain
func ToDomain(pokeball *Pokeball) pokeballs.Domain {
	return pokeballs.Domain{
		ID:          pokeball.ID,
		Name:        pokeball.Name,
		SuccessRate: pokeball.SuccessRate,
		CreatedAt:   pokeball.CreatedAt,
		UpdatedAt:   pokeball.UpdatedAt,
		DeletedAt:   pokeball.DeletedAt,
	}
}
