package birthday_greetings

import (
	"os"
	"strings"
	"testing"
	"time"

	smtpmock "github.com/mocktools/go-smtp-mock"
)

const NONSTANDARD_PORT = 9999

const testData = `last_name, first_name, date_of_birth, email
Doe, John, 1982/10/08, john.doe@foobar.com
Ann, Mary, 1975/03/11, mary.ann@foobar.com
`

func TestWillSendGreetings_WhenItsSomebodysBirthday(t *testing.T) {
	// Arrange
	server := smtpmock.New(smtpmock.ConfigurationAttr{
		PortNumber:        NONSTANDARD_PORT,
		LogToStdout:       false,
		LogServerActivity: false,
	})
	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start SMTP server: %v", err)
	}
	defer server.Stop()
	birthdayService := NewBirthdayService()
	createTestEmployeeDataFile(t)
	date, err := NewXDateFromString("2008/10/08")
	if err != nil {
		t.Fatalf("Failed to create date: %v", err)
	}

	// Act
	err = birthdayService.SendGreetings("employee_data.txt", date, "localhost", NONSTANDARD_PORT)
	if err != nil {
		t.Fatalf("Failed to send greetings: %v", err)
	}

	// Assert

	// Wait a bit for the message to be processed
	time.Sleep(100 * time.Millisecond)

	// Check that a message was sent
	if len(server.Messages()) != 1 {
		t.Errorf("Expected 1 message to be sent, got %d", len(server.Messages()))
	}

	// Get the message and check its contents
	messages := server.Messages()
	if len(messages) != 1 {
		t.Fatalf("Expected 1 message, got %d", len(messages))
	}
	message := messages[0]
	content := ExtractMessageContent(message.MsgRequest())
	if content.Body != "Happy Birthday, dear John" {
		t.Errorf("Expected message body to be 'Happy Birthday, dear John', got %s", message.MsgRequest())
	}
	if content.Subject != "Happy Birthday" {
		t.Errorf("Expected subject to be 'Happy Birthday!', got %s", message.MsgRequest())
	}
	if content.Recipient != "john.doe@foobar.com" {
		t.Errorf("Expected recipient to be 'john.doe@foobar.com', got %s", message.MsgRequest())
	}
}

func TestWillNotSendEmailsWhenNobodysBirthday(t *testing.T) {
	// Setup mock SMTP server
	server := smtpmock.New(smtpmock.ConfigurationAttr{
		PortNumber:        NONSTANDARD_PORT,
		LogToStdout:       false,
		LogServerActivity: false,
	})

	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start SMTP server: %v", err)
	}
	defer server.Stop()

	// Create a new BirthdayService
	birthdayService := NewBirthdayService()

	// Create test data file
	createTestEmployeeDataFile(t)

	// Create a date for testing
	date, err := NewXDateFromString("2008/01/01")
	if err != nil {
		t.Fatalf("Failed to create date: %v", err)
	}

	// Send greetings
	err = birthdayService.SendGreetings("employee_data.txt", date, "localhost", NONSTANDARD_PORT)
	if err != nil {
		t.Fatalf("Failed to send greetings: %v", err)
	}

	// Wait a bit for any messages to be processed
	time.Sleep(100 * time.Millisecond)

	// Check that no messages were sent
	if len(server.Messages()) != 0 {
		t.Errorf("Expected 0 messages to be sent, got %d", len(server.Messages()))
	}
}

// Helper function to create the test employee data file
func createTestEmployeeDataFile(t *testing.T) {
	err := os.WriteFile("employee_data.txt", []byte(testData), 0644)
	if err != nil {
		t.Fatalf("Failed to create test employee data file: %v", err)
	}
}

// MessageContent represents the parsed content of an email message
type MessageContent struct {
	Subject   string
	Body      string
	Recipient string
}

// ExtractMessageContent parses a message content string and extracts the subject and body
func ExtractMessageContent(msgContent string) MessageContent {
	lines := strings.Split(msgContent, "\r\n")
	var subject, body, recipient string
	var isBody bool

	for _, line := range lines {
		if strings.HasPrefix(line, "Subject: ") {
			subject = strings.TrimPrefix(line, "Subject: ")
		} else if strings.HasPrefix(line, "To: ") {
			recipient = strings.TrimPrefix(line, "To: ")
		} else if line == "" {
			isBody = true
		} else if isBody {
			body = line
			break // We only want the first line of the body
		}
	}

	return MessageContent{
		Subject:   subject,
		Body:      body,
		Recipient: recipient,
	}
}
