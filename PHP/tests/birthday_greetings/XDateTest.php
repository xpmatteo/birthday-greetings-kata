<?php

namespace birthday_greetings;

use PHPUnit\Framework\TestCase;

class XDateTest extends TestCase
{
    public function testGetters(): void
    {
        $date = new XDate("1789/01/14");
        $this->assertEquals(14, $date->getDay());
        $this->assertEquals(1, $date->getMonth());
    }

    public function testIsSameDayWhenIsTrue(): void
    {
        $date = new XDate("1789/01/14");
        $sameDay = new XDate("2001/01/14");
        $differentDay = new XDate("1789/01/15");
        $differentMonth = new XDate("1789/02/14");

        $this->assertTrue($date->isSameDay($sameDay));
        $this->assertFalse($date->isSameDay($differentDay));
        $this->assertFalse($date->isSameDay($differentMonth));
    }

    public function testIsSameDayWhenIsFalse(): void
    {
        $date = new XDate("1789/01/14");
        $differentDay = new XDate("1789/01/15");
        $differentMonth = new XDate("1789/02/14");

        $this->assertFalse($date->isSameDay($differentDay));
        $this->assertFalse($date->isSameDay($differentMonth));
    }
} 