package it.xpug.kata.birthday_greetings;

import java.io.*;
import java.text.ParseException;

import jakarta.mail.*;
import jakarta.mail.internet.*;

public class Main {

	public static void main(String[] args) throws AddressException, IOException, ParseException, MessagingException {
		BirthdayService service = new BirthdayService();
		service.sendGreetings("employee_data.txt", new XDate(), "localhost", 25);
	}

}
