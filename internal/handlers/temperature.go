package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"willianszwy/FC-Cloud-Run/configs"
	"willianszwy/FC-Cloud-Run/internal/interfaces"
	"willianszwy/FC-Cloud-Run/internal/temperature"
	"willianszwy/FC-Cloud-Run/internal/viacep"
	"willianszwy/FC-Cloud-Run/internal/weather"
)

func GetTemperature(config *configs.Config, httpClient interfaces.HTTPClient) func(writer http.ResponseWriter, request *http.Request) {
	log.Println("starting get temp handlers")
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("starting request")
		viaCepClient := viacep.New(httpClient)
		weatherClient := weather.New(httpClient, config.WeatherAPIKey)

		zipcode := request.URL.Query().Get("zipcode")
		log.Println(fmt.Sprintf("[zipcode:%s]", zipcode))

		city, err := viaCepClient.FindByZipCode(request.Context(), zipcode)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		tempByCity, err := weatherClient.FindTempByCity(request.Context(), city.Name)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := temperature.New(tempByCity.Current.TempC, tempByCity.Current.TempF)
		writer.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(writer).Encode(resp); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}
