package it.xpug.kata.birthday_greetings;

import java.text.ParseException;

public class Employee {

	private XDate birthDate;
	private String lastName;
	private String firstName;
	private String email;

	public Employee(String firstName, String lastName, String birthDate, String email) throws ParseException {
		this.firstName = firstName;
		this.lastName = lastName;
		this.birthDate = new XDate(birthDate);
		this.email = email;
	}

	public boolean isBirthday(XDate today) {
		return today.isSameDay(birthDate);
	}

	public String getEmail() {
		return email;
	}

	public String getFirstName() {
		return firstName;
	}

	@Override
	public String toString() {
		return "Employee " + firstName + " " + lastName + " <" + email + "> born " + birthDate;
	}
}
