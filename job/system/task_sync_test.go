package system

import (
	"context"
	"testing"
	"time"
)

func TestTaskSync(t *testing.T) {
	task := NewTaskSyncJob(context.TODO(), "* * * * *")
	task.Start()
	time.Sleep(5 * time.Minute)
}
