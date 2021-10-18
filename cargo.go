package qcmobile

// CargoClass -
type CargoClass struct {
	CargoClassDesc string  `json:"cargoClassDesc"`
	ID             CargoID `json:"id"`
}

// CargoID -
type CargoID struct {
	CargoClassID int `json:"cargoClassId"`
	DOTNumber    int `json:"dotNumber"`
}
