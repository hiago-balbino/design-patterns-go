package main

import (
	"fmt"
	"strings"
)

// Email is a model that contains all the information to build an email to be sent.
type Email struct {
	From, To, Subject, Body string
}

// Send is a fake function to send an email.
func (e Email) Send() {
	content := "EMAIL CONTENT\nFrom: %s\nTo: %s\nSubject: %s\nBody: %s\n"
	fmt.Printf(content, e.From, e.To, e.Subject, e.Body)
}

// EmailBuilder is a struct that contains functions to create an email.
type EmailBuilder struct {
	email Email
}

// From is a function to set the email sender.
func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	eb.email.From = from
	return eb
}

// To is a function to set the email recipient.
func (eb *EmailBuilder) To(to string) *EmailBuilder {
	eb.email.To = to
	return eb
}

// Subject is a function to set the email subject.
func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.email.Subject = subject
	return eb
}

// Body is a function to set the email body.
func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.email.Body = body
	return eb
}

type buildEmail func(*EmailBuilder)

// SendEmail is a function that receives an action to create an email and send it.
func SendEmail(action buildEmail) {
	builder := EmailBuilder{}
	action(&builder)
	builder.email.Send()
}

func main() {
	SendEmail(func(eb *EmailBuilder) {
		eb.
			From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
}
