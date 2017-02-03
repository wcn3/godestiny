package bungie

type ActivityHistory struct {
	ErrorCode   int64    `json:"ErrorCode"`
	ErrorStatus string   `json:"ErrorStatus"`
	Message     string   `json:"Message"`
	MessageData struct{} `json:"MessageData"`
	Response    struct {
		Data struct {
			Activities []struct {
				ActivityDetails struct {
					ActivityTypeHashOverride int64  `json:"activityTypeHashOverride"`
					InstanceID               string `json:"instanceId"`
					IsPrivate                bool   `json:"isPrivate"`
					Mode                     int64  `json:"mode"`
					ReferenceID              int64  `json:"referenceId"`
				} `json:"activityDetails"`
				Period string `json:"period"`
				Values struct {
					ActivityDurationSeconds struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"activityDurationSeconds"`
					Assists struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"assists"`
					AverageScorePerKill struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"averageScorePerKill"`
					AverageScorePerLife struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"averageScorePerLife"`
					Completed struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"completed"`
					CompletionReason struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"completionReason"`
					Deaths struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"deaths"`
					FireTeamID struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"fireTeamId"`
					Kills struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"kills"`
					KillsDeathsAssists struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"killsDeathsAssists"`
					KillsDeathsRatio struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"killsDeathsRatio"`
					LeaveRemainingSeconds struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"leaveRemainingSeconds"`
					PlayerCount struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"playerCount"`
					Score struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"score"`
					Standing struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"standing"`
					Team struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"team"`
					TeamScore struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
						StatID string `json:"statId"`
					} `json:"teamScore"`
				} `json:"values"`
			} `json:"activities"`
		} `json:"data"`
	} `json:"Response"`
	ThrottleSeconds int64 `json:"ThrottleSeconds"`
}
