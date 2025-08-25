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
	"github.com/babelcloud/gbox-sdk-go/packages/respjson"
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
func (r *V1BoxProxyService) Get(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxProxyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/proxy", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set the proxy for the box
func (r *V1BoxProxyService) Set(ctx context.Context, boxID string, body V1BoxProxySetParams, opts ...option.RequestOption) (res *V1BoxProxySetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/proxy", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Box Http Proxy
type V1BoxProxyGetResponse struct {
	// The host address of the proxy server
	Host string `json:"host,required"`
	// The port number of the proxy server
	Port float64 `json:"port,required"`
	// Box Proxy Auth
	Auth V1BoxProxyGetResponseAuth `json:"auth"`
	// List of IP addresses and domains that should bypass the proxy. These addresses
	// will be accessed directly without going through the proxy server. Default is
	// ['127.0.0.1', 'localhost']
	Excludes []string `json:"excludes"`
	// PAC (Proxy Auto-Configuration) URL.
	PacURL string `json:"pacUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Host        respjson.Field
		Port        respjson.Field
		Auth        respjson.Field
		Excludes    respjson.Field
		PacURL      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxProxyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxProxyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Box Proxy Auth
type V1BoxProxyGetResponseAuth struct {
	// Password for the proxy
	Password string `json:"password,required"`
	// Username for the proxy
	Username string `json:"username,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Password    respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxProxyGetResponseAuth) RawJSON() string { return r.JSON.raw }
func (r *V1BoxProxyGetResponseAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Box Http Proxy
type V1BoxProxySetResponse struct {
	// The host address of the proxy server
	Host string `json:"host,required"`
	// The port number of the proxy server
	Port float64 `json:"port,required"`
	// Box Proxy Auth
	Auth V1BoxProxySetResponseAuth `json:"auth"`
	// List of IP addresses and domains that should bypass the proxy. These addresses
	// will be accessed directly without going through the proxy server. Default is
	// ['127.0.0.1', 'localhost']
	Excludes []string `json:"excludes"`
	// PAC (Proxy Auto-Configuration) URL.
	PacURL string `json:"pacUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Host        respjson.Field
		Port        respjson.Field
		Auth        respjson.Field
		Excludes    respjson.Field
		PacURL      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxProxySetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxProxySetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Box Proxy Auth
type V1BoxProxySetResponseAuth struct {
	// Password for the proxy
	Password string `json:"password,required"`
	// Username for the proxy
	Username string `json:"username,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Password    respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxProxySetResponseAuth) RawJSON() string { return r.JSON.raw }
func (r *V1BoxProxySetResponseAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxProxySetParams struct {
	// The host address of the proxy server
	Host string `json:"host,required"`
	// The port number of the proxy server
	Port float64 `json:"port,required"`
	// PAC (Proxy Auto-Configuration) URL.
	PacURL param.Opt[string] `json:"pacUrl,omitzero"`
	// Box Proxy Auth
	Auth V1BoxProxySetParamsAuth `json:"auth,omitzero"`
	// List of IP addresses and domains that should bypass the proxy. These addresses
	// will be accessed directly without going through the proxy server. Default is
	// ['127.0.0.1', 'localhost']
	Excludes []string `json:"excludes,omitzero"`
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
