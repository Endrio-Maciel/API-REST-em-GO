package main

import (
	"api_rest/api"
	"api_rest/service"
	"net/http"
)

func main() {
	svc := &service.Application{
		Data: make(map[service.Id]service.User),
	}

	app := &api.Application{
		Service: svc,
	}

	http.ListenAndServe(":8080", app.Routes())
}
