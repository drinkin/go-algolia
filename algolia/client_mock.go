package algolia

type ClientMock struct {
	indexes map[string]Index
}

func NewClientMock() *ClientMock {
	return &ClientMock{
		indexes: make(map[string]Index),
	}
}

func (c *ClientMock) Index(name string) Index {
	if idx, ok := c.indexes[name]; ok {
		return idx
	}
	idx := NewIndexMock(name)
	c.indexes[name] = idx

	return idx
}
