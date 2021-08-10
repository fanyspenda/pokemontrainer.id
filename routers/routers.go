package routers

import (
	"pokemontrainer/controllers"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// InitRoute initiate Route
func InitRoute() {
	e := echo.New()
	e.GET("/trainers", controllers.GetTrainerController)
	e.POST("/trainers/register", controllers.InsertTrainerController)
	e.Start(":" + viper.Get("PORT").(string))
}
