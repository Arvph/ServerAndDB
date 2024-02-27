package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Создвние структуры Storage
type Storage struct {
	config *Config
	// Дескриптор для БД
	db *sql.DB
	// Subfield for repo interfacing
	UserRepository    *UserRepository
	ArticleRepository *ArticleRepository
}

// Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Метод открывает связь с БД
func (storage *Storage) Open() error {
	// Открытие и инициализация соединения с БД
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	// Проверка подключения к БД
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Databade connection created successfully")
	return nil
}

// Закрыть связь с БД
func (storage *Storage) Close() {
	storage.db.Close()
}

// Public Repo for Articles
func (s *Storage) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		storage: s,
	}
	return s.UserRepository
}

// Public Repo for Users
func (s *Storage) Article() *ArticleRepository {
	if s.ArticleRepository != nil {
		return s.ArticleRepository
	}
	s.ArticleRepository = &ArticleRepository{
		storage: s,
	}
	return s.ArticleRepository
}
