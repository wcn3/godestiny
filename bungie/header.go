package bungie

type Header struct {
	ErrorCode       int64    `json:"ErrorCode"`
	ErrorStatus     string   `json:"ErrorStatus"`
	Message         string   `json:"Message"`
	MessageData     struct{} `json:"MessageData"`
	ThrottleSeconds int64    `json:"ThrottleSeconds"`
}
