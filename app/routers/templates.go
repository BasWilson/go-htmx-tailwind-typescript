package routers

import (
	"github.com/a-h/templ"
	"github.com/baswilson/go-htmx-tailwind-typescript/app/framework"
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

	framework.RegisterStaticComponent(e, "/pokemon", func(c echo.Context) templ.Component {
		var result types.PokemonList
		err := utils.Get("https://pokeapi.co/api/v2/pokemon", &result)

		if err != nil {
			return pages.Error("Error getting data from the Pokemon API")
		}

		return pages.Pokemon(result.Results)
	}, -1)

	framework.RegisterStaticComponent(e, "/pokemon/:name", func(c echo.Context) templ.Component {
		var result types.Pokemon
		err := utils.Get("https://pokeapi.co/api/v2/pokemon/"+c.Param("name"), &result)

		if err != nil {
			return pages.Error("Error getting data from the Pokemon API")
		}

		return pages.PokemonSlug(result)
	}, -1)

	e.POST("/clicked", func(c echo.Context) error {
		Globals["Clicked"] = Globals["Clicked"].(int) + 1
		return utils.Rt(c, components.Clicked(Globals["Clicked"].(int)))
	})
}
