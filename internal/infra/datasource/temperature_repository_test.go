package datasource

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"testing"
	"weatherCheck/configs"
)

type MockHTTPTemperatureClient struct {
	mock.Mock
}

func (m *MockHTTPTemperatureClient) Get(url string) (*http.Response, error) {
	args := m.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestMockTestFetchTemperatureByCity(t *testing.T) {

	mockHTTPClient := new(MockHTTPTemperatureClient)
	mockHTTPClient.On("Get", "http://api.weatherapi.com/v1/current.json?key=abc123&q=Itajub%C3%A1").Return(&http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(`{"current": {"temp_c": 25.0}}`)),
	}, nil)

	var httpClient HTTPClient = mockHTTPClient

	conf := &configs.Config{WeatherAPIKey: "abc123"}

	repositoryMocked := NewTemperatureRepositoryForTest(httpClient, conf)
	result, err := repositoryMocked.FetchTemperatureByCity(map[string]interface{}{"localidade": "Itajubá"})

	mockHTTPClient.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 25.0, result)
}
