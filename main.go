package main

import (

	//"encoding/json"

	//"net/http"

	//Absolute path of development environment packges, refactor for production

	"github.com/dhruvshah/go_crash_course/aimail/cdn"
	"github.com/dhruvshah/go_crash_course/aimail/secure" //secure init
)

func main() {
	secure.InitSecureEnvironment()
	x := "Write a poem:"
	response := cdn.GetPoem(x)
	cdn.SendEmail(response)
}
