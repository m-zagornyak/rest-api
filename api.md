# todo-app

## Authentication

### POST
JSON Web Tokens (JWT).
* Мы должны отправить поле «Токен» в заголовке.

* Значение токена должно генерироваться на основе парольной фразы.

* Парольная фраза, размещенная на сервере REST API.

* Чтобы получить токен аутентификации, мы должны использовать этот запрос
```
/login
```
И отрпавить Имя, пароль

```json
{
    "Username": "user",
    "Password": "pass"
}
```

Если юзер и пароль верны мы получаем 200 ок response с нашим ключом, и после этого мы должны его сохранить в localStorage:

```
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Date: Sun, 12 Jul 2020 15:20:25 GMT
Content-Length: 8
Connection: close
"error"
```



responce:

```
[
  {
    "id": 131,
    "name": "AGRO",
    "ticker": "AGRO",
    "amount": 32,
    "priceperitem": 103.48,
    "purchaseprice": 31360,
    "pricecurrent": 33113.6,
    "percentchanges": 5.5918367346938735,
    "yearlyinvestment": 10.327319061645275,
    "clearmoney": 1721.3631999999986,
    "datepurchase": "2019-12-09T13:59:53.66277+03:00",
    "datelastupdate": "2020-06-30T18:00:49+03:00",
    "type": "share"
  },
  {
    "id": 145,
    "name": "Детский мир",
    "ticker": "DSKY",
    "amount": 450,
    "priceperitem": 103.48,
    "purchaseprice": 41885,
    "pricecurrent": 46566,
    "percentchanges": 11.1758386057061,
    "yearlyinvestment": 40.81456061148672,
    "clearmoney": 4636.7745,
    "datepurchase": "2020-03-13T13:59:53.66277+03:00",
    "datelastupdate": "2020-06-19T23:54:26.655833+03:00",
    "type": "share"
  },
  {
    "id": 138,
    "name": "Газпром",
    "ticker": "GAZP",
    "amount": 300,
    "priceperitem": 195.81,
    "purchaseprice": 68100,
    "pricecurrent": 58743,
    "percentchanges": -13.740088105726873,
    "yearlyinvestment": -35.809394273127744,
    "clearmoney": -9420.4215,
    "datepurchase": "2020-01-31T13:59:53.66277+03:00",
    "datelastupdate": "2020-06-19T23:59:27.189249+03:00",
    "type": "share"
  }
]
```