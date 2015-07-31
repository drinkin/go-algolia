package algolia

import (
	"fmt"
	"time"

	"github.com/drinkin/di/random"
)

type TaskStatus struct {
	Status  string `json:"status"`
	Pending bool   `json:pendingTask`
}

func (ts *TaskStatus) IsPublished() bool {
	return ts.Status == "published"
}

// Task represents an algolia task response
type Task struct {
	Id        int64     `json:"taskID"`
	ObjectId  string    `json:"ObjectId"`
	ObjectIds []string  `json:"objectIDs"`
	UpdatedAt time.Time `json:"updatedAt"`

	index Index
}

// Wait blocks until task status = "published"
func (t *Task) Wait() error {
	// check first
	isPub, err := t.IsPublished()
	if err != nil {
		return err
	}
	if isPub {
		return nil
	}

	pollingInterval := time.Millisecond * 100
	timeout := time.After(3 * time.Second)
	for {
		select {
		case <-time.After(pollingInterval):
			isPub, err := t.IsPublished()
			if err != nil {
				return err
			}
			if isPub {
				return nil
			}
		case <-timeout:
			return fmt.Errorf("Wait timeout")
		}
	}
}

// IsPublished hits the algolia api to check if the task is published
func (t *Task) IsPublished() (bool, error) {
	status, err := t.GetStatus()
	if err != nil {
		return false, err
	}

	return status.IsPublished(), nil
}

// GetStatus checks
func (t *Task) GetStatus() (*TaskStatus, error) {
	return t.index.GetTaskStatus(t.Id)
}

func randomTask(idx Index) *Task {
	return &Task{
		Id:    random.Int64(1, 9999999999),
		index: idx,
	}

}

func NewTask(idx Index, v Value) (*Task, error) {
	task := &Task{
		index: idx,
	}

	return task, v.Scan(task)
}
