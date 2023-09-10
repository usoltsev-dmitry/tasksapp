package memdb

import "tasksapp/pkg/storage/postgres"

// Пользовательский тип данных - реализация БЛ в памяти
type DB []postgres.Task

// Выполнение контракта интрефейса storage.Interface
func (db DB) GetTasks(int, int) ([]postgres.Task, error) {
	return db, nil
}
func (db DB) SetTask(postgres.Task) (int, error) {
	return 0, nil
}
