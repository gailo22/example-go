package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Unbind struct {
	Message string `json:"message"`
}

func main() {
	resp, err := request()
	if err != nil {
		log.Panic(err)
	}

	log.Println(resp)
}

func request() (out Unbind, err error) {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/404", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		log.Println(resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	var unbind Unbind
	err = json.Unmarshal(body, &unbind)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}

	out = unbind
	return
}
