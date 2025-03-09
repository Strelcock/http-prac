package main

import (
	"flag"
	"fmt"
	"log"
	"myHttp/geo"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	// format := flag.Int("format", 1, "Формат вывода")
	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		log.Fatalf("getMyLocation: %s", err)
	}
	fmt.Printf("%+v", geoData)
}
