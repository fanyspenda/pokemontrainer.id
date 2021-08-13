package gyms

import (
	"net/http"
	"pokemontrainer/business/gyms"
	"pokemontrainer/controllers"
	"pokemontrainer/controllers/gyms/requests"
	"pokemontrainer/controllers/gyms/responses"

	"github.com/labstack/echo/v4"
)

// GymController ...
type GymController struct {
	GymUseCases gyms.UseCases
}

// NewGymController collection of controller of Gym
func NewGymController(e *echo.Echo, gymUC gyms.UseCases) {
	controller := &GymController{
		GymUseCases: gymUC,
	}

	gyms := e.Group("gyms")
	gyms.POST("", controller.AddGym)
}

// AddGym controller for AddGym useCase
func (controller *GymController) AddGym(c echo.Context) error {
	var addGymData requests.GymAdd
	if err := c.Bind(&addGymData); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	gym, err := controller.GymUseCases.AddGym(c.Request().Context(), addGymData.Name, addGymData.Address)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(gym))
}
