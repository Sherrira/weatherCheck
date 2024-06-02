package datasource

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TemperatureRepository struct {
	HTTPClient HTTPClient
}

func NewTemperatureRepository() *TemperatureRepository {
	return &TemperatureRepository{HTTPClient: &http.Client{}}
}

func NewTemperatureRepositoryForTest(client HTTPClient) *TemperatureRepository {
	return &TemperatureRepository{HTTPClient: client}
}

type TemperatureAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func (t *TemperatureRepository) FetchTemperatureByCity(city map[string]interface{}) (float64, error) {
	local := city["localidade"]
	resp, err := t.HTTPClient.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=36903aee082b4a3f9c5214023242505&q=%s", url.QueryEscape(local.(string))))
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
