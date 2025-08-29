// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk

import (
	"context"
	"net/http"

	"github.com/babelcloud/gbox-sdk-go/internal/apijson"
	"github.com/babelcloud/gbox-sdk-go/internal/requestconfig"
	"github.com/babelcloud/gbox-sdk-go/option"
	"github.com/babelcloud/gbox-sdk-go/packages/respjson"
)

// V1DeviceService contains methods and other services that help with interacting
// with the gbox-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1DeviceService] method instead.
type V1DeviceService struct {
	Options []option.RequestOption
}

// NewV1DeviceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1DeviceService(opts ...option.RequestOption) (r V1DeviceService) {
	r = V1DeviceService{}
	r.Options = opts
	return
}

// Get device list
func (r *V1DeviceService) List(ctx context.Context, opts ...option.RequestOption) (res *GetDeviceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DeviceInfo struct {
	// Device ID
	DeviceID string `json:"deviceId,required"`
	// Device enable status
	Enable string `json:"enable,required"`
	// Whether device is idle
	IsIdle bool `json:"isIdle,required"`
	// Product model from ro.product.model
	ProductModel string `json:"productModel,required"`
	// Provider ID
	ProviderID string `json:"providerId,required"`
	// Provider type
	ProviderType string `json:"providerType,required"`
	// Device status
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceID     respjson.Field
		Enable       respjson.Field
		IsIdle       respjson.Field
		ProductModel respjson.Field
		ProviderID   respjson.Field
		ProviderType respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceInfo) RawJSON() string { return r.JSON.raw }
func (r *DeviceInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetDeviceListResponse struct {
	// List of devices
	Data []DeviceInfo `json:"data,required"`
	// Response message
	Message string `json:"message,required"`
	// Total number of devices
	Total float64 `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Message     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetDeviceListResponse) RawJSON() string { return r.JSON.raw }
func (r *GetDeviceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
