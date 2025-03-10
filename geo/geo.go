package geo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity, err := checkCity(city)
		if err != nil {
			return nil, fmt.Errorf("checkCity: %w", err)
		}
		if !isCity {
			return nil, errors.New("такого города нет")
		}
		return &GeoData{
			City: city,
		}, nil
	}
	resp, err := http.Get("https://freegeoip.app/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("NOT200: %w", err)
	}
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

func checkCity(city string) (bool, error) {
	baseUrl, err := url.Parse("https://nominatim.openstreetmap.org/search")
	if err != nil {
		return false, err
	}
	data := url.Values{}
	data.Add("city", city)
	data.Add("format", "json")
	baseUrl.RawQuery = data.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if len(body) < 3 {
		return false, nil
	} else {
		return true, nil
	}
}
