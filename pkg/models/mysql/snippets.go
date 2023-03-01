package mysql

import (
	"databse/sql"

	// Импорт только что созданный пакет models. Нам нужно добавить к нему префикс
	// независимо от того, какой путь к модулю вы установили
	// чтобы оператор импорта выглядил как
	// "{ваш-путь-модуля}/pkg/models"..
	"golangify.com/snippetbox/pkg/models"
)

// Snippetbox - Определяем тип который обертывает пул подключения sql.DB
type SnippetModel struct {
	DB *sql.DB
}

// Insert - метод для создания новой заметки в базе данных.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get - метод для возвращения данных заметки по её идентификатору ID.
func (m *SnippetModel) Get(id int) (*model.Snippet, error) {
	return nil, nil
}

// Latest - метод возварщает 10 наиболее часто используемой заметки.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
