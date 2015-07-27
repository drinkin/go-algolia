package algolia

import (
	"fmt"
	"time"
)

type Index struct {
	Name    string
	service *Service
}
type BatchTaskResp struct {
	TaskId    int64    `json:"taskID"`
	ObjectIds []string `json:"objectIDs"`
}

type TaskResp struct {
	TaskId    int64     `json:"taskID"`
	ObjectId  string    `json:"ObjectId"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (idx *Index) pathFor(objectId string) string {
	return fmt.Sprintf("/1/indexes/%s/%s", idx.Name, objectId)
}

func (idx *Index) UpdateObject(obj Indexable) (*TaskResp, error) {
	obj.AlgoliaBeforeIndex()
	return idx.service.Put(idx.pathFor(obj.AlgoliaId()), obj).asTaskResp()
}

func (idx *Index) BatchUpdate(objs []Indexable) (*BatchTaskResp, error) {
	requests := make([]*BatchItem, len(objs))

	for i, obj := range objs {
		obj.AlgoliaBeforeIndex()
		requests[i] = &BatchItem{
			Action: "updateObject",
			Body:   obj,
		}
	}
	tr := new(BatchTaskResp)
	return tr, idx.service.Post(idx.pathFor("batch"), map[string]interface{}{
		"requests": requests,
	}).Scan(tr)

}

func (idx *Index) GetObject(id string, attrs ...string) *Value {
	return idx.service.Get(idx.pathFor(id))
}
