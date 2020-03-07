package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Basic weather client
type HttpClient struct {
	// empty for now
}

// Grab weather data from OpenWeather api
// deconstruct response, if resp-code is not 404,
// print Temp, Humidity, Wind, and Local prediction,
// otherwise, print city not found...
// return nil
func (client *HttpClient) getWeather(city string) {
	// fetch weather data
	appid := "key-goes-here"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, appid)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Print(err)
		return
	}
	// read response body
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return
	}
	var raw map[string]interface{}
	// decode response body
	if err := json.Unmarshal(text, &raw); err != nil {
		fmt.Println(err)
	}
	// if return code is 404, city was not found
	if raw["cod"] == "404" {
		fmt.Printf("City not found\n")
		return
	}
	// parse json data
	features := raw["main"].(map[string]interface{})
	fmt.Printf("Current Temperature: %.1f f\n", kelvToF(features["temp"].(float64)))
	fmt.Printf("Humidity: %.1f \n", features["humidity"].(float64))
	wind := raw["wind"].(map[string]interface{})
	fmt.Printf("Wind: %.1f \n", wind["speed"].(float64))
	pred := raw["weather"].([]interface{})
	descr := pred[0].(map[string]interface{})
	fmt.Printf("Prediction: %s \n", descr["description"])

}

// convert kelvin to f
// expects float64 as arguments...
// returns converted temp as float64
func kelvToF(k float64) float64 {
	return (9 / 5) * (k - 273) + 32
}