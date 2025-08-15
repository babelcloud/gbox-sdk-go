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

// V1BoxStorageService contains methods and other services that help with
// interacting with the gbox-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BoxStorageService] method instead.
type V1BoxStorageService struct {
	Options []option.RequestOption
}

// NewV1BoxStorageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1BoxStorageService(opts ...option.RequestOption) (r V1BoxStorageService) {
	r = V1BoxStorageService{}
	r.Options = opts
	return
}

// Create a presigned url for a storage key. This endpoint provides a presigned url
// for a storage key, which can be used to download the file from the storage.
func (r *V1BoxStorageService) PresignedURL(ctx context.Context, boxID string, body V1BoxStoragePresignedURLParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/storage/presigned-url", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1BoxStoragePresignedURLParams struct {
	// Storage key
	StorageKey string `json:"storageKey,required"`
	// Presigned url expires in
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m Maximum allowed: 6h
	ExpiresIn param.Opt[string] `json:"expiresIn,omitzero"`
	paramObj
}

func (r V1BoxStoragePresignedURLParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxStoragePresignedURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxStoragePresignedURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
