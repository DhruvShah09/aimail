package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ParsedWeatherData struct {
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

func GetWeather(coords []float64) (bool, string) {
	apikey := os.Getenv("WEATHER_KEY")
	callString := "https://api.openweathermap.org/data/2.5/onecall?lat=" + fmt.Sprintf("%f", coords[0]) + "&lon=" + fmt.Sprintf("%f", coords[1]) + "&exclude=minutely,hourly,alerts,current&units=imperial&appid=" + apikey
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
		var result ParsedWeatherData
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
		return false, "Network Error"
	}

}

type ParsedGeoData []struct {
	Name       string `json:"name"`
	LocalNames struct {
		Ar          string `json:"ar"`
		ASCII       string `json:"ascii"`
		Bg          string `json:"bg"`
		Da          string `json:"da"`
		De          string `json:"de"`
		El          string `json:"el"`
		En          string `json:"en"`
		Fa          string `json:"fa"`
		FeatureName string `json:"feature_name"`
		Fi          string `json:"fi"`
		Fr          string `json:"fr"`
		He          string `json:"he"`
		ID          string `json:"id"`
		Ja          string `json:"ja"`
		Lt          string `json:"lt"`
		Mk          string `json:"mk"`
		Nl          string `json:"nl"`
		No          string `json:"no"`
		Pl          string `json:"pl"`
		Pt          string `json:"pt"`
		Ru          string `json:"ru"`
		Sr          string `json:"sr"`
		Th          string `json:"th"`
	} `json:"local_names"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

func GetLoc(city string) []float64 {
	requestString := "http://api.openweathermap.org/geo/1.0/direct?q=" + city + "&limit=1&appid=" + os.Getenv("WEATHER_KEY")
	resp, err := http.Get(requestString)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var geos ParsedGeoData
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error Parsing Body of Location API Request")
		}
		err2 := json.Unmarshal(body, &geos)
		if err2 != nil {
			fmt.Println(err2)
		}
		coords := []float64{geos[0].Lat, geos[0].Lon}
		return coords
	}
	//temp fix
	broken := GetLoc("Atlanta")
	return broken
}
