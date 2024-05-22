package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Temperature    float64 `json:"temperature"`
		Humidity       float64 `json:"humidity"`
		Rain_Intensity float64 `json:"rain_intensity"`
	} `json:"locality_weather_data"`
}

func main() {

	// baseURL := "https://www.weatherunion.com/gw/weather/external/v0"

	//https://b.zmtcdn.com/data/file_assets/65fa362da3aa560a92f0b8aeec0dfda31713163042.pdf

	m := map[string]string{

		"Koundampalayam": "ZWL007600",
		"RS Puram":       "ZWL008653",
		"Saibaba Colony": "ZWL009668",
		"Koramangala":    "ZWL001156",
	}

	var city string
	fmt.Print("Name of the city: ")
	fmt.Scanln(&city)

	if m[city] == "" {
		panic("City not found! Try: Koundampalayam | RS Puram | Saibaba Colony | Koramangala ")
	}

	URL := "https://www.weatherunion.com/gw/weather/external/v0/get_locality_weather_data?locality_id="
	id := m[city]

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL+id, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Add("x-zomato-api-key", "Your API KEY")

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		// panic(err)
		fmt.Println("Error reading HTTP response body:", err)
		return
	}

	// Print the response body
	// fmt.Println(string(body))

	var weather Weather

	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}

	cityName := fmt.Sprintf("%s: \n", city)

	message := fmt.Sprintf(
		"Temperature: %.2f \nHumidity: %.2f \nRain Intensity: %.2f \n",
		weather.Location.Temperature,
		weather.Location.Humidity,
		weather.Location.Rain_Intensity,
	)

	time := time.Now().Format("02-01-2006 | 15:04")

	heading_color := color.New(color.FgCyan).Add(color.Bold)
	heading_color.Println(cityName)

	body_color := color.New(color.FgHiBlue)
	body_color.Println(message)

	time_color := color.New(color.BgHiYellow)
	time_color.Println(time)

}
