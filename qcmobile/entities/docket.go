package entities

// Docket -
type Docket struct {
	DocketNumber   int    `json:"docketNumber"`
	DocketNumberID int    `json:"docketNumberId"`
	DotNumber      int    `json:"dotNumber"`
	Prefix         string `json:"prefix"`
}
