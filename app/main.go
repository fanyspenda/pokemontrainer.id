package main

import (
	_dbHelper "pokemontrainer/helpers/databases"
	"time"

	_trainerUseCase "pokemontrainer/business/trainers"
	_trainerController "pokemontrainer/controllers/trainers"
	_trainerRepo "pokemontrainer/drivers/databases/trainers"

	_gymUseCase "pokemontrainer/business/gyms"
	_gymController "pokemontrainer/controllers/gyms"
	_gymRepo "pokemontrainer/drivers/databases/gyms"

	_pokeballUseCase "pokemontrainer/business/pokeballs"
	_pokeballController "pokemontrainer/controllers/pokeballs"
	_pokeballRepo "pokemontrainer/drivers/databases/pokeballs"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func importEnvironment() {
	viper.SetConfigFile("app/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}

func main() {
	importEnvironment()

	dbParams := _dbHelper.ConfigParams{
		Charset: viper.GetString("database.params.charset"),
	}
	configDB := _dbHelper.Config{
		User:      viper.GetString("database.user"),
		Passwd:    viper.GetString("database.pass"),
		Addr:      viper.GetString("database.host"),
		Port:      viper.GetString("database.port"),
		DBName:    viper.GetString("database.name"),
		Params:    dbParams,
		Loc:       viper.GetString("database.loc"),
		ParseTime: viper.GetString("database.parseTime"),
	}
	configMongoDB := _dbHelper.MongoDBConfig{}

	db := configDB.InitDB()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()
	_dbHelper.InitMongoDB()

	trainerRepo := _trainerRepo.NewMysqlTrainerRepository(db)
	trainerUseCase := _trainerUseCase.NewTrainerUseCase(trainerRepo, timeoutContext)
	_trainerController.NewTrainerController(e, trainerUseCase)

	gymRepo := _gymRepo.NewMysqlGymRepository(db)
	gymUseCase := _gymUseCase.NewGymUseCases(gymRepo, timeoutContext)
	_gymController.NewGymController(e, gymUseCase)

	pokeballRepo := _pokeballRepo.NewMysqlPokeballRepository(db)
	pokeballUseCase := _pokeballUseCase.NewPokeballUseCase(pokeballRepo, timeoutContext)
	_pokeballController.NewControllerPokeball(e, pokeballUseCase)

	err := e.Start(viper.GetString("server.address"))
	if err != nil {
		panic(err)
	}
}
