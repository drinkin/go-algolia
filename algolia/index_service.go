package algolia

import (
	"fmt"
	"path/filepath"
	"strconv"
)

type IndexService struct {
	name    string
	service *Service
}

func (idx *IndexService) pathFor(parts ...string) string {
	return fmt.Sprintf("/1/indexes/%s/%s", idx.name, filepath.Join(parts...))
}

func (idx *IndexService) Name() string {
	return idx.name
}

func (idx *IndexService) GetTaskStatus(taskId int64) (*TaskStatus, error) {
	tr := new(TaskStatus)

	return tr, idx.service.Get(idx.pathFor("task", strconv.FormatInt(taskId, 10))).Scan(tr)
}

func (idx *IndexService) UpdateObject(obj Indexable) (*Task, error) {
	obj.AlgoliaBeforeIndex()
	v := idx.service.Put(idx.pathFor(obj.AlgoliaId()), obj)

	return NewTask(idx, v)
}

func (idx *IndexService) BatchUpdate(objs []Indexable) (*Task, error) {
	requests := make([]*BatchItem, len(objs))

	for i, obj := range objs {
		obj.AlgoliaBeforeIndex()
		requests[i] = &BatchItem{
			Action: "updateObject",
			Body:   obj,
		}
	}
	v := idx.service.Post(idx.pathFor("batch"), map[string]interface{}{
		"requests": requests,
	})

	return NewTask(idx, v)
}

func (idx *IndexService) GetObject(id string, attrs ...string) Value {
	return idx.service.Get(idx.pathFor(id))
}

func (idx *IndexService) DeleteObject(id string) Value {
	return idx.service.Delete(idx.pathFor(id))
}

func (idx *IndexService) SetSettings(s *Settings) (*Task, error) {
	v := idx.service.Put(idx.pathFor("settings"), s)
	return NewTask(idx, v)
}

func (idx *IndexService) Clear() (*Task, error) {
	v := idx.service.Post(idx.pathFor("clear"), nil)
	return NewTask(idx, v)
}

func (idx *IndexService) Settings() *SettingsBuilder {
	return NewSettingsBuilder(idx)
}
