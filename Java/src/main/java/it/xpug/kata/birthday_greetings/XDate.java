package it.xpug.kata.birthday_greetings;

import java.text.ParseException;
import java.time.LocalDate;
import java.time.format.DateTimeFormatter;

// Modern record using LocalDate instead of java.util.Date
public record XDate(LocalDate date) {

	private static final DateTimeFormatter FORMATTER = DateTimeFormatter.ofPattern("yyyy/MM/dd");

	public XDate() {
		this(LocalDate.now());
	}

	public XDate(String yyyyMMdd) throws ParseException {
		this(LocalDate.parse(yyyyMMdd, FORMATTER));
	}

	public int getDay() {
		return date.getDayOfMonth();
	}

	public int getMonth() {
		return date.getMonthValue();
	}

	public boolean isSameDay(XDate anotherDate) {
		return anotherDate.getDay() == this.getDay() && anotherDate.getMonth() == this.getMonth();
	}
}
