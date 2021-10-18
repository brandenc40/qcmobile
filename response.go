package qcmobile

type errorResponse struct {
	ErrMsg string `json:"content"`
}

type searchResponse struct {
	Content     []*CarrierDetails `json:"content"`
	RetrievedAt Timestamp         `json:"retrievalDate"`
}

type getCarriersByDocketResponse struct {
	Content     []*CarrierDetails `json:"content"`
	RetrievedAt Timestamp         `json:"retrievalDate"`
}

type carrierResponse struct {
	Content     *CarrierDetails `json:"content"`
	RetrievedAt Timestamp       `json:"retrievalDate"`
}

type cargoCarriedResponse struct {
	Content     []*CargoClass `json:"content"`
	RetrievedAt Timestamp     `json:"retrievalDate"`
}

type operationClassificationResponse struct {
	Content     []*OperationClass `json:"content"`
	RetrievedAt Timestamp         `json:"retrievalDate"`
}

type docketNumbersResponse struct {
	Content     []*Docket `json:"content"`
	RetrievedAt Timestamp `json:"retrievalDate"`
}

type authorityResponse struct {
	Content     []*AuthorityDetails `json:"content"`
	RetrievedAt Timestamp           `json:"retrievalDate"`
}

type basicsResponse struct {
	Content     []*BasicsDetails `json:"content"`
	RetrievedAt Timestamp        `json:"retrievalDate"`
}

type oosResponse struct {
	Content     []*OOSDetails `json:"content"`
	RetrievedAt Timestamp     `json:"retrievalDate"`
}
