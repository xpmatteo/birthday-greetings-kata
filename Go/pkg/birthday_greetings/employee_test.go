package birthday_greetings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeBirthday(t *testing.T) {
	employee, err := NewEmployee("John", "Doe", "1982/10/08", "john.doe@foobar.com")
	assert.NoError(t, err)

	date, err := NewXDateFromString("2008/10/08")
	assert.NoError(t, err)

	assert.True(t, employee.IsBirthday(date), "Expected it to be the employee's birthday")
}

func TestEmployeeIsNotBirthday_DifferentDay(t *testing.T) {
	employee, err := NewEmployee("John", "Doe", "1982/10/08", "john.doe@foobar.com")
	assert.NoError(t, err)

	date, err := NewXDateFromString("2008/10/07")
	assert.NoError(t, err)

	assert.False(t, employee.IsBirthday(date), "Expected it not to be the employee's birthday")
}

func TestEmployeeIsNotBirthday_DifferentMonth(t *testing.T) {
	employee, err := NewEmployee("John", "Doe", "1982/10/08", "john.doe@foobar.com")
	assert.NoError(t, err)

	date, err := NewXDateFromString("2008/11/08")
	assert.NoError(t, err)

	assert.False(t, employee.IsBirthday(date), "Expected it not to be the employee's birthday")
}
