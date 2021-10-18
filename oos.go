package qcmobile

// OOSDetails - Out of Service details
type OOSDetails struct {
	OOS OOS `json:"oos"`
}

// OOS - Out of Service
type OOS struct {
	DOTNumber            int    `json:"dotNumber"`
	ID                   int    `json:"id"`
	OOSDate              Date   `json:"oosDate"`
	OOSReason            string `json:"oosReason"`
	OOSReasonDescription string `json:"oosReasonDescription"`
}
