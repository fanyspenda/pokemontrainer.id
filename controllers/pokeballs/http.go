package pokeballs

import (
	"net/http"
	"pokemontrainer/business/pokeballs"
	"pokemontrainer/controllers"
	"pokemontrainer/controllers/pokeballs/requests"
	"pokemontrainer/controllers/pokeballs/responses"
	"pokemontrainer/helpers/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ControllerPokeball make NewController
type ControllerPokeball struct {
	pokeballUseCase pokeballs.UseCases
}

// NewControllerPokeball collection of controller with usecase
func NewControllerPokeball(e *echo.Echo, pokeballUC pokeballs.UseCases) {
	controller := &ControllerPokeball{
		pokeballUseCase: pokeballUC,
	}

	pokeballs := e.Group("/pokeballs")
	pokeballs.POST("", controller.AddPokeball, middleware.JWT(middlewares.KeyToByte()))
}

// AddPokeball add pokeball controller
func (controller *ControllerPokeball) AddPokeball(c echo.Context) error {
	reqData := requests.AddPokeball{}
	if err := c.Bind(&reqData); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	resp, err := controller.pokeballUseCase.AddPokeball(c.Request().Context(), reqData.Name, reqData.SuccessRate)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(resp))
}
