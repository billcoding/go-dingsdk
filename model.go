package ding

type authorizeResp struct {
	ErrCode     int
	ErrMsg      string
	AccessToken string `json:"access_token"`
}

type asyncSendResp struct {
	ErrCode int
	ErrMsg  string
	TaskId  int `json:"task_id"`
}
