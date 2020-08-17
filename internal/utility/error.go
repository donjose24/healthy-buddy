package utility

type HttpError struct {
	StatusCode int
	Message    string
	Errors     map[string]string
}

func (h *HttpError) Error() string {
	return h.Message
}
