package main

import "github.com/jmramos02/healthy-buddy/internal/api"

func main() {
	router := api.Initialize()

	router.Run()
}
