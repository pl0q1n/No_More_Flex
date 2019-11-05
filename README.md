# No More Flex

[![Build Status](https://travis-ci.com/pl0q1n/No_More_Flex.svg?branch=master)](https://travis-ci.com/pl0q1n/No_More_Flex)

## How to build
* go build ./cmd/nmf-server

## How to run
* docker pull pl0q1n/nmf
* docker run -d -p 8080:8080 nmf

## How to find API manual:
* go to http://petstore.swagger.io/
* run your nmf server
* enter address of your server in input box with /swagger.json at the end (eg. http://127.0.0.1:44586/swagger.json)