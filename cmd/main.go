package main

import (
	"JWT-Authentications-RestAPI/config"
	"JWT-Authentications-RestAPI/routes"
)

func main() {
	// Initialize Database
	config.Connect("root:Hamza@10@tcp(localhost:3306)/jwt_demo?parseTime=true")
	config.Migrate()
	// Initialize Router
	routes.InitRouter()

}
