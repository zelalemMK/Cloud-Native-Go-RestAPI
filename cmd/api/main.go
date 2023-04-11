package main

import (
	"log"
	"net/http"

	"myapp/api/router"
	"myapp/config"
)

//  @title          MYAPP API
//  @version        1.0
//  @description    This is a sample RESTful API with a CRUD

//  @license.name   MIT License
//  @license.url    https://github.com/learning-cloud-native-go/myapp/blob/master/LICENSE

// @host       localhost:8080
// @basePath   /v1
func main() {
	c := config.New()
	r := router.New()
	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  c.Server.TimeOurRead,
		WriteTimeout: c.Server.TimeOutWrite,
		IdleTimeout:  c.Server.TimeOutIdle,
	}

	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println(err)
		log.Fatal("Server startup failed")
	}

}
