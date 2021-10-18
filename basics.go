package qcmobile

// BasicsDetails -
type BasicsDetails struct {
	Basic     Basic  `json:"basic"`
	DOTNumber string `json:"dotNumber"`
}

// Basic -
type Basic struct {
	BasicsPercentile                                      string     `json:"basicsPercentile"`
	BasicsRunDate                                         Timestamp  `json:"basicsRunDate"`
	BasicsType                                            BasicsType `json:"basicsType"`
	BasicsViolationThreshold                              string     `json:"basicsViolationThreshold"`
	ExceededFMCSAInterventionThreshold                    string     `json:"exceededFMCSAInterventionThreshold"`
	ID                                                    BasicsID   `json:"id"`
	MeasureValue                                          string     `json:"measureValue"`
	OnRoadPerformanceThresholdViolationIndicator          string     `json:"onRoadPerformanceThresholdViolationIndicator"`
	SeriousViolationFromInvestigationPast12MonthIndicator string     `json:"seriousViolationFromInvestigationPast12MonthIndicator"`
	TotalInspectionWithViolation                          int        `json:"totalInspectionWithViolation"`
	TotalViolation                                        int        `json:"totalViolation"`
}

// BasicsType -
type BasicsType struct {
	BasicsCode      string `json:"basicsCode"`
	BasicsCodeMCMIS string `json:"basicsCodeMcmis"`
	BasicsID        int    `json:"basicsId"`
	BasicsLongDesc  string `json:"basicsLongDesc"`
	BasicsShortDesc string `json:"basicsShortDesc"`
}

// BasicsID -
type BasicsID struct {
	BasicsID  int `json:"basicsId"`
	DotNumber int `json:"dotNumber"`
}
