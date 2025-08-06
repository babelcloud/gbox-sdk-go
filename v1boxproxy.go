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

// V1BoxProxyService contains methods and other services that help with interacting
// with the gbox-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BoxProxyService] method instead.
type V1BoxProxyService struct {
	Options []option.RequestOption
}

// NewV1BoxProxyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1BoxProxyService(opts ...option.RequestOption) (r V1BoxProxyService) {
	r = V1BoxProxyService{}
	r.Options = opts
	return
}

// Clear the proxy for the box
func (r *V1BoxProxyService) Clear(ctx context.Context, boxID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/proxy", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get the proxy for the box
func (r *V1BoxProxyService) Get(ctx context.Context, boxID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/proxy", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Set the proxy for the box
func (r *V1BoxProxyService) Set(ctx context.Context, boxID string, body V1BoxProxySetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/proxy", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type V1BoxProxySetParams struct {
	// Box Proxy Auth
	Auth V1BoxProxySetParamsAuth `json:"auth,omitzero,required"`
	// Exclude IPs from the proxy. Default is ['127.0.0.1', 'localhost']
	Excludes []string `json:"excludes,omitzero,required"`
	// Proxy URL
	URL string `json:"url,required"`
	paramObj
}

func (r V1BoxProxySetParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxProxySetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxProxySetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Box Proxy Auth
//
// The properties Password, Username are required.
type V1BoxProxySetParamsAuth struct {
	// Password for the proxy
	Password string `json:"password,required"`
	// Username for the proxy
	Username string `json:"username,required"`
	paramObj
}

func (r V1BoxProxySetParamsAuth) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxProxySetParamsAuth
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxProxySetParamsAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
