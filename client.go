package qcmobile

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"golang.org/x/sync/errgroup"
)

const (
	_uri               = "https://mobile.fmcsa.dot.gov/qc/services/carriers/"
	_searchDocketPath  = "docket-number/"
	_searchPath        = "name/"
	_cargoPath         = "/cargo-carried"
	_opClassPath       = "/operation-classification"
	_carrierDocketPath = "/docket-numbers"
	_authPath          = "/authority"
	_oosPath           = "/oos"
	_basicsPath        = "/basics"
)

var (
	// ErrSystemMaintenance is returned when the api is failing due to system maintenance
	ErrSystemMaintenance = errors.New("FMCSA Portal Unavailable due to Scheduled System Maintenance")

	_maintenanceIndicator = []byte("<title>FMCSA System Maintenance Page</title>")
)

// Client - QC Mobile API Client interface
type Client interface {
	SearchCarriersByName(ctx context.Context, name string, start, size int) ([]*CarrierDetails, error)
	GetCompleteCarrierDetails(ctx context.Context, dotNumber string) (*CompleteCarrierDetails, error)
	GetCarriersByDocket(ctx context.Context, docketNumber string) ([]*CarrierDetails, error)
	GetCarrier(ctx context.Context, dotNumber string) (*CarrierDetails, error)
	GetCargoCarried(ctx context.Context, dotNumber string) ([]*CargoClass, error)
	GetOperationClassification(ctx context.Context, dotNumber string) ([]*OperationClass, error)
	GetDocketNumbers(ctx context.Context, dotNumber string) ([]*Docket, error)
	GetAuthority(ctx context.Context, dotNumber string) ([]*AuthorityDetails, error)
	GetOOS(ctx context.Context, dotNumber string) ([]*OOSDetails, error)
	GetBasics(ctx context.Context, dotNumber string) ([]*BasicsDetails, error)
}

// NewClient -
func NewClient(cfg Config) Client {
	client := &client{
		http:      cfg.HTTPClient,
		uri:       _uri,
		baseQuery: "?webKey=" + cfg.Key,
	}
	if client.http == nil {
		client.http = &http.Client{}
	}
	return client
}

type client struct {
	http      *http.Client
	uri       string
	baseQuery string
}

// SearchCarriersByName -
func (c *client) SearchCarriersByName(ctx context.Context, carrierName string, start, size int) ([]*CarrierDetails, error) {
	path := _searchPath + carrierName
	query := "start=" + strconv.Itoa(start) + "&size=" + strconv.Itoa(size)
	var response searchResponse
	if err := c.doGet(ctx, path, query, &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetCarriersByDocket -
func (c *client) GetCarriersByDocket(ctx context.Context, docketNumber string) ([]*CarrierDetails, error) {
	path := _searchDocketPath + docketNumber
	var response getCarriersByDocketResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

func (c *client) GetCompleteCarrierDetails(ctx context.Context, dotNumber string) (*CompleteCarrierDetails, error) {
	var (
		output = new(CompleteCarrierDetails)
		group  errgroup.Group
	)
	group.Go(func() error {
		c, err := c.GetCarrier(ctx, dotNumber)
		if err == nil {
			output.Carrier = c.Carrier
		}
		return err
	})
	group.Go(func() error {
		var err error
		output.CargosCarried, err = c.GetCargoCarried(ctx, dotNumber)
		return err
	})
	group.Go(func() error {
		var err error
		output.OperationClassifications, err = c.GetOperationClassification(ctx, dotNumber)
		return err
	})
	group.Go(func() error {
		var err error
		output.Dockets, err = c.GetDocketNumbers(ctx, dotNumber)
		return err
	})
	group.Go(func() error {
		var err error
		output.AuthorityDetails, err = c.GetAuthority(ctx, dotNumber)
		return err
	})
	group.Go(func() error {
		var err error
		output.BasicsDetails, err = c.GetBasics(ctx, dotNumber)
		return err
	})
	group.Go(func() error {
		var err error
		output.OOSDetails, err = c.GetOOS(ctx, dotNumber)
		return err
	})
	if err := group.Wait(); err != nil {
		return nil, err
	}
	return output, nil
}

// GetCarrier -
func (c *client) GetCarrier(ctx context.Context, dotNumber string) (*CarrierDetails, error) {
	var response carrierResponse
	if err := c.doGet(ctx, dotNumber, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetCargoCarried -
func (c *client) GetCargoCarried(ctx context.Context, dotNumber string) ([]*CargoClass, error) {
	path := dotNumber + _cargoPath
	var response cargoCarriedResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetOperationClassification -
func (c *client) GetOperationClassification(ctx context.Context, dotNumber string) ([]*OperationClass, error) {
	path := dotNumber + _opClassPath
	var response operationClassificationResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetDocketNumbers -
func (c *client) GetDocketNumbers(ctx context.Context, dotNumber string) ([]*Docket, error) {
	path := dotNumber + _carrierDocketPath
	var response docketNumbersResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetAuthority -
func (c *client) GetAuthority(ctx context.Context, dotNumber string) ([]*AuthorityDetails, error) {
	path := dotNumber + _authPath
	var response authorityResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetOOS -
func (c *client) GetOOS(ctx context.Context, dotNumber string) ([]*OOSDetails, error) {
	path := dotNumber + _oosPath
	var response oosResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetBasics -
func (c *client) GetBasics(ctx context.Context, dotNumber string) ([]*BasicsDetails, error) {
	path := dotNumber + _basicsPath
	var response basicsResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

func (c *client) doGet(ctx context.Context, path, query string, output interface{}) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.buildURL(path, query), nil)
	if err != nil {
		return err
	}
	resp, err := c.http.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return tryExtractError(resp)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(resp.Status)
	}
	if err := json.Unmarshal(body, output); err != nil {
		if bytes.Contains(body, _maintenanceIndicator) {
			return ErrSystemMaintenance
		}
		return err
	}
	return nil
}

func (c *client) buildURL(path, query string) string {
	if query != "" {
		return c.uri + path + c.baseQuery + "&" + query
	}
	return c.uri + path + c.baseQuery
}

func tryExtractError(resp *http.Response) error {
	if body, err := io.ReadAll(resp.Body); err == nil {
		var errResponse errorResponse
		if err := json.Unmarshal(body, &errResponse); err == nil {
			return errors.New(resp.Status + ": " + errResponse.ErrMsg)
		}
	}
	return errors.New(resp.Status)
}
