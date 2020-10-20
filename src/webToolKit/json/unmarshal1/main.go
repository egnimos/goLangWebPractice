package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type city struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	City      string  `json:"City"`
}

type cities []city

func main() {
	var data cities

	jsonFile := `[
		{
			"precision": "zip",
			"Latitude":  37.7668,
			"Longitude": -122.3959,
			"Address":   "",
			"City":      "SAN FRANCISCO",
			"State":     "CA",
			"Zip":       "94107",
			"Country":   "US"
		},
		{
			"precision": "zip",
			"Latitude":  37.371991,
			"Longitude": -122.026020,
			"Address":   "",
			"City":      "SUNNYVALE",
			"State":     "CA",
			"Zip":       "94085",
			"Country":   "US"
		}
	]`

	err := json.Unmarshal([]byte(jsonFile), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}