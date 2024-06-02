package main

import (
	"net/http"
	"weatherCheck/internal/infra/web"
)

func main() {
	http.HandleFunc("/cep/", web.GetTemperaturesHandler)
	http.ListenAndServe(":8080", nil)
}
