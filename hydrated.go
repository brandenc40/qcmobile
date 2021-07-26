package qcmobile

// CompleteCarrierDetails -
type CompleteCarrierDetails struct {
	Carrier                  *Carrier
	CargosCarried            []*CargoClass
	OperationClassifications []*OperationClass
	Dockets                  []*Docket
	AuthorityDetails         []*AuthorityDetails
	BasicsDetails            []*BasicsDetails
	OOSDetails               []*OOSDetails
}
