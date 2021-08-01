package src

import "net/http"

type MessageItem struct {
	EN []string `json:"en"`
	FA []string `json:"fa"`
}
const ErrorResponseResult string = "ERROR"
type ResponseData []interface{}
type ApiResponse struct {
	StatusCode int                    `json:"status_code"`
	Result   string                 `json:"result"`
	Messages map[string]MessageItem `json:"messages"`
	Data     ResponseData           `json:"data"`
}
func GetNotFoundErrorResponse (enResource string, faResource string) ApiResponse {
	messages := map[string]MessageItem{
		enResource: {
			EN: []string{enResource + " not found!"},
			FA: []string{faResource + " یافت نشد"},
		},
	}
	return ApiResponse{
		StatusCode: http.StatusNotFound,
		Result:     ErrorResponseResult,
		Messages:   messages,
		Data:       ResponseData{},
	}
}