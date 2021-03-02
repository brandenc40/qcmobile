package qcmobile

import "net/http"

// Config -
type Config struct {
	// Key - (required) Your QCMobile API WebKey
	Key string

	// HTTPClient - (optional) Defaults to &http.Client{}
	HTTPClient *http.Client
}
