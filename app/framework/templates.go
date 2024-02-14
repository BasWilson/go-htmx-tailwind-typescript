package framework

import (
	"go-htmx-tailwind-typescript/app/templ/components"
	"go-htmx-tailwind-typescript/app/templ/pages"
	"go-htmx-tailwind-typescript/app/types"
	"go-htmx-tailwind-typescript/app/utils"

	"github.com/a-h/templ"
	"github.com/labstack/echo"
)

var Globals = map[string]interface{}{
	"Clicked": 0,
}

func Templates(e *echo.Echo) {
	RegisterStaticComponent(e, "/", func(c echo.Context) templ.Component {
		return pages.Index()
	}, 60) // will revalidate after 60 seconds

	RegisterStaticComponent(e, "/pokemon", func(c echo.Context) templ.Component {
		var result types.PokemonList
		err := utils.Get("https://pokeapi.co/api/v2/pokemon", &result)

		if err != nil {
			return nil
		}

		return pages.Pokemon(result.Results)
	}, -1) // -1 means never revalidate

	RegisterStaticComponent(e, "/pokemon/:name", func(c echo.Context) templ.Component {
		var result types.Pokemon
		err := utils.Get("https://pokeapi.co/api/v2/pokemon/"+c.Param("name"), &result)

		if err != nil {
			return nil
		}

		return pages.PokemonSlug(result)
	}, -1) // -1 means never revalidate

	RegisterStaticComponent(e, "/clicked", func(c echo.Context) templ.Component {
		Globals["Clicked"] = Globals["Clicked"].(int) + 1
		return components.Clicked(Globals["Clicked"].(int))
	}, 0) // 0 means revalidate on every request
}
