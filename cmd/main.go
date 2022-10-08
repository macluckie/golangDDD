package main

import (
	"test/golang/api"
	d "test/golang/domain"
	r "test/golang/repository"
)

// var wg sync.WaitGroup

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /v1
func main() {
	repo, er := r.NewRepo()
	if er != nil {
		panic("error repo")
	}
	services := d.Services{}
	r := api.NewRouter(*repo, services)
	r.Run(":8080")
}
