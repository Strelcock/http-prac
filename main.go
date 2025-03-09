package main

import (
	"flag"
	"fmt"
	"log"
	"myHttp/geo"
	"myHttp/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода")
	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		log.Fatalf("getMyLocation: %s", err)
	}
	fmt.Printf("%+v\n", geoData)

	weatherData, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		log.Fatalf("GetWeather: %s", err)
	}
	fmt.Println(weatherData)
}
