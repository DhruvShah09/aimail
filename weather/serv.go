package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ParsedResponse struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Daily          []struct {
		Dt        int     `json:"dt"`
		Sunrise   int     `json:"sunrise"`
		Sunset    int     `json:"sunset"`
		Moonrise  int     `json:"moonrise"`
		Moonset   int     `json:"moonset"`
		MoonPhase float64 `json:"moon_phase"`
		Temp      struct {
			Day   float64 `json:"day"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"temp"`
		FeelsLike struct {
			Day   float64 `json:"day"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		DewPoint  float64 `json:"dew_point"`
		WindSpeed float64 `json:"wind_speed"`
		WindDeg   int     `json:"wind_deg"`
		WindGust  float64 `json:"wind_gust"`
		Weather   []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds int     `json:"clouds"`
		Pop    float64 `json:"pop"`
		Rain   float64 `json:"rain,omitempty"`
		Uvi    float64 `json:"uvi"`
	} `json:"daily"`
}

func GetWeather(coords []string) (bool, string) {
	apikey := os.Getenv("WEATHER_KEY")
	callString := "https://api.openweathermap.org/data/2.5/onecall?lat=" + coords[0] + "&lon=" + coords[1] + "&exclude=minutely,hourly,alerts,current&units=imperial&appid=" + apikey
	resp, err := http.Get(callString)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Parse error")
		}
		var result ParsedResponse
		err2 := json.Unmarshal(body, &result)
		if err2 != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}
		//Prompt Generation
		fmt.Println(result.Daily[0].Pop)
		fmt.Println(result.Daily[0].Clouds)
		b := []string{"Write a poem about the weather on a day with a high of ", " degrees Fahrenheit, a low of ", " degrees Fahrenheit, ", "% humidity, ", "and a rain chance of ", "%"}
		prompt_z := b[0] + fmt.Sprintf("%g", result.Daily[0].Temp.Max) + b[1] + fmt.Sprintf("%g", result.Daily[0].Temp.Min) + b[2] + fmt.Sprintf("%d", result.Daily[0].Humidity) + b[3] + b[4] + fmt.Sprintf("%g", result.Daily[0].Pop*100) + b[5]
		return true, prompt_z
	} else {
		return false, "Error recieving weather data"
	}

}
