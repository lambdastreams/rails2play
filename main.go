package main

import (
	"github.com/subbarao/transformer/pkg/http"
)

func main() {
	body, _ := http.GetMovie("nineth-gate")
	print(body)
}
