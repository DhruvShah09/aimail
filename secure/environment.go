package secure

import (
	"os"
)

//Insert Relavent API KEYS and Passwords Here

/*
	OPENAI_KEY: Corresponds to OPEN AI API key. Make sure to rotate before production
	EMAIL_USER: Email address, GMAIL (not sure if others work)
	EMAIL_PWD: Password for above email
	WEATHER_KEY: API Key for weather API -- Open Weather API.
*/
func InitSecureEnvironment() {
	os.Setenv("OPENAI_KEY", "")
	os.Setenv("EMAIL_USER", "")
	os.Setenv("EMAIL_PWD", "")
	os.Setenv("WEATHER_KEY", "x")
}
