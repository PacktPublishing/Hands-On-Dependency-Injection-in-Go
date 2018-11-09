# Hands-On-Dependency-Injection-in-Go
Hands-On Dependency Injection in Go, published by Packt

## Requirements

* Go - [link](https://golang.org/dl/)
* MySQL - [link](https://dev.mysql.com/doc/refman/5.7/en/installing.html)
* Free account from Currency Layer - [link](https://currencylayer.com/)

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

## Running tests for the sample service

To run sample service for a particular chapter you can use a command similar to the follow (which is for ch04):

1. First make sure you are in the base of this repository:
`cd $GOPATH/src/github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/` 
1. Use a command similar to the following (which is for ch04):
`ACME_CONFIG=$GOPATH/src/github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/config.json go test ./ch04/acme/...` 
