package geo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Status)
	// if resp.StatusCode != 200 {
	// 	return nil, fmt.Errorf("NOT200: %s", err)
	// }
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cant read body: %w", err)
	}
	var geo GeoData
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshall: %w", err)
	}
	return &geo, nil

}
