package bungie

type PGCR struct {
	ErrorCode   int64    `json:"ErrorCode"`
	ErrorStatus string   `json:"ErrorStatus"`
	Message     string   `json:"Message"`
	MessageData struct{} `json:"MessageData"`
	Response    struct {
		Data struct {
			ActivityDetails struct {
				ActivityTypeHashOverride int64  `json:"activityTypeHashOverride"`
				InstanceID               string `json:"instanceId"`
				IsPrivate                bool   `json:"isPrivate"`
				Mode                     int64  `json:"mode"`
				ReferenceID              int64  `json:"referenceId"`
			} `json:"activityDetails"`
			Entries []struct {
				CharacterID string `json:"characterId"`
				Extended    struct {
					Values struct {
						AllParticipantsCount struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"allParticipantsCount"`
						AllParticipantsScore struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"allParticipantsScore"`
						AllParticipantsTimePlayed struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"allParticipantsTimePlayed"`
						AverageKillDistance struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"averageKillDistance"`
						AverageLifespan struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"averageLifespan"`
						AverageScorePerKill struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"averageScorePerKill"`
						AverageScorePerLife struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"averageScorePerLife"`
						Deaths struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"deaths"`
						DeathsFromAiExpUltraTakenTitanArenaChallengeReef struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"deathsFromAiExpUltraTakenTitanArenaChallengeReef"`
						DeathsFromTaken struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"deathsFromTaken"`
						DeathsFromTakenThrall struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"deathsFromTakenThrall"`
						FireTeamID struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"fireTeamId"`
						Kills struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"kills"`
						KillsOfAiExpUltraTakenTitanArenaChallengeReef struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"killsOfAiExpUltraTakenTitanArenaChallengeReef"`
						KillsOfTaken struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"killsOfTaken"`
						KillsOfTakenGoblin struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"killsOfTakenGoblin"`
						KillsOfTakenHobgoblin struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"killsOfTakenHobgoblin"`
						KillsOfTakenPhalanx struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"killsOfTakenPhalanx"`
						KillsOfTakenPsion struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"killsOfTakenPsion"`
						KillsOfTakenThrall struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"killsOfTakenThrall"`
						KillsOfTakenVandal struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"killsOfTakenVandal"`
						LongestKillSpree struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"longestKillSpree"`
						LongestSingleLife struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"longestSingleLife"`
						OrbsDropped struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"orbsDropped"`
						PrecisionKillOfAiExpUltraTakenTitanArenaChallengeReef struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"precisionKillOfAiExpUltraTakenTitanArenaChallengeReef"`
						PrecisionKillOfTakenGoblin struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"precisionKillOfTakenGoblin"`
						PrecisionKillOfTakenHobgoblin struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"precisionKillOfTakenHobgoblin"`
						PrecisionKillOfTakenPhalanx struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"precisionKillOfTakenPhalanx"`
						PrecisionKillOfTakenPsion struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"precisionKillOfTakenPsion"`
						PrecisionKillOfTakenThrall struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"precisionKillOfTakenThrall"`
						PrecisionKillOfTakenVandal struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"precisionKillOfTakenVandal"`
						PrecisionKills struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"precisionKills"`
						Score struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"score"`
						SecondsPlayed struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"secondsPlayed"`
						TotalActivityDurationSeconds struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"totalActivityDurationSeconds"`
						TotalKillDistance struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"totalKillDistance"`
						WeaponBestType struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"weaponBestType"`
						WeaponKillsGrenade struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"weaponKillsGrenade"`
						WeaponKillsMachinegun struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"weaponKillsMachinegun"`
						WeaponKillsMelee struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"weaponKillsMelee"`
						WeaponKillsPrecisionKillsMachinegun struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"weaponKillsPrecisionKillsMachinegun"`
						WeaponKillsPrecisionKillsPulseRifle struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"weaponKillsPrecisionKillsPulseRifle"`
						WeaponKillsPulseRifle struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"weaponKillsPulseRifle"`
						WeaponPrecisionKillsMachinegun struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"weaponPrecisionKillsMachinegun"`
						WeaponPrecisionKillsPulseRifle struct {
							Basic struct {
								DisplayValue string  `json:"displayValue"`
								Value        float64 `json:"value"`
							} `json:"basic"`
						} `json:"weaponPrecisionKillsPulseRifle"`
					} `json:"values"`
					Weapons []struct {
						ReferenceID int64 `json:"referenceId"`
						Values      struct {
							UniqueWeaponKills struct {
								Basic struct {
									DisplayValue string  `json:"displayValue"`
									Value        float64 `json:"value"`
								} `json:"basic"`
							} `json:"uniqueWeaponKills"`
							UniqueWeaponKillsPrecisionKills struct {
								Basic struct {
									DisplayValue string  `json:"displayValue"`
									Value        float64 `json:"value"`
								} `json:"basic"`
							} `json:"uniqueWeaponKillsPrecisionKills"`
							UniqueWeaponPrecisionKills struct {
								Basic struct {
									DisplayValue string  `json:"displayValue"`
									Value        float64 `json:"value"`
								} `json:"basic"`
							} `json:"uniqueWeaponPrecisionKills"`
						} `json:"values"`
					} `json:"weapons"`
				} `json:"extended"`
				Player struct {
					BungieNetUserInfo struct {
						DisplayName    string `json:"displayName"`
						IconPath       string `json:"iconPath"`
						MembershipID   string `json:"membershipId"`
						MembershipType int64  `json:"membershipType"`
					} `json:"bungieNetUserInfo"`
					CharacterClass  string `json:"characterClass"`
					CharacterLevel  int64  `json:"characterLevel"`
					DestinyUserInfo struct {
						DisplayName    string `json:"displayName"`
						IconPath       string `json:"iconPath"`
						MembershipID   string `json:"membershipId"`
						MembershipType int64  `json:"membershipType"`
					} `json:"destinyUserInfo"`
					LightLevel int64 `json:"lightLevel"`
				} `json:"player"`
				Score struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
				} `json:"score"`
				Standing int64 `json:"standing"`
				Values   struct {
					ActivityDurationSeconds struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"activityDurationSeconds"`
					Assists struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"assists"`
					AverageScorePerKill struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"averageScorePerKill"`
					AverageScorePerLife struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"averageScorePerLife"`
					Completed struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"completed"`
					CompletionReason struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"completionReason"`
					Deaths struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"deaths"`
					FireTeamID struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"fireTeamId"`
					Kills struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"kills"`
					KillsDeathsAssists struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"killsDeathsAssists"`
					KillsDeathsRatio struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"killsDeathsRatio"`
					LeaveRemainingSeconds struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"leaveRemainingSeconds"`
					PlayerCount struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"playerCount"`
					Score struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"score"`
					TeamScore struct {
						Basic struct {
							DisplayValue string  `json:"displayValue"`
							Value        float64 `json:"value"`
						} `json:"basic"`
					} `json:"teamScore"`
				} `json:"values"`
			} `json:"entries"`
			Period string        `json:"period"`
			Teams  []interface{} `json:"teams"`
		} `json:"data"`
	} `json:"Response"`
	ThrottleSeconds int64 `json:"ThrottleSeconds"`
}
