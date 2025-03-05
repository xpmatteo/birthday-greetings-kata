package birthday_greetings

import (
	"os"
	"strings"
	"testing"
	"time"

	smtpmock "github.com/mocktools/go-smtp-mock"
)

const NONSTANDARD_PORT = 9999

func TestWillSendGreetings_WhenItsSomebodysBirthday(t *testing.T) {
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
	date, err := NewXDateFromString("2008/10/08")
	if err != nil {
		t.Fatalf("Failed to create date: %v", err)
	}

	// Send greetings
	err = birthdayService.SendGreetings("employee_data.txt", date, "localhost", NONSTANDARD_PORT)
	if err != nil {
		t.Fatalf("Failed to send greetings: %v", err)
	}

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

	// Check the message body
	msgContent := message.MsgRequest()
	if !strings.Contains(msgContent, "Happy Birthday, dear John") {
		t.Errorf("Expected message to contain 'Happy Birthday, dear John', got %s", msgContent)
	}

	// Check the subject
	if !strings.Contains(msgContent, "Subject: Happy Birthday!") {
		t.Errorf("Expected subject to be 'Happy Birthday!', got %s", msgContent)
	}

	// Check the recipient
	if !strings.Contains(msgContent, "To: john.doe@foobar.com") {
		t.Errorf("Expected recipient to be 'john.doe@foobar.com', got %s", msgContent)
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
	content := `last_name, first_name, date_of_birth, email
Doe, John, 1982/10/08, john.doe@foobar.com
Ann, Mary, 1975/03/11, mary.ann@foobar.com
`
	err := os.WriteFile("employee_data.txt", []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test employee data file: %v", err)
	}
}
