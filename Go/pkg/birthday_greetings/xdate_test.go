package birthday_greetings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDay(t *testing.T) {
	date, err := NewXDateFromString("2008/10/08")
	assert.NoError(t, err)

	assert.Equal(t, 8, date.GetDay())
}

func TestGetMonth(t *testing.T) {
	date, err := NewXDateFromString("2008/10/08")
	assert.NoError(t, err)

	assert.Equal(t, 10, date.GetMonth())
}

func TestIsSameDay(t *testing.T) {
	date1, err := NewXDateFromString("2008/10/08")
	assert.NoError(t, err)

	date2, err := NewXDateFromString("2007/10/08")
	assert.NoError(t, err)

	assert.True(t, date1.IsSameDay(date2), "Expected dates to be the same day")
}

func TestIsNotSameDay_DifferentDay(t *testing.T) {
	date1, err := NewXDateFromString("2008/10/08")
	assert.NoError(t, err)

	date2, err := NewXDateFromString("2008/10/07")
	assert.NoError(t, err)

	assert.False(t, date1.IsSameDay(date2), "Expected dates to be different days")
}

func TestIsNotSameDay_DifferentMonth(t *testing.T) {
	date1, err := NewXDateFromString("2008/10/08")
	assert.NoError(t, err)

	date2, err := NewXDateFromString("2008/11/08")
	assert.NoError(t, err)

	assert.False(t, date1.IsSameDay(date2), "Expected dates to be different days due to different months")
}
