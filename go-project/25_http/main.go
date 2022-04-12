package main

import (
	"encoding/json"
	"bytes"
	"io"
	"os"
	"net/http"
	"log"
)

// Get request

type Job struct {
	User string `json:"user"`
	Action string `json:"action"`
	Count int `json:"count"`
}

func main() {
	response, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("ERROR: can't call httpbin.org")
	}
	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)

// Post request

	job := &Job{
		User: "mohammad",
		Action: "punch",
		Count: 1,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("ERROR: can't encode job - %s", err)
	}

	response, err = http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("ERROR: can't call httpbin.org")
	}
	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)
}