package qcmobile

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	errRes          = `{"content":"Webkey not found","retrievalDate":"2021-02-28T07:35:25.991+0000","_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc"},"searchByName":{"href":"https://mobile.fmcsa.dot.gov/qc/name/:name"},"lookupBydotNumber":{"href":"https://mobile.fmcsa.dot.gov/qc/id/:dotNumber"}}}`
	carrierRes      = `{"content":{"_links":{"basics":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/basics"},"cargo carried":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/cargo-carried"},"operation classification":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/operation-classification"},"docket numbers":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/docket-numbers"},"carrier active-For-hire authority":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/authority"},"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121"}},"carrier":{"allowedToOperate":"Y","bipdInsuranceOnFile":"1000","bipdInsuranceRequired":"Y","bipdRequiredAmount":"750","bondInsuranceOnFile":"0","bondInsuranceRequired":"u","brokerAuthorityStatus":"N","cargoInsuranceOnFile":"0","cargoInsuranceRequired":"u","carrierOperation":{"carrierOperationCode":"A","carrierOperationDesc":"Interstate"},"censusTypeId":{"censusType":"C","censusTypeDesc":"CARRIER","censusTypeId":1},"commonAuthorityStatus":"A","contractAuthorityStatus":"A","crashTotal":15,"dbaName":null,"dotNumber":158121,"driverInsp":259,"driverOosInsp":3,"driverOosRate":1.15830115830115830115830115830115830116,"driverOosRateNationalAverage":"5.51","ein":391474414,"fatalCrash":0,"hazmatInsp":10,"hazmatOosInsp":1,"hazmatOosRate":10,"hazmatOosRateNationalAverage":"4.5","injCrash":3,"isPassengerCarrier":"N","issScore":null,"legalName":"VERIHA TRUCKING INC","mcs150Outdated":"N","oosDate":null,"oosRateNationalAverageYear":"2009-2010","phyCity":"MARINETTE","phyCountry":"US","phyState":"WI","phyStreet":"2830 CLEVELAND AVE","phyZipcode":"54143","reviewDate":"1996-04-22","reviewType":"C","safetyRating":"S","safetyRatingDate":"1996-04-25","safetyReviewDate":"1996-04-22","safetyReviewType":"C","snapshotDate":null,"statusCode":"A","totalDrivers":213,"totalPowerUnits":213,"towawayCrash":12,"vehicleInsp":124,"vehicleOosInsp":17,"vehicleOosRate":13.70967741935483870967741935483870967742,"vehicleOosRateNationalAverage":"20.72"}},"retrievalDate":"2021-02-28T06:33:31.490+0000"}`
	multiCarrierRes = `{"content":[{"_links":{"basics":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/753076/basics"},"cargo carried":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/753076/cargo-carried"},"operation classification":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/753076/operation-classification"},"docket numbers":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/753076/docket-numbers"},"carrier active-For-hire authority":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/753076/authority"}},"carrier":{"allowedToOperate":"Y","bipdInsuranceOnFile":"0","bipdInsuranceRequired":"Y","bipdRequiredAmount":"750","bondInsuranceOnFile":"0","bondInsuranceRequired":"u","brokerAuthorityStatus":"N","cargoInsuranceOnFile":"0","cargoInsuranceRequired":"u","carrierOperation":{"carrierOperationCode":"C","carrierOperationDesc":"Intrastate Non-Hazmat"},"censusTypeId":{"censusType":"C","censusTypeDesc":"CARRIER","censusTypeId":1},"commonAuthorityStatus":"N","contractAuthorityStatus":"N","crashTotal":0,"dbaName":"LEE VERIHA TRUCKING AND EXCAVATING","dotNumber":753076,"driverInsp":0,"driverOosInsp":0,"driverOosRate":0,"driverOosRateNationalAverage":"5.51","ein":391821630,"fatalCrash":0,"hazmatInsp":0,"hazmatOosInsp":0,"hazmatOosRate":0,"hazmatOosRateNationalAverage":"4.5","injCrash":0,"isPassengerCarrier":null,"issScore":null,"legalName":"AMERICAN ELM SAWMILL INC","mcs150Outdated":"N","oosDate":null,"oosRateNationalAverageYear":"2009-2010","phyCity":"PORTERFIELD","phyCountry":"US","phyState":"WI","phyStreet":"W 3065 VERIHA RD","phyZipcode":"54159","reviewDate":null,"reviewType":null,"safetyRating":null,"safetyRatingDate":null,"safetyReviewDate":null,"safetyReviewType":null,"snapshotDate":null,"statusCode":"A","totalDrivers":2,"totalPowerUnits":2,"towawayCrash":0,"vehicleInsp":0,"vehicleOosInsp":0,"vehicleOosRate":0,"vehicleOosRateNationalAverage":"20.72"}},{"_links":{"basics":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/basics"},"cargo carried":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/cargo-carried"},"operation classification":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/operation-classification"},"docket numbers":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/docket-numbers"},"carrier active-For-hire authority":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/authority"}},"carrier":{"allowedToOperate":"Y","bipdInsuranceOnFile":"1000","bipdInsuranceRequired":"Y","bipdRequiredAmount":"750","bondInsuranceOnFile":"0","bondInsuranceRequired":"u","brokerAuthorityStatus":"N","cargoInsuranceOnFile":"0","cargoInsuranceRequired":"u","carrierOperation":{"carrierOperationCode":"A","carrierOperationDesc":"Interstate"},"censusTypeId":{"censusType":"C","censusTypeDesc":"CARRIER","censusTypeId":1},"commonAuthorityStatus":"A","contractAuthorityStatus":"A","crashTotal":15,"dbaName":null,"dotNumber":158121,"driverInsp":259,"driverOosInsp":3,"driverOosRate":1.15830115830115830115830115830115830116,"driverOosRateNationalAverage":"5.51","ein":391474414,"fatalCrash":0,"hazmatInsp":10,"hazmatOosInsp":1,"hazmatOosRate":10,"hazmatOosRateNationalAverage":"4.5","injCrash":3,"isPassengerCarrier":null,"issScore":null,"legalName":"VERIHA TRUCKING INC","mcs150Outdated":"N","oosDate":null,"oosRateNationalAverageYear":"2009-2010","phyCity":"MARINETTE","phyCountry":"US","phyState":"WI","phyStreet":"2830 CLEVELAND AVE","phyZipcode":"54143","reviewDate":"1996-04-22","reviewType":"C","safetyRating":"S","safetyRatingDate":"1996-04-25","safetyReviewDate":"1996-04-22","safetyReviewType":"C","snapshotDate":null,"statusCode":"A","totalDrivers":213,"totalPowerUnits":213,"towawayCrash":12,"vehicleInsp":124,"vehicleOosInsp":17,"vehicleOosRate":13.70967741935483870967741935483870967742,"vehicleOosRateNationalAverage":"20.72"}}],"retrievalDate":"2021-02-28T07:25:05.638+0000"}`
	cargoRes        = `{"content":[{"cargoClassDesc":"General Freight","id":{"cargoClassId":1,"dotNumber":53467}},{"cargoClassDesc":"Metal; Sheets, Coils, Rolls","id":{"cargoClassId":3,"dotNumber":53467}},{"cargoClassDesc":"Logs, Poles, Beams, Lumber","id":{"cargoClassId":6,"dotNumber":53467}},{"cargoClassDesc":"Building Materials","id":{"cargoClassId":7,"dotNumber":53467}},{"cargoClassDesc":"Machinery, Large Objects","id":{"cargoClassId":9,"dotNumber":53467}},{"cargoClassDesc":"Fresh Produce","id":{"cargoClassId":10,"dotNumber":53467}},{"cargoClassDesc":"Liquids/Gases","id":{"cargoClassId":11,"dotNumber":53467}},{"cargoClassDesc":"Intermodal Containers","id":{"cargoClassId":12,"dotNumber":53467}},{"cargoClassDesc":"Passengers","id":{"cargoClassId":13,"dotNumber":53467}},{"cargoClassDesc":"Grain, Feed, Hay","id":{"cargoClassId":16,"dotNumber":53467}},{"cargoClassDesc":"Meat","id":{"cargoClassId":18,"dotNumber":53467}},{"cargoClassDesc":"Chemicals","id":{"cargoClassId":21,"dotNumber":53467}},{"cargoClassDesc":"Commodities Dry Bulk","id":{"cargoClassId":22,"dotNumber":53467}},{"cargoClassDesc":"Refrigerated Food","id":{"cargoClassId":23,"dotNumber":53467}},{"cargoClassDesc":"Beverages","id":{"cargoClassId":24,"dotNumber":53467}},{"cargoClassDesc":"Paper Products","id":{"cargoClassId":25,"dotNumber":53467}},{"cargoClassDesc":"Construction","id":{"cargoClassId":28,"dotNumber":53467}}],"retrievalDate":"2021-03-02T04:21:13.489+0000","_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/53467/cargo-carried"}}}`
	opClassRes      = `{"content":[{"id":{"dotNumber":53467,"operationClassId":1},"operationClassDesc":"Authorized For Hire"},{"id":{"dotNumber":53467,"operationClassId":3},"operationClassDesc":"Private Property"},{"id":{"dotNumber":53467,"operationClassId":4},"operationClassDesc":"Private Passenger, Business"}],"retrievalDate":"2021-03-02T04:25:07.415+0000","_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/53467/operation-classification"}}}`
	docketRes       = `{"content":[{"docketNumber":1515,"docketNumberId":125044,"dotNumber":44110,"prefix":"MC"}],"retrievalDate":"2021-03-02T04:42:16.385+0000","_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/44110/mc-numbers"}}}`
	authorityRes    = `{"content":[{"carrierAuthority":{"applicantID":8960,"authority":"N","authorizedForBroker":"Y","authorizedForHouseholdGoods":"N","authorizedForPassenger":"N","authorizedForProperty":"Y","brokerAuthorityStatus":"A","commonAuthorityStatus":"A","contractAuthorityStatus":"A","docketNumber":138328,"dotNumber":53467,"prefix":"MC"},"_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/53467/authority/8960"}}}],"retrievalDate":"2021-03-02T04:31:25.740+0000"}`
	oosRes          = `{"content":[{"oos":{"dotNumber":885213,"id":2992,"oosDate":"2004-06-04","oosReason":"NOP","oosReasonDescription":"90 day failure to pay fine"},"_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/885213/oos/2992"}}},{"oos":{"dotNumber":885213,"id":2993,"oosDate":"2004-06-04","oosReason":"NOP","oosReasonDescription":"90 day failure to pay fine"},"_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/885213/oos/2993"}}}],"retrievalDate":"2021-03-02T05:01:11.462+0000"}`
	basicsRes       = `{"content":[{"basic":{"basicsPercentile":"18%","basicsRunDate":"2017-01-27T05:00:00.000+0000","basicsType":{"basicsCode":"Unsafe Driving","basicsCodeMcmis":null,"basicsId":11,"basicsLongDesc":null,"basicsShortDesc":"Unsafe Driving"},"basicsViolationThreshold":"50","exceededFMCSAInterventionThreshold":"N","id":{"basicsId":11,"dotNumber":44110},"measureValue":"1.02","onRoadPerformanceThresholdViolationIndicator":"N","seriousViolationFromInvestigationPast12MonthIndicator":"N","totalInspectionWithViolation":155,"totalViolation":170},"dotNumber":null,"_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/44110/basics/11"}}},{"basic":{"basicsPercentile":"57%","basicsRunDate":"2017-01-27T05:00:00.000+0000","basicsType":{"basicsCode":"HOS Compliance","basicsCodeMcmis":null,"basicsId":12,"basicsLongDesc":null,"basicsShortDesc":"Hours-of-Service Compliance"},"basicsViolationThreshold":"50","exceededFMCSAInterventionThreshold":"Y","id":{"basicsId":12,"dotNumber":44110},"measureValue":"0.18","onRoadPerformanceThresholdViolationIndicator":"Y","seriousViolationFromInvestigationPast12MonthIndicator":"N","totalInspectionWithViolation":106,"totalViolation":119},"dotNumber":null,"_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/44110/basics/12"}}},{"basic":{"basicsPercentile":"28%","basicsRunDate":"2017-01-27T05:00:00.000+0000","basicsType":{"basicsCode":"Driver Fitness","basicsCodeMcmis":null,"basicsId":13,"basicsLongDesc":null,"basicsShortDesc":"Driver Fitness"},"basicsViolationThreshold":"65","exceededFMCSAInterventionThreshold":"N","id":{"basicsId":13,"dotNumber":44110},"measureValue":"0.03","onRoadPerformanceThresholdViolationIndicator":"N","seriousViolationFromInvestigationPast12MonthIndicator":"N","totalInspectionWithViolation":13,"totalViolation":14},"dotNumber":null,"_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/44110/basics/13"}}},{"basic":{"basicsPercentile":"0%","basicsRunDate":"2017-01-27T05:00:00.000+0000","basicsType":{"basicsCode":"Drugs/Alcohol","basicsCodeMcmis":null,"basicsId":14,"basicsLongDesc":null,"basicsShortDesc":"Controlled Substances/​Alcohol"},"basicsViolationThreshold":"65","exceededFMCSAInterventionThreshold":"N","id":{"basicsId":14,"dotNumber":44110},"measureValue":"0","onRoadPerformanceThresholdViolationIndicator":"N","seriousViolationFromInvestigationPast12MonthIndicator":"N","totalInspectionWithViolation":0,"totalViolation":0},"dotNumber":null,"_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/44110/basics/14"}}},{"basic":{"basicsPercentile":"19%","basicsRunDate":"2017-01-27T05:00:00.000+0000","basicsType":{"basicsCode":"Vehicle Maint.","basicsCodeMcmis":null,"basicsId":15,"basicsLongDesc":null,"basicsShortDesc":"Vehicle Maintenance"},"basicsViolationThreshold":"65","exceededFMCSAInterventionThreshold":"N","id":{"basicsId":15,"dotNumber":44110},"measureValue":"1.37","onRoadPerformanceThresholdViolationIndicator":"N","seriousViolationFromInvestigationPast12MonthIndicator":"N","totalInspectionWithViolation":381,"totalViolation":611},"dotNumber":null,"_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/44110/basics/15"}}}],"retrievalDate":"2021-03-02T04:32:03.523+0000"}`

	successDot        = "101"
	sysMaintenanceDot = "102"
	errDot            = "103"
)

func MockHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.Contains(r.URL.Path, "/carriers/name/"),
		strings.Contains(r.URL.Path, "/carriers/docket-number"):
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(multiCarrierRes))
	case strings.Contains(r.URL.Path, sysMaintenanceDot):
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(_maintenanceIndicator)
	case strings.Contains(r.URL.Path, errDot):
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintln(w, errRes)
	case r.URL.Path == successDot:
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(carrierRes))
	case strings.Contains(r.URL.Path, _cargoPath):
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(cargoRes))
	case strings.Contains(r.URL.Path, _opClassPath):
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(opClassRes))
	case strings.Contains(r.URL.Path, _carrierDocketPath):
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(docketRes))
	case strings.Contains(r.URL.Path, _authPath):
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(authorityRes))
	case strings.Contains(r.URL.Path, _oosPath):
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(oosRes))
	case strings.Contains(r.URL.Path, _basicsPath):
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(basicsRes))
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("404 Not Found"))
	}
}

type QCMobileClientTestSuite struct {
	suite.Suite

	client     Client
	testServer *httptest.Server
}

func (s *QCMobileClientTestSuite) SetupTest() {
	s.testServer = httptest.NewServer(http.HandlerFunc(MockHandler))
	testURL, err := url.Parse(s.testServer.URL)
	s.NoError(err)
	s.client = &client{
		http:      &http.Client{},
		uri:       testURL.Scheme + "://" + testURL.Host,
		baseQuery: "?webKey=" + "a-fake-key",
	}
}

func (s *QCMobileClientTestSuite) TearDownTest() {
	s.testServer.Close()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(QCMobileClientTestSuite))
}

func (s *QCMobileClientTestSuite) TestNewClient() {
	s.NotNil(NewClient(Config{Key: "a-fake-key"}))
}

func (s *QCMobileClientTestSuite) TestSearchCarriersByName() {
	res, err := s.client.SearchCarriersByName(context.Background(), "test", 1, 100)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(2, len(res))
	s.Equal(158121, res[1].Carrier.DOTNumber)
}

func (s *QCMobileClientTestSuite) TestSearchCarrierByDocket() {
	res, err := s.client.GetCarriersByDocket(context.Background(), "123")
	s.NoError(err)
	s.NotNil(res)
	s.Equal(2, len(res))
	s.Equal(158121, res[1].Carrier.DOTNumber)
}

func (s *QCMobileClientTestSuite) TestGetCarrier() {
	res, err := s.client.GetCarrier(context.Background(), successDot)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(158121, res.Carrier.DOTNumber)
}

func (s *QCMobileClientTestSuite) TestGetCargoCarried() {
	res, err := s.client.GetCargoCarried(context.Background(), successDot)
	s.NoError(err)
	s.NotNil(res)
	s.Equal("General Freight", res[0].CargoClassDesc)
}

func (s *QCMobileClientTestSuite) TestGetOperationClassification() {
	res, err := s.client.GetOperationClassification(context.Background(), successDot)
	s.NoError(err)
	s.NotNil(res)
	s.Equal("Authorized For Hire", res[0].OperationClassDesc)
}

func (s *QCMobileClientTestSuite) TestGetDocketNumbers() {
	res, err := s.client.GetDocketNumbers(context.Background(), successDot)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(1515, res[0].DocketNumber)
}

func (s *QCMobileClientTestSuite) TestGetAuthority() {
	res, err := s.client.GetAuthority(context.Background(), successDot)
	s.NoError(err)
	s.NotNil(res)
	s.Equal("N", res[0].CarrierAuthority.Authority)
}

func (s *QCMobileClientTestSuite) TestGetOOS() {
	res, err := s.client.GetOOS(context.Background(), successDot)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(Date("2004-06-04"), res[0].OOS.OOSDate)
}

func (s *QCMobileClientTestSuite) TestGetGetBasics() {
	res, err := s.client.GetBasics(context.Background(), successDot)
	s.NoError(err)
	s.NotNil(res)
	s.Equal("18%", res[0].Basic.BasicsPercentile)
}

func (s *QCMobileClientTestSuite) TestGetCompleteCarrierDetails() {
	res, err := s.client.GetCompleteCarrierDetails(context.Background(), successDot)
	s.NoError(err)
	s.NotNil(res)
}

func (s *QCMobileClientTestSuite) TestErrRes() {
	res, err := s.client.GetCarrier(context.Background(), errDot)
	s.Error(err)
	s.Nil(res)
	s.Equal("404 Not Found: Webkey not found", err.Error())
}

func (s *QCMobileClientTestSuite) TestMaintenanceErr() {
	res, err := s.client.GetCarrier(context.Background(), sysMaintenanceDot)
	s.Error(err)
	s.Nil(res)
	s.Equal(ErrSystemMaintenance, err)
}

func (s *QCMobileClientTestSuite) TestBuildURL() {
	c := NewClient(Config{}).(*client)
	path := _searchPath + "carrierName"
	query := "start=" + strconv.Itoa(1) + "&size=" + strconv.Itoa(2)

	expected := "https://mobile.fmcsa.dot.gov/qc/services/carriers/name/carrierName?webKey=my-key&start=1&size=2"
	s.Equal(expected, c.buildURL(path, query))
}

func BenchmarkClient_GetCarrier(b *testing.B) {
	testServer := httptest.NewServer(http.HandlerFunc(MockHandler))
	testURL, _ := url.Parse(testServer.URL)
	client := NewClient(Config{}).(*client)
	client.uri = testURL.Scheme + "://" + testURL.Host
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = client.GetCarrier(context.Background(), successDot)
	}
}
