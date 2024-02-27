package api

import (
	"net/http"

	"github.com/arvph/ServerAndDB/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// структура с настройками для api
type API struct {
	config  *Config
	logger  *logrus.Logger // берем из библиотеки logrus
	router  *mux.Router    // берем из бибиотеки gorilla/mux
	storage *storage.Storage
}

// API constructor: начальные настройки API
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server, configure loggers, router, database connection and etc...
func (api *API) Start() error {
	// Базовая конфигурация логгера
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	// Подтверждение, что логгер сконфигурирован. Медод Info берется из библиотеки logrus
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	// Базовая конфигурация маршрутизатора
	api.configureRouterField()

	// Базовая конфигурация хранилища
	if err := api.configureStorageField(); err != nil {
		return nil
	}
	// На этапе валидного завершения стартует http server. Передаем порт и mux.NewRouter
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
