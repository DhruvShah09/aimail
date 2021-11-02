package main

import (

	//"encoding/json"

	//"net/http"

	//Absolute path of development environment packges, refactor for production

	"fmt"

	"github.com/dhruvshah/go_crash_course/aimail/cdn"
	"github.com/dhruvshah/go_crash_course/aimail/secure" //secure init
	"github.com/dhruvshah/go_crash_course/aimail/weather"
)

func main() {
	secure.InitSecureEnvironment()
	location := weather.GetLoc("Atlanta")
	rec, prompt := weather.GetWeather(location)
	if rec {
		fmt.Println("Recieved Weather Data!")
		fmt.Println("The prompt is: " + prompt)
	} else {
		fmt.Println(prompt)
	}
	response := cdn.GetPoem(prompt)
	cdn.SendEmail(response)
}
