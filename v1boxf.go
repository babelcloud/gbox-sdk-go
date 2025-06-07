// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gbox-sdk-go/internal/apijson"
	"github.com/stainless-sdks/gbox-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/gbox-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/gbox-sdk-go/option"
	"github.com/stainless-sdks/gbox-sdk-go/packages/param"
	"github.com/stainless-sdks/gbox-sdk-go/packages/respjson"
)

// V1BoxFService contains methods and other services that help with interacting
// with the gbox-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BoxFService] method instead.
type V1BoxFService struct {
	Options []option.RequestOption
}

// NewV1BoxFService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1BoxFService(opts ...option.RequestOption) (r V1BoxFService) {
	r = V1BoxFService{}
	r.Options = opts
	return
}

// List box files
func (r *V1BoxFService) List(ctx context.Context, id string, query V1BoxFListParams, opts ...option.RequestOption) (res *V1BoxFListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/fs/list", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Read box file
func (r *V1BoxFService) Read(ctx context.Context, id string, query V1BoxFReadParams, opts ...option.RequestOption) (res *V1BoxFReadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/fs/read", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Write box file
func (r *V1BoxFService) Write(ctx context.Context, id string, body V1BoxFWriteParams, opts ...option.RequestOption) (res *V1BoxFWriteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/fs/write", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response containing directory listing results
type V1BoxFListResponse struct {
	// Array of files and directories
	Data []V1BoxFListResponseDataUnion `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxFListResponseDataUnion contains all possible properties and values from
// [V1BoxFListResponseDataFile], [V1BoxFListResponseDataDirectory].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxFListResponseDataUnion struct {
	// This field is from variant [V1BoxFListResponseDataFile].
	LastModified time.Time `json:"lastModified"`
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	// This field is from variant [V1BoxFListResponseDataFile].
	Size string `json:"size"`
	Type string `json:"type"`
	JSON struct {
		LastModified respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		Type         respjson.Field
		raw          string
	} `json:"-"`
}

func (u V1BoxFListResponseDataUnion) AsFile() (v V1BoxFListResponseDataFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxFListResponseDataUnion) AsDirectory() (v V1BoxFListResponseDataDirectory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxFListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxFListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File system file representation
type V1BoxFListResponseDataFile struct {
	// Last modified time of the file
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// Name of the file
	Name string `json:"name,required"`
	// Full path to the file
	Path string `json:"path,required"`
	// Size of the file
	Size string `json:"size,required"`
	// File type indicator
	//
	// Any of "file".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFListResponseDataFile) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFListResponseDataFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File system directory representation
type V1BoxFListResponseDataDirectory struct {
	// Name of the directory
	Name string `json:"name,required"`
	// Full path to the directory
	Path string `json:"path,required"`
	// Directory type indicator
	//
	// Any of "dir".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFListResponseDataDirectory) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFListResponseDataDirectory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing file content
type V1BoxFReadResponse struct {
	// Content of the file
	Content string `json:"content,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFReadResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFReadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response after writing file content
type V1BoxFWriteResponse struct {
	// Success message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFWriteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFWriteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxFListParams struct {
	// Path to the directory
	Path string `query:"path,required" json:"-"`
	// Depth of the directory
	Depth param.Opt[float64] `query:"depth,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BoxFListParams]'s query parameters as `url.Values`.
func (r V1BoxFListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1BoxFReadParams struct {
	// Path to the file
	Path string `query:"path,required" json:"-"`
	paramObj
}

// URLQuery serializes [V1BoxFReadParams]'s query parameters as `url.Values`.
func (r V1BoxFReadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1BoxFWriteParams struct {
	// Content of the file
	Content string `json:"content,required"`
	// Path to the file
	Path string `json:"path,required"`
	paramObj
}

func (r V1BoxFWriteParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxFWriteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxFWriteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
