package gyms

import (
	"context"
	"pokemontrainer/business/gyms"

	"gorm.io/gorm"
)

// MysqlGymRepository struct for NewMysqlGymRepository
type MysqlGymRepository struct {
	Conn *gorm.DB
}

// NewMysqlGymRepository ...
func NewMysqlGymRepository(newConn *gorm.DB) gyms.Repositories {
	return &MysqlGymRepository{
		Conn: newConn,
	}
}

// AddGym to the database
func (repo *MysqlGymRepository) AddGym(ctx context.Context, name, address string) (gyms.Domain, error) {
	addGymData := &Gym{
		Name:    name,
		Address: address,
	}
	result := repo.Conn.Create(&addGymData)
	if result.Error != nil {
		return gyms.Domain{}, result.Error
	}
	return ToDomain(addGymData), nil
}
