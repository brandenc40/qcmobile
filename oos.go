package qcmobile

type OOSDetails struct {
	Oos *Oos `json:"oos"`
}
type Oos struct {
	DotNumber            int    `json:"dotNumber"`
	ID                   int    `json:"id"`
	OosDate              Date   `json:"oosDate"`
	OosReason            string `json:"oosReason"`
	OosReasonDescription string `json:"oosReasonDescription"`
}