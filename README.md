# FMCSA QC Mobile API Client

Unofficial API client

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

## New Client

```go
cfg := qcmobile.Config{
    Key:        "your-key-here",
    HTTPClient: &http.Client{}, // (optional) so you can customize your HTTP client object
}
client := qcmobile.NewClient(cfg)
```
