package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/smtp"

	"github.com/hibiken/asynq"
)

const TypeEmailDelivery = "email:send"

type EmailDeliveryPayload struct {
	Email string
	Code  string
}

func NewEmailDeliveryTask(email, code string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{Email: email, Code: code})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	// Example: sending email via Mailtrap SMTP
	// auth := smtp.PlainAuth("", "64250fc141eb4d", "9877c867c77055", "smtp.mailtrap.io")
	// msg := []byte("Subject: Your OTP Code\n\nYour OTP is " + p.Code)
	// err := smtp.SendMail("smtp.mailtrap.io:2525", auth, "no-reply@example.com", []string{p.Email}, msg)
	// if err != nil {
	// 	return fmt.Errorf("mail error: %w", err)
	// }

	auth := smtp.PlainAuth("", "rahmani@orhan.sh", "fige wejw raik kuua", "smtp.gmail.com")
	msg := []byte("Subject: Your OTP Code\n\nYour OTP is " + p.Code)
	err := smtp.SendMail("smtp.gmail.com:587", auth, "no-reply@example.com", []string{p.Email}, msg)
	if err != nil {
		return fmt.Errorf("mail error: %w", err)
	}
	return nil
}
