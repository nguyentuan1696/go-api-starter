package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net"
	"net/smtp"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	// RFC 5322 compliant email regex
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	// Global email config
	globalEmailConfig *EmailConfig
	globalEmailOnce   sync.Once
)

// EmailConfig holds SMTP configuration
type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	FromName string
}

// EmailMessage represents an email message
type EmailMessage struct {
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
	IsHTML  bool
}

// IsValidEmail checks if email format is valid and within length limits
func IsValidEmail(email string) bool {
	if len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}

// IsValidEmailDomain checks if email domain has valid MX records
func IsValidEmailDomain(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	domain := parts[1]
	if len(domain) > 253 {
		return false
	}

	mx, err := net.LookupMX(domain)
	if err != nil || len(mx) == 0 {
		return false
	}

	return true
}

// SendEmail sends an email using SMTP
func SendEmail(config EmailConfig, message EmailMessage) error {
	// Validate email addresses
	for _, email := range message.To {
		if !IsValidEmail(email) {
			return fmt.Errorf("invalid recipient email: %s", email)
		}
	}

	// Create SMTP auth
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)

	// Build email content
	emailContent, err := buildEmailContent(config, message)
	if err != nil {
		return fmt.Errorf("failed to build email content: %w", err)
	}

	// Get all recipients
	allRecipients := append(message.To, message.Cc...)
	allRecipients = append(allRecipients, message.Bcc...)

	// Send email
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	err = smtp.SendMail(addr, auth, config.From, allRecipients, emailContent)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

// SendEmailTLS sends an email using SMTP with TLS
func SendEmailTLS(config EmailConfig, message EmailMessage) error {
	// Validate email addresses
	for _, email := range message.To {
		if !IsValidEmail(email) {
			return fmt.Errorf("invalid recipient email: %s", email)
		}
	}

	// Connect to SMTP server
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	conn, err := tls.Dial("tcp", addr, &tls.Config{
		ServerName: config.Host,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer conn.Close()

	// Create SMTP client
	client, err := smtp.NewClient(conn, config.Host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Quit()

	// Authenticate
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP authentication failed: %w", err)
	}

	// Set sender
	if err = client.Mail(config.From); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}

	// Set recipients
	allRecipients := append(message.To, message.Cc...)
	allRecipients = append(allRecipients, message.Bcc...)
	for _, recipient := range allRecipients {
		if err = client.Rcpt(recipient); err != nil {
			return fmt.Errorf("failed to set recipient %s: %w", recipient, err)
		}
	}

	// Send email content
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to get data writer: %w", err)
	}
	defer w.Close()

	// Build and write email content
	emailContent, err := buildEmailContent(config, message)
	if err != nil {
		return fmt.Errorf("failed to build email content: %w", err)
	}

	if _, err = w.Write(emailContent); err != nil {
		return fmt.Errorf("failed to write email content: %w", err)
	}

	return nil
}

// SendTemplateEmail sends an email using HTML template
func SendTemplateEmail(config EmailConfig, to []string, subject string, templatePath string, data interface{}) error {
	// Parse template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute template
	var body bytes.Buffer
	if err = tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Create email message
	message := EmailMessage{
		To:      to,
		Subject: subject,
		Body:    body.String(),
		IsHTML:  true,
	}

	// Send email
	return SendEmailTLS(config, message)
}

// SendResetPasswordEmail sends password reset email
func SendResetPasswordEmail(config EmailConfig, to string, resetToken string, resetURL string) error {
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>Password Reset Request</h2>
			<p>You have requested to reset your password. Click the link below to reset your password:</p>
			<p><a href="%s?token=%s">Reset Password</a></p>
			<p>If you did not request this, please ignore this email.</p>
			<p>This link will expire in 1 hour.</p>
		</body>
		</html>
	`, resetURL, resetToken)

	message := EmailMessage{
		To:      []string{to},
		Subject: "Password Reset Request",
		Body:    body,
		IsHTML:  true,
	}

	return SendEmailTLS(config, message)
}

// SendVerificationEmail sends email verification
func SendVerificationEmail(config EmailConfig, to string, verificationToken string, verificationURL string) error {
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>Email Verification</h2>
			<p>Please verify your email address by clicking the link below:</p>
			<p><a href="%s?token=%s">Verify Email</a></p>
			<p>If you did not create an account, please ignore this email.</p>
		</body>
		</html>
	`, verificationURL, verificationToken)

	message := EmailMessage{
		To:      []string{to},
		Subject: "Email Verification",
		Body:    body,
		IsHTML:  true,
	}

	return SendEmailTLS(config, message)
}

// buildEmailContent builds the complete email content with headers
func buildEmailContent(config EmailConfig, message EmailMessage) ([]byte, error) {
	var buffer bytes.Buffer

	// Email headers
	fromHeader := config.From
	if config.FromName != "" {
		fromHeader = fmt.Sprintf("%s <%s>", config.FromName, config.From)
	}

	buffer.WriteString(fmt.Sprintf("From: %s\r\n", fromHeader))
	buffer.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(message.To, ", ")))

	if len(message.Cc) > 0 {
		buffer.WriteString(fmt.Sprintf("Cc: %s\r\n", strings.Join(message.Cc, ", ")))
	}

	buffer.WriteString(fmt.Sprintf("Subject: %s\r\n", message.Subject))
	buffer.WriteString("MIME-Version: 1.0\r\n")

	if message.IsHTML {
		buffer.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
	} else {
		buffer.WriteString("Content-Type: text/plain; charset=UTF-8\r\n")
	}

	buffer.WriteString("\r\n")
	buffer.WriteString(message.Body)

	return buffer.Bytes(), nil
}

// InitEmailConfig khởi tạo global email config một lần
func InitEmailConfig(config EmailConfig) {
	globalEmailOnce.Do(func() {
		globalEmailConfig = &config
	})
}

// GetEmailConfig trả về global email config
func GetEmailConfig() *EmailConfig {
	if globalEmailConfig == nil {
		panic("email config not initialized. Call InitEmailConfig first")
	}
	return globalEmailConfig
}

// SendTemplateEmailFromTemplatesDir sends email using template from templates directory with global config
func SendTemplateEmailFromTemplatesDir(to []string, subject string, templateName string, data interface{}) error {
	config := GetEmailConfig()
	return SendTemplateEmailFromTemplatesDirWithConfig(*config, to, subject, templateName, data)
}

// SendTemplateEmailFromTemplatesDirWithConfig sends email using template from templates directory with custom config
func SendTemplateEmailFromTemplatesDirWithConfig(config EmailConfig, to []string, subject string, templateName string, data interface{}) error {
	// Build template path from templates directory
	templatePath := filepath.Join("templates", templateName)

	// Parse template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	// Execute template
	var body bytes.Buffer
	if err = tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", templatePath, err)
	}

	// Create email message
	message := EmailMessage{
		To:      to,
		Subject: subject,
		Body:    body.String(),
		IsHTML:  true,
	}

	// Send email
	return SendEmailTLS(config, message)
}

// SendResetPasswordEmailWithTemplate sends password reset email using HTML template with global config
func SendResetPasswordEmailWithTemplate(to string, name string, username string, resetToken string, resetURL string) error {
	config := GetEmailConfig()
	return SendResetPasswordEmailWithTemplateAndConfig(*config, to, name, username, resetToken, resetURL)
}

// SendResetPasswordEmailWithTemplateAndConfig sends password reset email using HTML template with custom config
func SendResetPasswordEmailWithTemplateAndConfig(config EmailConfig, to string, name string, username string, resetToken string, resetURL string) error {
	// Prepare template data
	data := TemplateData{
		ResetLink: fmt.Sprintf("%s?token=%s", resetURL, resetToken),
	}

	// Send email using template
	return SendTemplateEmailFromTemplatesDirWithConfig(config, []string{to}, "Password Reset Request", "reset_password.html", data)
}

// TemplateData represents data for email templates
type TemplateData struct {
	ResetLink string
	Token     string
	URL       string
	OTPCode   string
}

// LoadTemplateFromDir loads and parses template from templates directory
func LoadTemplateFromDir(templateName string) (*template.Template, error) {
	templatePath := filepath.Join("templates", templateName)
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load template %s: %w", templatePath, err)
	}
	return tmpl, nil
}

// RenderTemplate renders template with data
func RenderTemplate(tmpl *template.Template, data interface{}) (string, error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}
	return buf.String(), nil
}
