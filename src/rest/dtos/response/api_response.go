package response

import "time"

type ApiResponse struct {
	HttpStatus      int       `json:"-"`
	ErrorCode       string    `json:"error_code,omitempty"`
	ErrorMessage    string    `json:"error_message,omitempty"`
	TimeStamp       time.Time `json:"time_stamp,omitempty"`
	Data            any       `json:"data,omitempty"`
	ResponseCode    string    `json:"response_code,omitempty"`
	ResponseMessage string    `json:"response_message,omitempty"`
}
