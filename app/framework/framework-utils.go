package framework

import (
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/baswilson/go-htmx-tailwind-typescript/app/templ/pages"
	"github.com/baswilson/go-htmx-tailwind-typescript/app/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

type CachedComponent struct {
	Component    templ.Component
	RevalidateAt int64 // Unix timestamp
}

var cache = map[string]CachedComponent{}

func RegisterStaticComponent(e *echo.Echo, url string, renderMethod func(ctx echo.Context) templ.Component, revalidate int64) {
	e.GET(url, func(ctx echo.Context) error {
		path := ctx.Request().URL.Path
		var template templ.Component
		shouldRender := true
		timestamp := int64(time.Now().Unix())
		if cachedTemplate, ok := cache[path]; ok {
			template = cachedTemplate.Component
			if revalidate == -1 || timestamp < cachedTemplate.RevalidateAt {
				shouldRender = false
			}
		}

		// revalidation period has passed or not cached yet, render the component
		if shouldRender {
			println("[rendering-static-page]: " + path)
			cache[path] = CachedComponent{
				Component:    renderMethod(ctx),
				RevalidateAt: timestamp + revalidate,
			}
			template = cache[path].Component
		} else {
			println("[cached]: " + path)
		}

		err := template.Render(ctx.Request().Context(), ctx.Response())
		if err != nil {
			return pages.Error("Error getting data from the Pokemon API").Render(ctx.Request().Context(), ctx.Response())
		}

		return nil
	})
}

func Load() (*echo.Echo, string) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	e := echo.New()
	e.HideBanner = true
	e.Static("/", "public")

	if os.Getenv("ENV") == "development" {
		e.GET("/ws", utils.WsHotReload)
	}

	Templates(e)
	Api(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return e, ":" + port
}
