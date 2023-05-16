package httpc

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	StatusBadRequest          = 40000
	StatusInternalServerError = 50000
	StatusOk                  = 20000
)

func Result(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func Success(data interface{}) *Response {
	return &Response{
		Code:    StatusOk,
		Message: "",
		Data:    data,
	}
}

func NewError(code int, message string) *ResponseError {
	return &ResponseError{
		Code:    code,
		Message: message,
	}
}

func Error(code int, err error) *ResponseError {
	switch e := err.(type) {
	case *ResponseError:
		return e
	default:
		return &ResponseError{
			Code:    code,
			Message: e.Error(),
		}
	}
}

func DefaultError(message string) *ResponseError {
	return &ResponseError{
		Code:    StatusInternalServerError,
		Message: message,
	}
}

func (err *ResponseError) Error() string {
	return err.Message
}

func (err *ResponseError) Data() *Response {
	return &Response{
		Code:    err.Code,
		Message: err.Message,
		Data:    nil,
	}
}
