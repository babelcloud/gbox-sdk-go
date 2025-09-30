// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

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

// Simulates a click action on the box
func (r *V1BoxActionService) Click(ctx context.Context, boxID string, body V1BoxActionClickParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/click", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the clipboard content
func (r *V1BoxActionService) ClipboardGet(ctx context.Context, boxID string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/clipboard", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set the clipboard content
func (r *V1BoxActionService) ClipboardSet(ctx context.Context, boxID string, body V1BoxActionClipboardSetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/clipboard", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Simulates a drag gesture, moving from a start point to an end point over a set
// duration. Supports simple start/end coordinates, multi-point drag paths, and
// natural-language targets.
func (r *V1BoxActionService) Drag(ctx context.Context, boxID string, body V1BoxActionDragParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/drag", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detect and identify interactive UI elements in the current screen. Note: This
// feature currently only supports element detection within a running browser. If
// the browser is not running, the Elements array will be empty.
func (r *V1BoxActionService) ElementsDetect(ctx context.Context, boxID string, body V1BoxActionElementsDetectParams, opts ...option.RequestOption) (res *V1BoxActionElementsDetectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/elements/detect", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Extract data from the UI interface using a JSON schema.
func (r *V1BoxActionService) Extract(ctx context.Context, boxID string, body V1BoxActionExtractParams, opts ...option.RequestOption) (res *V1BoxActionExtractResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/extract", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a long press action at specified coordinates for a specified duration.
// Useful for triggering context menus, drag operations, or other long-press
// interactions.
func (r *V1BoxActionService) LongPress(ctx context.Context, boxID string, body V1BoxActionLongPressParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/long-press", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Moves the focus to a specific coordinate on the box without performing a click
// or tap. Use this endpoint to position the cursor, hover over elements, or
// prepare for chained actions such as drag or swipe.
func (r *V1BoxActionService) Move(ctx context.Context, boxID string, body V1BoxActionMoveParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/move", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Press device buttons like power, volume, home, back, etc.
func (r *V1BoxActionService) PressButton(ctx context.Context, boxID string, body V1BoxActionPressButtonParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/press-button", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Simulates pressing a specific key by triggering the complete keyboard key event
// chain (keydown, keypress, keyup). Use this to activate keyboard key event
// listeners such as shortcuts or form submissions.
func (r *V1BoxActionService) PressKey(ctx context.Context, boxID string, body V1BoxActionPressKeyParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/press-key", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Start recording the box screen. Only one recording can be active at a time. If a
// recording is already in progress, starting a new recording will stop the
// previous one and keep only the latest recording.
func (r *V1BoxActionService) RecordingStart(ctx context.Context, boxID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/recording/start", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Stop recording the box screen
func (r *V1BoxActionService) RecordingStop(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxActionRecordingStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/recording/stop", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Disable the device's background screen rewind recording.
func (r *V1BoxActionService) RewindDisable(ctx context.Context, boxID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/recording/rewind", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Enable the device's background screen rewind recording.
func (r *V1BoxActionService) RewindEnable(ctx context.Context, boxID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/recording/rewind", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Rewind and capture the device's background screen recording from a specified
// time period.
func (r *V1BoxActionService) RewindExtract(ctx context.Context, boxID string, body V1BoxActionRewindExtractParams, opts ...option.RequestOption) (res *V1BoxActionRewindExtractResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/recording/rewind/extract", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the current structured screen layout information. This endpoint returns
// detailed structural information about the UI elements currently displayed on the
// screen, which can be used for UI automation, element analysis, and accessibility
// purposes. The format varies by box type: Android boxes return XML format with
// detailed UI hierarchy information including element bounds, text content,
// resource IDs, and properties, while other box types may return different
// structured formats.
func (r *V1BoxActionService) ScreenLayout(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxActionScreenLayoutResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/screen-layout", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Rotates the screen orientation. Note that even after rotating the screen,
// applications or system layouts may not automatically adapt to the gravity sensor
// changes, so visual changes may not always occur.
func (r *V1BoxActionService) ScreenRotation(ctx context.Context, boxID string, body V1BoxActionScreenRotationParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/screen-rotation", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Captures a screenshot of the current box screen
func (r *V1BoxActionService) Screenshot(ctx context.Context, boxID string, body V1BoxActionScreenshotParams, opts ...option.RequestOption) (res *V1BoxActionScreenshotResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/screenshot", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Performs a scroll action. Supports both advanced scroll with coordinates and
// simple scroll with direction.
func (r *V1BoxActionService) Scroll(ctx context.Context, boxID string, body V1BoxActionScrollParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/scroll", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the action settings for the box
func (r *V1BoxActionService) Settings(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxActionSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/settings", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resets the box settings to default
func (r *V1BoxActionService) SettingsReset(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxActionSettingsResetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/settings", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Update the action settings for the box
func (r *V1BoxActionService) SettingsUpdate(ctx context.Context, boxID string, body V1BoxActionSettingsUpdateParams, opts ...option.RequestOption) (res *V1BoxActionSettingsUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/settings", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Performs a swipe in the specified direction
func (r *V1BoxActionService) Swipe(ctx context.Context, boxID string, body V1BoxActionSwipeParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/swipe", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Tap action for Android devices using ADB input tap command
func (r *V1BoxActionService) Tap(ctx context.Context, boxID string, body V1BoxActionTapParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/tap", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Performs more advanced touch gestures. Use this endpoint to simulate realistic
// behaviors.
func (r *V1BoxActionService) Touch(ctx context.Context, boxID string, body V1BoxActionTouchParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/touch", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Directly inputs text content without triggering physical key events (keydown,
// etc.), ideal for quickly filling large amounts of text when intermediate input
// events aren't needed.
func (r *V1BoxActionService) Type(ctx context.Context, boxID string, body V1BoxActionTypeParams, opts ...option.RequestOption) (res *ActionResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/type", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Action common options
type ActionCommonOptionsParam struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot ActionCommonOptionsScreenshotUnionParam `json:"screenshot,omitzero"`
	paramObj
}

func (r ActionCommonOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow ActionCommonOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionCommonOptionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ActionCommonOptionsScreenshotUnionParam struct {
	OfActionScreenshotOptions *ActionScreenshotOptionsParam `json:",omitzero,inline"`
	OfBool                    param.Opt[bool]               `json:",omitzero,inline"`
	paramUnion
}

func (u ActionCommonOptionsScreenshotUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfActionScreenshotOptions, u.OfBool)
}
func (u *ActionCommonOptionsScreenshotUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ActionCommonOptionsScreenshotUnionParam) asAny() any {
	if !param.IsOmitted(u.OfActionScreenshotOptions) {
		return u.OfActionScreenshotOptions
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Result of an UI action execution with optional screenshots
type ActionResult struct {
	// Unique identifier for each action. Use this ID to locate the action and report
	// issues.
	ActionID string `json:"actionId,required"`
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot ActionResultScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		Message     respjson.Field
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
	After ActionResultScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before ActionResultScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace ActionResultScreenshotTrace `json:"trace"`
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
	// Presigned url of the screenshot before the action
	PresignedURL string `json:"presignedUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri          respjson.Field
		PresignedURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
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
	// Presigned url of the screenshot before the action
	PresignedURL string `json:"presignedUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri          respjson.Field
		PresignedURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
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

// Action screenshot options
type ActionScreenshotOptionsParam struct {
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	Delay param.Opt[string] `json:"delay,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat ActionScreenshotOptionsOutputFormat `json:"outputFormat,omitzero"`
	// Specify which screenshot phases to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three phases. Can specify one or multiple in an array. If
	// empty array is provided, no screenshots will be taken.
	//
	// Any of "before", "after", "trace".
	Phases []string `json:"phases,omitzero"`
	paramObj
}

func (r ActionScreenshotOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow ActionScreenshotOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionScreenshotOptionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the URI. default is base64.
type ActionScreenshotOptionsOutputFormat string

const (
	ActionScreenshotOptionsOutputFormatBase64     ActionScreenshotOptionsOutputFormat = "base64"
	ActionScreenshotOptionsOutputFormatStorageKey ActionScreenshotOptionsOutputFormat = "storageKey"
)

// Detected UI element
type DetectedElement struct {
	// Element id
	ID string `json:"id,required"`
	// Element center x coordinate relative to screen
	CenterX float64 `json:"centerX,required"`
	// Element center y coordinate relative to screen
	CenterY float64 `json:"centerY,required"`
	// Element height
	Height float64 `json:"height,required"`
	// A human-readable identifier generated from the element's visible attributes to
	// help understand what this element represents. For images, it uses alt text or
	// filename; for links, it uses text content or href; for buttons, it uses text
	// content or aria-label; for inputs, it uses placeholder or value; etc.
	Label string `json:"label,required"`
	// Element path
	Path string `json:"path,required"`
	// Element source
	Source string `json:"source,required"`
	// Element type
	Type string `json:"type,required"`
	// Element width
	Width float64 `json:"width,required"`
	// Element x coordinate relative to screen
	X float64 `json:"x,required"`
	// Element y coordinate relative to screen
	Y float64 `json:"y,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CenterX     respjson.Field
		CenterY     respjson.Field
		Height      respjson.Field
		Label       respjson.Field
		Path        respjson.Field
		Source      respjson.Field
		Type        respjson.Field
		Width       respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetectedElement) RawJSON() string { return r.JSON.raw }
func (r *DetectedElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DetectedElement to a DetectedElementParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DetectedElementParam.Overrides()
func (r DetectedElement) ToParam() DetectedElementParam {
	return param.Override[DetectedElementParam](json.RawMessage(r.RawJSON()))
}

// Detected UI element
//
// The properties ID, CenterX, CenterY, Height, Label, Path, Source, Type, Width,
// X, Y are required.
type DetectedElementParam struct {
	// Element id
	ID string `json:"id,required"`
	// Element center x coordinate relative to screen
	CenterX float64 `json:"centerX,required"`
	// Element center y coordinate relative to screen
	CenterY float64 `json:"centerY,required"`
	// Element height
	Height float64 `json:"height,required"`
	// A human-readable identifier generated from the element's visible attributes to
	// help understand what this element represents. For images, it uses alt text or
	// filename; for links, it uses text content or href; for buttons, it uses text
	// content or aria-label; for inputs, it uses placeholder or value; etc.
	Label string `json:"label,required"`
	// Element path
	Path string `json:"path,required"`
	// Element source
	Source string `json:"source,required"`
	// Element type
	Type string `json:"type,required"`
	// Element width
	Width float64 `json:"width,required"`
	// Element x coordinate relative to screen
	X float64 `json:"x,required"`
	// Element y coordinate relative to screen
	Y float64 `json:"y,required"`
	paramObj
}

func (r DetectedElementParam) MarshalJSON() (data []byte, err error) {
	type shadow DetectedElementParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DetectedElementParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result containing original screenshot, annotated screenshot, and detected
// elements
type V1BoxActionElementsDetectResponse struct {
	// Detected UI elements
	Elements []DetectedElement `json:"elements,required"`
	// Detected elements screenshot
	Screenshot V1BoxActionElementsDetectResponseScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Elements    respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionElementsDetectResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionElementsDetectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detected elements screenshot
type V1BoxActionElementsDetectResponseScreenshot struct {
	// Result of screenshot capture action
	Marked V1BoxActionElementsDetectResponseScreenshotMarked `json:"marked,required"`
	// Result of screenshot capture action
	Source V1BoxActionElementsDetectResponseScreenshotSource `json:"source,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Marked      respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionElementsDetectResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionElementsDetectResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of screenshot capture action
type V1BoxActionElementsDetectResponseScreenshotMarked struct {
	// URL of the screenshot
	Uri string `json:"uri,required"`
	// Presigned url of the screenshot
	PresignedURL string `json:"presignedUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri          respjson.Field
		PresignedURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionElementsDetectResponseScreenshotMarked) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionElementsDetectResponseScreenshotMarked) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of screenshot capture action
type V1BoxActionElementsDetectResponseScreenshotSource struct {
	// URL of the screenshot
	Uri string `json:"uri,required"`
	// Presigned url of the screenshot
	PresignedURL string `json:"presignedUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri          respjson.Field
		PresignedURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionElementsDetectResponseScreenshotSource) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionElementsDetectResponseScreenshotSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of extract action execution
type V1BoxActionExtractResponse struct {
	// The extracted data structure that conforms to the provided JSON schema. The
	// actual structure and content depend on the schema defined in the extract action
	// request.
	Data map[string]any `json:"data,required"`
	// Base64-encoded screenshot of the UI interface at the time of extraction
	Screenshot string `json:"screenshot,required" format:"byte"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionExtractResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionExtractResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recording stop result
type V1BoxActionRecordingStopResponse struct {
	// Presigned URL of the recording. This is a temporary downloadable URL with an
	// expiration time for accessing the recording file.
	PresignedURL string `json:"presignedUrl,required"`
	// Storage key of the recording. Before the box is deleted, you can use this
	// storageKey with the endpoint `box/:boxId/storage/presigned-url` to get a
	// downloadable URL for the recording.
	StorageKey string `json:"storageKey,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PresignedURL respjson.Field
		StorageKey   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionRecordingStopResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionRecordingStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of extracting the recording rewind
type V1BoxActionRewindExtractResponse struct {
	// Presigned URL of the recording. This is a temporary downloadable URL with an
	// expiration time for accessing the recording file.
	PresignedURL string `json:"presignedUrl,required"`
	// Storage key of the recording. Before the box is deleted, you can use this
	// storageKey with the endpoint `box/:boxId/storage/presigned-url` to get a
	// downloadable URL for the recording.
	StorageKey string `json:"storageKey,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PresignedURL respjson.Field
		StorageKey   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionRewindExtractResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionRewindExtractResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screen layout content.
//
// Android boxes (XML):
//
// <?xml version='1.0' encoding='UTF-8' standalone='yes'?>
// <hierarchy rotation="0">
//
//	<node ... />
//
// </hierarchy>
//
// Browser (Linux) boxes (HTML):
//
// <html>
//
//	<head><title>Example</title></head>
//	<body>
//	  <h1>Hello World</h1>
//	</body>
//
// </html>
type V1BoxActionScreenLayoutResponse struct {
	// Screen layout content.
	//
	// Android boxes (XML):
	//
	// ```xml
	// <?xml version='1.0' encoding='UTF-8' standalone='yes'?>
	// <hierarchy rotation="0">
	//
	//	<node ... />
	//
	// </hierarchy>
	// ```
	//
	// Browser (Linux) boxes (HTML):
	//
	// ```html
	// <html>
	//
	//	<head>
	//	  <title>Example</title>
	//	</head>
	//	<body>
	//	  <h1>Hello World</h1>
	//	</body>
	//
	// </html>
	// ```
	Content string `json:"content,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionScreenLayoutResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScreenLayoutResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of screenshot capture action
type V1BoxActionScreenshotResponse struct {
	// URL of the screenshot
	Uri string `json:"uri,required"`
	// Presigned url of the screenshot
	PresignedURL string `json:"presignedUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri          respjson.Field
		PresignedURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionScreenshotResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScreenshotResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action setting
type V1BoxActionSettingsResponse struct {
	// The scale of the action to be performed. Must be greater than 0.1 and less than
	// or equal to 1.
	//
	// Notes:
	//
	//   - Scale does not change the box's actual screen resolution.
	//   - It affects the size of the output screenshot and the coordinates/distances of
	//     actions. Coordinates and distances are scaled by this factor. Example: when
	//     scale = 1, Click({x:100, y:100}); when scale = 0.5, the equivalent position is
	//     Click({x:50, y:50}).
	Scale float64 `json:"scale,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scale       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action setting
type V1BoxActionSettingsResetResponse struct {
	// The scale of the action to be performed. Must be greater than 0.1 and less than
	// or equal to 1.
	//
	// Notes:
	//
	//   - Scale does not change the box's actual screen resolution.
	//   - It affects the size of the output screenshot and the coordinates/distances of
	//     actions. Coordinates and distances are scaled by this factor. Example: when
	//     scale = 1, Click({x:100, y:100}); when scale = 0.5, the equivalent position is
	//     Click({x:50, y:50}).
	Scale float64 `json:"scale,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scale       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionSettingsResetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSettingsResetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action setting
type V1BoxActionSettingsUpdateResponse struct {
	// The scale of the action to be performed. Must be greater than 0.1 and less than
	// or equal to 1.
	//
	// Notes:
	//
	//   - Scale does not change the box's actual screen resolution.
	//   - It affects the size of the output screenshot and the coordinates/distances of
	//     actions. Coordinates and distances are scaled by this factor. Example: when
	//     scale = 1, Click({x:100, y:100}); when scale = 0.5, the equivalent position is
	//     Click({x:50, y:50}).
	Scale float64 `json:"scale,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scale       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionSettingsUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSettingsUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionClickParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Mouse
	// click action configuration
	OfClickAction *V1BoxActionClickParamsBodyClickAction `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Click
	// action configuration with natural language
	OfClickActionWithNaturalLanguage *V1BoxActionClickParamsBodyClickActionWithNaturalLanguage `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Click
	// action configuration by element
	OfClickActionByElement *V1BoxActionClickParamsBodyClickActionByElement `json:",inline"`

	paramObj
}

func (u V1BoxActionClickParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfClickAction, u.OfClickActionWithNaturalLanguage, u.OfClickActionByElement)
}
func (r *V1BoxActionClickParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mouse click action configuration
//
// The properties X, Y are required.
type V1BoxActionClickParamsBodyClickAction struct {
	// X coordinate of the click
	X float64 `json:"x,required"`
	// Y coordinate of the click
	Y float64 `json:"y,required"`
	// Whether to perform a double click
	Double param.Opt[bool] `json:"double,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Mouse button to click
	//
	// Any of "left", "right", "middle".
	Button string `json:"button,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionClickParamsBodyClickAction) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionClickParamsBodyClickAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionClickParamsBodyClickAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionClickParamsBodyClickAction](
		"button", "left", "right", "middle",
	)
	apijson.RegisterFieldValidator[V1BoxActionClickParamsBodyClickAction](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionClickParamsBodyClickAction](
		"outputFormat", "base64", "storageKey",
	)
}

// Click action configuration with natural language
//
// The property Target is required.
type V1BoxActionClickParamsBodyClickActionWithNaturalLanguage struct {
	// Describe the target to operate using natural language, e.g., 'login button' or
	// 'Chrome'.
	Target string `json:"target,required"`
	// Whether to perform a double click
	Double param.Opt[bool] `json:"double,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Mouse button to click
	//
	// Any of "left", "right", "middle".
	Button string `json:"button,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionClickParamsBodyClickActionWithNaturalLanguage) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionClickParamsBodyClickActionWithNaturalLanguage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionClickParamsBodyClickActionWithNaturalLanguage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionClickParamsBodyClickActionWithNaturalLanguage](
		"button", "left", "right", "middle",
	)
	apijson.RegisterFieldValidator[V1BoxActionClickParamsBodyClickActionWithNaturalLanguage](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionClickParamsBodyClickActionWithNaturalLanguage](
		"outputFormat", "base64", "storageKey",
	)
}

// Click action configuration by element
//
// The property Target is required.
type V1BoxActionClickParamsBodyClickActionByElement struct {
	// Detected UI element
	Target DetectedElementParam `json:"target,omitzero,required"`
	// Whether to perform a double click
	Double param.Opt[bool] `json:"double,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Mouse button to click
	//
	// Any of "left", "right", "middle".
	Button string `json:"button,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionClickParamsBodyClickActionByElement) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionClickParamsBodyClickActionByElement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionClickParamsBodyClickActionByElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionClickParamsBodyClickActionByElement](
		"button", "left", "right", "middle",
	)
	apijson.RegisterFieldValidator[V1BoxActionClickParamsBodyClickActionByElement](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionClickParamsBodyClickActionByElement](
		"outputFormat", "base64", "storageKey",
	)
}

type V1BoxActionClipboardSetParams struct {
	// The content to set the clipboard content
	Content string `json:"content,required"`
	paramObj
}

func (r V1BoxActionClipboardSetParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionClipboardSetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionClipboardSetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionDragParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Drag
	// action configuration with start and end points.
	//
	// Operation flow:
	//
	// 1. Touch finger at "start" coordinates
	// 2. Move to "end" coordinates within the "duration" time and lift finger
	OfDragSimple *V1BoxActionDragParamsBodyDragSimple `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Drag
	// action configuration with path points
	OfDragAdvanced *V1BoxActionDragParamsBodyDragAdvanced `json:",inline"`

	paramObj
}

func (u V1BoxActionDragParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDragSimple, u.OfDragAdvanced)
}
func (r *V1BoxActionDragParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Drag action configuration with start and end points.
//
// Operation flow:
//
// 1. Touch finger at "start" coordinates
// 2. Move to "end" coordinates within the "duration" time and lift finger
//
// The properties End, Start are required.
type V1BoxActionDragParamsBodyDragSimple struct {
	// End point of the drag path (coordinates or natural language)
	End V1BoxActionDragParamsBodyDragSimpleEndUnion `json:"end,omitzero,required"`
	// Start point of the drag path (coordinates or natural language)
	Start V1BoxActionDragParamsBodyDragSimpleStartUnion `json:"start,omitzero,required"`
	// Duration to complete the movement from start to end coordinates
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionDragParamsBodyDragSimple) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsBodyDragSimple
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsBodyDragSimple) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionDragParamsBodyDragSimple](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionDragParamsBodyDragSimple](
		"outputFormat", "base64", "storageKey",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionDragParamsBodyDragSimpleEndUnion struct {
	OfDragPathPoint *V1BoxActionDragParamsBodyDragSimpleEndDragPathPoint `json:",omitzero,inline"`
	OfString        param.Opt[string]                                    `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionDragParamsBodyDragSimpleEndUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDragPathPoint, u.OfString)
}
func (u *V1BoxActionDragParamsBodyDragSimpleEndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionDragParamsBodyDragSimpleEndUnion) asAny() any {
	if !param.IsOmitted(u.OfDragPathPoint) {
		return u.OfDragPathPoint
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Single point in a drag path
//
// The properties X, Y are required.
type V1BoxActionDragParamsBodyDragSimpleEndDragPathPoint struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionDragParamsBodyDragSimpleEndDragPathPoint) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsBodyDragSimpleEndDragPathPoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsBodyDragSimpleEndDragPathPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionDragParamsBodyDragSimpleStartUnion struct {
	OfDragPathPoint *V1BoxActionDragParamsBodyDragSimpleStartDragPathPoint `json:",omitzero,inline"`
	OfString        param.Opt[string]                                      `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionDragParamsBodyDragSimpleStartUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDragPathPoint, u.OfString)
}
func (u *V1BoxActionDragParamsBodyDragSimpleStartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionDragParamsBodyDragSimpleStartUnion) asAny() any {
	if !param.IsOmitted(u.OfDragPathPoint) {
		return u.OfDragPathPoint
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Single point in a drag path
//
// The properties X, Y are required.
type V1BoxActionDragParamsBodyDragSimpleStartDragPathPoint struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionDragParamsBodyDragSimpleStartDragPathPoint) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsBodyDragSimpleStartDragPathPoint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsBodyDragSimpleStartDragPathPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Drag action configuration with path points
//
// The property Path is required.
type V1BoxActionDragParamsBodyDragAdvanced struct {
	// Path of the drag action as a series of coordinates
	Path []V1BoxActionDragParamsBodyDragAdvancedPath `json:"path,omitzero,required"`
	// Time interval between points (e.g. "50ms")
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 50ms
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionDragParamsBodyDragAdvanced) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsBodyDragAdvanced
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsBodyDragAdvanced) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionDragParamsBodyDragAdvanced](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionDragParamsBodyDragAdvanced](
		"outputFormat", "base64", "storageKey",
	)
}

// Single point in a drag path
//
// The properties X, Y are required.
type V1BoxActionDragParamsBodyDragAdvancedPath struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionDragParamsBodyDragAdvancedPath) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsBodyDragAdvancedPath
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsBodyDragAdvancedPath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionElementsDetectParams struct {
	// Detect elements screenshot options
	Screenshot V1BoxActionElementsDetectParamsScreenshot `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionElementsDetectParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionElementsDetectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionElementsDetectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detect elements screenshot options
type V1BoxActionElementsDetectParamsScreenshot struct {
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionElementsDetectParamsScreenshot) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionElementsDetectParamsScreenshot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionElementsDetectParamsScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionElementsDetectParamsScreenshot](
		"outputFormat", "base64", "storageKey",
	)
}

type V1BoxActionExtractParams struct {
	// The instruction of the action to extract data from the UI interface
	Instruction string `json:"instruction,required"`
	// JSON Schema defining the structure of data to extract. Supports object, array,
	// string, number, boolean types with validation rules.
	//
	// Common use cases:
	//
	//   - Extract text content: { "type": "string" }
	//   - Extract structured data: { "type": "object", "properties": {...} }
	//   - Extract lists: { "type": "array", "items": {...} }
	//   - Extract with validation: Add constraints like "required", "enum", "pattern",
	//     etc.
	Schema any `json:"schema,omitzero"`
	paramObj
}

func (r V1BoxActionExtractParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionExtractParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionExtractParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionLongPressParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Long
	// press action configuration.
	//
	// Operation flow:
	//
	// 1. Touch finger at specified coordinates
	// 2. Hold for the specified duration
	// 3. Release finger
	//
	// This is useful for triggering context menus, drag operations, or other
	// long-press interactions.
	OfLongPressAction *V1BoxActionLongPressParamsBodyLongPressAction `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Long
	// press action configuration using natural language target
	OfLongPressActionWithNaturalLanguage *V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguage `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Long
	// press action configuration by element
	OfLongPressActionByElement *V1BoxActionLongPressParamsBodyLongPressActionByElement `json:",inline"`

	paramObj
}

func (u V1BoxActionLongPressParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLongPressAction, u.OfLongPressActionWithNaturalLanguage, u.OfLongPressActionByElement)
}
func (r *V1BoxActionLongPressParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Long press action configuration.
//
// Operation flow:
//
// 1. Touch finger at specified coordinates
// 2. Hold for the specified duration
// 3. Release finger
//
// This is useful for triggering context menus, drag operations, or other
// long-press interactions.
//
// The properties X, Y are required.
type V1BoxActionLongPressParamsBodyLongPressAction struct {
	// X coordinate of the long press
	X float64 `json:"x,required"`
	// Y coordinate of the long press
	Y float64 `json:"y,required"`
	// Duration to hold the press (e.g. '1s', '500ms')
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 1s
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionLongPressParamsBodyLongPressAction) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionLongPressParamsBodyLongPressAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionLongPressParamsBodyLongPressAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionLongPressParamsBodyLongPressAction](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionLongPressParamsBodyLongPressAction](
		"outputFormat", "base64", "storageKey",
	)
}

// Long press action configuration using natural language target
//
// The property Target is required.
type V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguage struct {
	// Describe the target to operate using natural language, e.g., 'Chrome icon',
	// 'login button'
	Target string `json:"target,required"`
	// Duration to hold the press (e.g. '1s', '500ms')
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 1s
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguage) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguage](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguage](
		"outputFormat", "base64", "storageKey",
	)
}

// Long press action configuration by element
//
// The property Target is required.
type V1BoxActionLongPressParamsBodyLongPressActionByElement struct {
	// Detected UI element
	Target DetectedElementParam `json:"target,omitzero,required"`
	// Duration to hold the press (e.g. '1s', '500ms')
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 1s
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionLongPressParamsBodyLongPressActionByElement) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionLongPressParamsBodyLongPressActionByElement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionLongPressParamsBodyLongPressActionByElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionLongPressParamsBodyLongPressActionByElement](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionLongPressParamsBodyLongPressActionByElement](
		"outputFormat", "base64", "storageKey",
	)
}

type V1BoxActionMoveParams struct {
	// X coordinate to move to
	X float64 `json:"x,required"`
	// Y coordinate to move to
	Y float64 `json:"y,required"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model V1BoxActionMoveParamsModel `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
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

// Model to use for natural-language target resolution. Defaults to 'uitars'.
type V1BoxActionMoveParamsModel string

const (
	V1BoxActionMoveParamsModelGpt5   V1BoxActionMoveParamsModel = "gpt-5"
	V1BoxActionMoveParamsModelGpt4o  V1BoxActionMoveParamsModel = "gpt-4o"
	V1BoxActionMoveParamsModelUitars V1BoxActionMoveParamsModel = "uitars"
	V1BoxActionMoveParamsModelCua    V1BoxActionMoveParamsModel = "cua"
)

// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
// default is base64. This field will be ignored when `options.screenshot` is
// provided.
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
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model V1BoxActionPressButtonParamsModel `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
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

// Model to use for natural-language target resolution. Defaults to 'uitars'.
type V1BoxActionPressButtonParamsModel string

const (
	V1BoxActionPressButtonParamsModelGpt5   V1BoxActionPressButtonParamsModel = "gpt-5"
	V1BoxActionPressButtonParamsModelGpt4o  V1BoxActionPressButtonParamsModel = "gpt-4o"
	V1BoxActionPressButtonParamsModelUitars V1BoxActionPressButtonParamsModel = "uitars"
	V1BoxActionPressButtonParamsModelCua    V1BoxActionPressButtonParamsModel = "cua"
)

// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
// default is base64. This field will be ignored when `options.screenshot` is
// provided.
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
	// Whether to press keys as combination (simultaneously) or sequentially. When
	// true, all keys are pressed together as a shortcut (e.g., Ctrl+C). When false,
	// keys are pressed one by one in sequence.
	Combination param.Opt[bool] `json:"combination,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model V1BoxActionPressKeyParamsModel `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
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

// Model to use for natural-language target resolution. Defaults to 'uitars'.
type V1BoxActionPressKeyParamsModel string

const (
	V1BoxActionPressKeyParamsModelGpt5   V1BoxActionPressKeyParamsModel = "gpt-5"
	V1BoxActionPressKeyParamsModelGpt4o  V1BoxActionPressKeyParamsModel = "gpt-4o"
	V1BoxActionPressKeyParamsModelUitars V1BoxActionPressKeyParamsModel = "uitars"
	V1BoxActionPressKeyParamsModelCua    V1BoxActionPressKeyParamsModel = "cua"
)

// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
// default is base64. This field will be ignored when `options.screenshot` is
// provided.
type V1BoxActionPressKeyParamsOutputFormat string

const (
	V1BoxActionPressKeyParamsOutputFormatBase64     V1BoxActionPressKeyParamsOutputFormat = "base64"
	V1BoxActionPressKeyParamsOutputFormatStorageKey V1BoxActionPressKeyParamsOutputFormat = "storageKey"
)

type V1BoxActionRewindExtractParams struct {
	// How far back in time to rewind for extracting recorded video. This specifies the
	// duration to go back from the current moment (e.g., '30s' rewinds 30 seconds to
	// get recent recorded activity). Default is 30s, max is 5m.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Maximum allowed: 5m
	Duration param.Opt[string] `json:"duration,omitzero"`
	paramObj
}

func (r V1BoxActionRewindExtractParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionRewindExtractParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionRewindExtractParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionScreenRotationParams struct {
	// Target screen orientation
	//
	// Any of "portrait", "landscapeLeft", "portraitUpsideDown", "landscapeRight".
	Orientation V1BoxActionScreenRotationParamsOrientation `json:"orientation,omitzero,required"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model V1BoxActionScreenRotationParamsModel `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionScreenRotationParamsOutputFormat `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionScreenRotationParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScreenRotationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScreenRotationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target screen orientation
type V1BoxActionScreenRotationParamsOrientation string

const (
	V1BoxActionScreenRotationParamsOrientationPortrait           V1BoxActionScreenRotationParamsOrientation = "portrait"
	V1BoxActionScreenRotationParamsOrientationLandscapeLeft      V1BoxActionScreenRotationParamsOrientation = "landscapeLeft"
	V1BoxActionScreenRotationParamsOrientationPortraitUpsideDown V1BoxActionScreenRotationParamsOrientation = "portraitUpsideDown"
	V1BoxActionScreenRotationParamsOrientationLandscapeRight     V1BoxActionScreenRotationParamsOrientation = "landscapeRight"
)

// Model to use for natural-language target resolution. Defaults to 'uitars'.
type V1BoxActionScreenRotationParamsModel string

const (
	V1BoxActionScreenRotationParamsModelGpt5   V1BoxActionScreenRotationParamsModel = "gpt-5"
	V1BoxActionScreenRotationParamsModelGpt4o  V1BoxActionScreenRotationParamsModel = "gpt-4o"
	V1BoxActionScreenRotationParamsModelUitars V1BoxActionScreenRotationParamsModel = "uitars"
	V1BoxActionScreenRotationParamsModelCua    V1BoxActionScreenRotationParamsModel = "cua"
)

// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
// default is base64. This field will be ignored when `options.screenshot` is
// provided.
type V1BoxActionScreenRotationParamsOutputFormat string

const (
	V1BoxActionScreenRotationParamsOutputFormatBase64     V1BoxActionScreenRotationParamsOutputFormat = "base64"
	V1BoxActionScreenRotationParamsOutputFormatStorageKey V1BoxActionScreenRotationParamsOutputFormat = "storageKey"
)

type V1BoxActionScreenshotParams struct {
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// Whether to save the screenshot to the device screenshot album
	SaveToAlbum param.Opt[bool] `json:"saveToAlbum,omitzero"`
	// Clipping region for screenshot capture
	Clip V1BoxActionScreenshotParamsClip `json:"clip,omitzero"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionScreenshotParamsOutputFormat `json:"outputFormat,omitzero"`
	// Scroll capture parameters
	ScrollCapture V1BoxActionScreenshotParamsScrollCapture `json:"scrollCapture,omitzero"`
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

// Scroll capture parameters
type V1BoxActionScreenshotParamsScrollCapture struct {
	// Maximum height of the screenshot in pixels. Limits the maximum height of the
	// automatically scrolled content. Useful for managing memory usage when capturing
	// tall content like long web pages. Default: 4000px
	MaxHeight param.Opt[float64] `json:"maxHeight,omitzero"`
	// Whether to scroll back to the original position after capturing the screenshot
	ScrollBack param.Opt[bool] `json:"scrollBack,omitzero"`
	paramObj
}

func (r V1BoxActionScreenshotParamsScrollCapture) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScreenshotParamsScrollCapture
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScreenshotParamsScrollCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionScrollParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Advanced scroll action configuration. The scroll will be performed from the
	// specified coordinates with the given scroll amounts. Use positive scrollY to
	// scroll content downward (reveal content below), negative scrollY to scroll
	// content upward (reveal content above). Use positive scrollX to scroll content
	// rightward (reveal content on the right), negative scrollX to scroll content
	// leftward (reveal content on the left).
	OfScrollAdvanced *V1BoxActionScrollParamsBodyScrollAdvanced `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Simple
	// scroll action configuration. The scroll will be performed from the center of the
	// screen towards the specified direction.
	OfScrollSimple *V1BoxActionScrollParamsBodyScrollSimple `json:",inline"`

	paramObj
}

func (u V1BoxActionScrollParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScrollAdvanced, u.OfScrollSimple)
}
func (r *V1BoxActionScrollParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced scroll action configuration. The scroll will be performed from the
// specified coordinates with the given scroll amounts. Use positive scrollY to
// scroll content downward (reveal content below), negative scrollY to scroll
// content upward (reveal content above). Use positive scrollX to scroll content
// rightward (reveal content on the right), negative scrollX to scroll content
// leftward (reveal content on the left).
//
// The properties ScrollX, ScrollY, X, Y are required.
type V1BoxActionScrollParamsBodyScrollAdvanced struct {
	// Horizontal scroll amount. Positive values scroll content rightward (reveals
	// content on the right), negative values scroll content leftward (reveals content
	// on the left).
	ScrollX float64 `json:"scrollX,required"`
	// Vertical scroll amount. Positive values scroll content downward (reveals content
	// below), negative values scroll content upward (reveals content above).
	ScrollY float64 `json:"scrollY,required"`
	// X coordinate of the scroll position
	X float64 `json:"x,required"`
	// Y coordinate of the scroll position
	Y float64 `json:"y,required"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionScrollParamsBodyScrollAdvanced) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScrollParamsBodyScrollAdvanced
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScrollParamsBodyScrollAdvanced) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionScrollParamsBodyScrollAdvanced](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionScrollParamsBodyScrollAdvanced](
		"outputFormat", "base64", "storageKey",
	)
}

// Simple scroll action configuration. The scroll will be performed from the center
// of the screen towards the specified direction.
//
// The property Direction is required.
type V1BoxActionScrollParamsBodyScrollSimple struct {
	// Direction to scroll. The scroll will be performed from the center of the screen
	// towards this direction. 'up' scrolls content upward (reveals content below),
	// 'down' scrolls content downward (reveals content above), 'left' scrolls content
	// leftward (reveals content on the right), 'right' scrolls content rightward
	// (reveals content on the left).
	//
	// Any of "up", "down", "left", "right".
	Direction string `json:"direction,omitzero,required"`
	// Duration of the scroll
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Natural language description of the location where the scroll should originate.
	// If not provided, the scroll will be performed from the center of the screen.
	Location param.Opt[string] `json:"location,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Distance of the scroll. Can be either a number (in pixels) or a predefined enum
	// value (tiny, short, medium, long). If not provided, the scroll will be performed
	// from the center of the screen to the screen edge
	Distance V1BoxActionScrollParamsBodyScrollSimpleDistanceUnion `json:"distance,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionScrollParamsBodyScrollSimple) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScrollParamsBodyScrollSimple
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScrollParamsBodyScrollSimple) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionScrollParamsBodyScrollSimple](
		"direction", "up", "down", "left", "right",
	)
	apijson.RegisterFieldValidator[V1BoxActionScrollParamsBodyScrollSimple](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionScrollParamsBodyScrollSimple](
		"outputFormat", "base64", "storageKey",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionScrollParamsBodyScrollSimpleDistanceUnion struct {
	OfFloat param.Opt[float64] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfV1BoxActionScrollsBodyScrollSimpleDistanceString)
	OfV1BoxActionScrollsBodyScrollSimpleDistanceString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionScrollParamsBodyScrollSimpleDistanceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfV1BoxActionScrollsBodyScrollSimpleDistanceString)
}
func (u *V1BoxActionScrollParamsBodyScrollSimpleDistanceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionScrollParamsBodyScrollSimpleDistanceUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfV1BoxActionScrollsBodyScrollSimpleDistanceString) {
		return &u.OfV1BoxActionScrollsBodyScrollSimpleDistanceString
	}
	return nil
}

type V1BoxActionScrollParamsBodyScrollSimpleDistanceString string

const (
	V1BoxActionScrollParamsBodyScrollSimpleDistanceStringTiny   V1BoxActionScrollParamsBodyScrollSimpleDistanceString = "tiny"
	V1BoxActionScrollParamsBodyScrollSimpleDistanceStringShort  V1BoxActionScrollParamsBodyScrollSimpleDistanceString = "short"
	V1BoxActionScrollParamsBodyScrollSimpleDistanceStringMedium V1BoxActionScrollParamsBodyScrollSimpleDistanceString = "medium"
	V1BoxActionScrollParamsBodyScrollSimpleDistanceStringLong   V1BoxActionScrollParamsBodyScrollSimpleDistanceString = "long"
)

type V1BoxActionSettingsUpdateParams struct {
	// The scale of the action to be performed. Must be greater than 0.1 and less than
	// or equal to 1.
	//
	// Notes:
	//
	//   - Scale does not change the box's actual screen resolution.
	//   - It affects the size of the output screenshot and the coordinates/distances of
	//     actions. Coordinates and distances are scaled by this factor. Example: when
	//     scale = 1, Click({x:100, y:100}); when scale = 0.5, the equivalent position is
	//     Click({x:50, y:50}).
	Scale float64 `json:"scale,required"`
	paramObj
}

func (r V1BoxActionSettingsUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSettingsUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSettingsUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionSwipeParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Simple
	// swipe action configuration. The gesture will be performed from the center of the
	// screen towards the specified direction.
	OfSwipeSimple *V1BoxActionSwipeParamsBodySwipeSimple `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Swipe
	// action configuration. The gesture will start from the specified start point and
	// move towards the end point.
	OfSwipeAdvanced *V1BoxActionSwipeParamsBodySwipeAdvanced `json:",inline"`

	paramObj
}

func (u V1BoxActionSwipeParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSwipeSimple, u.OfSwipeAdvanced)
}
func (r *V1BoxActionSwipeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Simple swipe action configuration. The gesture will be performed from the center
// of the screen towards the specified direction.
//
// The property Direction is required.
type V1BoxActionSwipeParamsBodySwipeSimple struct {
	// Direction to swipe. The gesture will be performed from the center of the screen
	// towards this direction.
	//
	// Any of "up", "down", "left", "right", "upLeft", "upRight", "downLeft",
	// "downRight".
	Direction string `json:"direction,omitzero,required"`
	// Duration of the swipe
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Natural language description of the location where the swipe should originate.
	// If not provided, the swipe will be performed from the center of the screen.
	Location param.Opt[string] `json:"location,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Distance of the swipe. Can be either a number (in pixels) or a predefined enum
	// value (tiny, short, medium, long). If not provided, the swipe will be performed
	// from the center of the screen to the screen edge
	Distance V1BoxActionSwipeParamsBodySwipeSimpleDistanceUnion `json:"distance,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
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
	apijson.RegisterFieldValidator[V1BoxActionSwipeParamsBodySwipeSimple](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionSwipeParamsBodySwipeSimple](
		"outputFormat", "base64", "storageKey",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionSwipeParamsBodySwipeSimpleDistanceUnion struct {
	OfFloat param.Opt[float64] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfV1BoxActionSwipesBodySwipeSimpleDistanceString)
	OfV1BoxActionSwipesBodySwipeSimpleDistanceString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionSwipeParamsBodySwipeSimpleDistanceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfV1BoxActionSwipesBodySwipeSimpleDistanceString)
}
func (u *V1BoxActionSwipeParamsBodySwipeSimpleDistanceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionSwipeParamsBodySwipeSimpleDistanceUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfV1BoxActionSwipesBodySwipeSimpleDistanceString) {
		return &u.OfV1BoxActionSwipesBodySwipeSimpleDistanceString
	}
	return nil
}

type V1BoxActionSwipeParamsBodySwipeSimpleDistanceString string

const (
	V1BoxActionSwipeParamsBodySwipeSimpleDistanceStringTiny   V1BoxActionSwipeParamsBodySwipeSimpleDistanceString = "tiny"
	V1BoxActionSwipeParamsBodySwipeSimpleDistanceStringShort  V1BoxActionSwipeParamsBodySwipeSimpleDistanceString = "short"
	V1BoxActionSwipeParamsBodySwipeSimpleDistanceStringMedium V1BoxActionSwipeParamsBodySwipeSimpleDistanceString = "medium"
	V1BoxActionSwipeParamsBodySwipeSimpleDistanceStringLong   V1BoxActionSwipeParamsBodySwipeSimpleDistanceString = "long"
)

// Swipe action configuration. The gesture will start from the specified start
// point and move towards the end point.
//
// The properties End, Start are required.
type V1BoxActionSwipeParamsBodySwipeAdvanced struct {
	// End point of the swipe path (coordinates or natural language)
	End V1BoxActionSwipeParamsBodySwipeAdvancedEndUnion `json:"end,omitzero,required"`
	// Start point of the swipe path (coordinates or natural language)
	Start V1BoxActionSwipeParamsBodySwipeAdvancedStartUnion `json:"start,omitzero,required"`
	// Duration of the swipe
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration param.Opt[string] `json:"duration,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionSwipeParamsBodySwipeAdvanced) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSwipeParamsBodySwipeAdvanced
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSwipeParamsBodySwipeAdvanced) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionSwipeParamsBodySwipeAdvanced](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionSwipeParamsBodySwipeAdvanced](
		"outputFormat", "base64", "storageKey",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionSwipeParamsBodySwipeAdvancedEndUnion struct {
	OfSwipePath *V1BoxActionSwipeParamsBodySwipeAdvancedEndSwipePath `json:",omitzero,inline"`
	OfString    param.Opt[string]                                    `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionSwipeParamsBodySwipeAdvancedEndUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSwipePath, u.OfString)
}
func (u *V1BoxActionSwipeParamsBodySwipeAdvancedEndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionSwipeParamsBodySwipeAdvancedEndUnion) asAny() any {
	if !param.IsOmitted(u.OfSwipePath) {
		return u.OfSwipePath
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Swipe path
//
// The properties X, Y are required.
type V1BoxActionSwipeParamsBodySwipeAdvancedEndSwipePath struct {
	// Start/end x coordinate of the swipe path
	X float64 `json:"x,required"`
	// Start/end y coordinate of the swipe path
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionSwipeParamsBodySwipeAdvancedEndSwipePath) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSwipeParamsBodySwipeAdvancedEndSwipePath
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSwipeParamsBodySwipeAdvancedEndSwipePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionSwipeParamsBodySwipeAdvancedStartUnion struct {
	OfSwipePath *V1BoxActionSwipeParamsBodySwipeAdvancedStartSwipePath `json:",omitzero,inline"`
	OfString    param.Opt[string]                                      `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionSwipeParamsBodySwipeAdvancedStartUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSwipePath, u.OfString)
}
func (u *V1BoxActionSwipeParamsBodySwipeAdvancedStartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionSwipeParamsBodySwipeAdvancedStartUnion) asAny() any {
	if !param.IsOmitted(u.OfSwipePath) {
		return u.OfSwipePath
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Swipe path
//
// The properties X, Y are required.
type V1BoxActionSwipeParamsBodySwipeAdvancedStartSwipePath struct {
	// Start/end x coordinate of the swipe path
	X float64 `json:"x,required"`
	// Start/end y coordinate of the swipe path
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionSwipeParamsBodySwipeAdvancedStartSwipePath) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSwipeParamsBodySwipeAdvancedStartSwipePath
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSwipeParamsBodySwipeAdvancedStartSwipePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionTapParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Tap
	// action configuration
	OfTapAction *V1BoxActionTapParamsBodyTapAction `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Tap
	// action configuration with natural language
	OfTapActionWithNaturalLanguage *V1BoxActionTapParamsBodyTapActionWithNaturalLanguage `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Tap
	// action configuration by element
	OfTapActionByElement *V1BoxActionTapParamsBodyTapActionByElement `json:",inline"`

	paramObj
}

func (u V1BoxActionTapParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTapAction, u.OfTapActionWithNaturalLanguage, u.OfTapActionByElement)
}
func (r *V1BoxActionTapParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tap action configuration
//
// The properties X, Y are required.
type V1BoxActionTapParamsBodyTapAction struct {
	// X coordinate of the tap
	X float64 `json:"x,required"`
	// Y coordinate of the tap
	Y float64 `json:"y,required"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionTapParamsBodyTapAction) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTapParamsBodyTapAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTapParamsBodyTapAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionTapParamsBodyTapAction](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionTapParamsBodyTapAction](
		"outputFormat", "base64", "storageKey",
	)
}

// Tap action configuration with natural language
//
// The property Target is required.
type V1BoxActionTapParamsBodyTapActionWithNaturalLanguage struct {
	// Describe the target to operate using natural language, e.g., 'login button' or
	// 'Chrome'.
	Target string `json:"target,required"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionTapParamsBodyTapActionWithNaturalLanguage) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTapParamsBodyTapActionWithNaturalLanguage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTapParamsBodyTapActionWithNaturalLanguage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionTapParamsBodyTapActionWithNaturalLanguage](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionTapParamsBodyTapActionWithNaturalLanguage](
		"outputFormat", "base64", "storageKey",
	)
}

// Tap action configuration by element
//
// The property Target is required.
type V1BoxActionTapParamsBodyTapActionByElement struct {
	// Detected UI element
	Target DetectedElementParam `json:"target,omitzero,required"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	//
	// Deprecated: deprecated
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	//
	// Deprecated: deprecated
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	//
	// Deprecated: deprecated
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model string `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	//
	// Deprecated: deprecated
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionTapParamsBodyTapActionByElement) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTapParamsBodyTapActionByElement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTapParamsBodyTapActionByElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionTapParamsBodyTapActionByElement](
		"model", "gpt-5", "gpt-4o", "uitars", "cua",
	)
	apijson.RegisterFieldValidator[V1BoxActionTapParamsBodyTapActionByElement](
		"outputFormat", "base64", "storageKey",
	)
}

type V1BoxActionTouchParams struct {
	// Array of touch points and their actions
	Points []V1BoxActionTouchParamsPoint `json:"points,omitzero,required"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model V1BoxActionTouchParamsModel `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
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
	Actions []V1BoxActionTouchParamsPointActionUnion `json:"actions,omitzero"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionTouchParamsPointActionUnion struct {
	OfTouchPointMoveAction *V1BoxActionTouchParamsPointActionTouchPointMoveAction `json:",omitzero,inline"`
	OfTouchPointWaitAction *V1BoxActionTouchParamsPointActionTouchPointWaitAction `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionTouchParamsPointActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTouchPointMoveAction, u.OfTouchPointWaitAction)
}
func (u *V1BoxActionTouchParamsPointActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionTouchParamsPointActionUnion) asAny() any {
	if !param.IsOmitted(u.OfTouchPointMoveAction) {
		return u.OfTouchPointMoveAction
	} else if !param.IsOmitted(u.OfTouchPointWaitAction) {
		return u.OfTouchPointWaitAction
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1BoxActionTouchParamsPointActionUnion) GetX() *float64 {
	if vt := u.OfTouchPointMoveAction; vt != nil {
		return &vt.X
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1BoxActionTouchParamsPointActionUnion) GetY() *float64 {
	if vt := u.OfTouchPointMoveAction; vt != nil {
		return &vt.Y
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1BoxActionTouchParamsPointActionUnion) GetDuration() *string {
	if vt := u.OfTouchPointMoveAction; vt != nil {
		return (*string)(&vt.Duration)
	} else if vt := u.OfTouchPointWaitAction; vt != nil {
		return (*string)(&vt.Duration)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1BoxActionTouchParamsPointActionUnion) GetType() *string {
	if vt := u.OfTouchPointMoveAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTouchPointWaitAction; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Touch point movement action configuration
//
// The properties Duration, Type, X, Y are required.
type V1BoxActionTouchParamsPointActionTouchPointMoveAction struct {
	// Duration of the movement (e.g. "200ms")
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 200ms
	Duration string `json:"duration,required"`
	// Type of the action
	Type string `json:"type,required"`
	// Target X coordinate
	X float64 `json:"x,required"`
	// Target Y coordinate
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionTouchParamsPointActionTouchPointMoveAction) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTouchParamsPointActionTouchPointMoveAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTouchParamsPointActionTouchPointMoveAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Touch point wait action configuration
//
// The properties Duration, Type are required.
type V1BoxActionTouchParamsPointActionTouchPointWaitAction struct {
	// Duration to wait (e.g. "500ms")
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration,required"`
	// Type of the action
	Type string `json:"type,required"`
	paramObj
}

func (r V1BoxActionTouchParamsPointActionTouchPointWaitAction) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTouchParamsPointActionTouchPointWaitAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTouchParamsPointActionTouchPointWaitAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model to use for natural-language target resolution. Defaults to 'uitars'.
type V1BoxActionTouchParamsModel string

const (
	V1BoxActionTouchParamsModelGpt5   V1BoxActionTouchParamsModel = "gpt-5"
	V1BoxActionTouchParamsModelGpt4o  V1BoxActionTouchParamsModel = "gpt-4o"
	V1BoxActionTouchParamsModelUitars V1BoxActionTouchParamsModel = "uitars"
	V1BoxActionTouchParamsModelCua    V1BoxActionTouchParamsModel = "cua"
)

// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
// default is base64. This field will be ignored when `options.screenshot` is
// provided.
type V1BoxActionTouchParamsOutputFormat string

const (
	V1BoxActionTouchParamsOutputFormatBase64     V1BoxActionTouchParamsOutputFormat = "base64"
	V1BoxActionTouchParamsOutputFormatStorageKey V1BoxActionTouchParamsOutputFormat = "storageKey"
)

type V1BoxActionTypeParams struct {
	// Text to type
	Text string `json:"text,required"`
	// ⚠️ DEPRECATED: Use `options.screenshot.phases` instead. This field will be
	// ignored when `options.screenshot` is provided. Whether to include screenshots in
	// the action response. If false, the screenshot object will still be returned but
	// with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.presignedExpiresIn` instead. Presigned
	// url expires in. Only takes effect when outputFormat is storageKey. This field
	// will be ignored when `options.screenshot` is provided.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// Whether to press Enter after typing the text
	PressEnter param.Opt[bool] `json:"pressEnter,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.delay` instead. This field will be
	// ignored when `options.screenshot` is provided.
	//
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
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms Maximum allowed: 30s
	ScreenshotDelay param.Opt[string] `json:"screenshotDelay,omitzero"`
	// Text input mode: 'append' to add text to existing content, 'replace' to replace
	// all existing text
	//
	// Any of "append", "replace".
	Mode V1BoxActionTypeParamsMode `json:"mode,omitzero"`
	// Model to use for natural-language target resolution. Defaults to 'uitars'.
	//
	// Any of "gpt-5", "gpt-4o", "uitars", "cua".
	Model V1BoxActionTypeParamsModel `json:"model,omitzero"`
	// Action common options
	Options ActionCommonOptionsParam `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
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

// Text input mode: 'append' to add text to existing content, 'replace' to replace
// all existing text
type V1BoxActionTypeParamsMode string

const (
	V1BoxActionTypeParamsModeAppend  V1BoxActionTypeParamsMode = "append"
	V1BoxActionTypeParamsModeReplace V1BoxActionTypeParamsMode = "replace"
)

// Model to use for natural-language target resolution. Defaults to 'uitars'.
type V1BoxActionTypeParamsModel string

const (
	V1BoxActionTypeParamsModelGpt5   V1BoxActionTypeParamsModel = "gpt-5"
	V1BoxActionTypeParamsModelGpt4o  V1BoxActionTypeParamsModel = "gpt-4o"
	V1BoxActionTypeParamsModelUitars V1BoxActionTypeParamsModel = "uitars"
	V1BoxActionTypeParamsModelCua    V1BoxActionTypeParamsModel = "cua"
)

// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
// default is base64. This field will be ignored when `options.screenshot` is
// provided.
type V1BoxActionTypeParamsOutputFormat string

const (
	V1BoxActionTypeParamsOutputFormatBase64     V1BoxActionTypeParamsOutputFormat = "base64"
	V1BoxActionTypeParamsOutputFormatStorageKey V1BoxActionTypeParamsOutputFormat = "storageKey"
)
