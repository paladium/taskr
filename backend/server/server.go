package server

import "taskr/config"

// Init function initalises the router
func Init() {
	config := config.GetConfig()
	r := router()
	r.Run(config.GetString("server.port"))
}
