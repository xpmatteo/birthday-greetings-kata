package birthday_greetings

import (
	"os"
	"strings"
	"testing"
	"time"

	smtpmock "github.com/mocktools/go-smtp-mock"
	"github.com/stretchr/testify/assert"
)

const NONSTANDARD_PORT = 9999

const testData = `last_name, first_name, date_of_birth, email
Doe, John, 1982/10/08, john.doe@foobar.com
Ann, Mary, 1975/03/11, mary.ann@foobar.com
`

func TestWillSendGreetings_WhenItsSomebodysBirthday(t *testing.T) {
	// Start mock SMTP server
	server := smtpmock.New(smtpmock.ConfigurationAttr{
		PortNumber:        NONSTANDARD_PORT,
		LogToStdout:       false,
		LogServerActivity: false,
	})
	err := server.Start()
	assert.NoError(t, err)
	defer server.Stop()

	// Arrange
	createTestEmployeeDataFile(t)
	birthdayService := NewBirthdayService()
	date, err := NewXDateFromString("2008/10/08")
	assert.NoError(t, err)

	// Act
	err = birthdayService.SendGreetings("employee_data.txt", date, "localhost", NONSTANDARD_PORT)
	assert.NoError(t, err)

	// Assert
	time.Sleep(100 * time.Millisecond)
	messages := server.Messages()
	assert.Len(t, messages, 1, "Expected exactly one message to be sent")

	message := messages[0]
	content := ExtractMessageContent(message.MsgRequest())
	assert.Equal(t, "Happy Birthday, dear John", content.Body)
	assert.Equal(t, "Happy Birthday", content.Subject)
	assert.Equal(t, "john.doe@foobar.com", content.Recipient)
}

func TestWillNotSendEmailsWhenNobodysBirthday(t *testing.T) {
	// Start mock SMTP server
	server := smtpmock.New(smtpmock.ConfigurationAttr{
		PortNumber:        NONSTANDARD_PORT,
		LogToStdout:       false,
		LogServerActivity: false,
	})
	err := server.Start()
	assert.NoError(t, err)
	defer server.Stop()

	// Arrange
	createTestEmployeeDataFile(t)
	birthdayService := NewBirthdayService()
	date, err := NewXDateFromString("2008/01/01")
	assert.NoError(t, err)

	// Act
	err = birthdayService.SendGreetings("employee_data.txt", date, "localhost", NONSTANDARD_PORT)
	assert.NoError(t, err)

	// Assert
	time.Sleep(100 * time.Millisecond)
	messages := server.Messages()
	assert.Empty(t, messages, "Expected no messages to be sent")
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
