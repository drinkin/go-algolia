package algolia

import "time"

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
