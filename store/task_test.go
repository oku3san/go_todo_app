package store

import (
  "context"
  "github.com/google/go-cmp/cmp"
  "github.com/oku3san/go_todo_app/clock"
  "github.com/oku3san/go_todo_app/entity"
  "github.com/oku3san/go_todo_app/testutil"
  "testing"
)

func TestRepository_ListTask(t *testing.T) {
  ctx := context.Background()

  tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)

  t.Cleanup(func() { _ = tx.Rollback() })
  if err != nil {
    t.Fatal(err)
  }
  wants := prepareTasks(ctx, t, tx)

  sut := &Repository{}
  gots, err := sut.ListTask(ctx, tx)
  if err != nil {
    t.Fatalf("unexected error: %v", err)
  }
  if d := cmp.Diff(gots, wants); len(d) != 0 {
    t.Errorf("differs: (-got +want)\n%s", d)
  }
}

func prepareTasks(ctx context.Context, t *testing.T, con Execer) entity.Tasks {
  t.Helper()
  if _, err := con.ExecContext(ctx, "DELETE FROM task;"); err != nil {
    t.Logf("failed to initiallize task: %v", err)
  }
  c := clock.FixedClocker{}
  wants := entity.Tasks{
    {
      Title: "want task 1", Status: "todo",
      Created: c.Now(), Modified: c.Now(),
    },
    {
      Title: "want task 2", Status: "done",
      Created: c.Now(), Modified: c.Now(),
    },
  }
  tasks := entity.Tasks{
    wants[0],
    {
      Title: "not want task", Status: "todo",
      Created: c.Now(), Modified: c.Now(),
    },
    wants[1],
  }
  result, err := con.ExecContext(ctx,
    `INSERT INTO task (title, status, created, modified)
			VALUES
			    (?, ?, ?, ?),
			    (?, ?, ?, ?),
			    (?, ?, ?, ?);`,
    tasks[0].Title, tasks[0].Status, tasks[0].Created, tasks[0].Modified,
    tasks[1].Title, tasks[1].Status, tasks[1].Created, tasks[1].Modified,
    tasks[2].Title, tasks[2].Status, tasks[2].Created, tasks[2].Modified,
  )
  if err != nil {
    t.Fatal(err)
  }
  id, err := result.LastInsertId()
  if err != nil {
    t.Fatal(err)
  }
  tasks[0].ID = entity.TaskID(id)
  tasks[1].ID = entity.TaskID(id + 1)
  tasks[2].ID = entity.TaskID(id + 2)
  return wants
}
