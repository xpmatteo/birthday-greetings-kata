# Birthday Greetings Kata - Go Version

This is a Go translation of the original Java Birthday Greetings Kata by Matteo Vaccari.

## What's this?

This is a simple refactoring exercise that is meant to teach something about dependency inversion and dependency injection.

The original documentation is in [this blog post](http://matteo.vaccari.name/blog/archives/154) and in the presentation in the doc directory.

## Setup

To run the tests:

```
go test ./...
```

To run the application:

```
go run cmd/birthday_greetings/main.go
```

## How to start

Run all the tests. One test will fail. Change production code so that all tests pass.

## How to continue

Read the [blog post](http://matteo.vaccari.name/blog/archives/154) and/or the included presentation.

## Notes

This is a direct translation of the Java code to Go, preserving the original structure and code smells. The code is intentionally not idiomatic Go to maintain the educational value of the refactoring exercise.

The SMTP testing uses the `github.com/mocktools/go-smtp-mock` package instead of the Dumbster SMTP server used in the Java version.



