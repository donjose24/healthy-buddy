package main

import "github.com/jmramos02/healthybuddy/internal/api"

func main() {
	router := api.Initialize()

	router.Run()
}
