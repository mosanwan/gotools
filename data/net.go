package data

type NetResult struct {
	Success bool        `json:"success,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Msg     string      `json:"msg,omitempty"`
}
