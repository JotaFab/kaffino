package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

// SESV2API defines the interface for the SESV2 client.  This allows us to mock it in tests.
type SESV2API interface {
	SendEmail(ctx context.Context, params *sesv2.SendEmailInput, optFns ...func(*sesv2.Options)) (*sesv2.SendEmailOutput, error)
}

// SendEmailAWS sends an email using AWS SES V2.
func SendEmailAWS(svc SESV2API, to, subject, body string) error {
	from := "no-reply@sessioninit-kafff.jota-fab.com" // Replace with your verified no-reply email
	charSet := "UTF-8"

	// Assemble the email.
	input := &sesv2.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{
				to,
			},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Html: &types.Content{
						Charset: aws.String(charSet),
						Data:    aws.String(body), // Use HTML body
					},
					Text: &types.Content{
						Charset: aws.String(charSet),
						Data:    aws.String(body), // Provide a plain text version as well
					},
				},
				Subject: &types.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(subject),
				},
			},
		},
		FromEmailAddress: aws.String(from), // Set the FromEmailAddress
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(context.TODO(), input)

	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	fmt.Println("Email sent with message ID:", *result.MessageId)
	return nil
}

// NewSESV2Client creates a new SESV2 client using the default AWS configuration.
func NewSESV2Client() (SESV2API, error) {


	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("Unable to load SDK config, ", err)
		return nil, err
	}
	// Create an SES session.
	return sesv2.NewFromConfig(cfg), nil
}

// SendEmail sends an email with the given subject and body to the recipient using AWS SES V2.
func SendEmail(to, subject, body string) error {
	svc, err := NewSESV2Client()
	if err != nil {
		return err
	}
	return SendEmailAWS(svc, to, subject, body)
}