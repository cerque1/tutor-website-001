package main

import "github.com/cerque1/tutor-website-001/internal/app"

// @title Tutor API
// @version 1.0
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app.Run()
}
