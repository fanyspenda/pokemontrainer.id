package trainers

import (
	"fmt"
	"net/http"
	"pokemontrainer/business/pokeballs"
	"pokemontrainer/business/trainers"
	"pokemontrainer/controllers"
	pokeballResponses "pokemontrainer/controllers/pokeballs/responses"
	"pokemontrainer/controllers/trainers/requests"
	"pokemontrainer/controllers/trainers/responses"
	"pokemontrainer/helpers/middlewares"
	"strconv"

	"github.com/labstack/echo/v4/middleware"

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
	trainers.GET("/", controller.GetTrainers, middleware.JWT(middlewares.KeyToByte()))
	trainers.GET("", controller.GetTrainers)
	trainers.POST("/catch", controller.CatchPokemon)
	trainers.PUT("/:id", controller.TrainerUpdate)
	trainers.POST("/gyms/register", controller.TrainerRegisterGym)
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

	trainer, err := controller.TrainerUseCase.Register(c.Request().Context(),
		registerData.Name,
		registerData.Address,
		registerData.Username,
		registerData.Password,
	)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	var pokeball pokeballs.Domain
	pokeball, err = controller.TrainerUseCase.GetFirstBall(c.Request().Context(), trainer.ID)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToRegisterResponse(trainer, pokeballResponses.FromDomain(pokeball)))
}

// GetTrainers controller for GetTrainers useCase
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

	token := c.Request().Header.Get("Authorization")
	fmt.Println("catchData", catchData)
	fmt.Println("token", token)
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

// TrainerRegisterGym add trainer to gym which is add trainerId and GymId to the join table
func (controller *TrainerController) TrainerRegisterGym(c echo.Context) error {
	var gymRegisterData requests.TrainerRegisterGym
	if err := c.Bind(&gymRegisterData); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	domainRes, err := controller.TrainerUseCase.AddGym(c.Request().Context(), gymRegisterData.TrainerID, gymRegisterData.GymID)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(domainRes))
}
