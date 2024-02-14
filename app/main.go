package main

import (
	"go-htmx-tailwind-typescript/app/framework"
)

func main() {
	e, address := framework.Load()
	e.Logger.Fatal(e.Start(address))
}
