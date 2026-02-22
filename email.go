package main

import (
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

	// Connexion directe au serveur SMTP
	c, err := smtp.Dial(addr)
	if err != nil {
		s.logger.Error("Failed to connect to SMTP server", "addr", addr, "error", err)
		return err
	}
	defer func(c *smtp.Client) {
		err := c.Quit()
		if err != nil {

		}
	}(c)

	// Authentification si nécessaire
	if s.username != "" && s.password != "" {
		s.logger.Debug("Authenticating with SMTP server", "user", s.username)
		auth := smtp.PlainAuth("", s.username, s.password, s.host)
		if err = c.Auth(auth); err != nil {
			s.logger.Error("SMTP authentication failed", "error", err)
			return err
		}
	}

	// Définition de l'expéditeur et du destinataire
	if err = c.Mail(s.username); err != nil {
		s.logger.Error("Failed to set SMTP sender", "from", s.username, "error", err)
		return err
	}
	if err = c.Rcpt(to); err != nil {
		s.logger.Error("Failed to set SMTP recipient", "to", to, "error", err)
		return err
	}

	// Envoi du corps du message
	s.logger.Debug("Sending email data", "to", to)
	w, err := c.Data()
	if err != nil {
		s.logger.Error("Failed to open SMTP data writer", "error", err)
		return err
	}
	_, err = w.Write([]byte(msg))
	if err != nil {
		s.logger.Error("Failed to write email data", "error", err)
		return err
	}
	err = w.Close()
	if err != nil {
		s.logger.Error("Failed to close SMTP data writer", "error", err)
		return err
	}

	s.logger.Info("Email sent successfully", "to", to)
	return nil
}

func (s *EmailService) SendPasswordResetEmail(to string, token string) error {
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", os.Getenv("FRONTEND_URL"), token)
	subject := "Réinitialisation de votre mot de passe"
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px;">
			<div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 30px; border-radius: 10px; box-shadow: 0 4px 6px rgba(0,0,0,0.1);">
				<h2 style="color: #4f46e5; text-align: center;">Réinitialisation de mot de passe</h2>
				<p>Bonjour,</p>
				<p>Vous avez demandé la réinitialisation de votre mot de passe pour votre compte <strong>threadStocks</strong>.</p>
				<p>Cliquez sur le bouton ci-dessous pour changer votre mot de passe. Ce lien expirera dans 1 heure.</p>
				<div style="text-align: center; margin: 30px 0;">
					<a href="%s" style="background-color: #4f46e5; color: white; padding: 12px 24px; text-decoration: none; border-radius: 5px; font-weight: bold;">Réinitialiser mon mot de passe</a>
				</div>
				<p>Si vous n'avez pas demandé ce changement, vous pouvez ignorer cet email en toute sécurité.</p>
				<hr style="border: 0; border-top: 1px solid #eeeeee; margin: 20px 0;">
				<p style="font-size: 12px; color: #888888; text-align: center;">&copy; 2026 threadStocks. Tous droits réservés.</p>
			</div>
		</body>
		</html>
	`, resetLink)

	return s.SendEmail(to, subject, body)
}

func (s *EmailService) SendContactEmail(name, email, subject, message string) error {
	to := os.Getenv("CONTACT_EMAIL")
	emailSubject := fmt.Sprintf("Nouveau message contact: %s", subject)
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px;">
			<div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 30px; border-radius: 10px; box-shadow: 0 4px 6px rgba(0,0,0,0.1);">
				<h2 style="color: #4f46e5;">Nouveau message de contact</h2>
				<p><strong>Nom:</strong> %s</p>
				<p><strong>Email:</strong> %s</p>
				<p><strong>Sujet:</strong> %s</p>
				<p><strong>Message:</strong></p>
				<div style="background-color: #f9fafb; padding: 15px; border-radius: 5px; border: 1px solid #e5e7eb;">
					%s
				</div>
				<hr style="border: 0; border-top: 1px solid #eeeeee; margin: 20px 0;">
				<p style="font-size: 12px; color: #888888; text-align: center;">&copy; 2026 threadStocks. Message envoyé depuis le formulaire de contact.</p>
			</div>
		</body>
		</html>
	`, name, email, subject, message)

	return s.SendEmail(to, emailSubject, body)
}
