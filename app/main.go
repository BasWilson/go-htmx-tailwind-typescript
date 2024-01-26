package main

import (
	"os"

	"github.com/baswilson/go-htmx-tailwind-typescript/app/routers"
	"github.com/baswilson/go-htmx-tailwind-typescript/app/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
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

	routers.Templates(e)
	routers.Api(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
