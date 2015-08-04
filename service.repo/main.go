package main

import "github.com/quangbuule/go-microservices-example/service.repo/service"

func main() {
	s := service.New()
	s.Initialize()
	s.Run()
}
