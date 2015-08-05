package main

import "github.com/agencyrevolution/go-microservices-example/service.repo/service"

func main() {
	s := service.New()
	s.Initialize()
	s.KeepAlive()
	s.Run()
}
