// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gbox-sdk-go/internal/apijson"
	"github.com/stainless-sdks/gbox-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/gbox-sdk-go/option"
	"github.com/stainless-sdks/gbox-sdk-go/packages/respjson"
)

// V1BoxBrowserService contains methods and other services that help with
// interacting with the gbox-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BoxBrowserService] method instead.
type V1BoxBrowserService struct {
	Options []option.RequestOption
}

// NewV1BoxBrowserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1BoxBrowserService(opts ...option.RequestOption) (r V1BoxBrowserService) {
	r = V1BoxBrowserService{}
	r.Options = opts
	return
}

// Get browser CDP url
func (r *V1BoxBrowserService) CdpURL(ctx context.Context, id string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/browser/connect-url/cdp", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get browser connect url
func (r *V1BoxBrowserService) ConnectURL(ctx context.Context, id string, opts ...option.RequestOption) (res *V1BoxBrowserConnectURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/browser/connect-url", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Browser connection address information
type V1BoxBrowserConnectURLResponse struct {
	// CDP URL
	CdpURL string `json:"cdpUrl,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CdpURL      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxBrowserConnectURLResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxBrowserConnectURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
