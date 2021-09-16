package secure

import (
	"os"
)

func SecureEnvironment() {
	os.Setenv("OPENAI_KEY", "")
	os.Setenv("EMAIL_USER", "")
	os.Setenv("EMAIL_PWD", "")
	os.Setenv("WEATHER_KEY", "")
}
