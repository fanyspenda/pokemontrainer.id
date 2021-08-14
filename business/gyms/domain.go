package gyms

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Domain for a Gym
type Domain struct {
	ID        uint
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// UseCases for Gym about what Gym can do
type UseCases interface {
	AddGym(ctx context.Context, name, address string) (Domain, error)
	UpdateGym(ctx context.Context, gymID int, name, address string) (Domain, error)
}

// Repositories for Gym
type Repositories interface {
	AddGym(ctx context.Context, name, address string) (Domain, error)
	UpdateGym(ctx context.Context, gymID int, name, address string) (Domain, error)
}
