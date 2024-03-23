package handlers

import "net/http"

func GetTemperature(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}
