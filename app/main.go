package main

import (
	"goauth/models"
	"goauth/routes"
)

func main() {
	models.Init()
	routes.Setup()
}
