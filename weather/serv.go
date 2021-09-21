package weather

type Coords struct {
	Latitude  float64
	Longitude float64
}

func GetWeather(c string) {

}

func GeneratePrompt(hu string, hi string, lo string, r string) string {
	a := []string{"Write a poem about the weather on a day with ", "% humidity, a high of ", " F, a low of ", ", and rain chance of ", "%:"}
	prompt := a[0] + hu + a[1] + hi + a[2] + lo + a[3] + r + a[4]
	return prompt
}
