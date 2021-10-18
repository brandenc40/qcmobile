package qcmobile

// Docket -
type Docket struct {
	DocketNumber   int    `json:"docketNumber"`
	DocketNumberID int    `json:"docketNumberId"`
	DOTNumber      int    `json:"dotNumber"`
	Prefix         string `json:"prefix"`
}
