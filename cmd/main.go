package main

import (
	"channel-contents-collector/api"
)

func main() {
	router := api.NewRouter()
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
