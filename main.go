package main

import (

	//"encoding/json"

	//"net/http"

	//Absolute path of devops packges, refactoring necessary for production

	"fmt"

	"github.com/dhruvshah/go_crash_course/aimail/cdn"
	"github.com/dhruvshah/go_crash_course/aimail/secure" //secure init
	"github.com/dhruvshah/go_crash_course/aimail/weather"
)

func main() {
	secure.InitSecureEnvironment()
	locInput := "Orlando"
	location := weather.GetLoc(locInput)
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
