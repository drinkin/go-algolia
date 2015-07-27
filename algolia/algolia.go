package algolia

import "time"

// Indexable represents objects that can be saved to the search index.
type Indexable interface {
	AlgoliaId() string

	AlgoliaBeforeIndex()
}

// Index represents a backend.
type Index interface {
	Name() string
	Must() *MustIndex
	GetTaskStatus(taskId int64) (*TaskStatus, error)
	UpdateObject(obj Indexable) (*Task, error)
	BatchUpdate(objs []Indexable) (*BatchTask, error)
	GetObject(id string, attrs ...string) Value
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type MustIndex struct {
	index Index
}

func (m *MustIndex) BatchUpdate(objs []Indexable) *BatchTask {
	ts, err := m.index.BatchUpdate(objs)
	check(err)
	return ts
}

func (m *MustIndex) GetTaskStatus(taskId int64) *TaskStatus {
	ts, err := m.index.GetTaskStatus(taskId)
	check(err)
	return ts
}

type Value interface {
	Scan(obj interface{}) error
}

type TaskStatus struct {
	Status  string `json:"status"`
	Pending bool   `json:pendingTask`
}

func (ts *TaskStatus) IsPublished() bool {
	return ts.Status == "published"
}

type BatchTask struct {
	TaskId    int64    `json:"taskID"`
	ObjectIds []string `json:"objectIDs"`
}

type Task struct {
	TaskId    int64     `json:"taskID"`
	ObjectId  string    `json:"ObjectId"`
	UpdatedAt time.Time `json:"updatedAt"`
}
