package qcmobile

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/brandenc40/fmcsa-qc-mobile/qcmobile/entities"
	"io"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

const (
	_scheme            = "https"
	_host              = "mobile.fmcsa.dot.gov"
	_basePath          = "/qc/services/carriers/"
	_searchDocketPath  = _basePath + "docket-number/"
	_searchPath        = _basePath + "name/"
	_cargoPath         = "/cargo-carried"
	_opClassPath       = "/operation-classification"
	_carrierDocketPath = "/docket-numbers"
	_authPath          = "/authority"
	_oosPath           = "/oos"
	_basicsPath        = "/basics"

	_maintenanceIndicator = "<title>FMCSA System Maintenance Page</title>"
)

// ErrSystemMaintenance -
var ErrSystemMaintenance = errors.New("FMCSA Portal Unavailable due to Scheduled System Maintenance")

// Client -
type Client interface {
	SearchCarriersByName(ctx context.Context, name string, start, size int) ([]*entities.CarrierDetails, error)
	GetCompleteCarrierDetails(ctx context.Context, dotNumber int) (*entities.CompleteCarrierDetails, error)
	GetCarriersByDocket(ctx context.Context, docketNumber int) ([]*entities.CarrierDetails, error)
	GetCarrier(ctx context.Context, dotNumber int) (*entities.CarrierDetails, error)
	GetCargoCarried(ctx context.Context, dotNumber int) ([]*entities.CargoClass, error)
	GetOperationClassification(ctx context.Context, dotNumber int) ([]*entities.OperationClass, error)
	GetDocketNumbers(ctx context.Context, dotNumber int) ([]*entities.Docket, error)
	GetAuthority(ctx context.Context, dotNumber int) ([]*entities.AuthorityDetails, error)
	GetOOS(ctx context.Context, dotNumber int) ([]*entities.OOSDetails, error)
	GetBasics(ctx context.Context, dotNumber int) ([]*entities.BasicsDetails, error)
}

// NewClient -
func NewClient(cfg Config) Client {
	client := &client{
		http:   cfg.HTTPClient,
		key:    cfg.Key,
		host:   _host,
		scheme: _scheme,
	}
	if client.http == nil {
		client.http = &http.Client{}
	}
	return client
}

type client struct {
	http   *http.Client
	key    string
	host   string
	scheme string
}

// SearchCarriersByName -
func (c *client) SearchCarriersByName(ctx context.Context, carrierName string, start, size int) ([]*entities.CarrierDetails, error) {
	path := _searchPath + carrierName
	query := "start=" + strconv.Itoa(start) + "&size=" + strconv.Itoa(size)
	var response entities.SearchResponse
	if err := c.doGet(ctx, path, query, &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetCarriersByDocket -
func (c *client) GetCarriersByDocket(ctx context.Context, docketNumber int) ([]*entities.CarrierDetails, error) {
	path := _searchDocketPath + strconv.Itoa(docketNumber)
	var response entities.GetCarriersByDocketResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

func (c *client) GetCompleteCarrierDetails(ctx context.Context, dotNumber int) (*entities.CompleteCarrierDetails, error) {
	var (
		carrier                  *entities.Carrier
		cargosCarried            []*entities.CargoClass
		operationClassifications []*entities.OperationClass
		dockets                  []*entities.Docket
		authorityDetails         []*entities.AuthorityDetails
		basicsDetails            []*entities.BasicsDetails
		oosDetails               []*entities.OOSDetails

		g errgroup.Group
	)
	g.Go(func() error {
		c, err := c.GetCarrier(ctx, dotNumber)
		if err == nil {
			carrier = c.Carrier
		}
		return err
	})
	g.Go(func() error {
		var err error
		cargosCarried, err = c.GetCargoCarried(ctx, dotNumber)
		return err
	})
	g.Go(func() error {
		var err error
		operationClassifications, err = c.GetOperationClassification(ctx, dotNumber)
		return err
	})
	g.Go(func() error {
		var err error
		dockets, err = c.GetDocketNumbers(ctx, dotNumber)
		return err
	})
	g.Go(func() error {
		var err error
		authorityDetails, err = c.GetAuthority(ctx, dotNumber)
		return err
	})
	g.Go(func() error {
		var err error
		basicsDetails, err = c.GetBasics(ctx, dotNumber)
		return err
	})
	g.Go(func() error {
		var err error
		oosDetails, err = c.GetOOS(ctx, dotNumber)
		return err
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return &entities.CompleteCarrierDetails{
		Carrier:                  carrier,
		CargosCarried:            cargosCarried,
		OperationClassifications: operationClassifications,
		Dockets:                  dockets,
		AuthorityDetails:         authorityDetails,
		BasicsDetails:            basicsDetails,
		OOSDetails:               oosDetails,
	}, nil
}

// GetCarrier -
func (c *client) GetCarrier(ctx context.Context, dotNumber int) (*entities.CarrierDetails, error) {
	path := _basePath + strconv.Itoa(dotNumber)
	var response entities.CarrierResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetCargoCarried -
func (c *client) GetCargoCarried(ctx context.Context, dotNumber int) ([]*entities.CargoClass, error) {
	path := _basePath + strconv.Itoa(dotNumber) + _cargoPath
	var response entities.CargoCarriedResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetOperationClassification -
func (c *client) GetOperationClassification(ctx context.Context, dotNumber int) ([]*entities.OperationClass, error) {
	path := _basePath + strconv.Itoa(dotNumber) + _opClassPath
	var response entities.OperationClassificationResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetDocketNumbers -
func (c *client) GetDocketNumbers(ctx context.Context, dotNumber int) ([]*entities.Docket, error) {
	path := _basePath + strconv.Itoa(dotNumber) + _carrierDocketPath
	var response entities.DocketNumbersResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetAuthority -
func (c *client) GetAuthority(ctx context.Context, dotNumber int) ([]*entities.AuthorityDetails, error) {
	path := _basePath + strconv.Itoa(dotNumber) + _authPath
	var response entities.AuthorityResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetOOS -
func (c *client) GetOOS(ctx context.Context, dotNumber int) ([]*entities.OOSDetails, error) {
	path := _basePath + strconv.Itoa(dotNumber) + _oosPath
	var response entities.OOSResponse
	if err := c.doGet(ctx, path, "", &response); err != nil {
		return nil, err
	}
	return response.Content, nil
}

// GetBasics -
func (c *client) GetBasics(ctx context.Context, dotNumber int) ([]*entities.BasicsDetails, error) {
	path := _basePath + strconv.Itoa(dotNumber) + _basicsPath
	var response entities.BasicsResponse
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
	if resp.StatusCode != 200 {
		return tryExtractError(resp)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(resp.Status)
	}
	if err := json.Unmarshal(body, output); err != nil {
		if strings.Contains(string(body), _maintenanceIndicator) {
			return ErrSystemMaintenance
		}
		return err
	}
	return nil
}

func (c *client) buildURL(path, query string) string {
	return c.scheme + "://" + c.host + path + "?webKey=" + c.key + "&" + query
}

func tryExtractError(resp *http.Response) error {
	if body, err := io.ReadAll(resp.Body); err == nil {
		var errResponse entities.ErrorResponse
		if err := json.Unmarshal(body, &errResponse); err == nil {
			return errors.New(resp.Status + ": " + errResponse.ErrMsg)
		}
	}
	return errors.New(resp.Status)
}
