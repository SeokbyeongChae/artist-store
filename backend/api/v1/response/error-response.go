package response

type ApiErrorResponse struct {
	Result  bool   `json:"result"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BuildErrorResponse(code int, err error) (int, ApiErrorResponse) {
	return code, ApiErrorResponse{
		Result:  false,
		Code:    code,
		Message: err.Error(),
	}
}
