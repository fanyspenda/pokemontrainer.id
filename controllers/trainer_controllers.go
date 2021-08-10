package controllers

import (
	"net/http"

	"pokemontrainer/configs"
	"pokemontrainer/models/base"
	"pokemontrainer/models/trainers"

	"github.com/labstack/echo/v4"
)

// GetTrainerController getting all trainers
func GetTrainerController(c echo.Context) error {
	var trainers []trainers.Trainer
	configs.DB.Find(&trainers)
	return c.JSON(http.StatusOK, trainers)
}

// InsertTrainerController register as new Trainer
func InsertTrainerController(c echo.Context) error {
	var newTrainer trainers.TrainerRegister
	c.Bind(&newTrainer)

	var newTrainerDB trainers.Trainer
	newTrainerDB.Name = newTrainer.Name
	newTrainerDB.Address = newTrainer.Address

	result := configs.DB.Create(&newTrainerDB)
	if result.Error != nil {
		response := base.Response{
			Code:    http.StatusInternalServerError,
			Message: "input invalid",
			Status:  "error",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := base.Response{
		Code:    http.StatusOK,
		Message: "success register",
		Status:  "success",
	}

	return c.JSON(http.StatusOK, response)

}
