# Lab 01 - Testing

In this lab we will get more hands on practice writing unit tests, and learn how to test HTTP handlers.

## HTTP Handler Testing
Learning how to test your http handlers is dead simple, and will save you a lot of time and problems as the assignments get harder. It will also make you a better developer :)

Go into the `test_handlers` folder, there is a simple program that listens on `localhost:4242` and returns a tailored random compliment for any user. By passing a parameter to `/compliment?name=<name>`, it will respond with something like: `Brendan, you are bodacious`.

We want to test the ComplimentHandler to make sure that it is behaving as it should. 


Links: [Source](https://elithrar.github.io/article/testing-http-handlers-go/ "Testing Your (HTTP) Handlers in Go "), 

## Unit Testing Practice
Now that you've been properly complimented, let's get to more practical matters. You recently started working for the FCC on a team that works on a `swears` program in go. It automatically detects swear words in text files.

This project recently hit a snag when their lead developer quit to go work for the CIA. The code is full of bugs, and he didn't write any unit tests! Your job is to write unit tests for the `swears.go` file to make sure that it is working properly, and handling edge cases!

In `/swears`, is a simple program that reads in two text files: a `known_swears.txt` file that has every single swearword ever in it, _totally_, and a file passed as a command line argument, that is then scanned for any swear words. The program then reports the tally of each swear word that it finds in the text file.

The previous developer was tired when he wrote this code, so there are _definitely_ no bugs in it.

Your job is to finish writing the unit tests for the four function stubs in `swears_test.go`.
Try and think of weird edge cases as you write these tests! How are the files being read in? What swear words is it checking for, etc?