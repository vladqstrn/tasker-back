# TASKER

Приложения для списка задач.

Для запуска необходим [сервис авторизации]("https://github.com/vladqstrn/tasker-auth")

# API

## POST /tasks/create HTTP/1.1

Создает новую задачу.
```
POST /tasks/create HTTP/1.1
Content-Type: application/json
{

    "Title":"test",
    "Text":"test",
    "Description":"test",
    "Executor":"test",
    "Status": "test"
}
```

## PUT /tasks/update/:id

Обновляет данные по задаче с заданным идентификатором.

```
PUT /tasks/update/:id HTTP/1.1
Content-Type: application/json
{
   "Title":"test",
    "Text":"test",
    "Description":"test",
    "Executor":"test",
    "Status": "test" 
}
```
## DELETE /tasks/delete/:id

Удаляет запись о задаче с переданным идентификатором.

```
DELETE /tasks/delete/:id HTTP/1.1
```
## GET /tasks/getalltask

Получает все сохраненные задачи.

```
GET /tasks/getalltask HTTP/1.1
```


## GET /tasks/gettask/:id

Получает задачу с конкретным идентификатором.
```
GET /tasks/gettask/:id HTTP/1.1
```

## GET /tasks/getuser/:id

Получает все задачи пользователя с данным идентификатором.

```
GET /tasks/getuser/:id HTTP/1.1
```


