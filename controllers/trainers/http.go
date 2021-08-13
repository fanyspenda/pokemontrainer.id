package trainers

import (
	"fmt"
	"net/http"
	"pokemontrainer/business/trainers"
	"pokemontrainer/controllers"
	"pokemontrainer/controllers/trainers/requests"
	"pokemontrainer/controllers/trainers/responses"
	"strconv"

	"github.com/labstack/echo/v4"
)

// TrainerController ...
type TrainerController struct {
	TrainerUseCase trainers.UseCase
}

// NewTrainerController collection of controller of Trainer
func NewTrainerController(e *echo.Echo, trainerUC trainers.UseCase) {
	controller := &TrainerController{
		TrainerUseCase: trainerUC,
	}

	trainers := e.Group("trainers")
	trainers.POST("/login", controller.Login)
	trainers.POST("/register", controller.Register)
	trainers.GET("/", controller.GetTrainers)
	trainers.GET("", controller.GetTrainers)
	trainers.POST("/catch", controller.CatchPokemon)
	trainers.PUT("/:id", controller.TrainerUpdate)
}

// Login controller for login useCase
func (controller *TrainerController) Login(c echo.Context) error {
	var loginData requests.TrainerLogin
	if err := c.Bind(&loginData); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	trainer, err := controller.TrainerUseCase.Login(c.Request().Context(), loginData.Username, loginData.Password)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, trainer)
}

// Register controller for Register useCase
func (controller *TrainerController) Register(c echo.Context) error {
	var registerData requests.TrainerRegister
	if err := c.Bind(&registerData); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	trainer, err := controller.TrainerUseCase.Register(c.Request().Context(), registerData.Name, registerData.Address, registerData.Username, registerData.Password)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, trainer)
}

// GetTrainers controller for Register useCase
func (controller *TrainerController) GetTrainers(c echo.Context) error {
	trainers, err := controller.TrainerUseCase.GetTrainers(c.Request().Context())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromSliceDomain(trainers))
}

// CatchPokemon ...
func (controller *TrainerController) CatchPokemon(c echo.Context) error {
	var catchData requests.TrainerCatchPokemon
	if err := c.Bind(&catchData); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	fmt.Println("catchData", catchData)
	trainer, err := controller.TrainerUseCase.CatchPokemon(c.Request().Context(), catchData.TrainerID, catchData.PokemonID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, trainer)
}

// TrainerUpdate controller to update trainer data including password if not empty
func (controller *TrainerController) TrainerUpdate(c echo.Context) error {
	updateTrainerData := requests.TrainerUpdate{}
	trainerID, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&updateTrainerData); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := controller.TrainerUseCase.UpdateTrainer(c.Request().Context(),
		trainerID,
		updateTrainerData.Name,
		updateTrainerData.Address,
		updateTrainerData.Username,
		updateTrainerData.Password)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(result))
}