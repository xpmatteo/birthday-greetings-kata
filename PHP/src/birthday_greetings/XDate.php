<?php

namespace birthday_greetings;


class XDate
{
    private $date;

    public function __construct(?string $yyyyMMdd = null)
    {
        if ($yyyyMMdd === null) {
            $this->date = new \DateTime();
        } else {
            $this->date = \DateTime::createFromFormat('Y/m/d', $yyyyMMdd);
            if ($this->date === false) {
                throw new \InvalidArgumentException("Invalid date format: $yyyyMMdd. Expected format: YYYY/MM/DD");
            }
        }
    }

    public function getDay(): int
    {
        return (int)$this->date->format('d');
    }

    public function getMonth(): int
    {
        return (int)$this->date->format('m');
    }

    public function isSameDay(XDate $anotherDate): bool
    {
        return $anotherDate->getDay() === $this->getDay() && 
               $anotherDate->getMonth() === $this->getMonth();
    }

    public function __toString(): string
    {
        return $this->date->format('Y-m-d');
    }

    public function equals($obj): bool
    {
        if (!($obj instanceof XDate)) {
            return false;
        }
        return $obj->date == $this->date;
    }

    public function hashCode(): int
    {
        return $this->date->getTimestamp();
    }
} 