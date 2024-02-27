package api

import (
	"net/http"

	"github.com/arvph/ServerAndDB/internal/app/middleware"
	"github.com/arvph/ServerAndDB/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

// Конфигурация поля api.logger
// метод берет данные из api.config.LoggerLevel и помещает их в api.logger
func (a *API) configureLoggerField() error {
	//logrus.ParseLevel используется для преобразования строкового представления уровня логирования в константу типа logrus.Level,
	//которая представляет собой уровень логирования:  "info" в logrus.InfoLevel
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}

	// SetLevel - используется для установки уровня логирования для логгера.
	// Уровень логирования определяет, какие сообщения будут записываться логгером.
	// Например, если уровень логирования установлен на logrus.InfoLevel,
	// то логгер будет записывать сообщения с уровнями логирования Info, Warn, Error, Fatal и Panic,
	// но не будет записывать сообщения с уровнями логирования Debug и Trace.
	a.logger.SetLevel(log_level)
	return nil
}

// Конфигурация маршрутизатора (поле router API)
func (a *API) configureRouterField() {
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	// Было до JWT
	// a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticlesById).Methods("GET")
	// Теперь требует наличие JWT
	a.router.Handle(prefix+"/articles/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.GetArticlesById),
	)).Methods("GET")
	//
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")

	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")
	a.router.HandleFunc(prefix+"/user/auth", a.PostToAuth).Methods("POST")
}

// Конфигурация поля api.storage
// метод берет данные из api.config.Storage и создает на их базе api.storage
func (a *API) configureStorageField() error {
	// создаем структуру storage.Storage и передаем в нее поле api.config.Storage
	storage := storage.New(a.config.Storage)

	// пытаемся соединись соединение, если невозможно возвращает ошибку
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
