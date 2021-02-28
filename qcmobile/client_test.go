package qcmobile

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

const (
	singleCarrierRes = `{"content":{"_links":{"basics":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/basics"},"cargo carried":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/cargo-carried"},"operation classification":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/operation-classification"},"docket numbers":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/docket-numbers"},"carrier active-For-hire authority":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/authority"},"self":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121"}},"carrier":{"allowedToOperate":"Y","bipdInsuranceOnFile":"1000","bipdInsuranceRequired":"Y","bipdRequiredAmount":"750","bondInsuranceOnFile":"0","bondInsuranceRequired":"u","brokerAuthorityStatus":"N","cargoInsuranceOnFile":"0","cargoInsuranceRequired":"u","carrierOperation":{"carrierOperationCode":"A","carrierOperationDesc":"Interstate"},"censusTypeId":{"censusType":"C","censusTypeDesc":"CARRIER","censusTypeId":1},"commonAuthorityStatus":"A","contractAuthorityStatus":"A","crashTotal":15,"dbaName":null,"dotNumber":158121,"driverInsp":259,"driverOosInsp":3,"driverOosRate":1.15830115830115830115830115830115830116,"driverOosRateNationalAverage":"5.51","ein":391474414,"fatalCrash":0,"hazmatInsp":10,"hazmatOosInsp":1,"hazmatOosRate":10,"hazmatOosRateNationalAverage":"4.5","injCrash":3,"isPassengerCarrier":"N","issScore":null,"legalName":"VERIHA TRUCKING INC","mcs150Outdated":"N","oosDate":null,"oosRateNationalAverageYear":"2009-2010","phyCity":"MARINETTE","phyCountry":"US","phyState":"WI","phyStreet":"2830 CLEVELAND AVE","phyZipcode":"54143","reviewDate":"1996-04-22","reviewType":"C","safetyRating":"S","safetyRatingDate":"1996-04-25","safetyReviewDate":"1996-04-22","safetyReviewType":"C","snapshotDate":null,"statusCode":"A","totalDrivers":213,"totalPowerUnits":213,"towawayCrash":12,"vehicleInsp":124,"vehicleOosInsp":17,"vehicleOosRate":13.70967741935483870967741935483870967742,"vehicleOosRateNationalAverage":"20.72"}},"retrievalDate":"2021-02-28T06:33:31.490+0000"}`
	multiCarrierRes  = `{"content":[{"_links":{"basics":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/basics"},"cargo carried":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/cargo-carried"},"operation classification":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/operation-classification"},"docket numbers":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/docket-numbers"},"carrier active-For-hire authority":{"href":"https://mobile.fmcsa.dot.gov/qc/services/carriers/158121/authority"}},"carrier":{"allowedToOperate":"Y","bipdInsuranceOnFile":"1000","bipdInsuranceRequired":"Y","bipdRequiredAmount":"750","bondInsuranceOnFile":"0","bondInsuranceRequired":"u","brokerAuthorityStatus":"N","cargoInsuranceOnFile":"0","cargoInsuranceRequired":"u","carrierOperation":{"carrierOperationCode":"A","carrierOperationDesc":"Interstate"},"censusTypeId":{"censusType":"C","censusTypeDesc":"CARRIER","censusTypeId":1},"commonAuthorityStatus":"A","contractAuthorityStatus":"A","crashTotal":15,"dbaName":null,"dotNumber":158121,"driverInsp":259,"driverOosInsp":3,"driverOosRate":1.15830115830115830115830115830115830116,"driverOosRateNationalAverage":"5.51","ein":391474414,"fatalCrash":0,"hazmatInsp":10,"hazmatOosInsp":1,"hazmatOosRate":10,"hazmatOosRateNationalAverage":"4.5","injCrash":3,"isPassengerCarrier":null,"issScore":null,"legalName":"VERIHA TRUCKING INC","mcs150Outdated":"N","oosDate":null,"oosRateNationalAverageYear":"2009-2010","phyCity":"MARINETTE","phyCountry":"US","phyState":"WI","phyStreet":"2830 CLEVELAND AVE","phyZipcode":"54143","reviewDate":"1996-04-22","reviewType":"C","safetyRating":"S","safetyRatingDate":"1996-04-25","safetyReviewDate":"1996-04-22","safetyReviewType":"C","snapshotDate":null,"statusCode":"A","totalDrivers":213,"totalPowerUnits":213,"towawayCrash":12,"vehicleInsp":124,"vehicleOosInsp":17,"vehicleOosRate":13.70967741935483870967741935483870967742,"vehicleOosRateNationalAverage":"20.72"}}],"retrievalDate":"2021-02-28T06:57:18.350+0000"}`
)

type QCMobileClientTestSuite struct {
	suite.Suite

	client     Client
	testServer *httptest.Server
}

func (s *QCMobileClientTestSuite) SetupTest() {
	s.testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/carriers/name/") || strings.Contains(r.URL.Path, "docket-number") {
			_, _ = fmt.Fprintln(w, multiCarrierRes)
		} else if strings.Contains(r.URL.Path, "/carriers/123") {
			_, _ = fmt.Fprintln(w, singleCarrierRes)
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
	res, err := s.client.GetCarrierByDOT(context.Background(), 123)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(158121, res.CarrierDetails.Carrier.DotNumber)
}

func (s *QCMobileClientTestSuite) TestGetCarrierByDocket() {
	res, err := s.client.GetCarrierByDocket(context.Background(), 123)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(1, len(res.MultiCarrierDetails))
	s.Equal(158121, res.MultiCarrierDetails[0].Carrier.DotNumber)
}

func (s *QCMobileClientTestSuite) TestSearchCarriersByName() {
	res, err := s.client.SearchCarriersByName(context.Background(), "test", 1, 100)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(1, len(res.MultiCarrierDetails))
	s.Equal(158121, res.MultiCarrierDetails[0].Carrier.DotNumber)
}
