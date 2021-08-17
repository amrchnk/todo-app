# REST API приложения для создания списков на Go

Данное приложение сделано по [видео-курсу](https://www.youtube.com/playlist?list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8) и охватывает обширный стек технологий

В нем реализованы следующие концепции:
* Построение структуры приложения с использованием подхода Чистой архитектуры и внедрением зависимостей
* БД Postgres: генерация файлов миграции (папка schema), написание SQL-запросов и работа с БД с использованием библиотеки [sqlx](https://github.com/jmoiron/sqlx)
* Работа с переменными окружения, использование библиотеки [viper](https://github.com/spf13/viper) для работы с файлами конфигурации
* Аутентификация пользователя с использованием JWT-токенов
* Хэширование паролей при регистрации нового пользователя
* Работа с фреймворков [gin](https://github.com/gin-gonic/gin) для настройки роутера

## Запуск приложения

Для запуска БД понадобится загрузить образ postgres в Docker:

`docker pull postgres`
 
Затем выполнить следующую команду:

`docker run --name=todo-db -e POSTGRES_PASSWORD='[пароль]' -p 5436 -d --rm postgres`

И применить файлы миграции к бд:

`migrate -path ./schema -database postgres://postgres:[пароль]@localhost:5436/postgres?sslmode=disable up
`

После чего запустить приложение:

`go run cmd/main.go`





