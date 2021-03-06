package databases

import (
	"fmt"
	"pokemontrainer/drivers/databases/gyms"
	"pokemontrainer/drivers/databases/pokeballs"
	"pokemontrainer/drivers/databases/pokemons"
	"pokemontrainer/drivers/databases/trainers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConfigParams define params of config
type ConfigParams struct {
	Charset string
}

// Config configure the Database
type Config struct {
	User      string
	Passwd    string
	Addr      string
	Port      string
	DBName    string
	Params    ConfigParams
	Loc       string
	ParseTime string
}

// InitDB initiate DB
func (db *Config) InitDB() *gorm.DB {

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		db.User, db.Passwd, db.Addr, db.Port, db.DBName, db.Params.Charset, db.ParseTime, db.Loc)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database")
	}
	Migrate(DB)

	return DB
}

//Migrate will automigrate database from struct
func Migrate(db *gorm.DB) {
	err := db.SetupJoinTable(&trainers.Trainer{}, "Pokemons", &trainers.TrainerPokemon{})
	if err != nil {

		panic(err)
	}
	err = db.SetupJoinTable(&trainers.Trainer{}, "Gyms", &trainers.TrainerGym{})
	if err != nil {

		panic(err)
	}

	err = db.SetupJoinTable(&trainers.Trainer{}, "Pokeballs", &trainers.TrainerPokeballs{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&trainers.Trainer{},
		&pokemons.Pokemon{},
		&gyms.Gym{},
		&pokeballs.Pokeball{},
	)
}
