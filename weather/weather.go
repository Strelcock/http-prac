package weather

import (
	"io"
	"myHttp/geo"
	"net/http"
	"net/url"
	"strconv"
)

func GetWeather(geo geo.GeoData, format int) (string, error) {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("format", strconv.Itoa(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
