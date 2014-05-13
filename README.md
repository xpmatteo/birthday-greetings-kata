# What's this?

This is a simple refactoring exercise that is meant to teach something about dependency inversion and dependency injection.

The documentation is in  [this blog bost](http://matteo.vaccari.name/blog/archives/154) and in the presentation in the doc directory.

## Setup

To import the project in Eclipse: select

    File -> Import -> Existing Project into Workspace

To import the project in Idea: run

    mvn idea:idea

## How to start

Run all the tests.  One test will fail.  Change production code so that all tests pass.

## How to continue

Read the [blog post](http://matteo.vaccari.name/blog/archives/154) and/or the included presentation.

## Notes for Windows

On Windows, the tests with the fake smtp server do not work.  If you are on Windows, then comment or delete the AcceptanceTest.java file.  If you are on Linux, Mac or other Unix, then delete the AcceptanceForWindowsTest.java



