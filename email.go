package main

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/smtp"
	"os"
)

type EmailService struct {
	host     string
	port     string
	username string
	password string
	from     string
	logger   *slog.Logger
}

func NewEmailService(logger *slog.Logger) *EmailService {
	return &EmailService{
		host:     os.Getenv("SMTP_HOST"),
		port:     os.Getenv("SMTP_PORT"),
		username: os.Getenv("SMTP_USER"),
		password: os.Getenv("SMTP_PASSWORD"),
		from:     os.Getenv("SMTP_FROM"),
		logger:   logger,
	}
}

func (s *EmailService) SendEmail(to string, subject string, body string) error {
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n%s",
		s.from, to, subject, body)

	s.logger.Info("Starting email sending", "to", to, "subject", subject, "host", s.host)

	// Connexion TLS implicite (port 465 / SMTPS)
	tlsConfig := &tls.Config{
		ServerName: s.host,
	}
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		s.logger.Error("Failed to connect to SMTP server (TLS)", "addr", addr, "error", err.Error())
		return err
	}

	c, err := smtp.NewClient(conn, s.host)
	if err != nil {
		err := conn.Close()
		if err != nil {
			s.logger.Error("Failed to create SMTP client", "error", err.Error())
			return err
		}
		return err
	}
	defer func() {
		_ = c.Quit()
	}()

	// Authentification si nécessaire
	if s.username != "" && s.password != "" {
		s.logger.Debug("Authenticating with SMTP server", "user", s.username)
		auth := smtp.PlainAuth("", s.username, s.password, s.host)
		if err = c.Auth(auth); err != nil {
			s.logger.Error("SMTP authentication failed", "error", err.Error())
			return err
		}
	}

	// Définition de l'expéditeur et du destinataire
	if err = c.Mail(s.from); err != nil {
		s.logger.Error("Failed to set SMTP sender", "from", s.from, "error", err.Error())
		return err
	}
	if err = c.Rcpt(to); err != nil {
		s.logger.Error("Failed to set SMTP recipient", "to", to, "error", err.Error())
		return err
	}

	// Envoi du corps du message
	s.logger.Debug("Sending email data", "to", to)
	w, err := c.Data()
	if err != nil {
		s.logger.Error("Failed to open SMTP data writer", "error", err.Error())
		return err
	}
	_, err = w.Write([]byte(msg))
	if err != nil {
		s.logger.Error("Failed to write email data", "error", err.Error())
		return err
	}
	err = w.Close()
	if err != nil {
		s.logger.Error("Failed to close SMTP data writer", "error", err.Error())
		return err
	}

	s.logger.Info("Email sent successfully", "to", to)
	return nil
}

func (s *EmailService) SendPasswordResetEmail(to string, token string) error {
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", os.Getenv("FRONTEND_URL"), token)
	subject := "Reset your password"
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px;">
			<div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 30px; border-radius: 10px; box-shadow: 0 4px 6px rgba(0,0,0,0.1);">
				<h2 style="color: #4f46e5; text-align: center;">Password Reset</h2>
				<p>Hello,</p>
				<p>You have requested a password reset for your <strong>threadStocks</strong> account.</p>
				<p>Click the button below to change your password. This link will expire in 1 hour.</p>
				<div style="text-align: center; margin: 30px 0;">
					<a href="%s" style="background-color: #4f46e5; color: white; padding: 12px 24px; text-decoration: none; border-radius: 5px; font-weight: bold;">Reset my password</a>
				</div>
				<p>If you did not request this change, you can safely ignore this email.</p>
				<hr style="border: 0; border-top: 1px solid #eeeeee; margin: 20px 0;">
				<p style="font-size: 12px; color: #888888; text-align: center;">&copy; 2026 threadStocks. All rights reserved.</p>
			</div>
		</body>
		</html>
	`, resetLink)

	return s.SendEmail(to, subject, body)
}

func (s *EmailService) SendContactEmail(name, email, subject, message string) error {
	to := os.Getenv("CONTACT_EMAIL")
	emailSubject := fmt.Sprintf("New contact message: %s", subject)
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px;">
			<div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 30px; border-radius: 10px; box-shadow: 0 4px 6px rgba(0,0,0,0.1);">
				<h2 style="color: #4f46e5;">New Contact Message</h2>
				<p><strong>Name:</strong> %s</p>
				<p><strong>Email:</strong> %s</p>
				<p><strong>Subject:</strong> %s</p>
				<p><strong>Message:</strong></p>
				<div style="background-color: #f9fafb; padding: 15px; border-radius: 5px; border: 1px solid #e5e7eb;">
					%s
				</div>
				<hr style="border: 0; border-top: 1px solid #eeeeee; margin: 20px 0;">
				<p style="font-size: 12px; color: #888888; text-align: center;">&copy; 2026 threadStocks. Sent from the contact form.</p>
			</div>
		</body>
		</html>
	`, name, email, subject, message)

	return s.SendEmail(to, emailSubject, body)
}
