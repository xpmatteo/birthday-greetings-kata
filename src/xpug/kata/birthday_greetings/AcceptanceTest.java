package xpug.kata.birthday_greetings;

import java.util.ArrayList;
import java.util.List;

import javax.mail.Message;
import javax.mail.MessagingException;

import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.*;


public class AcceptanceTest {

	private static final int NONSTANDARD_PORT = 3003;
	private List<Message> messagesSent;
	
	@Before
	public void setUp() throws Exception {
		messagesSent = new ArrayList<Message>();
	}
	
	@Test
	public void baseScenario() throws Exception {
		startBirthdayServiceFor("employee_data.txt", "2008/10/08");
		
		assertEquals("message not sent?", 1, messagesSent.size());
		Message message = messagesSent.get(0);
		assertEquals("Happy Birthday, dear John!", message.getContent());
		assertEquals("Happy Birthday!", message.getSubject());
		assertEquals(1, message.getAllRecipients().length);		
		assertEquals("john.doe@foobar.com", message.getAllRecipients()[0].toString());		
	}

	private void startBirthdayServiceFor(String employeeFileName, String date) throws Exception {
		BirthdayService service = new BirthdayService() {			
			@Override
			protected void sendMessage(Message msg) throws MessagingException {
				messagesSent.add(msg);
			}
		};
		service.sendGreetings(employeeFileName, new OurDate(date), "localhost", NONSTANDARD_PORT);
	}
}
