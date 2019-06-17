package base

type BaseResponseParam struct {
	ResponseCode      string `json:"response_code"`
	Message  string `json:"message"`
	ResponseData string `json:"response_data"`
}