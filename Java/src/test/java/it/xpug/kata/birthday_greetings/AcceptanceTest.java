package it.xpug.kata.birthday_greetings;

import static org.junit.jupiter.api.Assertions.*;
import org.junit.jupiter.api.*;

import com.icegreen.greenmail.util.GreenMail;
import com.icegreen.greenmail.util.ServerSetup;
import com.icegreen.greenmail.util.GreenMailUtil;
import jakarta.mail.internet.MimeMessage;

public class AcceptanceTest {

	private static final int NONSTANDARD_PORT = 9999;
	private BirthdayService birthdayService;
	private GreenMail mailServer;

	@BeforeEach
	public void setUp() throws Exception {
		mailServer = new GreenMail(new ServerSetup(NONSTANDARD_PORT, null, ServerSetup.PROTOCOL_SMTP));
		mailServer.start();
		birthdayService = new BirthdayService();
	}

	@AfterEach
	public void tearDown() throws Exception {
		mailServer.stop();
	}

	@Test
	public void willSendGreetings_whenItsSomebodysBirthday() throws Exception {
		birthdayService.sendGreetings("employee_data.txt", new XDate("2008/10/08"), "localhost", NONSTANDARD_PORT);

		assertEquals(1, mailServer.getReceivedMessages().length, "message not sent?");
		MimeMessage message = mailServer.getReceivedMessages()[0];
		assertEquals("Happy Birthday, dear John!", GreenMailUtil.getBody(message));
		assertEquals("Happy Birthday!", message.getSubject());
		assertEquals(1, message.getAllRecipients().length);
		assertEquals("john.doe@foobar.com", message.getAllRecipients()[0].toString());
	}

	@Test
	public void willNotSendEmailsWhenNobodysBirthday() throws Exception {
		birthdayService.sendGreetings("employee_data.txt", new XDate("2008/01/01"), "localhost", NONSTANDARD_PORT);

		assertEquals(0, mailServer.getReceivedMessages().length, "what? messages?");
	}
}
