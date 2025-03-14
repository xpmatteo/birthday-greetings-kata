package it.xpug.kata.birthday_greetings;

import jakarta.mail.MessagingException;
import jakarta.mail.internet.MimeMessage;
import java.io.IOException;

public class GreenMailUtil {
    public static String getBody(MimeMessage message) throws MessagingException, IOException {
        return message.getContent().toString().trim();
    }
} 