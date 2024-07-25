package responses

type Response interface {
	Err() error
	StatusCode() int
	Message() string
}

type GenericResponse struct {
	Msg    string
	Error  error
	Status int
}

func (r GenericResponse) Message() string { return r.Msg }
func (r GenericResponse) Err() error      { return r.Error }
func (r GenericResponse) StatusCode() int { return r.Status }
