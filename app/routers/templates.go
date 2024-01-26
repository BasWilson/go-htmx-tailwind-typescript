package routers

import (
	"github.com/baswilson/go-htmx-tailwind-typescript/app/components"
	"github.com/baswilson/go-htmx-tailwind-typescript/app/pages"
	"github.com/baswilson/go-htmx-tailwind-typescript/app/utils"
	"github.com/labstack/echo"
)

var Globals = map[string]interface{}{
	"Clicked": 0,
}

func Templates(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error { return utils.Rt(c, pages.Index()) })
	e.GET("/hello", func(c echo.Context) error { return utils.Rt(c, pages.Hello("Bas Wilson")) })
	e.POST("/clicked", func(c echo.Context) error {
		Globals["Clicked"] = Globals["Clicked"].(int) + 1
		return utils.Rt(c, components.Clicked(Globals["Clicked"].(int)))
	})
}
