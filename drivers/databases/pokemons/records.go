package pokemons

import (
	"time"

	"gorm.io/gorm"
)

// Record berisi model dari users

//Pokemon ...
type Pokemon struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
