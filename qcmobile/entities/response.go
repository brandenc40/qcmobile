package entities

// ErrorResponse -
type ErrorResponse struct {
	ErrMsg string `json:"content"`
}

// SearchResponse -
type SearchResponse struct {
	Content     []*CarrierDetails `json:"content"`
	RetrievedAt Timestamp         `json:"retrievalDate"`
}

// GetCarriersByDocketResponse -
type GetCarriersByDocketResponse struct {
	Content     []*CarrierDetails `json:"content"`
	RetrievedAt Timestamp         `json:"retrievalDate"`
}

// CarrierResponse -
type CarrierResponse struct {
	Content     *CarrierDetails `json:"content"`
	RetrievedAt Timestamp       `json:"retrievalDate"`
}

// CargoCarriedResponse -
type CargoCarriedResponse struct {
	Content     []*CargoClass `json:"content"`
	RetrievedAt Timestamp     `json:"retrievalDate"`
}

// OperationClassificationResponse -
type OperationClassificationResponse struct {
	Content     []*OperationClass `json:"content"`
	RetrievedAt Timestamp         `json:"retrievalDate"`
}

// DocketNumbersResponse -
type DocketNumbersResponse struct {
	Content     []*Docket `json:"content"`
	RetrievedAt Timestamp `json:"retrievalDate"`
}

// AuthorityResponse -
type AuthorityResponse struct {
	Content     []*AuthorityDetails `json:"content"`
	RetrievedAt Timestamp           `json:"retrievalDate"`
}

// BasicsResponse -
type BasicsResponse struct {
	Content     []*BasicsDetails `json:"content"`
	RetrievedAt Timestamp        `json:"retrievalDate"`
}

// OOSResponse -
type OOSResponse struct {
	Content     []*OOSDetails `json:"content"`
	RetrievedAt Timestamp     `json:"retrievalDate"`
}
