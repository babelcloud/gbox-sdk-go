// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/babelcloud/gbox-sdk-go/internal/apijson"
	"github.com/babelcloud/gbox-sdk-go/internal/requestconfig"
	"github.com/babelcloud/gbox-sdk-go/option"
	"github.com/babelcloud/gbox-sdk-go/packages/param"
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

// Generate pre-signed CDP url
func (r *V1BoxBrowserService) CdpURL(ctx context.Context, id string, body V1BoxBrowserCdpURLParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/browser/connect-url/cdp", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1BoxBrowserCdpURLParams struct {
	// The CDP url will be alive for the given duration (e.g. '120m')
	ExpiresIn param.Opt[string] `json:"expiresIn,omitzero"`
	paramObj
}

func (r V1BoxBrowserCdpURLParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxBrowserCdpURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxBrowserCdpURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
