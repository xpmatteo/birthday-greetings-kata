<?php

namespace birthday_greetings;

class Employee
{
    private $birthDate;
    private $lastName;
    private $firstName;
    private $email;

    public function __construct(string $firstName, string $lastName, string $birthDate, string $email)
    {
        $this->firstName = $firstName;
        $this->lastName = $lastName;
        $this->birthDate = new XDate($birthDate);
        $this->email = $email;
    }

    public function isBirthday(XDate $today): bool
    {
        return $today->isSameDay($this->birthDate);
    }

    public function getEmail(): string
    {
        return $this->email;
    }

    public function getFirstName(): string
    {
        return $this->firstName;
    }

    public function __toString(): string
    {
        return "Employee {$this->firstName} {$this->lastName} <{$this->email}> born {$this->birthDate}";
    }

    public function equals($obj): bool
    {
        if ($this === $obj) {
            return true;
        }
        if ($obj === null) {
            return false;
        }
        if (!($obj instanceof Employee)) {
            return false;
        }
        
        $other = $obj;
        if ($this->birthDate === null) {
            if ($other->birthDate !== null) {
                return false;
            }
        } elseif (!$this->birthDate->equals($other->birthDate)) {
            return false;
        }
        
        if ($this->email === null) {
            if ($other->email !== null) {
                return false;
            }
        } elseif ($this->email !== $other->email) {
            return false;
        }
        
        if ($this->firstName === null) {
            if ($other->firstName !== null) {
                return false;
            }
        } elseif ($this->firstName !== $other->firstName) {
            return false;
        }
        
        if ($this->lastName === null) {
            if ($other->lastName !== null) {
                return false;
            }
        } elseif ($this->lastName !== $other->lastName) {
            return false;
        }
        
        return true;
    }

    public function hashCode(): int
    {
        $prime = 31;
        $result = 1;
        $result = $prime * $result + ($this->birthDate === null ? 0 : $this->birthDate->hashCode());
        $result = $prime * $result + ($this->email === null ? 0 : crc32($this->email));
        $result = $prime * $result + ($this->firstName === null ? 0 : crc32($this->firstName));
        $result = $prime * $result + ($this->lastName === null ? 0 : crc32($this->lastName));
        return $result;
    }
} 