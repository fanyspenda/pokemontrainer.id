package trainers

import (
	"context"
	"time"

	"gorm.io/gorm"
)

//Domain trainer untuk mapping data
type Domain struct {
	ID        uint
	Name      string
	Address   string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// USE CASE
// merupakan interface terkait dengan fungsi-fungsi apasaja yang dapat dilakukan oleh Domain

// UseCase about what Trainer can do
type UseCase interface {
	Register(ctx context.Context, name, address, username, password string) (Domain, error)
	Login(ctx context.Context, username, password string) (Domain, error)
	GetTrainers(ctx context.Context) ([]Domain, error)
	CatchPokemon(ctx context.Context, ID, pokemonID int) (Domain, error)
	UpdateTrainer(ctx context.Context, trainerID int, name, address, username, password string) (Domain, error)
	// AddGym(ctx context.Context, ID, gymID int) (Domain, error)
}

// Repository is about what the UseCase to do with the Database
// Contoh: UseCase BuatNasiGoreng, reponya bisa siapin bumbu, siapin sayur, siapin nasi, campur, goreng, dll.
type Repository interface {
	Register(ctx context.Context, name, address, username, password string) (Domain, error)
	Login(ctx context.Context, username, password string) (Domain, error)
	GetTrainers(ctx context.Context) ([]Domain, error)
	CatchPokemon(ctx context.Context, ID, pokemonID int) (Domain, error)
	UpdateTrainer(ctx context.Context, trainerID int, name, address, username, password string) (Domain, error)
	// AddGym(ctx context.Context, ID, gymID int) (Domain, error)
}
