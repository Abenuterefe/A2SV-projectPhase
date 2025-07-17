package main

import (
	"task3/controllers"
	"task3/services"
)

func main() {
	library := services.NewLibrary()
	controllers.RunConsoleInterface(library, library)
}
