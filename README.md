## $5 Tech Unlocked 2021!
[Buy and download this Book for only $5 on PacktPub.com](https://www.packtpub.com/product/hands-on-dependency-injection-in-go/9781789132762)
-----
*If you have read this book, please leave a review on [Amazon.com](https://www.amazon.com/gp/product/1789132762).     Potential readers can then use your unbiased opinion to help them make purchase decisions. Thank you. The $5 campaign         runs from __December 15th 2020__ to __January 13th 2021.__*

# Hands-On Dependency Injection in Go

<a href="https://www.packtpub.com/application-development/hands-dependency-injection-go?utm_source=github&utm_medium=repository&utm_campaign=9781789132762 "><img src="https://d255esdrn735hr.cloudfront.net/sites/default/files/imagecache/ppv4_main_book_cover/B10763_MockupCover_new.png" alt="Hands-On Dependency Injection in Go" height="256px" align="right"></a>

This is the code repository for [Hands-On Dependency Injection in Go](https://www.packtpub.com/application-development/hands-dependency-injection-go?utm_source=github&utm_medium=repository&utm_campaign=9781789132762 ), published by Packt.

**Develop clean Go code that is easier to read, maintain, and test**

## What is this book about?
Hands-On Dependency Injection in Go takes you on a journey, refactoring existing code to adopt dependency injection (DI) using various methods available in Go.

This book covers the following exciting features:
* Understand the benefits of dependency injection 
* Explore SOLID design principles and how they relate to Go 
* Analyze various dependency injection patterns available in Go 
* Leverage DI to produce high quality, loosely coupled Go code 
* Refactor existing Go code to adopt dependency injection 
* Discover tools to improve your code's testability and test coverage 
* Generate and interpret Go dependency graphs 

If you feel this book is for you, get your [copy](https://www.amazon.com/dp/1789132762) today!

<a href="https://www.packtpub.com/?utm_source=github&utm_medium=banner&utm_campaign=GitHubBanner"><img src="https://raw.githubusercontent.com/PacktPublishing/GitHub/master/GitHub.png" 
alt="https://www.packtpub.com/" border="5" /></a>

## Instructions and Navigations
All of the code is organized into folders. For example, ch02.

The code will look like the following:
```
html, body, #map {
 height: 100%; 
 margin: 0;
 padding: 0
}
```

**Following is what you need for this book:**
Hands-On Dependency Injection in Go is for programmers with a few year s experience in any language and a basic understanding of Go. If you wish to produce clean, loosely coupled code that is inherently easier to test, this book is for you.

## Getting the source

The easiest way to obtain the source code is to use `go get`.  
This will ensure that the code is placed in the correct directory and should be then runnable and testable.

To download this repo use `go get github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/...`

## Code Organization

In this repository, there is 1 folder for every chapter of the book, named chXX where XX is the chapter number.

The code provided are expanded versions of the code presented in the book.  While it will compile and typically
will not throw an error when passed into `go test` it is not designed to be executed.

From chapter 4 onwards, there is an `acme` directory included with the code that chapter.
The `acme` directory is the code for the sample service presented in the book with the changes discussed in that chapter
already applied.

You will also find 2 additional directories in the root of the repository:
 
 * **resources** - this directory contains an SQL file that should be used to populate a MySQL database.  This database
 is used by the sample service 
 * **vendor** - this is standard go vendor directory which contains the external packages required by the sample service

## Setting up the MySQL database

The easiest way to create and populate the database required by the sample service is by running the following:

`mysql < ./resources/create.sql`

Depending on your settings you may want to provide a username and password like this:

`mysql -u [your username] -p < ./resources/create.sql`

This will create a database called `acme` with 1 table and 4 records.

## Creating a free account on CurrencyLayer

The sample service uses a free currency conversion service.  In order to successfully run all the examples, you will need
to sign up [here](https://currencylayer.com/) and obtain an API Key.

## Configuring the sample service

Now that you have your MySQL and CurrencyLayer credentials you can create a config for the sample service.

1. Copy `default-config.json` (found next to this file) to `config.json`
1. Open `config.json` in your favorite editor
1. Add your database credentials to the `"dsn"` setting.  Should be in the form: 
`"[username]:[password]@tcp(localhost:3306)/[database name]?autocommit=true"`
1. Add your API Key to the `"exchangeRateAPIKey"` setting.  Should be in the form: `"1234567890abcdef1234567890abcdef"`

## Running the sample service for a particular chapter

To run sample service for a particular chapter:
 
1. First make sure you are in the base of this repository:
`cd $GOPATH/src/github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/` 
1. Use a command similar to the following (which is for ch04):
`ACME_CONFIG=$GOPATH/src/github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/config.json go run ./ch04/acme/main.go` 

### Special instructions for chapters 10-12

As we have multiple files and tests in the `main` package, we cannot use the standard `go run ./ch10/acme/main.go` to run the service.

Instead we need to modify the command to `go run ./ch10/acme/main.go ./ch10/acme/wire_gen.go`

## Running tests for a chapter

To run sample service for a particular chapter you can use a command similar to the follow (which is for ch04):

1. First make sure you are in the base of this repository:
`cd $GOPATH/src/github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/` 
1. Use a command similar to the following (which is for ch04):
`ACME_CONFIG=$GOPATH/src/github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/config.json go test ./ch04/...` 

With the following software and hardware list you can run all code files present in the book (Chapter 1-12).
### Software and Hardware List
| Chapter  | Software required                    | OS required                        |
| -------- | ------------------------------------ | -----------------------------------|
| 1-12     | Go 1.10.x+                           | Windows, Mac OS X, and Linux (Any) |
| 4-12     | MySQL 5.7.x+                         | Windows, Mac OS X, and Linux (Any) |
| 4-12     | CurrencyLayer                        | Windows, Mac OS X, and Linux (Any) |

### Related products
* Mastering Go [[Packt]](https://www.packtpub.com/networking-and-servers/mastering-go?utm_source=github&utm_medium=repository&utm_campaign=9781788626545 ) [[Amazon]](https://www.amazon.com/dp/1788626540)

* Go Standard Library Cookbook [[Packt]](https://www.packtpub.com/application-development/go-standard-library-cookbook?utm_source=github&utm_medium=repository&utm_campaign=9781788475273 ) [[Amazon]](https://www.amazon.com/dp/1788475275)


## Get to Know the Author
**Corey Scott**
is a senior software engineer currently living in Melbourne, Australia. He’s been programming professionally since 2000, with the last 5 years spent building large-scale distributed services in Go.
An occasional technical speaker and blogger on a variety of software-related topics, he is passionate about designing and building quality software. He believes that software engineering is a craft that should be honed, debated, and continuously improved. He takes a pragmatic, non-zealot approach to coding and is always up for a good debate about software engineering, continuous delivery, testing, or clean coding.


## Other books by the authors


### Suggestions and Feedback
[Click here](https://docs.google.com/forms/d/e/1FAIpQLSdy7dATC6QmEL81FIUuymZ0Wy9vH1jHkvpY57OiMeKGqib_Ow/viewform) if you have any feedback or suggestions.


