package main

import (
	"flag"
	"fmt"

	"github.com/exolightor/golang-for-devops-course/assignment1/pkg/get"
)

type Words struct {
	Page         string             `json:"page"`
	Words        []string           `json:"words"`
	Percentages  map[string]float64 `json:"percentages"`
	Special      []string           `json:"special"`
	ExtraSpecial []interface{}      `json:"extraSpecial"`
}

func main() {
	var requestUrl string
	flag.StringVar(&requestUrl, "requestUrl", "", "enter url for GEt request")
	flag.Parse()

	a := get.New(requestUrl)
	words := a.MakeGetRequest()
	fmt.Println(words)
}
