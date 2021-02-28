package qcmobile

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

const (
	singleCarrierRes = `{"content":{"_links":{"basics":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/basics"},"cargo carried":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/cargo-carried"},"operation classification":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/operation-classification"},"docket numbers":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/docket-numbers"},"carrier active-For-hire authority":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/authority"},"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121"}},"carrier":{"allowedToOperate":"Y","bipdInsuranceOnFile":"1000","bipdInsuranceRequired":"Y","bipdRequiredAmount":"750","bondInsuranceOnFile":"0","bondInsuranceRequired":"u","brokerAuthorityStatus":"N","cargoInsuranceOnFile":"0","cargoInsuranceRequired":"u","carrierOperation":{"carrierOperationCode":"A","carrierOperationDesc":"Interstate"},"censusTypeId":{"censusType":"C","censusTypeDesc":"CARRIER","censusTypeId":1},"commonAuthorityStatus":"A","contractAuthorityStatus":"A","crashTotal":15,"dbaName":null,"dotNumber":158121,"driverInsp":259,"driverOosInsp":3,"driverOosRate":1.15830115830115830115830115830115830116,"driverOosRateNationalAverage":"5.51","ein":391474414,"fatalCrash":0,"hazmatInsp":10,"hazmatOosInsp":1,"hazmatOosRate":10,"hazmatOosRateNationalAverage":"4.5","injCrash":3,"isPassengerCarrier":"N","issScore":null,"legalName":"VERIHA TRUCKING INC","mcs150Outdated":"N","oosDate":null,"oosRateNationalAverageYear":"2009-2010","phyCity":"MARINETTE","phyCountry":"US","phyState":"WI","phyStreet":"2830 CLEVELAND AVE","phyZipcode":"54143","reviewDate":"1996-04-22","reviewType":"C","safetyRating":"S","safetyRatingDate":"1996-04-25","safetyReviewDate":"1996-04-22","safetyReviewType":"C","snapshotDate":null,"statusCode":"A","totalDrivers":213,"totalPowerUnits":213,"towawayCrash":12,"vehicleInsp":124,"vehicleOosInsp":17,"vehicleOosRate":13.70967741935483870967741935483870967742,"vehicleOosRateNationalAverage":"20.72"}},"retrievalDate":"2021-02-28T06:33:31.490+0000"}`
	multiCarrierRes  = `{"content":[{"_links":{"basics":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/753076/basics"},"cargo carried":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/753076/cargo-carried"},"operation classification":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/753076/operation-classification"},"docket numbers":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/753076/docket-numbers"},"carrier active-For-hire authority":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/753076/authority"}},"carrier":{"allowedToOperate":"Y","bipdInsuranceOnFile":"0","bipdInsuranceRequired":"Y","bipdRequiredAmount":"750","bondInsuranceOnFile":"0","bondInsuranceRequired":"u","brokerAuthorityStatus":"N","cargoInsuranceOnFile":"0","cargoInsuranceRequired":"u","carrierOperation":{"carrierOperationCode":"C","carrierOperationDesc":"Intrastate Non-Hazmat"},"censusTypeId":{"censusType":"C","censusTypeDesc":"CARRIER","censusTypeId":1},"commonAuthorityStatus":"N","contractAuthorityStatus":"N","crashTotal":0,"dbaName":"LEE VERIHA TRUCKING AND EXCAVATING","dotNumber":753076,"driverInsp":0,"driverOosInsp":0,"driverOosRate":0,"driverOosRateNationalAverage":"5.51","ein":391821630,"fatalCrash":0,"hazmatInsp":0,"hazmatOosInsp":0,"hazmatOosRate":0,"hazmatOosRateNationalAverage":"4.5","injCrash":0,"isPassengerCarrier":null,"issScore":null,"legalName":"AMERICAN ELM SAWMILL INC","mcs150Outdated":"N","oosDate":null,"oosRateNationalAverageYear":"2009-2010","phyCity":"PORTERFIELD","phyCountry":"US","phyState":"WI","phyStreet":"W 3065 VERIHA RD","phyZipcode":"54159","reviewDate":null,"reviewType":null,"safetyRating":null,"safetyRatingDate":null,"safetyReviewDate":null,"safetyReviewType":null,"snapshotDate":null,"statusCode":"A","totalDrivers":2,"totalPowerUnits":2,"towawayCrash":0,"vehicleInsp":0,"vehicleOosInsp":0,"vehicleOosRate":0,"vehicleOosRateNationalAverage":"20.72"}},{"_links":{"basics":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/basics"},"cargo carried":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/cargo-carried"},"operation classification":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/operation-classification"},"docket numbers":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/docket-numbers"},"carrier active-For-hire authority":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/authority"}},"carrier":{"allowedToOperate":"Y","bipdInsuranceOnFile":"1000","bipdInsuranceRequired":"Y","bipdRequiredAmount":"750","bondInsuranceOnFile":"0","bondInsuranceRequired":"u","brokerAuthorityStatus":"N","cargoInsuranceOnFile":"0","cargoInsuranceRequired":"u","carrierOperation":{"carrierOperationCode":"A","carrierOperationDesc":"Interstate"},"censusTypeId":{"censusType":"C","censusTypeDesc":"CARRIER","censusTypeId":1},"commonAuthorityStatus":"A","contractAuthorityStatus":"A","crashTotal":15,"dbaName":null,"dotNumber":158121,"driverInsp":259,"driverOosInsp":3,"driverOosRate":1.15830115830115830115830115830115830116,"driverOosRateNationalAverage":"5.51","ein":391474414,"fatalCrash":0,"hazmatInsp":10,"hazmatOosInsp":1,"hazmatOosRate":10,"hazmatOosRateNationalAverage":"4.5","injCrash":3,"isPassengerCarrier":null,"issScore":null,"legalName":"VERIHA TRUCKING INC","mcs150Outdated":"N","oosDate":null,"oosRateNationalAverageYear":"2009-2010","phyCity":"MARINETTE","phyCountry":"US","phyState":"WI","phyStreet":"2830 CLEVELAND AVE","phyZipcode":"54143","reviewDate":"1996-04-22","reviewType":"C","safetyRating":"S","safetyRatingDate":"1996-04-25","safetyReviewDate":"1996-04-22","safetyReviewType":"C","snapshotDate":null,"statusCode":"A","totalDrivers":213,"totalPowerUnits":213,"towawayCrash":12,"vehicleInsp":124,"vehicleOosInsp":17,"vehicleOosRate":13.70967741935483870967741935483870967742,"vehicleOosRateNationalAverage":"20.72"}}],"retrievalDate":"2021-02-28T07:25:05.638+0000"}`
	errRes           = `{"content":"Webkey not found","retrievalDate":"2021-02-28T07:35:25.991+0000","_links":{"self":{"href":"https://mobile.fmcsa.dot.gov/qc"},"searchByName":{"href":"https://mobile.fmcsa.dot.gov/qc/name/:name"},"lookupBydotNumber":{"href":"https://mobile.fmcsa.dot.gov/qc/id/:dotNumber"}}}`

	successDot        = 101
	errDot            = 102
	sysMaintenanceDot = 103
)

type QCMobileClientTestSuite struct {
	suite.Suite

	client     Client
	testServer *httptest.Server
}

func (s *QCMobileClientTestSuite) SetupTest() {
	s.testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/carriers/name/") || strings.Contains(r.URL.Path, "docket-number") {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(multiCarrierRes))
		} else if strings.Contains(r.URL.Path, "/carriers/"+strconv.Itoa(successDot)) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(singleCarrierRes))
		} else if strings.Contains(r.URL.Path, "/carriers/"+strconv.Itoa(sysMaintenanceDot)) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("some stuff" + _maintenanceIndicator))
		} else if strings.Contains(r.URL.Path, "/carriers/"+strconv.Itoa(errDot)) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = fmt.Fprintln(w, errRes)
		}
	}))
	testURL, err := url.Parse(s.testServer.URL)
	s.NoError(err)
	s.client = &client{
		http:   &http.Client{},
		key:    "a-fake-key",
		host:   testURL.Host,
		scheme: testURL.Scheme,
	}
}

func (s *QCMobileClientTestSuite) TearDownTest() {
	s.testServer.Close()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(QCMobileClientTestSuite))
}

func (s *QCMobileClientTestSuite) TestNewClient() {
	s.NotNil(NewClient("a-fake-key"))
}

func (s *QCMobileClientTestSuite) TestGetCarrierByDOT() {
	res, err := s.client.GetCarrierByDOT(context.Background(), successDot)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(158121, res.CarrierDetails.Carrier.DotNumber)
}

func (s *QCMobileClientTestSuite) TestGetCarrierByDocket() {
	res, err := s.client.GetCarrierByDocket(context.Background(), 123)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(2, len(res.MultiCarrierDetails))
	s.Equal(158121, res.MultiCarrierDetails[1].Carrier.DotNumber)
}

func (s *QCMobileClientTestSuite) TestSearchCarriersByName() {
	res, err := s.client.SearchCarriersByName(context.Background(), "test", 1, 100)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(2, len(res.MultiCarrierDetails))
	s.Equal(158121, res.MultiCarrierDetails[1].Carrier.DotNumber)
}

func (s *QCMobileClientTestSuite) TestErrRes() {
	res, err := s.client.GetCarrierByDOT(context.Background(), errDot)
	s.Error(err)
	s.Nil(res)
	s.Equal("404 Not Found: Webkey not found", err.Error())
}

func (s *QCMobileClientTestSuite) TestMaintenanceErr() {
	res, err := s.client.GetCarrierByDOT(context.Background(), sysMaintenanceDot)
	s.Error(err)
	s.Nil(res)
	s.Equal(ErrSystemMaintenance, err)
}
