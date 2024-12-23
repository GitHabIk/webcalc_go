# webcalc_go
Незамысловатый веб сервис калькулятора

# Что он из себя представляет?
> Небольшой веб сервис на который нужно отправить POST запрос с примером и он даст ответ

# **Как им пользоваться?**
> 1) Скачиваете
> 2) Переходите в папку **cmd** и запускаете файл **main.go** (```go run cmd/main.go```) (P.S. сервер запускается на "http://localhost:8080/api/v1/calculate")
> 3) Открываем терминал
> 4) Пишем туда POST запрос который должен выглядить так: **```Invoke-WebRequest -Uri "http://localhost:8080/api/v1/calculate" -Method Post -Body '{"expression": "3 + 5"}' -ContentType "application/json"```** (В некоторых случаях может выглядеть по другому)
> 5) В терминале должно вывести примерно это: 
```
StatusCode        : 200
StatusDescription : OK
Content           : {"result":"8.000000"}

RawContent        : HTTP/1.1 200 OK
                    Content-Length: 22
                    Content-Type: application/json
                    Date: Wed, 18 Dec 2024 17:52:29 GMT

                    {"result":"8.000000"}

Forms             : {}
Headers           : {[Content-Length, 22], [Content-Type, application/json], [Date, Wed, 18 Dec 2024 17:52:29 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 22
```
> Нам интересны первые 3 строчки
> 1) ```StatusCode        : 200``` - обозначает статус, в данном случае всё нормально
> 2) ```StatusDescription : OK``` - описание статуса, в данном случае OK или же всё нормально
> 3) ```Content           : {"result":"8.000000"}``` - обозначает результат примера, у нас это 8.000000 или же просто 8

Готово! Мы разобрались в запросе!

# Заключение
> Я очень старался сделать скрипты и ридми, это мой первый проект так что не судите строго пожалуйста :)
> Спасибо что дочитали до конца можете отписать в тг - **```Olomadness```**
