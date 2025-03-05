# Birthday Greetings Kata (PHP Version)

This is a PHP translation of the Birthday Greetings Kata, which is a refactoring exercise focused on teaching dependency inversion and dependency injection principles.

The original Java version and its documentation can be found in the [original repository](http://matteo.vaccari.name/blog/archives/154).

## Setup

1. Make sure you have PHP 7.4 or higher installed
2. Install Composer if you haven't already
3. Run `composer install` to install dependencies

## Running the Tests

1. Start MailDev in a separate terminal:
   ```bash
   docker run -p 1080:1080 -p 1025:1025 maildev/maildev
   ```
   This will start a test SMTP server on port 1025 and a web interface on port 1080.

2. Run the tests using PHPUnit:
   ```bash
   ./vendor/bin/phpunit
   ```

3. You can view received emails in your browser at http://localhost:1080

## Project Structure

- `src/birthday_greetings/` - Main application code
- `tests/` - Test code
- `employee_data.txt` - Sample employee data file

## The Exercise

The goal is to refactor the code to:
- Apply dependency inversion principle
- Implement dependency injection
- Separate concerns
- Make the code more testable and maintainable

## Notes

- The tests use MailDev as a mock SMTP server, which provides a web interface to inspect received emails
- The code intentionally contains code smells that should be addressed during the refactoring exercise



