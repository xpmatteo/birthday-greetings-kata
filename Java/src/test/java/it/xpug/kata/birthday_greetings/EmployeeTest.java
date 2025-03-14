package it.xpug.kata.birthday_greetings;

// Replace JUnit 4 imports with JUnit 5
import static org.junit.jupiter.api.Assertions.*;
import org.junit.jupiter.api.*;

// Import AssertJ for fluent assertions
import static org.assertj.core.api.Assertions.*;

public class EmployeeTest {

	@Test
	public void testBirthday() throws Exception {
		Employee employee = new Employee("foo", "bar", "1990/01/31", "a@b.c");
		
		assertThat(employee.isBirthday(new XDate("2008/01/30")))
			.as("not his birthday")
			.isFalse();
		assertThat(employee.isBirthday(new XDate("2008/01/31")))
			.as("his birthday")
			.isTrue();
	}

	@Test
	public void equality() throws Exception {
		Employee base = new Employee("First", "Last", "1999/09/01", "first@last.com");
		Employee same = new Employee("First", "Last", "1999/09/01", "first@last.com");
		Employee different = new Employee("First", "Last", "1999/09/01", "boom@boom.com");

		assertThat(base)
			.isNotEqualTo(null)
			.isNotEqualTo("")
			.isEqualTo(same)
			.isNotEqualTo(different);
	}
}
