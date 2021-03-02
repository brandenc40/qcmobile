package entities

// AuthorityDetails -
type AuthorityDetails struct {
	CarrierAuthority CarrierAuthority `json:"carrierAuthority"`
}

// CarrierAuthority -
type CarrierAuthority struct {
	ApplicantID                 int    `json:"applicantID"`
	Authority                   string `json:"authority"`
	AuthorizedForBroker         string `json:"authorizedForBroker"`
	AuthorizedForHouseholdGoods string `json:"authorizedForHouseholdGoods"`
	AuthorizedForPassenger      string `json:"authorizedForPassenger"`
	AuthorizedForProperty       string `json:"authorizedForProperty"`
	BrokerAuthorityStatus       string `json:"brokerAuthorityStatus"`
	CommonAuthorityStatus       string `json:"commonAuthorityStatus"`
	ContractAuthorityStatus     string `json:"contractAuthorityStatus"`
	DocketNumber                int    `json:"docketNumber"`
	DotNumber                   int    `json:"dotNumber"`
	Prefix                      string `json:"prefix"`
}
