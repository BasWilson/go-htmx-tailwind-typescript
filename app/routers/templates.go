package routers

import (
	"github.com/baswilson/go-htmx-tailwind-typescript/app/templ/components"
	"github.com/baswilson/go-htmx-tailwind-typescript/app/templ/pages"
	"github.com/baswilson/go-htmx-tailwind-typescript/app/types"
	"github.com/baswilson/go-htmx-tailwind-typescript/app/utils"
	"github.com/labstack/echo"
)

var Globals = map[string]interface{}{
	"Clicked": 0,
}

func Templates(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error { return utils.Rt(c, pages.Index()) })
	e.GET("/pokemon", func(c echo.Context) error {
		var result types.PokemonList
		err := utils.Get("https://pokeapi.co/api/v2/pokemon", &result)

		if err != nil {
			return utils.Rt(c, pages.Error("Error getting data from the Pokemon API"))
		}

		return utils.Rt(c, pages.Pokemon(result.Results))
	})
	e.GET("/pokemon/:name", func(c echo.Context) error {

		pokemonName := c.Param("name")
		var result types.Pokemon
		err := utils.Get("https://pokeapi.co/api/v2/pokemon/"+pokemonName, &result)

		if err != nil {
			return utils.Rt(c, pages.Error("Error getting data from the Pokemon API"))
		}

		return utils.Rt(c, pages.PokemonSlug(result))
	})
	e.POST("/clicked", func(c echo.Context) error {
		Globals["Clicked"] = Globals["Clicked"].(int) + 1
		return utils.Rt(c, components.Clicked(Globals["Clicked"].(int)))
	})
}
