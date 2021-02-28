package qcmobile

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/brandenc40/fmcsa-qc-mobile/qcmobile/entities"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	_scheme     = "https"
	_host       = "mobile.fmcsa.dot.gov"
	_basePath   = "/qc/services/carriers/"
	_docketPath = _basePath + "docket-number/"
	_searchPath = _basePath + "name/"

	_maintenanceIndicator = "<title>FMCSA System Maintenance Page</title>"
)

// ErrSystemMaintenance -
var ErrSystemMaintenance = errors.New("FMCSA Portal Unavailable due to Scheduled System Maintenance")

// Client -
type Client interface {
	GetCarrierByDOT(ctx context.Context, dot int) (*entities.SingleCarrierResponse, error)
	GetCarrierByDocket(ctx context.Context, docket int) (*entities.MultiCarrierResponse, error)
	SearchCarriersByName(ctx context.Context, name string, start, size int) (*entities.MultiCarrierResponse, error)
}

// NewClient -
func NewClient(apiKey string) Client {
	return &client{
		http:   &http.Client{},
		key:    apiKey,
		host:   _host,
		scheme: _scheme,
	}
}

type client struct {
	http   *http.Client
	key    string
	host   string
	scheme string
}

// GetCarrierByDOT -
func (c *client) GetCarrierByDOT(ctx context.Context, dotNumber int) (*entities.SingleCarrierResponse, error) {
	reqURL := url.URL{
		Scheme:   c.scheme,
		Host:     c.host,
		Path:     _basePath + strconv.Itoa(dotNumber),
		RawQuery: "webKey=" + c.key,
	}
	var response entities.SingleCarrierResponse
	if err := c.doGet(ctx, reqURL.String(), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCarrierByDocket -
func (c *client) GetCarrierByDocket(ctx context.Context, docketNumber int) (*entities.MultiCarrierResponse, error) {
	reqURL := url.URL{
		Scheme:   c.scheme,
		Host:     c.host,
		Path:     _docketPath + strconv.Itoa(docketNumber),
		RawQuery: "webKey=" + c.key,
	}
	var response entities.MultiCarrierResponse
	if err := c.doGet(ctx, reqURL.String(), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SearchCarriersByName -
func (c *client) SearchCarriersByName(ctx context.Context, carrierName string, start, size int) (*entities.MultiCarrierResponse, error) {
	reqURL := url.URL{
		Scheme:   c.scheme,
		Host:     c.host,
		Path:     _searchPath + carrierName,
		RawQuery: "webKey=" + c.key + "&start=" + strconv.Itoa(start) + "&size=" + strconv.Itoa(size),
	}
	var response entities.MultiCarrierResponse
	if err := c.doGet(ctx, reqURL.String(), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *client) doGet(ctx context.Context, url string, output interface{}) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

func tryExtractError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(resp.Status)
	}
	var errResponse entities.ErrorResponse
	if err := json.Unmarshal(body, &errResponse); err != nil {
		return errors.New(resp.Status)
	}
	return errors.New(resp.Status + ": " + errResponse.ErrMsg)
}
