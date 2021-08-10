package main

import (
	"net/http"
	"time"

	"gorm.io/driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

//Trainer ...
type Trainer struct {
	ID        uint           `gorm:"primarykey"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

//TrainerRegister used for register as Trainer
type TrainerRegister struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

//BaseResponse used as response format
type BaseResponse struct {
	Code    int
	Status  string
	Message string
	Data    interface{}
}

//DB to access DB
var DB *gorm.DB

func initDB() {
	dsn := "alkanahalida:123456@tcp(127.0.0.1:3306)/pokemontrainer?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database")
	}
}

//Migrate will automigrate database from struct
func Migrate() {
	DB.AutoMigrate(&Trainer{})
}

func main() {
	initDB()
	Migrate()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	e := echo.New()
	e.GET("/trainers", getTrainerController)
	e.POST("/trainers/register", insertTrainerController)
	e.Start(":" + viper.Get("PORT").(string))
}

func getTrainerController(c echo.Context) error {
	var trainers []Trainer
	DB.Find(&trainers)
	return c.JSON(http.StatusOK, trainers)
}

func insertTrainerController(c echo.Context) error {
	var newTrainer TrainerRegister
	c.Bind(&newTrainer)

	var newTrainerDB Trainer
	newTrainerDB.Name = newTrainer.Name
	newTrainerDB.Address = newTrainer.Address

	result := DB.Create(&newTrainerDB)
	if result.Error != nil {
		response := BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "input invalid",
			Status:  "error",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := BaseResponse{
		Code:    http.StatusOK,
		Message: "success register",
		Status:  "success",
	}

	return c.JSON(http.StatusOK, response)

}
