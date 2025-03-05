<?php

namespace birthday_greetings;

use PHPUnit\Framework\TestCase;

class AcceptanceTest extends TestCase
{
    private const SMTP_PORT = 1025; // MailDev's default SMTP port
    private const API_URL = 'http://localhost:1080';
    private BirthdayService $birthdayService;

    protected function setUp(): void
    {
        $this->birthdayService = new BirthdayService();
        $this->deleteAllEmails();
    }

    private function deleteAllEmails(): void
    {
        file_get_contents(self::API_URL . '/email/all', false, stream_context_create([
            'http' => ['method' => 'DELETE']
        ]));
    }

    private function getEmails(): array
    {
        $response = file_get_contents(self::API_URL . '/email');
        return json_decode($response, true) ?? [];
    }

    public function testWillSendGreetingsWhenItsSomebodysBirthday(): void
    {
        $this->birthdayService->sendGreetings(
            'employee_data.txt',
            new XDate('2008/10/08'),
            'localhost',
            self::SMTP_PORT
        );

        // Wait for email to be processed
        sleep(1);

        $emails = $this->getEmails();
        $this->assertCount(1, $emails, 'Expected exactly one email to be sent');

        $email = $emails[0];
        $this->assertEquals('Happy Birthday!', $email['subject']);
        $this->assertEquals('Happy Birthday, dear John!', $email['text']);
        $this->assertCount(1, $email['to']);
        $this->assertEquals('john.doe@foobar.com', $email['to'][0]['address']);
    }

    public function testWillNotSendEmailsWhenNobodysBirthday(): void
    {
        $this->birthdayService->sendGreetings(
            'employee_data.txt',
            new XDate('2008/01/01'),
            'localhost',
            self::SMTP_PORT
        );

        // Wait for any potential emails to be processed
        sleep(1);

        $emails = $this->getEmails();
        $this->assertCount(0, $emails, 'Expected no emails to be sent');
    }
} 
