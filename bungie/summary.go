package bungie

type SummaryResponse struct {
	ErrorCode   int64    `json:"ErrorCode"`
	ErrorStatus string   `json:"ErrorStatus"`
	Message     string   `json:"Message"`
	MessageData struct{} `json:"MessageData"`
	Response    struct {
		Data struct {
			Characters    []CharacterData `json:"characters"`
			GrimoireScore int64           `json:"grimoireScore"`
			Inventory     struct {
				Currencies []struct {
					ItemHash int64 `json:"itemHash"`
					Value    int64 `json:"value"`
				} `json:"currencies"`
				Items []interface{} `json:"items"`
			} `json:"inventory"`
			MembershipID   string `json:"membershipId"`
			MembershipType int64  `json:"membershipType"`
			Versions       int64  `json:"versions"`
		} `json:"data"`
	} `json:"Response"`
	ThrottleSeconds int64 `json:"ThrottleSeconds"`
}

type CharacterData struct {
	BackgroundPath     string `json:"backgroundPath"`
	BaseCharacterLevel int64  `json:"baseCharacterLevel"`
	CharacterBase      struct {
		BuildStatGroupHash  int64  `json:"buildStatGroupHash"`
		CharacterID         string `json:"characterId"`
		ClassHash           int64  `json:"classHash"`
		ClassType           int64  `json:"classType"`
		CurrentActivityHash int64  `json:"currentActivityHash"`
		Customization       struct {
			DecalColor   int64 `json:"decalColor"`
			DecalIndex   int64 `json:"decalIndex"`
			EyeColor     int64 `json:"eyeColor"`
			Face         int64 `json:"face"`
			FeatureColor int64 `json:"featureColor"`
			FeatureIndex int64 `json:"featureIndex"`
			HairColor    int64 `json:"hairColor"`
			HairIndex    int64 `json:"hairIndex"`
			LipColor     int64 `json:"lipColor"`
			Personality  int64 `json:"personality"`
			SkinColor    int64 `json:"skinColor"`
			WearHelmet   bool  `json:"wearHelmet"`
		} `json:"customization"`
		DateLastPlayed           string `json:"dateLastPlayed"`
		GenderHash               int64  `json:"genderHash"`
		GenderType               int64  `json:"genderType"`
		GrimoireScore            int64  `json:"grimoireScore"`
		LastCompletedStoryHash   int64  `json:"lastCompletedStoryHash"`
		MembershipID             string `json:"membershipId"`
		MembershipType           int64  `json:"membershipType"`
		MinutesPlayedThisSession string `json:"minutesPlayedThisSession"`
		MinutesPlayedTotal       string `json:"minutesPlayedTotal"`
		PeerView                 struct {
			Equipment []struct {
				Dyes []struct {
					ChannelHash int64 `json:"channelHash"`
					DyeHash     int64 `json:"dyeHash"`
				} `json:"dyes"`
				ItemHash int64 `json:"itemHash"`
			} `json:"equipment"`
		} `json:"peerView"`
		PowerLevel int64 `json:"powerLevel"`
		RaceHash   int64 `json:"raceHash"`
		Stats      struct {
			StatAgility struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_AGILITY"`
			StatArmor struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_ARMOR"`
			StatDefense struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_DEFENSE"`
			StatDiscipline struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_DISCIPLINE"`
			StatIntellect struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_INTELLECT"`
			StatLight struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_LIGHT"`
			StatOptics struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_OPTICS"`
			StatRecovery struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_RECOVERY"`
			StatStrength struct {
				MaximumValue int64 `json:"maximumValue"`
				StatHash     int64 `json:"statHash"`
				Value        int64 `json:"value"`
			} `json:"STAT_STRENGTH"`
		} `json:"stats"`
	} `json:"characterBase"`
	CharacterLevel   int64  `json:"characterLevel"`
	EmblemHash       int64  `json:"emblemHash"`
	EmblemPath       string `json:"emblemPath"`
	IsPrestigeLevel  bool   `json:"isPrestigeLevel"`
	LevelProgression struct {
		CurrentProgress     int64 `json:"currentProgress"`
		DailyProgress       int64 `json:"dailyProgress"`
		Level               int64 `json:"level"`
		NextLevelAt         int64 `json:"nextLevelAt"`
		ProgressToNextLevel int64 `json:"progressToNextLevel"`
		ProgressionHash     int64 `json:"progressionHash"`
		Step                int64 `json:"step"`
		WeeklyProgress      int64 `json:"weeklyProgress"`
	} `json:"levelProgression"`
	PercentToNextLevel float64 `json:"percentToNextLevel"`
}
