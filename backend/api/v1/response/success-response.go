package response

type ApiSusseccResponse[T any] struct {
	Result bool `json:"result"`
	Code   int  `json:"code"`
	Data   T    `json:"data"`
}

func BuildSuccessResponse[T any](code int, data T) (int, ApiSusseccResponse[T]) {
	return code, ApiSusseccResponse[T]{
		Result: true,
		Code:   code,
		Data:   data,
	}
}
