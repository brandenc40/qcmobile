package qcmobile

// OOSDetails - Out of Service details
type OOSDetails struct {
	Oos *OOS `json:"oos"`
}

// OOS - Out of Service
type OOS struct {
	DotNumber            int    `json:"dotNumber"`
	ID                   int    `json:"id"`
	OosDate              Date   `json:"oosDate"`
	OosReason            string `json:"oosReason"`
	OosReasonDescription string `json:"oosReasonDescription"`
}
