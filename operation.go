package qcmobile

// OperationClass -
type OperationClass struct {
	OperationClassDesc string      `json:"operationClassDesc"`
	ID                 OperationID `json:"id"`
}

// OperationID -
type OperationID struct {
	DotNumber        int `json:"dotNumber"`
	OperationClassID int `json:"operationClassId"`
}
