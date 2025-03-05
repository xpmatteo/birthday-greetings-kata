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

// Equals checks if this employee is equal to another employee
func (e *Employee) Equals(obj interface{}) bool {
	other, ok := obj.(*Employee)
	if !ok {
		return false
	}

	// Check birthDate
	if e.birthDate == nil {
		if other.birthDate != nil {
			return false
		}
	} else if !e.birthDate.Equals(other.birthDate) {
		return false
	}

	// Check email
	if e.email != other.email {
		return false
	}

	// Check firstName
	if e.firstName != other.firstName {
		return false
	}

	// Check lastName
	if e.lastName != other.lastName {
		return false
	}

	return true
}
