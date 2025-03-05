<?php

namespace birthday_greetings;

use PHPUnit\Framework\TestCase;

class EmployeeTest extends TestCase
{
    public function testBirthday(): void
    {
        $employee = new Employee("foo", "bar", "1990/01/31", "a@b.c");
        $this->assertTrue($employee->isBirthday(new XDate("2008/01/31")));
        $this->assertFalse($employee->isBirthday(new XDate("2008/01/30")));
        $this->assertFalse($employee->isBirthday(new XDate("2008/02/01")));
    }
} 