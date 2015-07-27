package algolia

type IndexMock struct {
}

func (idx *IndexMock) UpdateObject(obj Indexable) (*Task, error) {
	return nil, nil
}

func (idx *IndexMock) BatchUpdate(objs []Indexable) (*BatchTask, error) {
	return nil, nil

}

func (idx *IndexMock) GetObject(id string, attrs ...string) Value {
	return nil
}
