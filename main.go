package main

import "welcome-service/router"

func main() {
	r := router.InitRouters()

	port := "8080"

	r.Run(":" + port)
}
