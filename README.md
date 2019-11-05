# No More Flex

[![Build Status](https://travis-ci.com/pl0q1n/No_More_Flex.svg?branch=master)](https://travis-ci.com/pl0q1n/No_More_Flex)

## How to build
* go build ./cmd/nmf-server

## How to run
* docker pull pl0q1n/nmf
* docker run -d -p 8080:8080 nmf

## NMF API

### Add Transaction

Request to add transaction from user to db.

**Github issue** :

**URL** : `/add`

**Method** : `POST`

**Auth required** : YES

```json
[
    {
     "sender": "MasterCard #1",
     "receiver": "BANK OF AMERICA",
     "value": 1337,
     "int": "2016-06-22 19:10:25",
     "category": "Investing" 
    }
]
```

**Params**
<p>
sender (string, required) -- indentifier of money account (eg. card name) which send money.  <br>
receiver (string, required) -- identifier of money account (eg. company name) who get the money.   <br>
value (int, required) -- value of transaction. <br>
time (int, required) -- timestamp of transaction (or time of notification).   <br>
category (string, optional) -- name of category which transaction related to (eg. "grocery"), if not set, we should try to classificate by receiver field.   <br>
</p>

### Get Expenses

Request to get expenses (array of transactions) from db 

**Github issue** :

**URL** : `/transactions`

**Method** : `GET`

**Auth required** : YES

**Params**
<p>
sender (string, optional) -- indentifier of money account (eg. card name) which send money.   <br>
receiver (string, optional) -- identifier of money account (eg. company name) who get the money.   <br>
time_range_start (int, optional) -- start of time range for transaction (all transactions with time >= time_range_start should be presented), if omitted, time_range_start == 1970-01-01 00:00:00 (in timestamp).   <br>
time_range_end (int, optional) -- end of time range for transaction (all transactions with time <= time_range_start should be presented), if omitted, time_range_end == now.   <br>
category (string, optional) -- name of category which transaction related to (eg. "grocery"), if not set, we should try to classificate by receiver field. If omitted, return transactions for all catefories.   <br>
</p>
