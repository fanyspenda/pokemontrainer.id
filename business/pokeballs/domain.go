package pokeballs

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Domain for pokeball useCase
type Domain struct {
	ID          uint
	Name        string
	SuccessRate float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

// UseCases for Pokeball
type UseCases interface {
	AddPokeball(ctx context.Context, name string, successRate float32) (Domain, error)
}

// Repositories for Pokeball
type Repositories interface {
	AddPokeball(ctx context.Context, name string, successRate float32) (Domain, error)
}
