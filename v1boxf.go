// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/babelcloud/gbox-sdk-go/internal/apiform"
	"github.com/babelcloud/gbox-sdk-go/internal/apijson"
	"github.com/babelcloud/gbox-sdk-go/internal/apiquery"
	"github.com/babelcloud/gbox-sdk-go/internal/requestconfig"
	"github.com/babelcloud/gbox-sdk-go/option"
	"github.com/babelcloud/gbox-sdk-go/packages/param"
	"github.com/babelcloud/gbox-sdk-go/packages/respjson"
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
func (r *V1BoxFService) List(ctx context.Context, boxID string, query V1BoxFListParams, opts ...option.RequestOption) (res *V1BoxFListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/fs/list", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Check if file/dir exists
func (r *V1BoxFService) Exists(ctx context.Context, boxID string, body V1BoxFExistsParams, opts ...option.RequestOption) (res *V1BoxFExistsResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/fs/exists", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get file/dir
func (r *V1BoxFService) Info(ctx context.Context, boxID string, query V1BoxFInfoParams, opts ...option.RequestOption) (res *V1BoxFInfoResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/fs/info", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Read box file
func (r *V1BoxFService) Read(ctx context.Context, boxID string, query V1BoxFReadParams, opts ...option.RequestOption) (res *V1BoxFReadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/fs/read", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a file or dir. If target path is not exists, the delete will be failed.
func (r *V1BoxFService) Remove(ctx context.Context, boxID string, body V1BoxFRemoveParams, opts ...option.RequestOption) (res *V1BoxFRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/fs", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Rename a file or dir. If target newPath is already exists, the rename will be
// failed.
func (r *V1BoxFService) Rename(ctx context.Context, boxID string, body V1BoxFRenameParams, opts ...option.RequestOption) (res *V1BoxFRenameResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/fs/rename", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates or overwrites a file. Creates necessary directories in the path if they
// don't exist. If target path is already exists, the write will be failed.
func (r *V1BoxFService) Write(ctx context.Context, boxID string, body V1BoxFWriteParams, opts ...option.RequestOption) (res *V1BoxFWriteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/fs/write", boxID)
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
	LastModified time.Time `json:"lastModified"`
	Mode         string    `json:"mode"`
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	// This field is from variant [V1BoxFListResponseDataFile].
	Size string `json:"size"`
	Type string `json:"type"`
	JSON struct {
		LastModified respjson.Field
		Mode         respjson.Field
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
	// File metadata
	Mode string `json:"mode,required"`
	// Name of the file
	Name string `json:"name,required"`
	// Full path to the file in the box
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
		Mode         respjson.Field
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
	// Last modified time of the directory
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// Directory metadata
	Mode string `json:"mode,required"`
	// Name of the directory
	Name string `json:"name,required"`
	// Full path to the directory in the box
	Path string `json:"path,required"`
	// Directory type indicator
	//
	// Any of "dir".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Mode         respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFListResponseDataDirectory) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFListResponseDataDirectory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxFExistsResponseUnion contains all possible properties and values from
// [V1BoxFExistsResponseExistsFileDirectoryResult],
// [V1BoxFExistsResponseNotExistsFileDirectoryResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxFExistsResponseUnion struct {
	Exists bool `json:"exists"`
	// This field is from variant [V1BoxFExistsResponseExistsFileDirectoryResult].
	Type string `json:"type"`
	JSON struct {
		Exists respjson.Field
		Type   respjson.Field
		raw    string
	} `json:"-"`
}

func (u V1BoxFExistsResponseUnion) AsExistsFileDirectoryResult() (v V1BoxFExistsResponseExistsFileDirectoryResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxFExistsResponseUnion) AsNotExistsFileDirectoryResult() (v V1BoxFExistsResponseNotExistsFileDirectoryResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxFExistsResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxFExistsResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response after checking if a file/directory exists
type V1BoxFExistsResponseExistsFileDirectoryResult struct {
	// Exists
	Exists bool `json:"exists,required"`
	// Type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exists      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFExistsResponseExistsFileDirectoryResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFExistsResponseExistsFileDirectoryResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response after checking if a file/directory not exists
type V1BoxFExistsResponseNotExistsFileDirectoryResult struct {
	// Exists
	Exists bool `json:"exists,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exists      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFExistsResponseNotExistsFileDirectoryResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFExistsResponseNotExistsFileDirectoryResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxFInfoResponseUnion contains all possible properties and values from
// [V1BoxFInfoResponseFile], [V1BoxFInfoResponseDirectory].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxFInfoResponseUnion struct {
	LastModified time.Time `json:"lastModified"`
	Mode         string    `json:"mode"`
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	// This field is from variant [V1BoxFInfoResponseFile].
	Size string `json:"size"`
	Type string `json:"type"`
	JSON struct {
		LastModified respjson.Field
		Mode         respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		Type         respjson.Field
		raw          string
	} `json:"-"`
}

func (u V1BoxFInfoResponseUnion) AsFile() (v V1BoxFInfoResponseFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxFInfoResponseUnion) AsDirectory() (v V1BoxFInfoResponseDirectory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxFInfoResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxFInfoResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File system file representation
type V1BoxFInfoResponseFile struct {
	// Last modified time of the file
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// File metadata
	Mode string `json:"mode,required"`
	// Name of the file
	Name string `json:"name,required"`
	// Full path to the file in the box
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
		Mode         respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFInfoResponseFile) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFInfoResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File system directory representation
type V1BoxFInfoResponseDirectory struct {
	// Last modified time of the directory
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// Directory metadata
	Mode string `json:"mode,required"`
	// Name of the directory
	Name string `json:"name,required"`
	// Full path to the directory in the box
	Path string `json:"path,required"`
	// Directory type indicator
	//
	// Any of "dir".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Mode         respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFInfoResponseDirectory) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFInfoResponseDirectory) UnmarshalJSON(data []byte) error {
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

// Response after deleting file/directory
type V1BoxFRemoveResponse struct {
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
func (r V1BoxFRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxFRenameResponseUnion contains all possible properties and values from
// [V1BoxFRenameResponseFile], [V1BoxFRenameResponseDirectory].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxFRenameResponseUnion struct {
	LastModified time.Time `json:"lastModified"`
	Mode         string    `json:"mode"`
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	// This field is from variant [V1BoxFRenameResponseFile].
	Size string `json:"size"`
	Type string `json:"type"`
	JSON struct {
		LastModified respjson.Field
		Mode         respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		Type         respjson.Field
		raw          string
	} `json:"-"`
}

func (u V1BoxFRenameResponseUnion) AsFile() (v V1BoxFRenameResponseFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxFRenameResponseUnion) AsDirectory() (v V1BoxFRenameResponseDirectory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxFRenameResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxFRenameResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File system file representation
type V1BoxFRenameResponseFile struct {
	// Last modified time of the file
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// File metadata
	Mode string `json:"mode,required"`
	// Name of the file
	Name string `json:"name,required"`
	// Full path to the file in the box
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
		Mode         respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFRenameResponseFile) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFRenameResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File system directory representation
type V1BoxFRenameResponseDirectory struct {
	// Last modified time of the directory
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// Directory metadata
	Mode string `json:"mode,required"`
	// Name of the directory
	Name string `json:"name,required"`
	// Full path to the directory in the box
	Path string `json:"path,required"`
	// Directory type indicator
	//
	// Any of "dir".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Mode         respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxFRenameResponseDirectory) RawJSON() string { return r.JSON.raw }
func (r *V1BoxFRenameResponseDirectory) UnmarshalJSON(data []byte) error {
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
	// Target directory path in the box
	Path string `query:"path,required" json:"-"`
	// Depth of the directory
	Depth param.Opt[float64] `query:"depth,omitzero" json:"-"`
	// Working directory. If not provided, the file will be read from the
	// `box.config.workingDir` directory.
	WorkingDir param.Opt[string] `query:"workingDir,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BoxFListParams]'s query parameters as `url.Values`.
func (r V1BoxFListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1BoxFExistsParams struct {
	// Target path in the box. If the path does not start with '/', the file/directory
	// will be checked relative to the working directory
	Path string `json:"path,required"`
	// Working directory. If not provided, the file will be read from the
	// `box.config.workingDir` directory.
	WorkingDir param.Opt[string] `json:"workingDir,omitzero"`
	paramObj
}

func (r V1BoxFExistsParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxFExistsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxFExistsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxFInfoParams struct {
	// Target path in the box. If the path does not start with '/', the file/directory
	// will be checked relative to the working directory
	Path string `query:"path,required" json:"-"`
	// Working directory. If not provided, the file will be read from the
	// `box.config.workingDir` directory.
	WorkingDir param.Opt[string] `query:"workingDir,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BoxFInfoParams]'s query parameters as `url.Values`.
func (r V1BoxFInfoParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1BoxFReadParams struct {
	// Target path in the box. If the path does not start with '/', the file will be
	// read from the working directory.
	Path string `query:"path,required" json:"-"`
	// Working directory. If not provided, the file will be read from the
	// `box.config.workingDir` directory.
	WorkingDir param.Opt[string] `query:"workingDir,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BoxFReadParams]'s query parameters as `url.Values`.
func (r V1BoxFReadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1BoxFRemoveParams struct {
	// Target path in the box. If the path does not start with '/', the file/directory
	// will be deleted relative to the working directory. If the target path does not
	// exist, the delete will fail.
	Path string `json:"path,required"`
	// Working directory. If not provided, the file will be read from the
	// `box.config.workingDir` directory.
	WorkingDir param.Opt[string] `json:"workingDir,omitzero"`
	paramObj
}

func (r V1BoxFRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxFRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxFRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxFRenameParams struct {
	// New path in the box. If the path does not start with '/', the file/directory
	// will be renamed relative to the working directory. If the newPath already
	// exists, the rename will fail.
	NewPath string `json:"newPath,required"`
	// Old path in the box. If the path does not start with '/', the file/directory
	// will be renamed relative to the working directory. If the oldPath does not
	// exist, the rename will fail.
	OldPath string `json:"oldPath,required"`
	// Working directory. If not provided, the file will be read from the
	// `box.config.workingDir` directory.
	WorkingDir param.Opt[string] `json:"workingDir,omitzero"`
	paramObj
}

func (r V1BoxFRenameParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxFRenameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxFRenameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxFWriteParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Request
	// parameters for writing content to a file
	OfWriteFile *V1BoxFWriteParamsBodyWriteFile `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Request
	// parameters for writing binary content to a file
	OfWriteFileByBinary *V1BoxFWriteParamsBodyWriteFileByBinary `json:",inline"`

	paramObj
}

func (r V1BoxFWriteParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Request parameters for writing content to a file
//
// The properties Content, Path are required.
type V1BoxFWriteParamsBodyWriteFile struct {
	// Content of the file (Max size: 512MB)
	Content string `json:"content,required"`
	// Target path in the box. If the path does not start with '/', the file will be
	// written relative to the working directory. Creates necessary directories in the
	// path if they don't exist. If the target path already exists, the write will
	// fail.
	Path string `json:"path,required"`
	// Working directory. If not provided, the file will be read from the
	// `box.config.workingDir` directory.
	WorkingDir param.Opt[string] `json:"workingDir,omitzero"`
	paramObj
}

func (r V1BoxFWriteParamsBodyWriteFile) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxFWriteParamsBodyWriteFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxFWriteParamsBodyWriteFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request parameters for writing binary content to a file
//
// The properties Content, Path are required.
type V1BoxFWriteParamsBodyWriteFileByBinary struct {
	// Binary content of the file (Max file size: 512MB)
	Content io.Reader `json:"content,omitzero,required" format:"binary"`
	// Target path in the box. If the path does not start with '/', the file will be
	// written relative to the working directory. Creates necessary directories in the
	// path if they don't exist. If the target path already exists, the write will
	// fail.
	Path string `json:"path,required"`
	// Working directory. If not provided, the file will be read from the
	// `box.config.workingDir` directory.
	WorkingDir param.Opt[string] `json:"workingDir,omitzero"`
	paramObj
}

func (r V1BoxFWriteParamsBodyWriteFileByBinary) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxFWriteParamsBodyWriteFileByBinary
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxFWriteParamsBodyWriteFileByBinary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
