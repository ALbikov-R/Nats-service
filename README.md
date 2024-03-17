# NATS service
Сервис для получения данных при помощи брокера сообщения NATS. Получает на вход JSON файл следующего формата. Хранение данных происходит inmemory и PostgreSQL. Проект находится в папке nats-service, script - необходим для демонстрации его работы (симуляция получения данных со стороннего сервиса)
## Installing application
Установка приложения происходит при помощи Makefile. Для использования данного сервиса необходимо установить NATS. Порядок запуска:
```
nats-server
go mod download
make
```
## End points
Сервис содержит 2 end point'а:
+ localhost:8084/
+ localhost:8084/postfrom
### End point localhost:8084/
Содержит форму ввода для поиска товара по order_uid, после нажатия кнопки "отправить", то будет переадрисация на endpoint localhost:8084/postfrom?{order_uid}
### End point localhost:8084/postfrom
Содержит полный список данных по order_uid, но можно непосредственно сразу получить список данных по запросу localhost:8084/postfrom?{order_uid}
## JSON 
Передача данных через NATS происходит по следующему адресу:
```
0.0.0.0:4222
```
Формат JSON файла, который приходит с брокера NATS:
```
{
  "order_uid": "",
  "track_number": "",
  "entry": "",
  "delivery": {
    "name": "",
    "phone": "",
    "zip": "",
    "city": "",
    "address": "",
    "region": "",
    "email": ""
  },
  "payment": {
    "transaction": "",
    "request_id": "",
    "currency": "",
    "provider": "",
    "amount": ,
    "payment_dt": ,
    "bank": "",
    "delivery_cost": ,
    "goods_total": ,
    "custom_fee": 
  },
  "items": [
    {
      "chrt_id": ,
      "track_number": "",
      "price": ,
      "rid": "",
      "name": "",
      "sale": ,
      "size": "",
      "total_price": ,
      "nm_id": ,
      "brand": "",
      "status": 
    }
  ],
  "locale": "",
  "internal_signature": "",
  "customer_id": "",
  "delivery_service": "",
  "shardkey": "",
  "sm_id": ,
  "date_created": "",
  "oof_shard": ""
}
```
## Config File
Конфиг сервиса состоит из следующих настроек:
```
nats:
  subject: "subject"
  natsURL: "0.0.0.0:4222"
#PostgreSQL
postgresql: 
  database_URL: "host=localhost port=5432 user=postgres password=1234 dbname=wildtest sslmode=disable"
#Router
bind_addr: ":8084"
```
## Example of usage 
### Скрипт для примера
Во время работы программы для отправки данных будет использован скрипт. Он находится в корне проекте в папке (script). Сначала необходимо запустить сам nats-service, а только потом скрипт, который отпраляет даннные на сервис. 
```
.\script\script.exe
```
В результате чего, будут получен JSON файл. Содержание JSON файла можете посмотреть в script\data.json .
### Пример данных 
```
{
        "order_uid": "b563feb7b2b84b6test",
        "track_number": "TRACK",
        "entry": "WBIL",
        "delivery": {
          "name": "Test Testov",
          "phone": "+9720000000",
          "zip": "2639809",
          "city": "Kiryat Mozkin",
          "address": "Ploshad Mira 15",
          "region": "Kraiot",
          "email": "test@gmail.com"
        },
        "payment": {
          "transaction": "b563feb7b2b84b6test",
          "request_id": "",
          "currency": "USD",
          "provider": "wbpay",
          "amount": 1817,
          "payment_dt": 1637907727,
          "bank": "alpha",
          "delivery_cost": 1500,
          "goods_total": 317,
          "custom_fee": 0
        },
        "items": [
          {
            "chrt_id": 9934930,
            "track_number": "TRACK",
            "price": 453,
            "rid": "ab4219087a764ae0btest",
            "name": "Mascaras",
            "sale": 30,
            "size": "0",
            "total_price": 317,
            "nm_id": 2389212,
            "brand": "Vivienne Sabo",
            "status": 202
          }
        ],
        "locale": "en",
        "internal_signature": "",
        "customer_id": "test",
        "delivery_service": "meest",
        "shardkey": "9",
        "sm_id": 99,
        "date_created": "2021-11-26T06:22:19Z",
        "oof_shard": "1"
    }
```
###  Работа Nats-service 
Необходимо перейти по ссылке localhost:8084/ и в TextBox ввести order_uid заказа. После чего сервис выполнить POST запрос с order_uid на localhost:8084/postfrom?{order_uid}, где будет содержаться полная информация о заказе.