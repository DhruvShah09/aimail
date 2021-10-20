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
	os.Setenv("OPENAI_KEY", "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")
	os.Setenv("EMAIL_USER", "testporson@gmail.com")
	os.Setenv("EMAIL_PWD", "Jenko123!")
	os.Setenv("WEATHER_KEY", "10680dad62226eaf44bf939ce5109d55")
}
