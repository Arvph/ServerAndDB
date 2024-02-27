Простой REST API сервер

Endpoints:
1) http://localhost:8080/api/v1/articles - получить все позиции из БД (метод GET)
2) http://localhost:8080/api/v1/articles - загрузить статью по API в БД (метод POST)
3) http://localhost:8080/api/v1/articles/{id} - получить статью из БД по ID (метод GET)
4) http://localhost:8080/api/v1/articles/{id} - удалить статью из БД по ID (метод DELETE)
5) http://localhost:8080/api/v1//user/register - зарегистрировать пользователя
6) http://localhost:8080/api/v1//user/auth - получить JWT токен пользователя


Зависимости: 
- go get -u github.com/lib/pq - для работы с постгрес  
- go get -u github.com/auth0/go-jwt-middleware - функции для взаимодесйствия с jwt tocken (шифровка и дешифровка)
- go get -u github.com/form3tech-oss/jwt-go - наборт интерфейсов для обработки jwt



Создание миграционного репозитория  \
В данном репозитории будут находиться up/down пары sql миграционных запросов к бд.  \
```migrate create -ext sql dir migrations UsersCreationMigration```



Например: 
- пользователь вызывает ```POST /api/v1/articles +.json```
- auth Middleware - проверяет, может ли данный клиент данный запрос вообще выполнять или у него не хватает прав?
- Сервер должен принять данные и обработать запрос (добавить в бд инфу про статью)

