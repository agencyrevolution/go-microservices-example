package service

import "os"

type options struct {
	Port string
}

func getOptions() *options {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return &options{port}
}
