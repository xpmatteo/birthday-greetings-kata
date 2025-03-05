package birthday_greetings

import (
	"testing"
)

func TestEmployeeBirthday(t *testing.T) {
	employee, err := NewEmployee("John", "Doe", "1982/10/08", "john.doe@foobar.com")
	if err != nil {
		t.Fatalf("Failed to create employee: %v", err)
	}

	date, err := NewXDateFromString("2008/10/08")
	if err != nil {
		t.Fatalf("Failed to create date: %v", err)
	}

	if !employee.IsBirthday(date) {
		t.Errorf("Expected it to be the employee's birthday")
	}
}

func TestEmployeeIsNotBirthday_DifferentDay(t *testing.T) {
	employee, err := NewEmployee("John", "Doe", "1982/10/08", "john.doe@foobar.com")
	if err != nil {
		t.Fatalf("Failed to create employee: %v", err)
	}

	date, err := NewXDateFromString("2008/10/07")
	if err != nil {
		t.Fatalf("Failed to create date: %v", err)
	}

	if employee.IsBirthday(date) {
		t.Errorf("Expected it not to be the employee's birthday")
	}
}

func TestEmployeeIsNotBirthday_DifferentMonth(t *testing.T) {
	employee, err := NewEmployee("John", "Doe", "1982/10/08", "john.doe@foobar.com")
	if err != nil {
		t.Fatalf("Failed to create employee: %v", err)
	}

	date, err := NewXDateFromString("2008/11/08")
	if err != nil {
		t.Fatalf("Failed to create date: %v", err)
	}

	if employee.IsBirthday(date) {
		t.Errorf("Expected it not to be the employee's birthday")
	}
}
