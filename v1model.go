// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk

import (
	"context"
	"net/http"
	"slices"

	"github.com/babelcloud/gbox-sdk-go/internal/apijson"
	"github.com/babelcloud/gbox-sdk-go/internal/requestconfig"
	"github.com/babelcloud/gbox-sdk-go/option"
	"github.com/babelcloud/gbox-sdk-go/packages/param"
	"github.com/babelcloud/gbox-sdk-go/packages/respjson"
)

// V1ModelService contains methods and other services that help with interacting
// with the gbox-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ModelService] method instead.
type V1ModelService struct {
	Options []option.RequestOption
}

// NewV1ModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1ModelService(opts ...option.RequestOption) (r V1ModelService) {
	r = V1ModelService{}
	r.Options = opts
	return
}

// Generate coordinates for a model
func (r *V1ModelService) Call(ctx context.Context, body V1ModelCallParams, opts ...option.RequestOption) (res *V1ModelCallResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "model"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Model response data structure
type V1ModelCallResponse struct {
	// Unique ID of this request, can be used for issue reporting and feedback
	ID string `json:"id,required"`
	// Model response data
	Response any `json:"response,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ModelCallResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ModelCallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ModelCallParams struct {
	// Structured action object (click or drag)
	Action any `json:"action,omitzero,required"`
	// HTTP(S) URL to screenshot image
	Screenshot string `json:"screenshot,required"`
	// Model to use
	//
	// Any of "gbox-handy-1".
	Model V1ModelCallParamsModel `json:"model,omitzero"`
	paramObj
}

func (r V1ModelCallParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ModelCallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ModelCallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model to use
type V1ModelCallParamsModel string

const (
	V1ModelCallParamsModelGboxHandy1 V1ModelCallParamsModel = "gbox-handy-1"
)
