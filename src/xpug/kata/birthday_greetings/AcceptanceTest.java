package xpug.kata.birthday_greetings;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

import javax.mail.MessagingException;
import javax.mail.internet.AddressException;

import org.junit.Test;

import com.dumbster.smtp.SmtpMessage;

import static org.junit.Assert.*;


public class AcceptanceTest {

	private static final int NONSTANDARD_PORT = 3003;
	private Iterator<SmtpMessage> emailIterator;
	
	@Test
	public void baseScenario() throws Exception {
		startBirthdayServiceFor("employee_data.txt", "2008/10/08");
		
		expectEmailWithSubject_andBody_sentTo("Happy Birthday!", "Happy Birthday, dear John!", "john.doe@foobar.com");
	}

	private void expectEmailWithSubject_andBody_sentTo(String subject, String body, String recipient) {
		SmtpMessage message = emailIterator.next();
		assertNotNull("message not received", message);
		assertEquals(body, message.getBody());
		assertEquals(subject, message.getHeaderValue("Subject"));
		assertEquals(recipient, message.getHeaderValue("To"));		
	}

	private void startBirthdayServiceFor(String employeeFileName, String date) throws Exception {
		final List<SmtpMessage> messages = new ArrayList<SmtpMessage>();
		BirthdayService service = new BirthdayService() {
			@Override
			
			protected void sendMessage(String smtpHost, int smtpPort, String sender, final String subject, final String body, final String recipient) throws AddressException, MessagingException {
				SmtpMessage message = new SmtpMessage() {

					@Override
					public String getBody() {
						return body;
					}

					@Override
					public String getHeaderValue(String header) {
						if ("subject".equalsIgnoreCase(header)) {
							return subject;
						}
						if ("to".equalsIgnoreCase(header)) {
							return recipient;
						}
						return null;
					}
					
				};
				messages.add(message);
			}
		};
		service.sendGreetings(employeeFileName, new OurDate(date), "localhost", NONSTANDARD_PORT);
		emailIterator = messages.iterator();
	}
}
