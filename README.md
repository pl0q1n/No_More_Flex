# No More Flex

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
sender (string, required) -- indentifier of money account (eg. card name) which send money.  
receiver (string, required) -- identifier of money account (eg. company name) who get the money.  
value (int, required) -- value of transaction.
time (int, required) -- timestamp of transaction (or time of notification).  
category (string, optional) -- name of category which transaction related to (eg. "grocery"), if not set, we should try to classificate by receiver field.  
</p>

### Get Expenses

Request to get expenses (array of transactions) from db 

**Github issue** :

**URL** : `/transactions`

**Method** : `GET`

**Auth required** : YES

**Params**
<p>
sender (string, optional) -- indentifier of money account (eg. card name) which send money.  
receiver (string, optional) -- identifier of money account (eg. company name) who get the money.  
time_range_start (int, optional) -- start of time range for transaction (all transactions with time >= time_range_start should be presented), if omitted, time_range_start == 1970-01-01 00:00:00 (in timestamp).  
time_range_end (int, optional) -- end of time range for transaction (all transactions with time <= time_range_start should be presented), if omitted, time_range_end == now.  
category (string, optional) -- name of category which transaction related to (eg. "grocery"), if not set, we should try to classificate by receiver field. If omitted, return transactions for all catefories.  
</p>