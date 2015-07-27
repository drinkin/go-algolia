package algolia

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
