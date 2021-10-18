package qcmobile

// CarrierDetails -
type CarrierDetails struct {
	Carrier *Carrier `json:"carrier"`
}

// Carrier -
type Carrier struct {
	AllowedToOperate              string           `json:"allowedToOperate"`
	BIPDInsuranceOnFile           string           `json:"bipdInsuranceOnFile"`
	BIPDInsuranceRequired         string           `json:"bipdInsuranceRequired"`
	BIPDRequiredAmount            string           `json:"bipdRequiredAmount"`
	BondInsuranceOnFile           string           `json:"bondInsuranceOnFile"`
	BondInsuranceRequired         string           `json:"bondInsuranceRequired"`
	BrokerAuthorityStatus         string           `json:"brokerAuthorityStatus"`
	CargoInsuranceOnFile          string           `json:"cargoInsuranceOnFile"`
	CargoInsuranceRequired        string           `json:"cargoInsuranceRequired"`
	CarrierOperation              CarrierOperation `json:"carrierOperation"`
	CensusTypeID                  CensusTypeID     `json:"censusTypeId"`
	CommonAuthorityStatus         string           `json:"commonAuthorityStatus"`
	ContractAuthorityStatus       string           `json:"contractAuthorityStatus"`
	CrashTotal                    int              `json:"crashTotal"`
	DBAName                       string           `json:"dbaName"`
	DOTNumber                     int              `json:"dotNumber"`
	DriverInspections             int              `json:"driverInsp"`
	DriverOOSInspections          int              `json:"driverOosInsp"`
	DriverOOSRate                 float64          `json:"driverOosRate"`
	DriverOOSRateNationalAverage  string           `json:"driverOosRateNationalAverage"`
	EIN                           int              `json:"ein"`
	FatalCrash                    int              `json:"fatalCrash"`
	HazmatInspections             int              `json:"hazmatInsp"`
	HazmatOOSInspections          int              `json:"hazmatOosInsp"`
	HazmatOOSRate                 float64          `json:"hazmatOosRate"`
	HazmatOOSRateNationalAverage  string           `json:"hazmatOosRateNationalAverage"`
	InjCrash                      int              `json:"injCrash"`
	IsPassengerCarrier            string           `json:"isPassengerCarrier"`
	ISSScore                      interface{}      `json:"issScore"`
	LegalName                     string           `json:"legalName"`
	MCS150Outdated                string           `json:"mcs150Outdated"`
	OOSDate                       Date             `json:"oosDate"`
	OOSRateNationalAverageYear    string           `json:"oosRateNationalAverageYear"`
	PhyCity                       string           `json:"phyCity"`
	PhyCountry                    string           `json:"phyCountry"`
	PhyState                      string           `json:"phyState"`
	PhyStreet                     string           `json:"phyStreet"`
	PhyZipcode                    string           `json:"phyZipcode"`
	ReviewDate                    Date             `json:"reviewDate"`
	ReviewType                    string           `json:"reviewType"`
	SafetyRating                  string           `json:"safetyRating"`
	SafetyRatingDate              Date             `json:"safetyRatingDate"`
	SafetyReviewDate              Date             `json:"safetyReviewDate"`
	SafetyReviewType              string           `json:"safetyReviewType"`
	SnapshotDate                  Date             `json:"snapshotDate"`
	StatusCode                    string           `json:"statusCode"`
	TotalDrivers                  int              `json:"totalDrivers"`
	TotalPowerUnits               int              `json:"totalPowerUnits"`
	TowAwayCrash                  int              `json:"towawayCrash"`
	VehicleInspections            int              `json:"vehicleInsp"`
	VehicleOOSInspections         int              `json:"vehicleOosInsp"`
	VehicleOOSRate                float64          `json:"vehicleOosRate"`
	VehicleOOSRateNationalAverage string           `json:"vehicleOosRateNationalAverage"`
}

// CarrierOperation -
type CarrierOperation struct {
	CarrierOperationCode string `json:"carrierOperationCode"`
	CarrierOperationDesc string `json:"carrierOperationDesc"`
}

// CensusTypeID -
type CensusTypeID struct {
	CensusType     string `json:"censusType"`
	CensusTypeDesc string `json:"censusTypeDesc"`
	CensusTypeID   int    `json:"censusTypeId"`
}
