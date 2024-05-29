package http

func newBodyError(err error) map[string]string {
	return map[string]string{
		"message": err.Error(),
	}
}
