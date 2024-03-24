package handlers

import (
	"net/http"
	"willianszwy/FC-Cloud-Run/configs"
)

func GetTemperature(config *configs.Config) func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		writer.Write([]byte("hello world"))
	}
}
