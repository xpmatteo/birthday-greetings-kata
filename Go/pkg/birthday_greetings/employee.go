package birthday_greetings

import (
	"fmt"
)

// Employee represents an employee with birthday information
type Employee struct {
	birthDate *XDate
	lastName  string
	firstName string
	email     string
}

// NewEmployee creates a new Employee
func NewEmployee(firstName, lastName, birthDate, email string) (*Employee, error) {
	bd, err := NewXDateFromString(birthDate)
	if err != nil {
		return nil, fmt.Errorf("failed to create employee: %w", err)
	}

	return &Employee{
		firstName: firstName,
		lastName:  lastName,
		birthDate: bd,
		email:     email,
	}, nil
}

// IsBirthday checks if today is the employee's birthday
func (e *Employee) IsBirthday(today *XDate) bool {
	return today.IsSameDay(e.birthDate)
}

// GetEmail returns the employee's email
func (e *Employee) GetEmail() string {
	return e.email
}

// GetFirstName returns the employee's first name
func (e *Employee) GetFirstName() string {
	return e.firstName
}

// String returns a string representation of the employee
func (e *Employee) String() string {
	return fmt.Sprintf("Employee %s %s <%s> born %v", e.firstName, e.lastName, e.email, e.birthDate)
}
