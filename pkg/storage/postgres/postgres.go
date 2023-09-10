package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Задача
type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

// GetTasks возвращает список задач из БД по id задачи и автору
func (s *Storage) GetTasks(taskID, authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT t.id,
		       t.opened,
			   t.closed,
			   t.author_id,
			   t.assigned_id,
			   t.title,
			   t.content
		FROM tasks t
		WHERE ($1 = 0 OR t.id = $1)
		AND ($2 = 0 OR t.author_id = $2)
		ORDER BY t.id;
	`,
		taskID,
		authorID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

// SetTask создаёт новую задачу и возвращает её id.
func (s *Storage) SetTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, content)
		VALUES ($1, $2) RETURNING id;
		`,
		t.Title,
		t.Content,
	).Scan(&id)
	return id, err
}
