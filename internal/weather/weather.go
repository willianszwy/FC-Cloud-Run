package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"willianszwy/FC-Cloud-Run/internal/interfaces"
)

type Response struct {
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
	} `json:"current"`
}

type Weather struct {
	client interfaces.HTTPClient
	Apikey string
}

func New(client interfaces.HTTPClient, apikey string) *Weather {
	return &Weather{client: client, Apikey: apikey}
}

func (w *Weather) FindTempByCity(ctx context.Context, city string) (Response, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", w.Apikey, city)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Response{}, fmt.Errorf("error creating request %w", err)
	}
	resp, err := w.client.Do(req)
	if err != nil {
		return Response{}, fmt.Errorf("error doing request %w", err)
	}
	defer resp.Body.Close()
	var weatherResponse Response
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return Response{}, fmt.Errorf("error deconding request %w", err)
	}
	return weatherResponse, nil
}
