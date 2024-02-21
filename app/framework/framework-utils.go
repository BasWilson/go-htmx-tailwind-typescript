package framework

import (
	"fmt"
	"os"
	"time"

	"go-htmx-tailwind-typescript/app/templ/pages"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

type CachedComponent struct {
	Component    templ.Component
	RevalidateAt int64 // Unix timestamp
}

var cache = map[string]CachedComponent{}

func RegisterStaticComponent(e *echo.Echo, url string, renderMethod func(ctx echo.Context) templ.Component, revalidate int64, method string) {
	httpMethod := e.GET
	if method == "POST" {
		httpMethod = e.POST
	}
	httpMethod(url, func(ctx echo.Context) error {
		startTime := time.Now()
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
			println("[rendering-static-path]: " + path)
			renderedComponent := renderMethod(ctx)
			if renderedComponent != nil {
				cache[path] = CachedComponent{
					Component:    renderedComponent,
					RevalidateAt: timestamp + revalidate,
				}
				template = cache[path].Component
			} else {
				return pages.Error("Render method did not return a component").Render(ctx.Request().Context(), ctx.Response())
			}
		}

		err := template.Render(ctx.Request().Context(), ctx.Response())
		if err != nil {
			return pages.Error("Error rendering component: "+err.Error()).Render(ctx.Request().Context(), ctx.Response())
		}
		println("[rendered]: " + path + " in " + time.Since(startTime).String())
		return nil
	})
}

func Load() (*echo.Echo, string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	e := echo.New()
	e.HideBanner = true
	e.Static("/", "public")

	if os.Getenv("ENV") == "development" {
		e.GET("/ws", WsHotReload)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return e, ":" + port
}
