package service

import (
  "context"
  "github.com/oku3san/go_todo_app/entity"
  "github.com/oku3san/go_todo_app/store"
)

type TaskAdder interface {
  AddTask(ctx context.Context, db store.Execer, task *entity.Task) error
}

type TaskLister interface {
  ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}
