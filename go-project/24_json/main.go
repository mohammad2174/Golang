package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `
{
	"user": "mohammadreza khorrami",
	"type": "deposit",
	"amount": 1000000.5
}
`
type Request struct {
	Login string `json:"user"`
	Type string `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	reader := bytes.NewBufferString(data)
	decode := json.NewDecoder(reader)
	request := &Request{}

	if err := decode.Decode(request); err != nil {
		log.Fatalf("ERROR: can't decode - %s\n", err)
	}
	fmt.Printf("got: %+v\n", request)

	prevBalance := 8500000.0
	response := map[string]interface{}{
		"ok": true,
		"code": 200,
		"balance": prevBalance + request.Amount,

	}

	encode := json.NewEncoder(os.Stdout)
	if err := encode.Encode(response); err != nil {
		log.Fatalf("ERROR: can't encode - %s\n", err)
	}
}