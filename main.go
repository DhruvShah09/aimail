package main

import (
	"context"
	//"encoding/json"

	//"net/http"

	"os"
	//Absolute path of development environment packges, refactor for production

	"github.com/dhruvshah/go_crash_course/aimail/cdn"
	"github.com/dhruvshah/go_crash_course/aimail/secure" //secure init

	gogpt "github.com/sashabaranov/go-gpt3" //points to absolute path of GPT3 bindings, Change for distinct development environments
)

type WeData struct {
	temp_min      string
	temp_max      string
	Humidity      string
	Precipitation string
}

func main() {
	secure.SecureEnvironment()
	x := "Write a poem:"
	response := getPoem(x)
	cdn.SendEmail(response)
}

func getPoem(x string) string {
	c := gogpt.NewClient(os.Getenv("OPENAI_KEY"))
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		MaxTokens: 1024,
		Prompt:    x,
	}
	resp, err := c.CreateCompletion(ctx, "davinci-instruct-beta", req)
	if err != nil {
		return "Failed to gather text/body of email"
	}
	return resp.Choices[0].Text
}

/*
func generatePrompt() {
	a := []string{"Write a poem about the weather on a day with ", "% humidity, a high of ", " F, a low of 82 ", ", and rain chance of ", "%:"}

}
*/
/*
func relData(r *http.Response) []string {
	reader := json.NewDecoder(r.Body).Decode(target)
}
func getWeather(city string) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Atlanta&units=imperial&appid=921214e3238613cdd481cd104a3ec643")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data := relData(resp)

}
*/
