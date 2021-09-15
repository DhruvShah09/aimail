package secure

import (
	"os"
)

func SecureEnvironment() {
	os.Setenv("OPENAI_KEY", "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")
	os.Setenv("EMAIL_USER", "testporson@gmail.com")
	os.Setenv("EMAIL_PWD", "Jenko123!")
	os.Setenv("WEATHER_KEY", "84e7da84ac5041069e2151349210809")
}
