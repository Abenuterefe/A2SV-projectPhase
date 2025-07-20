package main

import (
	"A2SV-projectPhase/task4/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
