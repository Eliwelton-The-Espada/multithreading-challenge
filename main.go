package main

import (
	"fmt"
	"time"

	"github.com/Eliwelton-The-Espada/multithreading-challenge/client"
)

const (
	VIA_CEP     string = "ViaCep"
	CDN_API_CEP string = "CdnApiCep"
)

func main() {
	c1 := make(chan client.Response)
	c2 := make(chan client.Response)

	go func() {
		response := client.RequestApi("http://viacep.com.br/ws/55880-000/json/")
		c1 <- response
	}()

	go func() {
		response := client.RequestApi("https://cdn.apicep.com/file/apicep/55880-000.json")
		c2 <- response
	}()

	select {
	case response := <-c1:
		fmt.Printf("ApiName: %s\nStatusCode: %d\nResponseBody: %s\n", VIA_CEP, response.StatusCode, response.ResponseBody)
	case response := <-c2:
		fmt.Printf("ApiName: %s\nStatusCode: %d\nResponseBody: %s\n", CDN_API_CEP, response.StatusCode, response.ResponseBody)
	case <-time.After(time.Second):
		fmt.Println("Error: Timeout")
	}
}
