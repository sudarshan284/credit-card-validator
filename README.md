Introduction 🪩

In this project, we will build a credit card validator using Go and the Luhn Algorithm.

The Luhn algorithm is a mathematical formula used to validate credit card numbers. We are going to write a function that implements it and use it to validate a card number.

To achieve this, we will build a very simple server in Go which processes POST requests, extract the JSON payload with the card number, and returns a JSON response if the number is valid or not.

Let's start coding 🚀

Project set up

Create a new directory for your project, for example, credit-card-validation.

Open a terminal and navigate to the project directory: cd path/to/credit-card-validation

Inside your project directory, create a Go file where we will implement the Luhn algorithm: touch luhn_algorithm.go

To test if it's working, use Postman or Curl to send a POST request to the server at route "/", adding a credit card number to the request body:

Set the headers: Content-Type: application/json

Send a POST request to http://localhost:8080/

Set the request body to { "number": "4003600000000014" }

If all goes well, you should receive this response {"valid": true}
