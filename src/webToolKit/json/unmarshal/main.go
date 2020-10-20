package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// "fmt"

// "net/http"
// "html/template"

type thumbnail struct {
	URL           string `json:"Url"`
	Height, Width int
}

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

func main() {
	var data img
	jsonFile := `{
        "Width":  800,
        "Height": 600,
        "Title":  "View from 15th Floor",
        "Thumbnail": {
            "Url":    "http://www.example.com/image/481989943",
            "Height": 125,
            "Width":  100
        },
        "Animated" : false,
        "IDs": [116, 943, 234, 38793]
	}`

	err := json.Unmarshal([]byte(jsonFile), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)

	for i, v := range data.IDs {
		fmt.Println(i, v)
	}

	fmt.Println(data.Thumbnail.URL)
}
