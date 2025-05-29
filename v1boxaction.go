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
	"github.com/stainless-sdks/gbox-sdk-go/packages/param"
	"github.com/stainless-sdks/gbox-sdk-go/packages/respjson"
)

// V1BoxActionService contains methods and other services that help with
// interacting with the gbox-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BoxActionService] method instead.
type V1BoxActionService struct {
	Options []option.RequestOption
}

// NewV1BoxActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1BoxActionService(opts ...option.RequestOption) (r V1BoxActionService) {
	r = V1BoxActionService{}
	r.Options = opts
	return
}

func (r *V1BoxActionService) Click(ctx context.Context, id string, body V1BoxActionClickParams, opts ...option.RequestOption) (res *LinuxBox, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/boxes/%s/actions/click", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *V1BoxActionService) Drag(ctx context.Context, id string, body V1BoxActionDragParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/boxes/%s/actions/drag", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *V1BoxActionService) Move(ctx context.Context, id string, body V1BoxActionMoveParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/boxes/%s/actions/move", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *V1BoxActionService) Press(ctx context.Context, id string, body V1BoxActionPressParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/boxes/%s/actions/press", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *V1BoxActionService) Screenshot(ctx context.Context, id string, body V1BoxActionScreenshotParams, opts ...option.RequestOption) (res *V1BoxActionScreenshotResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/boxes/%s/actions/screenshot", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *V1BoxActionService) Scroll(ctx context.Context, id string, body V1BoxActionScrollParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/boxes/%s/actions/scroll", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *V1BoxActionService) Touch(ctx context.Context, id string, body V1BoxActionTouchParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/boxes/%s/actions/touch", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *V1BoxActionService) Type(ctx context.Context, id string, body V1BoxActionTypeParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/boxes/%s/actions/type", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ActionResult struct {
	// screenshot
	Screenshot ActionResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResult) RawJSON() string { return r.JSON.raw }
func (r *ActionResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// screenshot
type ActionResultScreenshot struct {
	// URI of the screenshot after the action
	After ActionResultScreenshotAfter `json:"after,required"`
	// URI of the screenshot before the action
	Before ActionResultScreenshotBefore `json:"before,required"`
	// URI of the screenshot before the action with highlight
	Highlight ActionResultScreenshotHighlight `json:"highlight,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Before      respjson.Field
		Highlight   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResultScreenshot) RawJSON() string { return r.JSON.raw }
func (r *ActionResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URI of the screenshot after the action
type ActionResultScreenshotAfter struct {
	// URI of the screenshot after the action
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResultScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *ActionResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URI of the screenshot before the action
type ActionResultScreenshotBefore struct {
	// URI of the screenshot before the action
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResultScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *ActionResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URI of the screenshot before the action with highlight
type ActionResultScreenshotHighlight struct {
	// URI of the screenshot before the action with highlight
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResultScreenshotHighlight) RawJSON() string { return r.JSON.raw }
func (r *ActionResultScreenshotHighlight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionScreenshotResponse struct {
	// URL of the screenshot
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionScreenshotResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScreenshotResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionClickParams struct {
	// Action type for mouse click
	Type any `json:"type,omitzero,required"`
	// X coordinate of the click
	X float64 `json:"x,required"`
	// Y coordinate of the click
	Y float64 `json:"y,required"`
	// Whether to perform a double click
	Double param.Opt[bool] `json:"double,omitzero"`
	// Mouse button to click
	//
	// Any of "left", "right", "middle".
	Button V1BoxActionClickParamsButton `json:"button,omitzero"`
	// Type of the URI
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionClickParamsOutputFormat `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionClickParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionClickParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionClickParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mouse button to click
type V1BoxActionClickParamsButton string

const (
	V1BoxActionClickParamsButtonLeft   V1BoxActionClickParamsButton = "left"
	V1BoxActionClickParamsButtonRight  V1BoxActionClickParamsButton = "right"
	V1BoxActionClickParamsButtonMiddle V1BoxActionClickParamsButton = "middle"
)

// Type of the URI
type V1BoxActionClickParamsOutputFormat string

const (
	V1BoxActionClickParamsOutputFormatBase64     V1BoxActionClickParamsOutputFormat = "base64"
	V1BoxActionClickParamsOutputFormatStorageKey V1BoxActionClickParamsOutputFormat = "storageKey"
)

type V1BoxActionDragParams struct {
	// Path of the drag action as a series of coordinates
	Path []V1BoxActionDragParamsPath `json:"path,omitzero,required"`
	// Action type for drag interaction
	Type any `json:"type,omitzero,required"`
	// Time interval between points (e.g. "50ms")
	Duration param.Opt[string] `json:"duration,omitzero"`
	// Type of the URI
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionDragParamsOutputFormat `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionDragParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties X, Y are required.
type V1BoxActionDragParamsPath struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionDragParamsPath) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsPath
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsPath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the URI
type V1BoxActionDragParamsOutputFormat string

const (
	V1BoxActionDragParamsOutputFormatBase64     V1BoxActionDragParamsOutputFormat = "base64"
	V1BoxActionDragParamsOutputFormatStorageKey V1BoxActionDragParamsOutputFormat = "storageKey"
)

type V1BoxActionMoveParams struct {
	// Action type for cursor movement
	Type any `json:"type,omitzero,required"`
	// X coordinate to move to
	X float64 `json:"x,required"`
	// Y coordinate to move to
	Y float64 `json:"y,required"`
	// Type of the URI
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionMoveParamsOutputFormat `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionMoveParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionMoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionMoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the URI
type V1BoxActionMoveParamsOutputFormat string

const (
	V1BoxActionMoveParamsOutputFormatBase64     V1BoxActionMoveParamsOutputFormat = "base64"
	V1BoxActionMoveParamsOutputFormatStorageKey V1BoxActionMoveParamsOutputFormat = "storageKey"
)

type V1BoxActionPressParams struct {
	// Array of keys to press
	Keys []string `json:"keys,omitzero,required"`
	// Action type for keyboard key press
	Type any `json:"type,omitzero,required"`
	// Type of the URI
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionPressParamsOutputFormat `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionPressParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionPressParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionPressParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the URI
type V1BoxActionPressParamsOutputFormat string

const (
	V1BoxActionPressParamsOutputFormatBase64     V1BoxActionPressParamsOutputFormat = "base64"
	V1BoxActionPressParamsOutputFormatStorageKey V1BoxActionPressParamsOutputFormat = "storageKey"
)

type V1BoxActionScreenshotParams struct {
	// clip of the screenshot
	Clip V1BoxActionScreenshotParamsClip `json:"clip,omitzero"`
	// Type of the URI
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionScreenshotParamsOutputFormat `json:"outputFormat,omitzero"`
	// Action type for screenshot
	//
	// Any of "png", "jpeg".
	Type V1BoxActionScreenshotParamsType `json:"type,omitzero"`
	paramObj
}

func (r V1BoxActionScreenshotParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScreenshotParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScreenshotParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// clip of the screenshot
//
// The properties Height, Width, X, Y are required.
type V1BoxActionScreenshotParamsClip struct {
	// Height of the clip
	Height float64 `json:"height,required"`
	// Width of the clip
	Width float64 `json:"width,required"`
	// X coordinate of the clip
	X float64 `json:"x,required"`
	// Y coordinate of the clip
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionScreenshotParamsClip) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScreenshotParamsClip
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScreenshotParamsClip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the URI
type V1BoxActionScreenshotParamsOutputFormat string

const (
	V1BoxActionScreenshotParamsOutputFormatBase64     V1BoxActionScreenshotParamsOutputFormat = "base64"
	V1BoxActionScreenshotParamsOutputFormatStorageKey V1BoxActionScreenshotParamsOutputFormat = "storageKey"
)

// Action type for screenshot
type V1BoxActionScreenshotParamsType string

const (
	V1BoxActionScreenshotParamsTypePng  V1BoxActionScreenshotParamsType = "png"
	V1BoxActionScreenshotParamsTypeJpeg V1BoxActionScreenshotParamsType = "jpeg"
)

type V1BoxActionScrollParams struct {
	// Horizontal scroll amount
	ScrollX float64 `json:"scrollX,required"`
	// Vertical scroll amount
	ScrollY float64 `json:"scrollY,required"`
	// Action type for scroll interaction
	Type any `json:"type,omitzero,required"`
	// X coordinate of the scroll position
	X float64 `json:"x,required"`
	// Y coordinate of the scroll position
	Y float64 `json:"y,required"`
	// Type of the URI
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionScrollParamsOutputFormat `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionScrollParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScrollParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScrollParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the URI
type V1BoxActionScrollParamsOutputFormat string

const (
	V1BoxActionScrollParamsOutputFormatBase64     V1BoxActionScrollParamsOutputFormat = "base64"
	V1BoxActionScrollParamsOutputFormatStorageKey V1BoxActionScrollParamsOutputFormat = "storageKey"
)

type V1BoxActionTouchParams struct {
	// Array of touch points and their actions
	Points []V1BoxActionTouchParamsPoint `json:"points,omitzero,required"`
	// Action type for touch interaction
	Type any `json:"type,omitzero,required"`
	// Type of the URI
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionTouchParamsOutputFormat `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionTouchParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTouchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTouchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Start is required.
type V1BoxActionTouchParamsPoint struct {
	// Starting position for touch
	Start V1BoxActionTouchParamsPointStart `json:"start,omitzero,required"`
	// Sequence of actions to perform after initial touch
	Actions []any `json:"actions,omitzero"`
	paramObj
}

func (r V1BoxActionTouchParamsPoint) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTouchParamsPoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTouchParamsPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Starting position for touch
//
// The properties X, Y are required.
type V1BoxActionTouchParamsPointStart struct {
	// Starting X coordinate
	X float64 `json:"x,required"`
	// Starting Y coordinate
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionTouchParamsPointStart) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTouchParamsPointStart
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTouchParamsPointStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the URI
type V1BoxActionTouchParamsOutputFormat string

const (
	V1BoxActionTouchParamsOutputFormatBase64     V1BoxActionTouchParamsOutputFormat = "base64"
	V1BoxActionTouchParamsOutputFormatStorageKey V1BoxActionTouchParamsOutputFormat = "storageKey"
)

type V1BoxActionTypeParams struct {
	// Text to type
	Text string `json:"text,required"`
	// Action type for typing text
	Type any `json:"type,omitzero,required"`
	// Type of the URI
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionTypeParamsOutputFormat `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionTypeParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTypeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTypeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the URI
type V1BoxActionTypeParamsOutputFormat string

const (
	V1BoxActionTypeParamsOutputFormatBase64     V1BoxActionTypeParamsOutputFormat = "base64"
	V1BoxActionTypeParamsOutputFormatStorageKey V1BoxActionTypeParamsOutputFormat = "storageKey"
)
