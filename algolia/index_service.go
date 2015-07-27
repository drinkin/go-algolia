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

func (idx *IndexService) Must() *MustIndex {
	return &MustIndex{idx}
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
	tr := new(Task)
	return tr, idx.service.Put(idx.pathFor(obj.AlgoliaId()), obj).Scan(tr)
}

func (idx *IndexService) BatchUpdate(objs []Indexable) (*BatchTask, error) {
	requests := make([]*BatchItem, len(objs))

	for i, obj := range objs {
		obj.AlgoliaBeforeIndex()
		requests[i] = &BatchItem{
			Action: "updateObject",
			Body:   obj,
		}
	}
	tr := new(BatchTask)
	return tr, idx.service.Post(idx.pathFor("batch"), map[string]interface{}{
		"requests": requests,
	}).Scan(tr)

}

func (idx *IndexService) GetObject(id string, attrs ...string) Value {
	return idx.service.Get(idx.pathFor(id))
}
