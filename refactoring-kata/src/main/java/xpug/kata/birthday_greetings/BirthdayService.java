package xpug.kata.birthday_greetings;

import static java.util.Arrays.asList;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.text.ParseException;

import javax.mail.Message;
import javax.mail.MessagingException;
import javax.mail.Session;
import javax.mail.Transport;
import javax.mail.internet.AddressException;
import javax.mail.internet.InternetAddress;
import javax.mail.internet.MimeMessage;

public class BirthdayService {

	public void sendGreetings(String fileName, OurDate ourDate, String smtpHost, int smtpPort) throws IOException, ParseException, AddressException, MessagingException {
		BufferedReader in = new BufferedReader(new FileReader(fileName));
		String str = "";
		str = in.readLine(); // skip header
		while ((str = in.readLine()) != null) {
			String[] employeeData = str.split(", ");
			Employee employee = new Employee(employeeData[1], employeeData[0], employeeData[2], employeeData[3]);
			if (employee.isBirthday(ourDate)) {
				String recipient = employee.getEmail();
				String body = "Happy Birthday, dear %NAME%!".replace("%NAME%", employee.getFirstName());
				String subject = "Happy Birthday!";
				sendMessage(smtpHost, smtpPort, "sender@here.com", subject, body, recipient);
			}
		}
	}

	private void sendMessage(String smtpHost, int smtpPort, String sender, String subject, String body, String recipient) throws AddressException, MessagingException {
		// Create a mail session
		java.util.Properties props = new java.util.Properties();
		props.put("mail.smtp.host", smtpHost);
		props.put("mail.smtp.port", "" + smtpPort);
		Session session = Session.getDefaultInstance(props, null);

		// Construct the message
		Message msg = new MimeMessage(session);
		msg.setFrom(new InternetAddress(sender));
		msg.setRecipient(Message.RecipientType.TO, new InternetAddress(recipient));
		msg.setSubject(subject);
		msg.setText(body);

		// Send the message
		sendMessage(msg);
	}

	// made protected for testing :-(
	protected void sendMessage(Message msg) throws MessagingException {
		Transport.send(msg);
	}

	public static void main(String[] args) {
		BirthdayService service = new BirthdayService();
		try {
			service.sendGreetings("employee_data.txt", new OurDate("2008/10/08"), "localhost", 25);
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}
