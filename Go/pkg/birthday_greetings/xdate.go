package birthday_greetings

import (
	"fmt"
	"time"
)

// XDate is a wrapper around Go's time.Time for date handling
type XDate struct {
	date time.Time
}

// NewXDate creates a new XDate with the current date
func NewXDate() *XDate {
	return &XDate{
		date: time.Now(),
	}
}

// NewXDateFromString creates a new XDate from a string in the format "yyyy/MM/dd"
func NewXDateFromString(yyyyMMdd string) (*XDate, error) {
	date, err := time.Parse("2006/01/02", yyyyMMdd)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %w", err)
	}
	return &XDate{
		date: date,
	}, nil
}

// GetDay returns the day of the month
func (x *XDate) GetDay() int {
	return x.date.Day()
}

// GetMonth returns the month (1-12)
func (x *XDate) GetMonth() int {
	return int(x.date.Month())
}

// IsSameDay checks if this date has the same day and month as another date
func (x *XDate) IsSameDay(anotherDate *XDate) bool {
	return anotherDate.GetDay() == x.GetDay() && anotherDate.GetMonth() == x.GetMonth()
}

// Equals checks if this date is equal to another date
func (x *XDate) Equals(obj interface{}) bool {
	other, ok := obj.(*XDate)
	if !ok {
		return false
	}
	return other.date.Equal(x.date)
}
