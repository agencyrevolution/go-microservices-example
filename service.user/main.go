package main

import "github.com/agencyrevolution/go-microservices-example/service.user/service"

func main() {
	s := service.New()
	s.Initialize()
	s.Run()
}
