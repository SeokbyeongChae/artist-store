package response

// type ApiSusseccResponse[T any] struct {
// 	Code int `json:"code"`
// 	Data T   `json:"data"`
// }

// func BuildSuccessResponse[T any](code int, message string, data T) (int, ApiSusseccResponse[T]) {
// 	return code, ApiSusseccResponse[T]{
// 		Code: code,
// 		Data: data,
// 	}
// }

// type ApiErrorResponse struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// }

// func BuildErrorResponse(code int, err error) (int, ApiErrorResponse) {
// 	return code, ApiErrorResponse{
// 		Code:    code,
// 		Message: err.Error(),
// 	}
// }

/*
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
*/
