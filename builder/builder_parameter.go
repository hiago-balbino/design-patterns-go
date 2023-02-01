package main

import (
	"fmt"
	"strings"
)

type Email struct {
	From, To, Subject, Body string
}

func (e Email) Send() {
	content := "EMAIL CONTENT\nFrom: %s\nTo: %s\nSubject: %s\nBody: %s\n"
	fmt.Printf(content, e.From, e.To, e.Subject, e.Body)
}

type EmailBuilder struct {
	email Email
}

func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	eb.email.From = from
	return eb
}

func (eb *EmailBuilder) To(to string) *EmailBuilder {
	eb.email.To = to
	return eb
}

func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.email.Subject = subject
	return eb
}

func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.email.Body = body
	return eb
}

type buildEmail func(*EmailBuilder)

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
