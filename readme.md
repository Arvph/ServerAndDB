# Простой REST API Сервер

Добро пожаловать в репозиторий нашего простого REST API сервера. Этот сервер представляет собой базовую реализацию RESTful веб-сервиса, предназначенного для работы с артикулами и пользовательскими данными.

### Основные Функции

Сервер предоставляет следующие endpoints для взаимодействия:

- <b>GET /api/v1/articles:</b> Получение всех статей из базы данных.
- <b>POST /api/v1/articles:</b> Загрузка статьи в базу данных через API.
- <b>GET /api/v1/articles/{id}:</b> Получение статьи из базы данных по ID.
- <b>DELETE /api/v1/articles/{id}:</b> Удаление статьи из базы данных по ID.
- <b>POST /api/v1/user/register:</b> Регистрация пользователя.
- <b>POST /api/v1/user/auth:</b> Получение JWT токена для пользователя.

### Зависимости

Проект использует следующие зависимости:

- PostgreSQL драйвер:
```go get -u github.com/lib/pq```

- JWT Middleware:
```go get -u github.com/auth0/go-jwt-middleware```

- Обработка JWT:
```go get -u github.com/form3tech-oss/jwt-go```

### Миграции Базы Данных

Для управления миграциями используется система миграций. Создание миграционного файла для структуры пользователей можно выполнить следующим образом:
```migrate create -ext sql -dir migrations UsersCreationMigration```

### Работа Сервера

- При отправке запроса на POST <b>/api/v1/articles</b> с JSON-телом, сервер через промежуточное ПО (middleware) проверяет права доступа пользователя.
В случае успешной авторизации, сервер обрабатывает запрос и добавляет информацию о статье в базу данных.

### Инструкции По Запуску

- Убедитесь, что все зависимости установлены.
- Запустите сервер командой <b>go build -v ./cmd/api</b>  или командой <b>make</b>.

#### Настройка PostgreSQL с Docker

Для работы с PostgreSQL можно использовать Docker. Ниже приведены команды для управления базой данных:
- Запуск PostgreSQL \
```docker run --name postgres16 -p 5434:5434 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:16-alpine```
- Создание Базы Данных \
```docker exec -it postgres16 createdb --username=admin --owner=admin test_db```
- Удаление Базы Данных \
```docker exec -it postgres16 dropdb --force test_db```
- Применение Миграций \
```migrate -path migration -database "postgres://localhost:5434/restapi?sslmode=disable&user=admin&password=admin" -verbose up```
- Откат Миграций \
```migrate -path migration -database "postgres://localhost:5434/restapi?sslmode=disable&user=admin&password=admin" -verbose down```

#### Сборка и Запуск
Указанные команды .PHONY используются для упрощения работы с базой данных и миграциями. Выполнение этих команд поможет вам управлять состоянием базы данных и применять необходимые изменения.

### Заключение

Этот проект является отличным стартовым пунктом для разработки и понимания работы REST API серверов. Он включает основные операции CRUD и работу с JWT для аутентификации и авторизации.
