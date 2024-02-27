package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/arvph/ServerAndDB/internal/app/middleware"
	"github.com/arvph/ServerAndDB/internal/app/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

// Вспомогательная структура для формирования сообщений
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

// Функция устанавливает заголовок ответа Content-Type в значение "application/json",
// указывая, что тело ответа будет представлять собой данные в формате JSON.
// Заголовок Content-Type устанавливается на "application/json", что гарантирует,
// что клиент ожидает получить данные в формате JSON.
func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	//инициализация хедеров
	initHeaders(writer)
	// логируем момент начала обработки
	// Метод Info() принадлежит объекту *logrus.Logger и
	// используется для записи сообщения уровня информации (info) в журнал
	api.logger.Info("Get All Articles GET /api/v1/articles")
	// получаем все статьи из БД
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		// ошибка на этапе подключения
		api.logger.Info("error while Articles.SelectAll: ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Access DB error",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) PostArticle(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Article POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid JSON received from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided JSON is invalid",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Trouble with creating new article: ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Database access error. Please try again",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(a)
}

func (api *API) GetArticlesById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Article by ID GET /api/v1/articles/{id}")

	// Функция mux.Vars(r *http.Request) является частью пакета mux,
	// который является популярным маршрутизатором HTTP для языка Go.
	// Она используется для извлечения переменных маршрута из HTTP-запроса.

	// mux.Vars(req)["id"]: Это извлечение значения переменной маршрута id из HTTP-запроса req,
	// который был обработан маршрутизатором mux.
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Trouble with parsing {id} new article: ", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Error id value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article, ok, err := api.storage.ArticleRepository.FindArticleById(id)
	if err != nil {
		api.logger.Info("Trouble with access db article with id. Err: ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Access DB error. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Cannot find article with id. Err: ", err)
		msg := Message{
			StatusCode: 404,
			Message:    "Article with such id does not exist",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)
}

func (api *API) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("DELETE Article by ID DELETE /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Trouble with deleting {id} article: ", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Error id value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.ArticleRepository.FindArticleById(id)
	if err != nil {
		api.logger.Info("Trouble with access db article with id. Err: ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Access DB error. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Cannot find article with id. Err: ", err)
		msg := Message{
			StatusCode: 404,
			Message:    "Article with such id does not exist",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Trouble with deleting article with id. Err: ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Delete article error. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Article wirh ID: %v deleted", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) PostUserRegister(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post user register request POST /api/v1/user/register")

	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid JSON file: ", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Error invalid JSON file",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Trouble with user access with id. Err: ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Access user error. Try again_2",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if ok {
		api.logger.Info("User with this ID already exists", err)
		msg := Message{
			StatusCode: 400,
			Message:    "User with this login already exists in DB",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	userAdded, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("Trouble with user access with id. Err: ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Access user error. Try again_3",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User with Login: %v registered", userAdded.Login),
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post to Auth POST /api/v1/user/auth")
	var userFromJson models.User
	err := json.NewDecoder(req.Body).Decode(&userFromJson)
	//Обрабатываем случай, если json - вовсе не json или в нем какие-либо пробелмы
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// обнаружить пользователя с таким login в БД
	userInDB, ok, err := api.storage.User().FindByLogin(userFromJson.Login)
	// Проблема доступа к бд
	if err != nil {
		api.logger.Info("Can not make user search in database:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles while accessing database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// если подключение удалось, но пользователя с таким логином нет
	if !ok {
		api.logger.Info("User with that login does not exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login does not exists in database. Try register first",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// если логин есть, проверяем совпадение паролей
	if userInDB.Password != userFromJson.Password {
		api.logger.Info("Invalid credetials to auth")
		msg := Message{
			StatusCode: 404,
			Message:    "Your password is invalid",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// теперь выбиваем токен как знак успешной аутентификации
	token := jwt.New(jwt.SigningMethodHS256)             // Тот же метод подписания токена, что и в JwtMiddleware.go
	claims := token.Claims.(jwt.MapClaims)               // Дополнительные действия (в формате мапы) для шифрования
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //Время жизни токена
	claims["admin"] = true
	claims["name"] = userInDB.Login
	tokenString, err := token.SignedString(middleware.SecretKey)
	// в случае если токен выдать не удалось
	if err != nil {
		api.logger.Info("Can not claim jwt-token")
		msg := Message{
			StatusCode: 500,
			Message:    "We have some trouble. Try again.",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// в случае если токен успешно выбит - отдаем его клиенту
	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}
