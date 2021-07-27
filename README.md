# FMCSA QC Mobile API Client

[![Go Reference](https://pkg.go.dev/badge/github.com/brandenc40/qcmobile.svg)](https://pkg.go.dev/github.com/brandenc40/qcmobile)
[![Go Report Card](https://goreportcard.com/badge/github.com/brandenc40/qcmobile)](https://goreportcard.com/report/github.com/brandenc40/qcmobile)

Unofficial API client

https://mobile.fmcsa.dot.gov/QCDevsite/home

```
go get github.com/brandenc40/qcmobile
```

## Client Interface

```go
// Client - QC Mobile API Client interface
type Client interface {
    SearchCarriersByName(ctx context.Context, name string, start, size int) ([]*CarrierDetails, error)
    GetCompleteCarrierDetails(ctx context.Context, dotNumber int) (*CompleteCarrierDetails, error)
    GetCarriersByDocket(ctx context.Context, docketNumber int) ([]*CarrierDetails, error)
    GetCarrier(ctx context.Context, dotNumber int) (*CarrierDetails, error)
    GetCargoCarried(ctx context.Context, dotNumber int) ([]*CargoClass, error)
    GetOperationClassification(ctx context.Context, dotNumber int) ([]*OperationClass, error)
    GetDocketNumbers(ctx context.Context, dotNumber int) ([]*Docket, error)
    GetAuthority(ctx context.Context, dotNumber int) ([]*AuthorityDetails, error)
    GetOOS(ctx context.Context, dotNumber int) ([]*OOSDetails, error)
    GetBasics(ctx context.Context, dotNumber int) ([]*BasicsDetails, error)
}
```

## Example Usage

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/brandenc40/qcmobile"
)

func main() {
	// build client
	cfg := qcmobile.Config{
		Key:        "YOUR_KEY",
		HTTPClient: &http.Client{}, // OPTIONAL - will default to &http.Client{} if nil
	}
	client := qcmobile.NewClient(cfg)
	
	// build context to handle function timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// get carrier from QC Mobile and print
	carrier, err := client.GetCarrier(ctx, 53467)
	if err != nil {
		// handle error
	}
	fmt.Println(carrier.Carrier)
}
```
