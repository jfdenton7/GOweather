package main

import (
	"fmt"
	"os"
	"strings"
)

// Main entry point to weather terminal
// requests OpenWeather API based on user
// entered city name
// spaces allowed in argument call
// examples:
// weather New York
// weather Blacksburg
func main() {
	client := new(HttpClient)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: weather <city-name>\n")
		return
	}
	city := strings.Join(os.Args[1:], " ")
	client.getWeather(city)
}
