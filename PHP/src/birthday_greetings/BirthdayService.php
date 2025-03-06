<?php

namespace birthday_greetings;

use PHPMailer\PHPMailer\PHPMailer;
use PHPMailer\PHPMailer\Exception;

class BirthdayService
{
    public function sendGreetings(string $fileName, XDate $xDate, string $smtpHost, int $smtpPort): void
    {
        $handle = fopen($fileName, 'r');
        if ($handle === false) {
            throw new \RuntimeException("Could not open file: $fileName");
        }

        // Skip header
        fgetcsv($handle, 0, ',', '"', '\\');

        while (($data = fgetcsv($handle, 0, ',', '"', '\\')) !== false) {
            if (count($data) < 4) {
                continue; // Skip invalid lines
            }

            try {
                $employee = new Employee(trim($data[1]), trim($data[0]), trim($data[2]), trim($data[3]));
                if ($employee->isBirthday($xDate)) {
                    $recipient = $employee->getEmail();
                    $body = str_replace('%NAME%', $employee->getFirstName(), 'Happy Birthday, dear %NAME%');
                    $subject = 'Happy Birthday!';
                    $this->sendMessage($smtpHost, $smtpPort, 'sender@here.com', $subject, $body, $recipient);
                }
            } catch (\Exception $e) {
                // Log error and continue with next employee
                error_log("Error processing employee data: " . $e->getMessage());
                continue;
            }
        }

        fclose($handle);
    }

    private function sendMessage(string $smtpHost, int $smtpPort, string $sender, string $subject, string $body, string $recipient): void
    {
        $mail = new PHPMailer(true);

        try {
            // Server settings
            $mail->isSMTP();
            $mail->Host = $smtpHost;
            $mail->Port = $smtpPort;
            $mail->SMTPAuth = false;

            // Recipients
            $mail->setFrom($sender);
            $mail->addAddress($recipient);

            // Content
            $mail->isHTML(false);
            $mail->Subject = $subject;
            $mail->Body = $body;

            $mail->send();
        } catch (Exception $e) {
            throw new \RuntimeException("Message could not be sent. Mailer Error: {$mail->ErrorInfo}", 0, $e);
        }
    }
}
