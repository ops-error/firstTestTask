# Привет, First Test Task
### Стек:
- Golang
- PostgreSQL
- Kafka 

Это мой первый проект на Go. Да и в целом я впервые работаю с Kafka и PostgreSQL. 
Кеша тут нет, поскольку я пока не до конца поняла эту тему, но все впереди.
На данный момент реализовано:
- GET-запрос `/orders/{uid}`

Принимает в параметрах uid заказа

- Топик `create-order`

Принимает сообщение и создает запись(_-и_) в БД

### Отправка сообщений на сервер через Kafka:
Обычно я использовала для этого файл в корневой папке `order.json` и в терминале 
производила следующие действия:
- Копирование файла в контейнер

`docker cp order.json firsttesttask-kafka-1:/order.json`

- Отправка сообщения

`docker exec -i firsttesttask-kafka-1 \` \
`kafka-console-producer \` \
`--bootstrap-server localhost:9092 \` \
`--topic create-order < /order.json`