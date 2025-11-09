package main

import (
	"Nov 3 - Nov 7/Task 5/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
