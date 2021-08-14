package trainers

import (
	"pokemontrainer/business/trainers"
	"pokemontrainer/drivers/databases/gyms"
	"pokemontrainer/drivers/databases/pokeballs"
	"pokemontrainer/drivers/databases/pokemons"
	"time"

	"gorm.io/gorm"
)

// Record berisi model dari users

//Trainer database table structure
type Trainer struct {
	ID        uint                 `gorm:"primarykey" json:"id"`
	Name      string               `json:"name"`
	Address   string               `json:"address"`
	Username  string               `json:"username"`
	Password  string               `json:"-"`
	Pokemons  []pokemons.Pokemon   `gorm:"many2many:trainer_pokemons;"`
	Gyms      []gyms.Gym           `gorm:"many2many:trainer_gyms;"`
	Pokeballs []pokeballs.Pokeball `gorm:"many2many:trainer_pokeballs;"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	DeletedAt gorm.DeletedAt       `gorm:"index" json:"deleted_at"`
}

//TrainerPokemon relation table structure
type TrainerPokemon struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	TrainerID int  `gorm:"primaryKey;autoIncrement:false" json:"trainer_id"`
	PokemonID int  `gorm:"primaryKey;autoIncrement:false" json:"pokemon_id"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//TrainerGym relation table structure
type TrainerGym struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	TrainerID int  `gorm:"primaryKey;autoIncrement:false" json:"trainer_id"`
	GymID     int  `gorm:"primaryKey;autoIncrement:false" json:"gym_id"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// TrainerPokeballs relation table structure
type TrainerPokeballs struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	TrainerID  int  `gorm:"primaryKey;autoIncrement:false" json:"trainer_id"`
	PokeballID int  `gorm:"primaryKey;autoIncrement:false" json:"pokeball_id"`
	Quantity   int  `json:"quantity"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (rec *Trainer) toDomain() trainers.Domain {
	return trainers.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Address:   rec.Address,
		Username:  rec.Username,
		Password:  rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func fromDomain(trainer trainers.Domain) *Trainer {
	return &Trainer{
		ID:        trainer.ID,
		Name:      trainer.Name,
		Address:   trainer.Address,
		Username:  trainer.Address,
		Password:  trainer.Password,
		CreatedAt: trainer.CreatedAt,
		UpdatedAt: trainer.UpdatedAt,
		DeletedAt: trainer.DeletedAt,
	}
}

// ToSliceDomain return Trainer Data as Slice
func ToSliceDomain(trainerData []Trainer) []trainers.Domain {
	var convertedTrainerData []trainers.Domain
	for _, v := range trainerData {
		convertedTrainerData = append(convertedTrainerData, v.toDomain())
	}
	return convertedTrainerData
}
