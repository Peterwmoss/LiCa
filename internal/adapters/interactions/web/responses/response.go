package responses

type Response struct {
	Data       string
	Err        error
	StatusCode int
}

func NewResponse(data string, err error, status int) Response {
	return Response{
		Data:   data,
		Err:    err,
		StatusCode: status,
	}
}
