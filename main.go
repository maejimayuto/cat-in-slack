package main

import (
	"cat-in-slack/weather"
	"fmt"
)

func main() {
	result := weather.GetWeather()
	fmt.Printf("%s\n", result)
}
