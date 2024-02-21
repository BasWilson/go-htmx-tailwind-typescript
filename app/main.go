package main

import (
	"go-htmx-tailwind-typescript/app/framework"
	"go-htmx-tailwind-typescript/app/routes"
)

func main() {
	e, address := framework.Load()
	routes.Pages(e)
	routes.Api(e)
	e.Logger.Fatal(e.Start(address))
}
