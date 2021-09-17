package cdn

import (
	"context"
	"fmt"
	"net/smtp"
	"os"

	"github.com/dhruvshah/go_crash_course/aimail/formfactor"
	gogpt "github.com/sashabaranov/go-gpt3" //points to absolute path of GPT3 bindings, Change for distinct development environments
)

func GetPoem(x string) string {
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

func SendEmail(z string) {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PWD")

	// Receiver email address.
	to := []string{
		"dhruv3037@gmail.com",
		"dhruvshah434@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Email Content.
	subject := "Some poetry to enrich your day."
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := formfactor.ConstructEmail(z)
	message := []byte(subject + mime + body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
