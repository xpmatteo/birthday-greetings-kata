package it.xpug.kata.birthday_greetings;

// Replace JUnit 4 imports with JUnit 5
import static org.junit.jupiter.api.Assertions.*;
import org.junit.jupiter.api.*;

public class EmployeeTest {

	@Test
	public void testBirthday() throws Exception {
		Employee employee = new Employee("foo", "bar", "1990/01/31", "a@b.c");
		assertFalse(employee.isBirthday(new XDate("2008/01/30")), "not his birthday");
		assertTrue(employee.isBirthday(new XDate("2008/01/31")), "his birthday");
	}

	@Test
	public void equality() throws Exception {
		Employee base = new Employee("First", "Last", "1999/09/01", "first@last.com");
		Employee same = new Employee("First", "Last", "1999/09/01", "first@last.com");
		Employee different = new Employee("First", "Last", "1999/09/01", "boom@boom.com");

		assertFalse(base.equals(null));
		assertFalse(base.equals(""));
		assertTrue(base.equals(same));
		assertFalse(base.equals(different));
	}
}
