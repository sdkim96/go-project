package main

import (
	router "sdkim-go-project/router"
)

func main() {
	router.SetupRouter().Run(":8080")
}
