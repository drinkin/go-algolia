package algolia

import "fmt"

type ClientMock struct {
	indexes map[string]Index
	prefix  string
}

func NewClientMock() *ClientMock {
	return &ClientMock{
		indexes: make(map[string]Index),
	}
}

func (c *ClientMock) SetIndexPrefix(p string) {
	c.prefix = p
}

func (c *ClientMock) Index(n string) Index {
	name := fmt.Sprintf("%s%s", c.prefix, n)
	if idx, ok := c.indexes[name]; ok {
		return idx
	}
	idx := NewIndexMock(name)
	c.indexes[name] = idx

	return idx
}
