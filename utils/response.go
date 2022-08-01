package utils

type ResponseStr struct {
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

func Response(data interface{}) *ResponseStr {
	return &ResponseStr{Data: data, Status: "success"}
}

func (r *ResponseStr) SetStatus(status string) *ResponseStr {
	r.Status = status
	return r
}
func (r *ResponseStr) SetMessage(message string) *ResponseStr {
	r.Message = message
	return r
}

func (r *ResponseStr) SetErrors(err interface{}) *ResponseStr {
	r.Status = "error"
	r.Errors = err
	return r
}

func (r *ResponseStr) SetMeta(meta interface{}) *ResponseStr {
	r.Meta = meta
	return r
}
