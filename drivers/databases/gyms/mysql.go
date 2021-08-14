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

// UpdateGym Do some DB logic and then return to as Domain data
func (repo *MysqlGymRepository) UpdateGym(ctx context.Context, gymID int, name, address string) (gyms.Domain, error) {
	resultGymData := &Gym{}
	result := repo.Conn.Find(resultGymData).Where("id = ?", gymID).Updates(&Gym{
		Name:    name,
		Address: address,
	})

	if result.Error != nil {
		return gyms.Domain{}, nil
	}
	return ToDomain(resultGymData), nil
}

// GetGyms get Gyms Data
func (repo *MysqlGymRepository) GetGyms(ctx context.Context) ([]gyms.Domain, error) {
	resultGymData := []Gym{}
	result := repo.Conn.Find(&resultGymData)
	if result.Error != nil {
		return []gyms.Domain{}, result.Error
	}
	return ToSliceDomain(resultGymData), nil
}
