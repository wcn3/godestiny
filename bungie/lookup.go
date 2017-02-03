package bungie

type LookupResponse struct {
	ErrorCode   int64    `json:"ErrorCode"`
	ErrorStatus string   `json:"ErrorStatus"`
	Message     string   `json:"Message"`
	MessageData struct{} `json:"MessageData"`
	Response    []struct {
		DisplayName    string `json:"displayName"`
		IconPath       string `json:"iconPath"`
		MembershipID   string `json:"membershipId"`
		MembershipType int64  `json:"membershipType"`
	} `json:"Response"`
	ThrottleSeconds int64 `json:"ThrottleSeconds"`
}
