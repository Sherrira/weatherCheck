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
	resp, err := t.HTTPClient.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s=%s", t.conf.WeatherAPIKey, url.QueryEscape(local.(string))))
	if err != nil {

		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var weatherAPIResponse TemperatureAPIResponse
	err = json.Unmarshal(body, &weatherAPIResponse)
	if err != nil {
		return 0, err
	}
	return weatherAPIResponse.Current.TempC, nil
}
