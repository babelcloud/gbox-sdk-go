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

// Use natural language instructions to perform UI operations on the box. The
// endpoint will stream progress events before and after the action is executed. If
// you don't need intermediate events, set stream to false.
func (r *V1BoxActionService) AI(ctx context.Context, boxID string, body V1BoxActionAIParams, opts ...option.RequestOption) (res *V1BoxActionAIResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/ai", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Click
func (r *V1BoxActionService) Click(ctx context.Context, boxID string, body V1BoxActionClickParams, opts ...option.RequestOption) (res *V1BoxActionClickResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/click", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Drag
func (r *V1BoxActionService) Drag(ctx context.Context, boxID string, body V1BoxActionDragParams, opts ...option.RequestOption) (res *V1BoxActionDragResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/drag", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Extract data from the UI interface using a JSON schema.
func (r *V1BoxActionService) Extract(ctx context.Context, boxID string, body V1BoxActionExtractParams, opts ...option.RequestOption) (res *V1BoxActionExtractResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *V1BoxActionService) LongPress(ctx context.Context, boxID string, body V1BoxActionLongPressParams, opts ...option.RequestOption) (res *V1BoxActionLongPressResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/long-press", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Move to position
func (r *V1BoxActionService) Move(ctx context.Context, boxID string, body V1BoxActionMoveParams, opts ...option.RequestOption) (res *V1BoxActionMoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/move", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Press device buttons like power, volume, home, back, etc.
func (r *V1BoxActionService) PressButton(ctx context.Context, boxID string, body V1BoxActionPressButtonParams, opts ...option.RequestOption) (res *V1BoxActionPressButtonResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *V1BoxActionService) PressKey(ctx context.Context, boxID string, body V1BoxActionPressKeyParams, opts ...option.RequestOption) (res *V1BoxActionPressKeyResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *V1BoxActionService) RecordingStart(ctx context.Context, boxID string, body V1BoxActionRecordingStartParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/recording/start", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Stop recording the box screen
func (r *V1BoxActionService) RecordingStop(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxActionRecordingStopResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/recording/stop", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Stop the device's background screen rewind recording.
func (r *V1BoxActionService) RewindDisable(ctx context.Context, boxID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/recording/rewind", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Start the device's background screen rewind recording.
func (r *V1BoxActionService) RewindEnable(ctx context.Context, boxID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/screen-layout", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Rotate the screen orientation. Note that even after rotating the screen,
// applications or system layouts may not automatically adapt to the gravity sensor
// changes, so visual changes may not always occur.
func (r *V1BoxActionService) ScreenRotation(ctx context.Context, boxID string, body V1BoxActionScreenRotationParams, opts ...option.RequestOption) (res *V1BoxActionScreenRotationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/screen-rotation", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Take screenshot
func (r *V1BoxActionService) Screenshot(ctx context.Context, boxID string, body V1BoxActionScreenshotParams, opts ...option.RequestOption) (res *V1BoxActionScreenshotResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *V1BoxActionService) Scroll(ctx context.Context, boxID string, body V1BoxActionScrollParams, opts ...option.RequestOption) (res *V1BoxActionScrollResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/scroll", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the box action settings
func (r *V1BoxActionService) Settings(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxActionSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/settings", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Reset the box settings to default
func (r *V1BoxActionService) SettingsReset(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxActionSettingsResetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/settings", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Update the box action settings
func (r *V1BoxActionService) SettingsUpdate(ctx context.Context, boxID string, body V1BoxActionSettingsUpdateParams, opts ...option.RequestOption) (res *V1BoxActionSettingsUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/settings", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Performs a swipe in the specified direction
func (r *V1BoxActionService) Swipe(ctx context.Context, boxID string, body V1BoxActionSwipeParams, opts ...option.RequestOption) (res *V1BoxActionSwipeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/swipe", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Tap action for Android devices using ADB input tap command
func (r *V1BoxActionService) Tap(ctx context.Context, boxID string, body V1BoxActionTapParams, opts ...option.RequestOption) (res *V1BoxActionTapResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/tap", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Touch
func (r *V1BoxActionService) Touch(ctx context.Context, boxID string, body V1BoxActionTouchParams, opts ...option.RequestOption) (res *V1BoxActionTouchResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *V1BoxActionService) Type(ctx context.Context, boxID string, body V1BoxActionTypeParams, opts ...option.RequestOption) (res *V1BoxActionTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/type", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Result of an UI action execution with optional screenshots
type V1BoxActionAIResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionAIResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionAIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionAIResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionAIResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionAIResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionAIResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionAIResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionAIResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionAIResponseScreenshotAfter struct {
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
func (r V1BoxActionAIResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionAIResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionAIResponseScreenshotBefore struct {
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
func (r V1BoxActionAIResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionAIResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionAIResponseScreenshotTrace struct {
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
func (r V1BoxActionAIResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionAIResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with optional screenshots
type V1BoxActionClickResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionClickResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionClickResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionClickResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionClickResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionClickResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionClickResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionClickResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionClickResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionClickResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionClickResponseScreenshotAfter struct {
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
func (r V1BoxActionClickResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionClickResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionClickResponseScreenshotBefore struct {
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
func (r V1BoxActionClickResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionClickResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionClickResponseScreenshotTrace struct {
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
func (r V1BoxActionClickResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionClickResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with optional screenshots
type V1BoxActionDragResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionDragResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionDragResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionDragResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionDragResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionDragResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionDragResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionDragResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionDragResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionDragResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionDragResponseScreenshotAfter struct {
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
func (r V1BoxActionDragResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionDragResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionDragResponseScreenshotBefore struct {
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
func (r V1BoxActionDragResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionDragResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionDragResponseScreenshotTrace struct {
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
func (r V1BoxActionDragResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionDragResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
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

// Result of an UI action execution with optional screenshots
type V1BoxActionLongPressResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionLongPressResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionLongPressResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionLongPressResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionLongPressResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionLongPressResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionLongPressResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionLongPressResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionLongPressResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionLongPressResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionLongPressResponseScreenshotAfter struct {
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
func (r V1BoxActionLongPressResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionLongPressResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionLongPressResponseScreenshotBefore struct {
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
func (r V1BoxActionLongPressResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionLongPressResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionLongPressResponseScreenshotTrace struct {
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
func (r V1BoxActionLongPressResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionLongPressResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with optional screenshots
type V1BoxActionMoveResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionMoveResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionMoveResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionMoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionMoveResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionMoveResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionMoveResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionMoveResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionMoveResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionMoveResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionMoveResponseScreenshotAfter struct {
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
func (r V1BoxActionMoveResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionMoveResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionMoveResponseScreenshotBefore struct {
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
func (r V1BoxActionMoveResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionMoveResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionMoveResponseScreenshotTrace struct {
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
func (r V1BoxActionMoveResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionMoveResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with optional screenshots
type V1BoxActionPressButtonResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionPressButtonResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionPressButtonResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressButtonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionPressButtonResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionPressButtonResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionPressButtonResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionPressButtonResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionPressButtonResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressButtonResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionPressButtonResponseScreenshotAfter struct {
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
func (r V1BoxActionPressButtonResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressButtonResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionPressButtonResponseScreenshotBefore struct {
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
func (r V1BoxActionPressButtonResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressButtonResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionPressButtonResponseScreenshotTrace struct {
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
func (r V1BoxActionPressButtonResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressButtonResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with optional screenshots
type V1BoxActionPressKeyResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionPressKeyResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionPressKeyResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressKeyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionPressKeyResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionPressKeyResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionPressKeyResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionPressKeyResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionPressKeyResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressKeyResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionPressKeyResponseScreenshotAfter struct {
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
func (r V1BoxActionPressKeyResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressKeyResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionPressKeyResponseScreenshotBefore struct {
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
func (r V1BoxActionPressKeyResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressKeyResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionPressKeyResponseScreenshotTrace struct {
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
func (r V1BoxActionPressKeyResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressKeyResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
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

// Result of an UI action execution with optional screenshots
type V1BoxActionScreenRotationResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionScreenRotationResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionScreenRotationResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScreenRotationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionScreenRotationResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionScreenRotationResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionScreenRotationResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionScreenRotationResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionScreenRotationResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScreenRotationResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionScreenRotationResponseScreenshotAfter struct {
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
func (r V1BoxActionScreenRotationResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScreenRotationResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionScreenRotationResponseScreenshotBefore struct {
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
func (r V1BoxActionScreenRotationResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScreenRotationResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionScreenRotationResponseScreenshotTrace struct {
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
func (r V1BoxActionScreenRotationResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScreenRotationResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
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

// Result of an UI action execution with optional screenshots
type V1BoxActionScrollResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionScrollResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionScrollResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScrollResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionScrollResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionScrollResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionScrollResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionScrollResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionScrollResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScrollResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionScrollResponseScreenshotAfter struct {
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
func (r V1BoxActionScrollResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScrollResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionScrollResponseScreenshotBefore struct {
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
func (r V1BoxActionScrollResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScrollResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionScrollResponseScreenshotTrace struct {
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
func (r V1BoxActionScrollResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScrollResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
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

// Result of an UI action execution with optional screenshots
type V1BoxActionSwipeResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionSwipeResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionSwipeResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSwipeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionSwipeResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionSwipeResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionSwipeResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionSwipeResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionSwipeResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSwipeResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionSwipeResponseScreenshotAfter struct {
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
func (r V1BoxActionSwipeResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSwipeResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionSwipeResponseScreenshotBefore struct {
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
func (r V1BoxActionSwipeResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSwipeResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionSwipeResponseScreenshotTrace struct {
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
func (r V1BoxActionSwipeResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSwipeResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with optional screenshots
type V1BoxActionTapResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionTapResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionTapResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTapResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionTapResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionTapResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionTapResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionTapResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionTapResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTapResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionTapResponseScreenshotAfter struct {
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
func (r V1BoxActionTapResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTapResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionTapResponseScreenshotBefore struct {
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
func (r V1BoxActionTapResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTapResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionTapResponseScreenshotTrace struct {
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
func (r V1BoxActionTapResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTapResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with optional screenshots
type V1BoxActionTouchResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionTouchResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionTouchResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTouchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionTouchResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionTouchResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionTouchResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionTouchResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionTouchResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTouchResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionTouchResponseScreenshotAfter struct {
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
func (r V1BoxActionTouchResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTouchResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionTouchResponseScreenshotBefore struct {
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
func (r V1BoxActionTouchResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTouchResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionTouchResponseScreenshotTrace struct {
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
func (r V1BoxActionTouchResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTouchResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with optional screenshots
type V1BoxActionTypeResponse struct {
	// message
	Message string `json:"message,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionTypeResponseScreenshot `json:"screenshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionTypeResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTypeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionTypeResponseScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionTypeResponseScreenshotAfter `json:"after"`
	// Screenshot taken before action execution
	Before V1BoxActionTypeResponseScreenshotBefore `json:"before"`
	// Screenshot with action operation trace
	Trace V1BoxActionTypeResponseScreenshotTrace `json:"trace"`
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
func (r V1BoxActionTypeResponseScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTypeResponseScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionTypeResponseScreenshotAfter struct {
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
func (r V1BoxActionTypeResponseScreenshotAfter) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTypeResponseScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionTypeResponseScreenshotBefore struct {
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
func (r V1BoxActionTypeResponseScreenshotBefore) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTypeResponseScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionTypeResponseScreenshotTrace struct {
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
func (r V1BoxActionTypeResponseScreenshotTrace) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTypeResponseScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionAIParams struct {
	// Direct instruction of the UI action to perform (e.g., 'click the login button',
	// 'input username in the email field', 'scroll down', 'swipe left')
	Instruction string `json:"instruction,required"`
	// The background of the UI action to perform. The purpose of background is to let
	// the action executor to understand the context of why the instruction is given
	// including important previous actions and observations
	Background param.Opt[string] `json:"background,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Whether to stream progress events using Server-Sent Events (SSE). When true, the
	// API returns an event stream. When false or omitted, the API returns a normal
	// JSON response.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// Action common option
	Options V1BoxActionAIParamsOptions `json:"options,omitzero"`
	// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
	// default is base64. This field will be ignored when `options.screenshot` is
	// provided.
	//
	// Any of "base64", "storageKey".
	OutputFormat V1BoxActionAIParamsOutputFormat `json:"outputFormat,omitzero"`
	// AI action settings
	Settings V1BoxActionAIParamsSettings `json:"settings,omitzero"`
	paramObj
}

func (r V1BoxActionAIParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionAIParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionAIParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action common option
type V1BoxActionAIParamsOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionAIParamsOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionAIParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionAIParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionAIParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionAIParamsOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                             `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionAIParamsOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionAIParamsOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionAIParamsOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionAIParamsOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionAIParamsOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionAIParamsOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionAIParamsOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionAIParamsOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionAIParamsOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
}

// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
// default is base64. This field will be ignored when `options.screenshot` is
// provided.
type V1BoxActionAIParamsOutputFormat string

const (
	V1BoxActionAIParamsOutputFormatBase64     V1BoxActionAIParamsOutputFormat = "base64"
	V1BoxActionAIParamsOutputFormatStorageKey V1BoxActionAIParamsOutputFormat = "storageKey"
)

// AI action settings
type V1BoxActionAIParamsSettings struct {
	// System prompt that defines the AI's behavior and capabilities when executing UI
	// actions. This prompt instructs the AI on how to interpret the screen, understand
	// user instructions, and determine the appropriate UI actions to take. A
	// well-crafted system prompt can significantly improve the accuracy and
	// reliability of AI-driven UI automation. If not provided, uses the default
	// computer use instruction template that includes basic screen interaction
	// guidelines.
	SystemPrompt param.Opt[string] `json:"systemPrompt,omitzero"`
	// Whether disable actions
	DisableActions []string `json:"disableActions,omitzero"`
	paramObj
}

func (r V1BoxActionAIParamsSettings) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionAIParamsSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionAIParamsSettings) UnmarshalJSON(data []byte) error {
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

	paramObj
}

func (u V1BoxActionClickParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfClickAction, u.OfClickActionWithNaturalLanguage)
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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionClickParamsBodyClickActionOptions `json:"options,omitzero"`
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
		"outputFormat", "base64", "storageKey",
	)
}

// Action common option
type V1BoxActionClickParamsBodyClickActionOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionClickParamsBodyClickActionOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionClickParamsBodyClickActionOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionClickParamsBodyClickActionOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionClickParamsBodyClickActionOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionClickParamsBodyClickActionOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                               `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionClickParamsBodyClickActionOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionClickParamsBodyClickActionOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionClickParamsBodyClickActionOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionClickParamsBodyClickActionOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionClickParamsBodyClickActionOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionClickParamsBodyClickActionOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionClickParamsBodyClickActionOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionClickParamsBodyClickActionOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionClickParamsBodyClickActionOptionsScreenshotActionScreenshotOption](
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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptions `json:"options,omitzero"`
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
		"outputFormat", "base64", "storageKey",
	)
}

// Action common option
type V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                                                  `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionClickParamsBodyClickActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionDragParamsBodyDragSimpleOptions `json:"options,omitzero"`
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

// Action common option
type V1BoxActionDragParamsBodyDragSimpleOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionDragParamsBodyDragSimpleOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionDragParamsBodyDragSimpleOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsBodyDragSimpleOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsBodyDragSimpleOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionDragParamsBodyDragSimpleOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                             `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionDragParamsBodyDragSimpleOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionDragParamsBodyDragSimpleOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionDragParamsBodyDragSimpleOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionDragParamsBodyDragSimpleOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionDragParamsBodyDragSimpleOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionDragParamsBodyDragSimpleOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsBodyDragSimpleOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsBodyDragSimpleOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionDragParamsBodyDragSimpleOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionDragParamsBodyDragAdvancedOptions `json:"options,omitzero"`
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

// Action common option
type V1BoxActionDragParamsBodyDragAdvancedOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionDragParamsBodyDragAdvancedOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionDragParamsBodyDragAdvancedOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsBodyDragAdvancedOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsBodyDragAdvancedOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionDragParamsBodyDragAdvancedOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                               `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionDragParamsBodyDragAdvancedOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionDragParamsBodyDragAdvancedOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionDragParamsBodyDragAdvancedOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionDragParamsBodyDragAdvancedOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionDragParamsBodyDragAdvancedOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionDragParamsBodyDragAdvancedOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsBodyDragAdvancedOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsBodyDragAdvancedOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionDragParamsBodyDragAdvancedOptionsScreenshotActionScreenshotOption](
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

	paramObj
}

func (u V1BoxActionLongPressParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLongPressAction, u.OfLongPressActionWithNaturalLanguage)
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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionLongPressParamsBodyLongPressActionOptions `json:"options,omitzero"`
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
		"outputFormat", "base64", "storageKey",
	)
}

// Action common option
type V1BoxActionLongPressParamsBodyLongPressActionOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionLongPressParamsBodyLongPressActionOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionLongPressParamsBodyLongPressActionOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionLongPressParamsBodyLongPressActionOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionLongPressParamsBodyLongPressActionOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionLongPressParamsBodyLongPressActionOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                                       `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionLongPressParamsBodyLongPressActionOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionLongPressParamsBodyLongPressActionOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionLongPressParamsBodyLongPressActionOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionLongPressParamsBodyLongPressActionOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionLongPressParamsBodyLongPressActionOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionLongPressParamsBodyLongPressActionOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionLongPressParamsBodyLongPressActionOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionLongPressParamsBodyLongPressActionOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionLongPressParamsBodyLongPressActionOptionsScreenshotActionScreenshotOption](
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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptions `json:"options,omitzero"`
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
		"outputFormat", "base64", "storageKey",
	)
}

// Action common option
type V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                                                          `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionLongPressParamsBodyLongPressActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
}

type V1BoxActionMoveParams struct {
	// X coordinate to move to
	X float64 `json:"x,required"`
	// Y coordinate to move to
	Y float64 `json:"y,required"`
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionMoveParamsOptions `json:"options,omitzero"`
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

// Action common option
type V1BoxActionMoveParamsOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionMoveParamsOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionMoveParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionMoveParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionMoveParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionMoveParamsOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                               `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionMoveParamsOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionMoveParamsOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionMoveParamsOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionMoveParamsOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionMoveParamsOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionMoveParamsOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionMoveParamsOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionMoveParamsOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionMoveParamsOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
}

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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionPressButtonParamsOptions `json:"options,omitzero"`
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

// Action common option
type V1BoxActionPressButtonParamsOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionPressButtonParamsOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionPressButtonParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionPressButtonParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionPressButtonParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionPressButtonParamsOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                      `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionPressButtonParamsOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionPressButtonParamsOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionPressButtonParamsOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionPressButtonParamsOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionPressButtonParamsOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionPressButtonParamsOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionPressButtonParamsOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionPressButtonParamsOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionPressButtonParamsOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
}

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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionPressKeyParamsOptions `json:"options,omitzero"`
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

// Action common option
type V1BoxActionPressKeyParamsOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionPressKeyParamsOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionPressKeyParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionPressKeyParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionPressKeyParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionPressKeyParamsOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                   `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionPressKeyParamsOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionPressKeyParamsOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionPressKeyParamsOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionPressKeyParamsOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionPressKeyParamsOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionPressKeyParamsOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionPressKeyParamsOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionPressKeyParamsOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionPressKeyParamsOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
}

// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
// default is base64. This field will be ignored when `options.screenshot` is
// provided.
type V1BoxActionPressKeyParamsOutputFormat string

const (
	V1BoxActionPressKeyParamsOutputFormatBase64     V1BoxActionPressKeyParamsOutputFormat = "base64"
	V1BoxActionPressKeyParamsOutputFormatStorageKey V1BoxActionPressKeyParamsOutputFormat = "storageKey"
)

type V1BoxActionRecordingStartParams struct {
	// Duration of the recording. Default is 30m, max is 30m. The recording will
	// automatically stop when the duration time is reached.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Maximum allowed: 30m
	Duration param.Opt[string] `json:"duration,omitzero"`
	paramObj
}

func (r V1BoxActionRecordingStartParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionRecordingStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionRecordingStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionScreenRotationParamsOptions `json:"options,omitzero"`
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

// Action common option
type V1BoxActionScreenRotationParamsOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionScreenRotationParamsOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionScreenRotationParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScreenRotationParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScreenRotationParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionScreenRotationParamsOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                         `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionScreenRotationParamsOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionScreenRotationParamsOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionScreenRotationParamsOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionScreenRotationParamsOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionScreenRotationParamsOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionScreenRotationParamsOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScreenRotationParamsOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScreenRotationParamsOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionScreenRotationParamsOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
}

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
	//   - If not provided, uses the scale value from UI action settings; otherwise uses
	//     the passed value.
	Scale param.Opt[float64] `json:"scale,omitzero"`
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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionScrollParamsBodyScrollAdvancedOptions `json:"options,omitzero"`
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
		"outputFormat", "base64", "storageKey",
	)
}

// Action common option
type V1BoxActionScrollParamsBodyScrollAdvancedOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionScrollParamsBodyScrollAdvancedOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionScrollParamsBodyScrollAdvancedOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScrollParamsBodyScrollAdvancedOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScrollParamsBodyScrollAdvancedOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionScrollParamsBodyScrollAdvancedOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                                   `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionScrollParamsBodyScrollAdvancedOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionScrollParamsBodyScrollAdvancedOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionScrollParamsBodyScrollAdvancedOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionScrollParamsBodyScrollAdvancedOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionScrollParamsBodyScrollAdvancedOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionScrollParamsBodyScrollAdvancedOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScrollParamsBodyScrollAdvancedOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScrollParamsBodyScrollAdvancedOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionScrollParamsBodyScrollAdvancedOptionsScreenshotActionScreenshotOption](
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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Distance of the scroll. Can be either a number (in pixels) or a predefined enum
	// value (tiny, short, medium, long). If not provided, the scroll will be performed
	// from the center of the screen to the screen edge
	Distance V1BoxActionScrollParamsBodyScrollSimpleDistanceUnion `json:"distance,omitzero"`
	// Action common option
	Options V1BoxActionScrollParamsBodyScrollSimpleOptions `json:"options,omitzero"`
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

// Action common option
type V1BoxActionScrollParamsBodyScrollSimpleOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionScrollParamsBodyScrollSimpleOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionScrollParamsBodyScrollSimpleOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScrollParamsBodyScrollSimpleOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScrollParamsBodyScrollSimpleOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionScrollParamsBodyScrollSimpleOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                                 `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionScrollParamsBodyScrollSimpleOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionScrollParamsBodyScrollSimpleOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionScrollParamsBodyScrollSimpleOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionScrollParamsBodyScrollSimpleOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionScrollParamsBodyScrollSimpleOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionScrollParamsBodyScrollSimpleOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScrollParamsBodyScrollSimpleOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScrollParamsBodyScrollSimpleOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionScrollParamsBodyScrollSimpleOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
}

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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionSwipeParamsBodySwipeSimpleOptions `json:"options,omitzero"`
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

// Action common option
type V1BoxActionSwipeParamsBodySwipeSimpleOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionSwipeParamsBodySwipeSimpleOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionSwipeParamsBodySwipeSimpleOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSwipeParamsBodySwipeSimpleOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSwipeParamsBodySwipeSimpleOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionSwipeParamsBodySwipeSimpleOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                               `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionSwipeParamsBodySwipeSimpleOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionSwipeParamsBodySwipeSimpleOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionSwipeParamsBodySwipeSimpleOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionSwipeParamsBodySwipeSimpleOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionSwipeParamsBodySwipeSimpleOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionSwipeParamsBodySwipeSimpleOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSwipeParamsBodySwipeSimpleOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSwipeParamsBodySwipeSimpleOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionSwipeParamsBodySwipeSimpleOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
}

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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionSwipeParamsBodySwipeAdvancedOptions `json:"options,omitzero"`
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

// Action common option
type V1BoxActionSwipeParamsBodySwipeAdvancedOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionSwipeParamsBodySwipeAdvancedOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionSwipeParamsBodySwipeAdvancedOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSwipeParamsBodySwipeAdvancedOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSwipeParamsBodySwipeAdvancedOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionSwipeParamsBodySwipeAdvancedOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                                 `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionSwipeParamsBodySwipeAdvancedOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionSwipeParamsBodySwipeAdvancedOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionSwipeParamsBodySwipeAdvancedOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionSwipeParamsBodySwipeAdvancedOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionSwipeParamsBodySwipeAdvancedOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionSwipeParamsBodySwipeAdvancedOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSwipeParamsBodySwipeAdvancedOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSwipeParamsBodySwipeAdvancedOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionSwipeParamsBodySwipeAdvancedOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
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

	paramObj
}

func (u V1BoxActionTapParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTapAction, u.OfTapActionWithNaturalLanguage)
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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionTapParamsBodyTapActionOptions `json:"options,omitzero"`
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
		"outputFormat", "base64", "storageKey",
	)
}

// Action common option
type V1BoxActionTapParamsBodyTapActionOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionTapParamsBodyTapActionOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionTapParamsBodyTapActionOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTapParamsBodyTapActionOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTapParamsBodyTapActionOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionTapParamsBodyTapActionOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                           `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionTapParamsBodyTapActionOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionTapParamsBodyTapActionOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionTapParamsBodyTapActionOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionTapParamsBodyTapActionOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionTapParamsBodyTapActionOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionTapParamsBodyTapActionOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTapParamsBodyTapActionOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTapParamsBodyTapActionOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionTapParamsBodyTapActionOptionsScreenshotActionScreenshotOption](
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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptions `json:"options,omitzero"`
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
		"outputFormat", "base64", "storageKey",
	)
}

// Action common option
type V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                                              `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionTapParamsBodyTapActionWithNaturalLanguageOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
}

type V1BoxActionTouchParams struct {
	// Array of touch points and their actions
	Points []V1BoxActionTouchParamsPoint `json:"points,omitzero,required"`
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionTouchParamsOptions `json:"options,omitzero"`
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

// Action common option
type V1BoxActionTouchParamsOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionTouchParamsOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionTouchParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTouchParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTouchParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionTouchParamsOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                                `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionTouchParamsOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionTouchParamsOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionTouchParamsOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionTouchParamsOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionTouchParamsOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionTouchParamsOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTouchParamsOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTouchParamsOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionTouchParamsOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
}

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
	// ⚠️ DEPRECATED: Use `options.screenshot.range` instead. This field will be
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
	// Action common option
	Options V1BoxActionTypeParamsOptions `json:"options,omitzero"`
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

// Action common option
type V1BoxActionTypeParamsOptions struct {
	// Screenshot options. Can be a boolean to enable/disable screenshots, or an object
	// to configure screenshot options.
	Screenshot V1BoxActionTypeParamsOptionsScreenshotUnion `json:"screenshot,omitzero"`
	paramObj
}

func (r V1BoxActionTypeParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTypeParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTypeParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxActionTypeParamsOptionsScreenshotUnion struct {
	OfBool                   param.Opt[bool]                                               `json:",omitzero,inline"`
	OfActionScreenshotOption *V1BoxActionTypeParamsOptionsScreenshotActionScreenshotOption `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionTypeParamsOptionsScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfActionScreenshotOption)
}
func (u *V1BoxActionTypeParamsOptionsScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionTypeParamsOptionsScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfActionScreenshotOption) {
		return u.OfActionScreenshotOption
	}
	return nil
}

// Action screenshot option
type V1BoxActionTypeParamsOptionsScreenshotActionScreenshotOption struct {
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
	OutputFormat string `json:"outputFormat,omitzero"`
	// Specify which screenshots to capture.
	//
	// Available options:
	//
	// - before: Screenshot before the action
	// - after: Screenshot after the action
	// - trace: Screenshot with operation trace
	//
	// Default captures all three types. Can specify one or multiple in an array.
	//
	// Any of "before", "after", "trace".
	Range []string `json:"range,omitzero"`
	paramObj
}

func (r V1BoxActionTypeParamsOptionsScreenshotActionScreenshotOption) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTypeParamsOptionsScreenshotActionScreenshotOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTypeParamsOptionsScreenshotActionScreenshotOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionTypeParamsOptionsScreenshotActionScreenshotOption](
		"outputFormat", "base64", "storageKey",
	)
}

// ⚠️ DEPRECATED: Use `options.screenshot.outputFormat` instead. Type of the URI.
// default is base64. This field will be ignored when `options.screenshot` is
// provided.
type V1BoxActionTypeParamsOutputFormat string

const (
	V1BoxActionTypeParamsOutputFormatBase64     V1BoxActionTypeParamsOutputFormat = "base64"
	V1BoxActionTypeParamsOutputFormatStorageKey V1BoxActionTypeParamsOutputFormat = "storageKey"
)
