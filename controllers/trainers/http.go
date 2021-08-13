package trainers

import (
	"net/http"
	"pokemontrainer/business/trainers"
	"pokemontrainer/controllers"
	"pokemontrainer/controllers/trainers/requests"

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
	trainer, err := controller.TrainerUseCase.GetTrainers(c.Request().Context())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, trainer)
}
