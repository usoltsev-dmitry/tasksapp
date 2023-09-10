package storage

import "tasksapp/pkg/storage/postgres"

// Интерфейс БД
type Interface interface {
	GetTasks(int, int) ([]postgres.Task, error)
	SetTask(postgres.Task) (int, error)
}
