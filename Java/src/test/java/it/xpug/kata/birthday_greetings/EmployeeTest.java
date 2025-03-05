package it.xpug.kata.birthday_greetings;
import static org.junit.Assert.*;

import org.junit.*;



public class EmployeeTest {

	@Test
	public void testBirthday() throws Exception {
		Employee employee = new Employee("foo", "bar", "1990/01/31", "a@b.c");
		assertFalse("not his birthday", employee.isBirthday(new XDate("2008/01/30")));
		assertTrue("his birthday", employee.isBirthday(new XDate("2008/01/31")));
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
