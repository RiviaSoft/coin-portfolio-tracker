# coin-portfolio-tracker
2021 Veritabanı Sistemleri Ödevi


To-do list: http://gg.gg/cpt-todo


API v0.1 documentation : https://www.notion.so/msrexe/COIN-PORTFOLIO-TRACKER-API-8436cc6c1db242bb8a113f6bdb0b14cd

### **API BASE URL : "xxx.com/api/"**

**-AUTH**

POST || */auth/register

request type : json

request :

{

"name" : "abc",

"email" : "abc@example.com",

"password" : "abcabc"

}

POST || */auth/login

request type : json

request :

{

"email" : "abc@example.com",

"password" : "abcabc"

}

**********************************************************

**-USER**

request type : json

request :

{

"id" : 12,

"name" : "abc",

"email" : "abc@example.com",

"password" : ""

}

GET || */users/getcurrentuser

POST || */users/delete

POST || */users/update

**********************************************************

**-RECENT BUY OPERATIONS**

request type : json

request :

{

"userid" : 12,

"coinsymbol" : "BTC/USDT",

"coinamount" : 12.2,

"buycost" : 23.32

}

GET || */operations/getall

GET || */operations/get/id:<id>

POST|| */operations/add

POST || */operations/delete

POST || */operations/update

**********************************************************

**-ARCHIVED OPERATIONS**

request type : json

request :

{

"userid" : 12,

"coinsymbol" : "BTC/USDT",

"coinamount" : 12.2,

"buycost" : 23.32,

"sellcost" : 35.45

}

GET || */archivedoperations/getall

GET || */operations/get/id:<id>

POST || */archivedoperations/add /// satma işlemi

POST || */archivedoperations/delete

**********************************************************

**-WALLET**

request type : json

request :

{

"userid" : 12,

"name" : "Kısa vade"

}

GET || */wallets/getall #kullanıcıya ait tüm cüzdanları getirir

GET || */wallets/getbyid?id=2 #kullanıcıya ait 2 nolu cüzdanını getirir

POST || */wallets/add #yeni cüzdan oluşturma veya cüzdana coin eklemeye yarar

POST || */wallets/delete #[c](//cüzdandaki)üzdanı siler

**********************************************************

**-WALLET OPERATIONS**

request type : json

request :

{

"walletid" : 6,

"operationid" : 2

}

GET || */walletoperations/getall #cüzdana ait tüm işlemleri(coinleri) getirir

POST || */walletoperations/add #cüzdana coin eklemeye yarar

POST || */walletoperations/delete #cüzdandaki coini siler
