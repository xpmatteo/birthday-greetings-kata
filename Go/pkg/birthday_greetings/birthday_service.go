package birthday_greetings

import (
	"bufio"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

// BirthdayService handles sending birthday greetings
type BirthdayService struct{}

// NewBirthdayService creates a new BirthdayService
func NewBirthdayService() *BirthdayService {
	return &BirthdayService{}
}

// SendGreetings sends birthday greetings to employees whose birthday is today
func (bs *BirthdayService) SendGreetings(fileName string, xDate *XDate, smtpHost string, smtpPort int) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Skip header
	if scanner.Scan() {
		_ = scanner.Text()
	}

	for scanner.Scan() {
		line := scanner.Text()
		employeeData := strings.Split(line, ", ")

		if len(employeeData) < 4 {
			continue // Skip invalid lines
		}

		employee, err := NewEmployee(employeeData[1], employeeData[0], employeeData[2], employeeData[3])
		if err != nil {
			return fmt.Errorf("failed to create employee: %w", err)
		}

		if employee.IsBirthday(xDate) {
			recipient := employee.GetEmail()
			body := strings.Replace("Happy Birthday, dear %NAME%", "%NAME%", employee.GetFirstName(), -1)
			subject := "Happy Birthday"
			err := bs.sendMessage(smtpHost, smtpPort, "sender@here.com", subject, body, recipient)
			if err != nil {
				return fmt.Errorf("failed to send message: %w", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

// sendMessage sends an email message
func (bs *BirthdayService) sendMessage(smtpHost string, smtpPort int, sender, subject, body, recipient string) error {
	// Construct the message
	message := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", sender, recipient, subject, body))

	// Connect to the server, authenticate, and send the message
	addr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)
	auth := smtp.PlainAuth("", "", "", smtpHost)

	// In a real application, we would use authentication, but for simplicity and to match
	// the Java version's behavior, we'll skip it for localhost
	var err error
	if smtpHost == "localhost" {
		err = smtp.SendMail(addr, nil, sender, []string{recipient}, message)
	} else {
		err = smtp.SendMail(addr, auth, sender, []string{recipient}, message)
	}

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
