package birthday_greetings

import (
	"testing"
)

func TestGetDay(t *testing.T) {
	date, err := NewXDateFromString("2008/10/08")
	if err != nil {
		t.Fatalf("Failed to create date: %v", err)
	}

	if date.GetDay() != 8 {
		t.Errorf("Expected day to be 8, got %d", date.GetDay())
	}
}

func TestGetMonth(t *testing.T) {
	date, err := NewXDateFromString("2008/10/08")
	if err != nil {
		t.Fatalf("Failed to create date: %v", err)
	}

	if date.GetMonth() != 10 {
		t.Errorf("Expected month to be 10, got %d", date.GetMonth())
	}
}

func TestIsSameDay(t *testing.T) {
	date1, err := NewXDateFromString("2008/10/08")
	if err != nil {
		t.Fatalf("Failed to create date1: %v", err)
	}

	date2, err := NewXDateFromString("2007/10/08")
	if err != nil {
		t.Fatalf("Failed to create date2: %v", err)
	}

	if !date1.IsSameDay(date2) {
		t.Errorf("Expected dates to be the same day")
	}
}

func TestIsNotSameDay_DifferentDay(t *testing.T) {
	date1, err := NewXDateFromString("2008/10/08")
	if err != nil {
		t.Fatalf("Failed to create date1: %v", err)
	}

	date2, err := NewXDateFromString("2008/10/07")
	if err != nil {
		t.Fatalf("Failed to create date2: %v", err)
	}

	if date1.IsSameDay(date2) {
		t.Errorf("Expected dates to be different days")
	}
}

func TestIsNotSameDay_DifferentMonth(t *testing.T) {
	date1, err := NewXDateFromString("2008/10/08")
	if err != nil {
		t.Fatalf("Failed to create date1: %v", err)
	}

	date2, err := NewXDateFromString("2008/11/08")
	if err != nil {
		t.Fatalf("Failed to create date2: %v", err)
	}

	if date1.IsSameDay(date2) {
		t.Errorf("Expected dates to be different days due to different months")
	}
}
