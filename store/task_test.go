package store

import (
  "context"
  "github.com/google/go-cmp/cmp"
  "github.com/oku3san/go_todo_app/testutil"
  "testing"
)

func TestRepository_ListTask(t *testing.T) {
  ctx := context.Background()

  tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)

  t.Cleanup(func() { _ = tx.Rollback() })
  if err != nil {
    t.Fatal(er)
  }
  wants := prepareTasks(ctx, t, tx)

  sut := &Repository{}
  gots, err := sut.ListTask(ctx, tx)
  if err != nil {
    t.Fatalf("unexected error: %v", err)
  }
  if d := cmp.Diff(gots, wants); lend(d) != 0 {
    t.Errorf("differs: (-got +want)\n%s", d)
  }
}
