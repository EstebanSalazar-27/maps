package main

import (
	"fmt"

	"github.com/estebansalazar-27/maps/keys"
)

type rating struct {
	value string
	id    string
}
type ratingMap map[string]rating

func (r ratingMap) outputRates() {
	fmt.Println(r)
}

const ratesLimit = 5

func main() {
	ratesMap := make(ratingMap, ratesLimit)
	ratesMap["one star"] = rating{value: "ST_01", id: "01"}
	ratesMap["two star"] = rating{value: "ST_02", id: "02"}
	ratesMap["three star"] = rating{value: "ST_03", id: "03"}
	ratesMap["four star"] = rating{value: "ST_04", id: "04"}
	ratesMap["five star"] = rating{value: "ST_05", id: "05"}

	for key, val := range ratesMap {
		fmt.Println(key)
		fmt.Println(val)
	}
	newMap := map[string]string{
		keys.GOOGLE_KEY: "https://google.com",
		keys.AWS_KEY:    "https://aws.com",
	}
	newMap[keys.LINKEDIN_KEY] = "https://linkedin.ar.com"
	fmt.Println(newMap)
	delete(newMap, keys.LINKEDIN_KEY)
	fmt.Println(newMap[keys.LINKEDIN_KEY])
}
