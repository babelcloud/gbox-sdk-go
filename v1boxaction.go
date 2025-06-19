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

// Click
func (r *V1BoxActionService) Click(ctx context.Context, id string, body V1BoxActionClickParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/click", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Drag
func (r *V1BoxActionService) Drag(ctx context.Context, id string, body V1BoxActionDragParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/drag", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Move to position
func (r *V1BoxActionService) Move(ctx context.Context, id string, body V1BoxActionMoveParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/move", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Press button on the device. like power button, volume up button, volume down
// button, etc.
func (r *V1BoxActionService) PressButton(ctx context.Context, id string, body V1BoxActionPressButtonParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/press-button", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates pressing a specific key by triggering the complete keyboard key event
// chain (keydown, keypress, keyup). Use this to activate keyboard key event
// listeners such as shortcuts or form submissions.
func (r *V1BoxActionService) PressKey(ctx context.Context, id string, body V1BoxActionPressKeyParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/press-key", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Take screenshot
func (r *V1BoxActionService) Screenshot(ctx context.Context, id string, body V1BoxActionScreenshotParams, opts ...option.RequestOption) (res *V1BoxActionScreenshotResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/screenshot", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scroll
func (r *V1BoxActionService) Scroll(ctx context.Context, id string, body V1BoxActionScrollParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/scroll", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Performs a swipe in the specified direction
func (r *V1BoxActionService) Swipe(ctx context.Context, id string, body V1BoxActionSwipeParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/swipe", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Touch
func (r *V1BoxActionService) Touch(ctx context.Context, id string, body V1BoxActionTouchParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/touch", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Directly inputs text content without triggering physical key events (keydown,
// etc.), ideal for quickly filling large amounts of text when intermediate input
// events aren't needed.
func (r *V1BoxActionService) Type(ctx context.Context, id string, body V1BoxActionTypeParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/type", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Result of an UI action execution
type ActionResult struct {
	// Complete screenshot result with operation trace, before and after images
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

// Complete screenshot result with operation trace, before and after images
type ActionResultScreenshot struct {
	// Screenshot taken after action execution
	After ActionResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before ActionResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace ActionResultScreenshotTrace `json:"trace,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Before      respjson.Field
		Trace       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResultScreenshot) RawJSON() string { return r.JSON.raw }
func (r *ActionResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
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

// Screenshot taken before action execution
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

// Screenshot with action operation trace
type ActionResultScreenshotTrace struct {
	// URI of the screenshot with operation trace
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResultScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *ActionResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of screenshot capture action
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
	// X coordinate of the click
	X float64 `json:"x,required"`
	// Y coordinate of the click
	Y float64 `json:"y,required"`
	// Whether to perform a double click
	Double param.Opt[bool] `json:"double,omitzero"`
	// Delay after performing the action, before taking the final screenshot.
	//
	// Execution flow:
	//
	// 1. Take screenshot before action
	// 2. Perform the action
	// 3. Wait for screenshotDelay (this parameter)
	// 4. Take screenshot after action
	//
	// Example: '500ms' means wait 500ms after the action before capturing the final
	// screenshot.
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Mouse button to click
	//
	// Any of "left", "right", "middle".
	Button V1BoxActionClickParamsButton `json:"button,omitzero"`
	// Type of the URI. default is base64.
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

// Type of the URI. default is base64.
type V1BoxActionClickParamsOutputFormat string

const (
	V1BoxActionClickParamsOutputFormatBase64     V1BoxActionClickParamsOutputFormat = "base64"
	V1BoxActionClickParamsOutputFormatStorageKey V1BoxActionClickParamsOutputFormat = "storageKey"
)

type V1BoxActionDragParams struct {
	// Path of the drag action as a series of coordinates
	Path []V1BoxActionDragParamsPath `json:"path,omitzero,required"`
	// Time interval between points (e.g. "50ms")
	Duration param.Opt[string] `json:"duration,omitzero"`
	// Delay after performing the action, before taking the final screenshot.
	//
	// Execution flow:
	//
	// 1. Take screenshot before action
	// 2. Perform the action
	// 3. Wait for screenshotDelay (this parameter)
	// 4. Take screenshot after action
	//
	// Example: '500ms' means wait 500ms after the action before capturing the final
	// screenshot.
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Type of the URI. default is base64.
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

// Single point in a drag path
//
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

// Type of the URI. default is base64.
type V1BoxActionDragParamsOutputFormat string

const (
	V1BoxActionDragParamsOutputFormatBase64     V1BoxActionDragParamsOutputFormat = "base64"
	V1BoxActionDragParamsOutputFormatStorageKey V1BoxActionDragParamsOutputFormat = "storageKey"
)

type V1BoxActionMoveParams struct {
	// X coordinate to move to
	X float64 `json:"x,required"`
	// Y coordinate to move to
	Y float64 `json:"y,required"`
	// Delay after performing the action, before taking the final screenshot.
	//
	// Execution flow:
	//
	// 1. Take screenshot before action
	// 2. Perform the action
	// 3. Wait for screenshotDelay (this parameter)
	// 4. Take screenshot after action
	//
	// Example: '500ms' means wait 500ms after the action before capturing the final
	// screenshot.
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Type of the URI. default is base64.
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

// Type of the URI. default is base64.
type V1BoxActionMoveParamsOutputFormat string

const (
	V1BoxActionMoveParamsOutputFormatBase64     V1BoxActionMoveParamsOutputFormat = "base64"
	V1BoxActionMoveParamsOutputFormatStorageKey V1BoxActionMoveParamsOutputFormat = "storageKey"
)

type V1BoxActionPressButtonParams struct {
	// Button to press
	//
	// Any of "power", "volumeUp", "volumeDown", "volumeMute", "home", "back", "menu",
	// "appSwitch".
	Buttons []string `json:"buttons,omitzero,required"`
	// Delay after performing the action, before taking the final screenshot.
	//
	// Execution flow:
	//
	// 1. Take screenshot before action
	// 2. Perform the action
	// 3. Wait for screenshotDelay (this parameter)
	// 4. Take screenshot after action
	//
	// Example: '500ms' means wait 500ms after the action before capturing the final
	// screenshot.
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionPressButtonParamsOutputFormat `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionPressButtonParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionPressButtonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionPressButtonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the URI. default is base64.
type V1BoxActionPressButtonParamsOutputFormat string

const (
	V1BoxActionPressButtonParamsOutputFormatBase64     V1BoxActionPressButtonParamsOutputFormat = "base64"
	V1BoxActionPressButtonParamsOutputFormatStorageKey V1BoxActionPressButtonParamsOutputFormat = "storageKey"
)

type V1BoxActionPressKeyParams struct {
	// This is an array of keyboard keys to press. Supports cross-platform
	// compatibility.
	//
	// Any of "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
	// "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3",
	// "4", "5", "6", "7", "8", "9", "f1", "f2", "f3", "f4", "f5", "f6", "f7", "f8",
	// "f9", "f10", "f11", "f12", "control", "alt", "shift", "meta", "win", "cmd",
	// "option", "arrowUp", "arrowDown", "arrowLeft", "arrowRight", "home", "end",
	// "pageUp", "pageDown", "enter", "space", "tab", "escape", "backspace", "delete",
	// "insert", "capsLock", "numLock", "scrollLock", "pause", "printScreen", ";", "=",
	// ",", "-", ".", "/", "`", "[", "\\", "]", "'", "numpad0", "numpad1", "numpad2",
	// "numpad3", "numpad4", "numpad5", "numpad6", "numpad7", "numpad8", "numpad9",
	// "numpadAdd", "numpadSubtract", "numpadMultiply", "numpadDivide",
	// "numpadDecimal", "numpadEnter", "numpadEqual", "volumeUp", "volumeDown",
	// "volumeMute", "mediaPlayPause", "mediaStop", "mediaNextTrack",
	// "mediaPreviousTrack".
	Keys []string `json:"keys,omitzero,required"`
	// Delay after performing the action, before taking the final screenshot.
	//
	// Execution flow:
	//
	// 1. Take screenshot before action
	// 2. Perform the action
	// 3. Wait for screenshotDelay (this parameter)
	// 4. Take screenshot after action
	//
	// Example: '500ms' means wait 500ms after the action before capturing the final
	// screenshot.
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionPressKeyParamsOutputFormat `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionPressKeyParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionPressKeyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionPressKeyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the URI. default is base64.
type V1BoxActionPressKeyParamsOutputFormat string

const (
	V1BoxActionPressKeyParamsOutputFormatBase64     V1BoxActionPressKeyParamsOutputFormat = "base64"
	V1BoxActionPressKeyParamsOutputFormatStorageKey V1BoxActionPressKeyParamsOutputFormat = "storageKey"
)

type V1BoxActionScreenshotParams struct {
	// Clipping region for screenshot capture
	Clip V1BoxActionScreenshotParamsClip `json:"clip,omitzero"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionScreenshotParamsOutputFormat `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionScreenshotParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScreenshotParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScreenshotParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Clipping region for screenshot capture
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

// Type of the URI. default is base64.
type V1BoxActionScreenshotParamsOutputFormat string

const (
	V1BoxActionScreenshotParamsOutputFormatBase64     V1BoxActionScreenshotParamsOutputFormat = "base64"
	V1BoxActionScreenshotParamsOutputFormatStorageKey V1BoxActionScreenshotParamsOutputFormat = "storageKey"
)

type V1BoxActionScrollParams struct {
	// Horizontal scroll amount
	ScrollX float64 `json:"scrollX,required"`
	// Vertical scroll amount
	ScrollY float64 `json:"scrollY,required"`
	// X coordinate of the scroll position
	X float64 `json:"x,required"`
	// Y coordinate of the scroll position
	Y float64 `json:"y,required"`
	// Delay after performing the action, before taking the final screenshot.
	//
	// Execution flow:
	//
	// 1. Take screenshot before action
	// 2. Perform the action
	// 3. Wait for screenshotDelay (this parameter)
	// 4. Take screenshot after action
	//
	// Example: '500ms' means wait 500ms after the action before capturing the final
	// screenshot.
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Type of the URI. default is base64.
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

// Type of the URI. default is base64.
type V1BoxActionScrollParamsOutputFormat string

const (
	V1BoxActionScrollParamsOutputFormatBase64     V1BoxActionScrollParamsOutputFormat = "base64"
	V1BoxActionScrollParamsOutputFormatStorageKey V1BoxActionScrollParamsOutputFormat = "storageKey"
)

type V1BoxActionSwipeParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Simple
	// swipe action configuration
	OfSwipeSimple *V1BoxActionSwipeParamsBodySwipeSimple `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Swipe
	// action configuration
	OfSwipeAction *V1BoxActionSwipeParamsBodySwipeAction `json:",inline"`

	paramObj
}

func (u V1BoxActionSwipeParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSwipeSimple, u.OfSwipeAction)
}
func (r *V1BoxActionSwipeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Simple swipe action configuration
//
// The property Direction is required.
type V1BoxActionSwipeParamsBodySwipeSimple struct {
	// Direction of the swipe
	//
	// Any of "up", "down", "left", "right", "upLeft", "upRight", "downLeft",
	// "downRight".
	Direction string `json:"direction,omitzero,required"`
	// Distance of the swipe in pixels. If not provided, will use a default distance
	// based on screen size
	Distance param.Opt[float64] `json:"distance,omitzero"`
	// Duration of the swipe
	Duration param.Opt[string] `json:"duration,omitzero"`
	paramObj
}

func (r V1BoxActionSwipeParamsBodySwipeSimple) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSwipeParamsBodySwipeSimple
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSwipeParamsBodySwipeSimple) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionSwipeParamsBodySwipeSimple](
		"direction", "up", "down", "left", "right", "upLeft", "upRight", "downLeft", "downRight",
	)
}

// Swipe action configuration
//
// The properties End, Start are required.
type V1BoxActionSwipeParamsBodySwipeAction struct {
	// End point of the swipe path
	End any `json:"end,omitzero,required"`
	// Start point of the swipe path
	Start any `json:"start,omitzero,required"`
	// Duration of the swipe
	Duration param.Opt[string] `json:"duration,omitzero"`
	// Delay after performing the action, before taking the final screenshot.
	//
	// Execution flow:
	//
	// 1. Take screenshot before action
	// 2. Perform the action
	// 3. Wait for screenshotDelay (this parameter)
	// 4. Take screenshot after action
	//
	// Example: '500ms' means wait 500ms after the action before capturing the final
	// screenshot.
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionSwipeParamsBodySwipeAction) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSwipeParamsBodySwipeAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSwipeParamsBodySwipeAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionSwipeParamsBodySwipeAction](
		"outputFormat", "base64", "storageKey",
	)
}

type V1BoxActionTouchParams struct {
	// Array of touch points and their actions
	Points []V1BoxActionTouchParamsPoint `json:"points,omitzero,required"`
	// Delay after performing the action, before taking the final screenshot.
	//
	// Execution flow:
	//
	// 1. Take screenshot before action
	// 2. Perform the action
	// 3. Wait for screenshotDelay (this parameter)
	// 4. Take screenshot after action
	//
	// Example: '500ms' means wait 500ms after the action before capturing the final
	// screenshot.
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Type of the URI. default is base64.
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

// Touch point configuration with start position and actions
//
// The property Start is required.
type V1BoxActionTouchParamsPoint struct {
	// Initial touch point position
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

// Initial touch point position
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

// Type of the URI. default is base64.
type V1BoxActionTouchParamsOutputFormat string

const (
	V1BoxActionTouchParamsOutputFormatBase64     V1BoxActionTouchParamsOutputFormat = "base64"
	V1BoxActionTouchParamsOutputFormatStorageKey V1BoxActionTouchParamsOutputFormat = "storageKey"
)

type V1BoxActionTypeParams struct {
	// Text to type
	Text string `json:"text,required"`
	// Delay after performing the action, before taking the final screenshot.
	//
	// Execution flow:
	//
	// 1. Take screenshot before action
	// 2. Perform the action
	// 3. Wait for screenshotDelay (this parameter)
	// 4. Take screenshot after action
	//
	// Example: '500ms' means wait 500ms after the action before capturing the final
	// screenshot.
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Type of the URI. default is base64.
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

// Type of the URI. default is base64.
type V1BoxActionTypeParamsOutputFormat string

const (
	V1BoxActionTypeParamsOutputFormatBase64     V1BoxActionTypeParamsOutputFormat = "base64"
	V1BoxActionTypeParamsOutputFormatStorageKey V1BoxActionTypeParamsOutputFormat = "storageKey"
)
