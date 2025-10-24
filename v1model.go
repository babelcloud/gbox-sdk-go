// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk

import (
	"context"
	"encoding/json"
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
	Response V1ModelCallResponseResponseUnion `json:"response,required"`
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

// V1ModelCallResponseResponseUnion contains all possible properties and values
// from [V1ModelCallResponseResponseModelClickResponseData],
// [V1ModelCallResponseResponseModelDragResponseData],
// [V1ModelCallResponseResponseModelScrollResponseData].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1ModelCallResponseResponseUnion struct {
	// This field is a union of
	// [V1ModelCallResponseResponseModelClickResponseDataCoordinates],
	// [V1ModelCallResponseResponseModelDragResponseDataCoordinates],
	// [V1ModelCallResponseResponseModelScrollResponseDataCoordinates]
	Coordinates V1ModelCallResponseResponseUnionCoordinates `json:"coordinates"`
	Type        string                                      `json:"type"`
	JSON        struct {
		Coordinates respjson.Field
		Type        respjson.Field
		raw         string
	} `json:"-"`
}

func (u V1ModelCallResponseResponseUnion) AsModelClickResponseData() (v V1ModelCallResponseResponseModelClickResponseData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1ModelCallResponseResponseUnion) AsModelDragResponseData() (v V1ModelCallResponseResponseModelDragResponseData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1ModelCallResponseResponseUnion) AsModelScrollResponseData() (v V1ModelCallResponseResponseModelScrollResponseData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1ModelCallResponseResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1ModelCallResponseResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1ModelCallResponseResponseUnionCoordinates is an implicit subunion of
// [V1ModelCallResponseResponseUnion]. V1ModelCallResponseResponseUnionCoordinates
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1ModelCallResponseResponseUnion].
type V1ModelCallResponseResponseUnionCoordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	// This field is from variant
	// [V1ModelCallResponseResponseModelDragResponseDataCoordinates].
	Destination V1ModelCallResponseResponseModelDragResponseDataCoordinatesDestination `json:"destination"`
	// This field is from variant
	// [V1ModelCallResponseResponseModelDragResponseDataCoordinates].
	Target V1ModelCallResponseResponseModelDragResponseDataCoordinatesTarget `json:"target"`
	// This field is from variant
	// [V1ModelCallResponseResponseModelScrollResponseDataCoordinates].
	ScrollX float64 `json:"scrollX"`
	// This field is from variant
	// [V1ModelCallResponseResponseModelScrollResponseDataCoordinates].
	ScrollY float64 `json:"scrollY"`
	JSON    struct {
		X           respjson.Field
		Y           respjson.Field
		Destination respjson.Field
		Target      respjson.Field
		ScrollX     respjson.Field
		ScrollY     respjson.Field
		raw         string
	} `json:"-"`
}

func (r *V1ModelCallResponseResponseUnionCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model click response data structure
type V1ModelCallResponseResponseModelClickResponseData struct {
	// Single click result with coordinates
	Coordinates V1ModelCallResponseResponseModelClickResponseDataCoordinates `json:"coordinates,required"`
	// Action type
	//
	// Any of "click", "drag", "scroll".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ModelCallResponseResponseModelClickResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ModelCallResponseResponseModelClickResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single click result with coordinates
type V1ModelCallResponseResponseModelClickResponseDataCoordinates struct {
	// X coordinate. Returns -1 if no valid target is found.
	X float64 `json:"x,required"`
	// Y coordinate. Returns -1 if no valid target is found.
	Y float64 `json:"y,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ModelCallResponseResponseModelClickResponseDataCoordinates) RawJSON() string {
	return r.JSON.raw
}
func (r *V1ModelCallResponseResponseModelClickResponseDataCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Drag response data structure
type V1ModelCallResponseResponseModelDragResponseData struct {
	// Single drag result with target and destination coordinates
	Coordinates V1ModelCallResponseResponseModelDragResponseDataCoordinates `json:"coordinates,required"`
	// Action type
	//
	// Any of "click", "drag", "scroll".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ModelCallResponseResponseModelDragResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ModelCallResponseResponseModelDragResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single drag result with target and destination coordinates
type V1ModelCallResponseResponseModelDragResponseDataCoordinates struct {
	// X and Y coordinates. Returns -1, -1 if no valid target is found.
	Destination V1ModelCallResponseResponseModelDragResponseDataCoordinatesDestination `json:"destination,required"`
	// X and Y coordinates. Returns -1, -1 if no valid target is found.
	Target V1ModelCallResponseResponseModelDragResponseDataCoordinatesTarget `json:"target,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Destination respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ModelCallResponseResponseModelDragResponseDataCoordinates) RawJSON() string {
	return r.JSON.raw
}
func (r *V1ModelCallResponseResponseModelDragResponseDataCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X and Y coordinates. Returns -1, -1 if no valid target is found.
type V1ModelCallResponseResponseModelDragResponseDataCoordinatesDestination struct {
	// X coordinate. Returns -1 if no valid target is found.
	X float64 `json:"x,required"`
	// Y coordinate. Returns -1 if no valid target is found.
	Y float64 `json:"y,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ModelCallResponseResponseModelDragResponseDataCoordinatesDestination) RawJSON() string {
	return r.JSON.raw
}
func (r *V1ModelCallResponseResponseModelDragResponseDataCoordinatesDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// X and Y coordinates. Returns -1, -1 if no valid target is found.
type V1ModelCallResponseResponseModelDragResponseDataCoordinatesTarget struct {
	// X coordinate. Returns -1 if no valid target is found.
	X float64 `json:"x,required"`
	// Y coordinate. Returns -1 if no valid target is found.
	Y float64 `json:"y,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ModelCallResponseResponseModelDragResponseDataCoordinatesTarget) RawJSON() string {
	return r.JSON.raw
}
func (r *V1ModelCallResponseResponseModelDragResponseDataCoordinatesTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scroll response data structure
type V1ModelCallResponseResponseModelScrollResponseData struct {
	// Single scroll result with location and direction
	Coordinates V1ModelCallResponseResponseModelScrollResponseDataCoordinates `json:"coordinates,required"`
	// Action type
	//
	// Any of "click", "drag", "scroll".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ModelCallResponseResponseModelScrollResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ModelCallResponseResponseModelScrollResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single scroll result with location and direction
type V1ModelCallResponseResponseModelScrollResponseDataCoordinates struct {
	// Horizontal scroll amount
	ScrollX float64 `json:"scrollX,required"`
	// Vertical scroll amount
	ScrollY float64 `json:"scrollY,required"`
	// X coordinate
	X float64 `json:"x,required"`
	// Y coordinate
	Y float64 `json:"y,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScrollX     respjson.Field
		ScrollY     respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ModelCallResponseResponseModelScrollResponseDataCoordinates) RawJSON() string {
	return r.JSON.raw
}
func (r *V1ModelCallResponseResponseModelScrollResponseDataCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ModelCallParams struct {
	// Structured action object (click or drag)
	Action V1ModelCallParamsActionUnion `json:"action,omitzero,required"`
	// Screenshot image as HTTP(S) URL or base64-encoded data URI. Supports both
	// formats: 1) HTTP(S) URL pointing to an image file; 2) Base64-encoded data URI
	// with format 'data:image/png;base64,[data]' or 'data:image/jpeg;base64,[data]'.
	// Only PNG and JPEG formats are supported for base64.
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1ModelCallParamsActionUnion struct {
	OfClickAction  *V1ModelCallParamsActionClickAction  `json:",omitzero,inline"`
	OfDragAction   *V1ModelCallParamsActionDragAction   `json:",omitzero,inline"`
	OfScrollAction *V1ModelCallParamsActionScrollAction `json:",omitzero,inline"`
	paramUnion
}

func (u V1ModelCallParamsActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfClickAction, u.OfDragAction, u.OfScrollAction)
}
func (u *V1ModelCallParamsActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1ModelCallParamsActionUnion) asAny() any {
	if !param.IsOmitted(u.OfClickAction) {
		return u.OfClickAction
	} else if !param.IsOmitted(u.OfDragAction) {
		return u.OfDragAction
	} else if !param.IsOmitted(u.OfScrollAction) {
		return u.OfScrollAction
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ModelCallParamsActionUnion) GetDestination() *string {
	if vt := u.OfDragAction; vt != nil {
		return &vt.Destination
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ModelCallParamsActionUnion) GetDirection() *string {
	if vt := u.OfScrollAction; vt != nil {
		return &vt.Direction
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ModelCallParamsActionUnion) GetLocation() *string {
	if vt := u.OfScrollAction; vt != nil {
		return &vt.Location
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ModelCallParamsActionUnion) GetTarget() *string {
	if vt := u.OfClickAction; vt != nil {
		return (*string)(&vt.Target)
	} else if vt := u.OfDragAction; vt != nil {
		return (*string)(&vt.Target)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1ModelCallParamsActionUnion) GetType() *string {
	if vt := u.OfClickAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDragAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfScrollAction; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Click action structure
//
// The properties Target, Type are required.
type V1ModelCallParamsActionClickAction struct {
	// Natural language description of what to click
	Target string `json:"target,required"`
	// Action type
	//
	// Any of "click", "drag", "scroll".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r V1ModelCallParamsActionClickAction) MarshalJSON() (data []byte, err error) {
	type shadow V1ModelCallParamsActionClickAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ModelCallParamsActionClickAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ModelCallParamsActionClickAction](
		"type", "click", "drag", "scroll",
	)
}

// Drag action structure
//
// The properties Destination, Target, Type are required.
type V1ModelCallParamsActionDragAction struct {
	// Natural language description of ending position
	Destination string `json:"destination,required"`
	// Natural language description of starting position
	Target string `json:"target,required"`
	// Action type
	//
	// Any of "click", "drag", "scroll".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r V1ModelCallParamsActionDragAction) MarshalJSON() (data []byte, err error) {
	type shadow V1ModelCallParamsActionDragAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ModelCallParamsActionDragAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ModelCallParamsActionDragAction](
		"type", "click", "drag", "scroll",
	)
}

// Scroll action structure
//
// The properties Direction, Location, Type are required.
type V1ModelCallParamsActionScrollAction struct {
	// Scroll direction
	//
	// Any of "up", "down", "left", "right".
	Direction string `json:"direction,omitzero,required"`
	// Natural language description of the location where the scroll should originate.
	Location string `json:"location,required"`
	// Action type
	//
	// Any of "click", "drag", "scroll".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r V1ModelCallParamsActionScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow V1ModelCallParamsActionScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ModelCallParamsActionScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ModelCallParamsActionScrollAction](
		"direction", "up", "down", "left", "right",
	)
	apijson.RegisterFieldValidator[V1ModelCallParamsActionScrollAction](
		"type", "click", "drag", "scroll",
	)
}

// Model to use
type V1ModelCallParamsModel string

const (
	V1ModelCallParamsModelGboxHandy1 V1ModelCallParamsModel = "gbox-handy-1"
)
