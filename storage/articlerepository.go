package storage

import (
	"fmt"
	"log"

	"github.com/arvph/ServerAndDB/internal/app/models"
)

// Структура для работы с ArticleRepository
type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle string = "articles"
)

// add the article to db
func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES ($1, $2, $3) RETURNING id", tableArticle)
	// Метод QueryRow() принадлежит объекту *sql.DB и используется для выполнения запроса,
	// который ожидает возвращения только одной строки результата.
	// Метод Scan() принадлежит объекту *sql.Row и используется для сканирования значений
	// из строки результата запроса в переменные, переданные в качестве аргументов.
	err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content).Scan(&a.ID)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// delete the article from db
func (ar *ArticleRepository) DeleteById(id int) (*models.Article, error) {
	article, ok, err := ar.FindArticleById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableArticle)
		// Метод Exec() принадлежит объекту *sql.DB и используется для выполнения запросов к базе данных,
		// которые не ожидают возвращения строк результата (например, запросы на обновление,
		// вставку или удаление данных)
		_, err := ar.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}
	return article, nil
}

// get the article from db
func (ar *ArticleRepository) FindArticleById(id int) (*models.Article, bool, error) {
	articles, err := ar.SelectAll()
	var founded bool

	if err != nil {
		return nil, founded, err
	}
	var articleFounded *models.Article
	for _, a := range articles {
		if a.ID == id {
			articleFounded = a
			founded = true
			break
		}
	}
	return articleFounded, founded, nil
}

// get all articles from db
func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableArticle)
	// Метод Query() принадлежит объекту *sql.DB и используется для выполнения запросов к базе данных,
	// которые ожидают возвращения строк результата (например, запросы на выборку данных).
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	// Метод Close() принадлежит объекту *sql.Rows и используется для закрытия набора строк результата запроса.
	// Вызов метода Close() освобождает ресурсы, связанные с объектом *sql.Rows,
	// включая ресурсы базы данных и память, занимаемую результатами запроса.
	defer rows.Close()

	articles := make([]*models.Article, 0)
	// Метод Next() принадлежит типу *sql.Rows, который представляет набор строк результата запроса.
	// Возвращаемое значение - это булево значение, которое указывает,
	// есть ли еще строки для сканирования в результате запроса.
	for rows.Next() {
		a := models.Article{}
		// Метод Scan() используется для сканирования значений из текущей строки результата
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	// Метод Err() принадлежит объекту *sql.Rows и используется для получения любой ошибки,
	// которая может произойти во время работы с набором строк результата запроса.
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}
