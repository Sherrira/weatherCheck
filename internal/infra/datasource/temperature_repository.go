package datasource

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weatherCheck/configs"
)

type TemperatureRepository struct {
	HTTPClient HTTPClient
	conf       configs.Config
}

func NewTemperatureRepository() *TemperatureRepository {
	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	return &TemperatureRepository{
		HTTPClient: &http.Client{},
		conf:       *conf,
	}
}

func NewTemperatureRepositoryForTest(client HTTPClient, conf *configs.Config) *TemperatureRepository {
	return &TemperatureRepository{
		HTTPClient: client,
		conf:       *conf,
	}
}

type TemperatureAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func (t *TemperatureRepository) FetchTemperatureByCity(city map[string]interface{}) (float64, error) {
	local := city["localidade"]
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", t.conf.WeatherAPIKey, url.QueryEscape(local.(string)))
	fmt.Printf("Calling GET weatherapi: %v\n", url)
	resp, err := t.HTTPClient.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather API: %v\n", err)
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading weather API response: %v\n", err)
		return 0, err
	}

	fmt.Printf("Success fetching temperature: %s\n", body)
	var weatherAPIResponse TemperatureAPIResponse
	err = json.Unmarshal(body, &weatherAPIResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling weather API response: %v\n", err)
		return 0, err
	}
	return weatherAPIResponse.Current.TempC, nil
}
