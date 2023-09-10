package postgres

import (
	"reflect"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestStorage_GetTasks(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		taskID   int
		authorID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.GetTasks(tt.args.taskID, tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_SetTask(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.SetTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.SetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.SetTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
