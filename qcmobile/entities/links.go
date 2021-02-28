package entities

// Links -
type Links struct {
	Basics                        Href `json:"basics"`
	CargoCarried                  Href `json:"cargo carried"`
	OperationClassification       Href `json:"operation classification"`
	DocketNumbers                 Href `json:"docket numbers"`
	CarrierActiveForHireAuthority Href `json:"carrier active-For-hire authority"`
	Self                          Href `json:"self"`
}

// Href -
type Href struct {
	Href string `json:"href"`
}
