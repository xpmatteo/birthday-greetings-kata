package it.xpug.kata.birthday_greetings;

import org.junit.jupiter.api.*;

// Import AssertJ for fluent assertions
import static org.assertj.core.api.Assertions.*;

public class XDateTest {
	@Test
	public void getters() throws Exception {
		XDate date = new XDate("1789/01/24");
		
		assertThat(date.getMonth()).isEqualTo(1);
		assertThat(date.getDay()).isEqualTo(24);
	}

	@Test
	public void isSameDate() throws Exception {
		XDate date = new XDate("1789/01/24");
		XDate sameDay = new XDate("2001/01/24");
		XDate notSameDay = new XDate("1789/01/25");
		XDate notSameMonth = new XDate("1789/02/25");

		assertThat(date.isSameDay(sameDay))
			.as("same")
			.isTrue();
		assertThat(date.isSameDay(notSameDay))
			.as("not same day")
			.isFalse();
		assertThat(date.isSameDay(notSameMonth))
			.as("not same month")
			.isFalse();
	}

	// No need for equality test - records handle it automatically!
}
