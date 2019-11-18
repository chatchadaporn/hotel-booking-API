package commons

type (
	Response struct {
		HttpStatus int64       `json:"status"`
		Body       interface{} `json:"body"`
	}

	SuccessResponse struct {
		Message string      `json:"message"`
		Detail  interface{} `json:"detail,omitempty"`
	}

	ErrorResponse struct {
		ErrorCode    int64  `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	}
)

func (r *Response) Success(message string, detail interface{}) {
	r.HttpStatus = 200
	r.Body = SuccessResponse{
		Message: message,
		Detail:  detail,
	}
}

func (r *Response) Fail(httpCode int64, code int64, message string) {
	r.HttpStatus = httpCode
	r.Body = ErrorResponse{
		ErrorCode:    code,
		ErrorMessage: message,
	}
}

func (r *Response) NotFound() {
	r.HttpStatus = 404
	r.Body = ErrorResponse{
		ErrorCode:    404,
		ErrorMessage: "api not found",
	}
}
