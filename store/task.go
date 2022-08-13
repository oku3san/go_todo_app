package store

import (
  "context"
  "github.com/oku3san/go_todo_app/entity"
)

func (r *Repository) ListTask(
  ctx context.Context, db Queryer, ) (entity.Tasks, error) {
  tasks := entity.Tasks{}
  sql := `Select id, title, status, created, modified FROM task;`
  if err := db.SelectContext(ctx, &tasks, sql); err != nil {
    return nil, err
  }
  return tasks, nil
}

func (r *Repository) AddTask(
  ctx context.Context, db Execer, t *entity.Task,
) error {
  t.Created = r.Clocker.Now()
  t.Modified = r.Clocker.Now()
  sql := `INSERT INTO task (title, status, created, modifyed)
  VALUES (?, ?, ?, ?)`
  result, err := db.ExecContext(
    ctx, sql, t.Title, t.Status, t.Created, t.Modified,
  )
  if err != nil {
    return err
  }
  id, err := result.LastInsertId()
  if err != nil {
    return err
  }
  t.ID = entity.TaskID(id)
  return nil
}
