package datasource

import (
	"encoding/json"
	"io"
	"net/http"
)

type CityRepository struct {
	HTTPClient HTTPClient
}

func NewCityRepository() *CityRepository {
	return &CityRepository{HTTPClient: &http.Client{}}
}

func NewCityRepositoryForTest(client HTTPClient) *CityRepository {
	return &CityRepository{HTTPClient: client}
}

func (c *CityRepository) FetchCityByCEP(cep string) (map[string]interface{}, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
