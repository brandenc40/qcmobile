package entities

// ErrorResponse -
type ErrorResponse struct {
	ErrMsg string `json:"content"`
}

// CarrierDetails -
type CarrierDetails struct {
	Links   *Links   `json:"_links"`
	Carrier *Carrier `json:"carrier"`
}

// SingleCarrierResponse -
type SingleCarrierResponse struct {
	CarrierDetails *CarrierDetails `json:"content"`
	RetrievalDate  TimestampString `json:"retrievalDate"`
}

// MultiCarrierResponse -
type MultiCarrierResponse struct {
	MultiCarrierDetails []*CarrierDetails `json:"content"`
	RetrievalDate       TimestampString   `json:"retrievalDate"`
}
