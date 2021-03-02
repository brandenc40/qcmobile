package entities

// CarrierDetails
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
	DbaName                       string           `json:"dbaName"`
	DotNumber                     int              `json:"dotNumber"`
	DriverInspections             int              `json:"driverInsp"`
	DriverOOSInspections          int              `json:"driverOosInsp"`
	DriverOOSRate                 float64          `json:"driverOosRate"`
	DriverOOSRateNationalAverage  string           `json:"driverOosRateNationalAverage"`
	EIN                           int              `json:"ein"`
	FatalCrash                    int              `json:"fatalCrash"`
	HazmatInspections             int              `json:"hazmatInsp"`
	HazmatOOSInspections          int              `json:"hazmatOosInsp"`
	HazmatOOSRate                 int              `json:"hazmatOosRate"`
	HazmatOOSRateNationalAverage  string           `json:"hazmatOosRateNationalAverage"`
	InjCrash                      int              `json:"injCrash"`
	IsPassengerCarrier            string           `json:"isPassengerCarrier"`
	IssScore                      interface{}      `json:"issScore"`
	LegalName                     string           `json:"legalName"`
	Mcs150Outdated                string           `json:"mcs150Outdated"`
	OosDate                       Date             `json:"oosDate"`
	OosRateNationalAverageYear    string           `json:"oosRateNationalAverageYear"`
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
	TowawayCrash                  int              `json:"towawayCrash"`
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
