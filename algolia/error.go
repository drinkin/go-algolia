package algolia

type Err struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (err *Err) Error() string {
	return err.Message
}
