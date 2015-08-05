package algolia

import (
	"encoding/json"
	"sync"
	"time"
)

type MockValue struct {
	b []byte
}

func NewMockValue(b []byte) *MockValue {
	return &MockValue{b}
}

func (v *MockValue) Scan(obj interface{}) error {
	return json.Unmarshal(v.b, obj)
}

type IndexMock struct {
	mu   sync.RWMutex
	name string

	objects  map[string][]byte
	tasks    map[int64]bool
	settings *Settings
}

func NewIndexMock(name string) *IndexMock {
	return &IndexMock{
		name:    name,
		objects: make(map[string][]byte),
		tasks:   make(map[int64]bool),
	}
}

func (idx *IndexMock) Name() string {
	return idx.name
}

func (idx *IndexMock) GetTaskStatus(taskId int64) (*TaskStatus, error) {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if _, ok := idx.tasks[taskId]; ok {
		return &TaskStatus{Status: "published", Pending: false}, nil
	}
	return &TaskStatus{Status: "notPublished", Pending: false}, nil
}

func (idx *IndexMock) doUpdate(obj Indexable) error {
	obj.AlgoliaBeforeIndex()
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	idx.objects[obj.AlgoliaId()] = b
	return nil
}

func (idx *IndexMock) UpdateObject(obj Indexable) (*Task, error) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	if err := idx.doUpdate(obj); err != nil {
		return nil, err
	}

	task := randomTask(idx)
	task.ObjectId = obj.AlgoliaId()
	task.UpdatedAt = time.Now()

	idx.tasks[task.Id] = true

	return task, nil
}

func (idx *IndexMock) BatchUpdate(objs []Indexable) (*Task, error) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	task := randomTask(idx)
	for _, obj := range objs {
		idx.doUpdate(obj)
		task.ObjectIds = append(task.ObjectIds, obj.AlgoliaId())
	}
	return task, nil
}

func (idx *IndexMock) GetObject(id string, attrs ...string) Value {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	b, ok := idx.objects[id]
	if !ok {
		return NewErrValue(&Err{
			Message: "ObjectID does not exist",
			Status:  404,
		})
	}

	return NewMockValue(b)
}

func (idx *IndexMock) SetSettings(s *Settings) (*Task, error) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	idx.settings = s
	return randomTask(idx), nil
}

func (idx *IndexMock) Clear() (*Task, error) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	idx.objects = make(map[string][]byte)
	return randomTask(idx), nil
}

func (idx *IndexMock) Settings() *SettingsBuilder {
	return NewSettingsBuilder(idx)
}
