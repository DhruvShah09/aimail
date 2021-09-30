package main

import (

	//"encoding/json"

	//"net/http"

	//Absolute path of development environment packges, refactor for production

	"fmt"

	"github.com/dhruvshah/go_crash_course/aimail/secure" //secure init
	"github.com/dhruvshah/go_crash_course/aimail/weather"
)

func main() {
	secure.InitSecureEnvironment()
	//x := "Write a poem:"
	//response := cdn.GetPoem(x)
	//cdn.SendEmail(response)
	location := []string{"33.44", "-94.04"}
	if weather.GetWeather(location) {
		fmt.Println("Retrieved Weather Data")
	} else {
		fmt.Println("Error collecting weather data")
	}
}
