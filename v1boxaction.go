// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk

import (
	"context"
	"encoding/json"
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
func (r *V1BoxActionService) AI(ctx context.Context, boxID string, body V1BoxActionAIParams, opts ...option.RequestOption) (res *V1BoxActionAIResponseUnion, err error) {
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
func (r *V1BoxActionService) Click(ctx context.Context, boxID string, body V1BoxActionClickParams, opts ...option.RequestOption) (res *V1BoxActionClickResponseUnion, err error) {
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
func (r *V1BoxActionService) Drag(ctx context.Context, boxID string, body V1BoxActionDragParams, opts ...option.RequestOption) (res *V1BoxActionDragResponseUnion, err error) {
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
func (r *V1BoxActionService) LongPress(ctx context.Context, boxID string, body V1BoxActionLongPressParams, opts ...option.RequestOption) (res *V1BoxActionLongPressResponseUnion, err error) {
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
func (r *V1BoxActionService) Move(ctx context.Context, boxID string, body V1BoxActionMoveParams, opts ...option.RequestOption) (res *V1BoxActionMoveResponseUnion, err error) {
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
func (r *V1BoxActionService) PressButton(ctx context.Context, boxID string, body V1BoxActionPressButtonParams, opts ...option.RequestOption) (res *V1BoxActionPressButtonResponseUnion, err error) {
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
func (r *V1BoxActionService) PressKey(ctx context.Context, boxID string, body V1BoxActionPressKeyParams, opts ...option.RequestOption) (res *V1BoxActionPressKeyResponseUnion, err error) {
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
func (r *V1BoxActionService) ScreenRotation(ctx context.Context, boxID string, body V1BoxActionScreenRotationParams, opts ...option.RequestOption) (res *V1BoxActionScreenRotationResponseUnion, err error) {
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
func (r *V1BoxActionService) Scroll(ctx context.Context, boxID string, body V1BoxActionScrollParams, opts ...option.RequestOption) (res *V1BoxActionScrollResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/scroll", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the box action setting
func (r *V1BoxActionService) Setting(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxActionSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/setting", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Reset the box setting
func (r *V1BoxActionService) SettingReset(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxActionSettingResetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/setting/reset", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Setting the box action setting
func (r *V1BoxActionService) SettingUpdate(ctx context.Context, boxID string, body V1BoxActionSettingUpdateParams, opts ...option.RequestOption) (res *V1BoxActionSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/setting", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Performs a swipe in the specified direction
func (r *V1BoxActionService) Swipe(ctx context.Context, boxID string, body V1BoxActionSwipeParams, opts ...option.RequestOption) (res *V1BoxActionSwipeResponseUnion, err error) {
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
func (r *V1BoxActionService) Tap(ctx context.Context, boxID string, body V1BoxActionTapParams, opts ...option.RequestOption) (res *V1BoxActionTapResponseUnion, err error) {
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
func (r *V1BoxActionService) Touch(ctx context.Context, boxID string, body V1BoxActionTouchParams, opts ...option.RequestOption) (res *V1BoxActionTouchResponseUnion, err error) {
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
func (r *V1BoxActionService) Type(ctx context.Context, boxID string, body V1BoxActionTypeParams, opts ...option.RequestOption) (res *V1BoxActionTypeResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/actions/type", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// V1BoxActionAIResponseUnion contains all possible properties and values from
// [V1BoxActionAIResponseAIActionScreenshotResult],
// [V1BoxActionAIResponseAIActionResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionAIResponseUnion struct {
	// This field is a union of
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponse],
	// [V1BoxActionAIResponseAIActionResultAIResponse]
	AIResponse V1BoxActionAIResponseUnionAIResponse `json:"aiResponse"`
	Output     string                               `json:"output"`
	// This field is from variant [V1BoxActionAIResponseAIActionScreenshotResult].
	Screenshot V1BoxActionAIResponseAIActionScreenshotResultScreenshot `json:"screenshot"`
	JSON       struct {
		AIResponse respjson.Field
		Output     respjson.Field
		Screenshot respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionAIResponseUnion) AsAIActionScreenshotResult() (v V1BoxActionAIResponseAIActionScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseUnion) AsAIActionResult() (v V1BoxActionAIResponseAIActionResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionAIResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseUnionAIResponse is an implicit subunion of
// [V1BoxActionAIResponseUnion]. V1BoxActionAIResponseUnionAIResponse provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxActionAIResponseUnion].
type V1BoxActionAIResponseUnionAIResponse struct {
	// This field is a union of
	// [[]V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion],
	// [[]V1BoxActionAIResponseAIActionResultAIResponseActionUnion]
	Actions   V1BoxActionAIResponseUnionAIResponseActions `json:"actions"`
	Messages  []string                                    `json:"messages"`
	Model     string                                      `json:"model"`
	Reasoning string                                      `json:"reasoning"`
	JSON      struct {
		Actions   respjson.Field
		Messages  respjson.Field
		Model     respjson.Field
		Reasoning respjson.Field
		raw       string
	} `json:"-"`
}

func (r *V1BoxActionAIResponseUnionAIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseUnionAIResponseActions is an implicit subunion of
// [V1BoxActionAIResponseUnion]. V1BoxActionAIResponseUnionAIResponseActions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxActionAIResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActions
// OfV1BoxActionAIResponseAIActionResultAIResponseActions]
type V1BoxActionAIResponseUnionAIResponseActions struct {
	// This field will be present if the value is a
	// [[]V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion] instead
	// of an object.
	OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActions []V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]V1BoxActionAIResponseAIActionResultAIResponseActionUnion] instead of an
	// object.
	OfV1BoxActionAIResponseAIActionResultAIResponseActions []V1BoxActionAIResponseAIActionResultAIResponseActionUnion `json:",inline"`
	JSON                                                   struct {
		OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActions respjson.Field
		OfV1BoxActionAIResponseAIActionResultAIResponseActions           respjson.Field
		raw                                                              string
	} `json:"-"`
}

func (r *V1BoxActionAIResponseUnionAIResponseActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of AI action execution with screenshot
type V1BoxActionAIResponseAIActionScreenshotResult struct {
	// Response of AI action execution
	AIResponse V1BoxActionAIResponseAIActionScreenshotResultAIResponse `json:"aiResponse,required"`
	// output
	Output string `json:"output,required"`
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionAIResponseAIActionScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIResponse  respjson.Field
		Output      respjson.Field
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionAIResponseAIActionScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response of AI action execution
type V1BoxActionAIResponseAIActionScreenshotResultAIResponse struct {
	// Actions to be executed by the AI with type identifier
	Actions []V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion `json:"actions,required"`
	// messages returned by the model
	Messages []string `json:"messages,required"`
	// The name of the model that processed this request
	Model string `json:"model,required"`
	// reasoning
	Reasoning string `json:"reasoning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions     respjson.Field
		Messages    respjson.Field
		Model       respjson.Field
		Reasoning   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion contains all
// possible properties and values from
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedClickAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressKeyAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressButtonAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedLongPressAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTypeAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedMoveAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenRotationAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedWaitAction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedClickAction].
	Button string `json:"button"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedClickAction].
	Double             bool   `json:"double"`
	IncludeScreenshot  bool   `json:"includeScreenshot"`
	OutputFormat       string `json:"outputFormat"`
	PresignedExpiresIn string `json:"presignedExpiresIn"`
	ScreenshotDelay    string `json:"screenshotDelay"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchAction].
	Points []V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPoint `json:"points"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedAction].
	Path     []V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedActionPath `json:"path"`
	Duration string                                                                                     `json:"duration"`
	// This field is a union of
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndUnion],
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndUnion]
	End V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionEnd `json:"end"`
	// This field is a union of
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartUnion],
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartUnion]
	Start V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionStart `json:"start"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction].
	ScrollX float64 `json:"scrollX"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction].
	ScrollY   float64 `json:"scrollY"`
	Direction string  `json:"direction"`
	// This field is a union of
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceUnion],
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceUnion]
	Distance V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionDistance `json:"distance"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleAction].
	Location string `json:"location"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressKeyAction].
	Keys []string `json:"keys"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressKeyAction].
	Combination bool `json:"combination"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressButtonAction].
	Buttons []string `json:"buttons"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTypeAction].
	Text string `json:"text"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTypeAction].
	Mode string `json:"mode"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTypeAction].
	PressEnter bool `json:"pressEnter"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenRotationAction].
	Orientation string `json:"orientation"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotAction].
	Clip V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotActionClip `json:"clip"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotAction].
	Scale float64 `json:"scale"`
	JSON  struct {
		X                  respjson.Field
		Y                  respjson.Field
		Button             respjson.Field
		Double             respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		Points             respjson.Field
		Path               respjson.Field
		Duration           respjson.Field
		End                respjson.Field
		Start              respjson.Field
		ScrollX            respjson.Field
		ScrollY            respjson.Field
		Direction          respjson.Field
		Distance           respjson.Field
		Location           respjson.Field
		Keys               respjson.Field
		Combination        respjson.Field
		Buttons            respjson.Field
		Text               respjson.Field
		Mode               respjson.Field
		PressEnter         respjson.Field
		Orientation        respjson.Field
		Clip               respjson.Field
		Scale              respjson.Field
		raw                string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedClickAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedClickAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedTouchAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedDragAdvancedAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedDragSimpleAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedScrollAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedScrollSimpleAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedSwipeSimpleAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedSwipeAdvancedAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedPressKeyAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressKeyAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedPressButtonAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressButtonAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedLongPressAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedLongPressAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedTypeAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTypeAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedMoveAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedMoveAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedScreenRotationAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenRotationAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedScreenshotAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) AsTypedWaitAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedWaitAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionEnd is an
// implicit subunion of
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion].
// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionEnd provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionEnd struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string  `json:",inline"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	JSON     struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionStart is an
// implicit subunion of
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion].
// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionStart provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionStart struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string  `json:",inline"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	JSON     struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionDistance is an
// implicit subunion of
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion].
// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionDistance
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat
// OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString]
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionDistance struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString string `json:",inline"`
	JSON                                                                                                struct {
		OfFloat                                                                                             respjson.Field
		OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString respjson.Field
		raw                                                                                                 string
	} `json:"-"`
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionDistance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed click action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedClickAction struct {
	// X coordinate of the click
	X float64 `json:"x,required"`
	// Y coordinate of the click
	Y float64 `json:"y,required"`
	// Mouse button to click
	//
	// Any of "left", "right", "middle".
	Button string `json:"button"`
	// Whether to perform a double click
	Double bool `json:"double"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X                  respjson.Field
		Y                  respjson.Field
		Button             respjson.Field
		Double             respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedClickAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedClickAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed touch action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchAction struct {
	// Array of touch points and their actions
	Points []V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPoint `json:"points,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Points             respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Touch point configuration with start position and actions
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPoint struct {
	// Initial touch point position
	Start V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointStart `json:"start,required"`
	// Sequence of actions to perform after initial touch
	Actions []V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionUnion `json:"actions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Start       respjson.Field
		Actions     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Initial touch point position
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointStart struct {
	// Starting X coordinate
	X float64 `json:"x,required"`
	// Starting Y coordinate
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
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointStart) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionUnion
// contains all possible properties and values from
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionUnion struct {
	Duration string `json:"duration"`
	Type     string `json:"type"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction].
	X float64 `json:"x"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction].
	Y    float64 `json:"y"`
	JSON struct {
		Duration respjson.Field
		Type     respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionUnion) AsTouchPointMoveAction() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionUnion) AsV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Touch point movement action configuration
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction struct {
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Type        respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto struct {
	// Duration to wait (e.g. "500ms")
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration,required"`
	// Type of the action
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed drag advanced action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedAction struct {
	// Path of the drag action as a series of coordinates
	Path []V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedActionPath `json:"path,required"`
	// Time interval between points (e.g. "50ms")
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 50ms
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path               respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single point in a drag path
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedActionPath struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
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
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedActionPath) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedActionPath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed drag simple action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleAction struct {
	// End point of the drag path (coordinates or natural language)
	End V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndUnion `json:"end,required"`
	// Start point of the drag path (coordinates or natural language)
	Start V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartUnion `json:"start,required"`
	// Duration to complete the movement from start to end coordinates
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End                respjson.Field
		Start              respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndUnion
// contains all possible properties and values from
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndDragPathPoint],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndDragPathPoint].
	X float64 `json:"x"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndDragPathPoint].
	Y    float64 `json:"y"`
	JSON struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndUnion) AsDragPathPoint() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndDragPathPoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single point in a drag path
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndDragPathPoint struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
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
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndDragPathPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEndDragPathPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartUnion
// contains all possible properties and values from
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartDragPathPoint],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartDragPathPoint].
	X float64 `json:"x"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartDragPathPoint].
	Y    float64 `json:"y"`
	JSON struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartUnion) AsDragPathPoint() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartDragPathPoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single point in a drag path
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartDragPathPoint struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
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
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartDragPathPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStartDragPathPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed scroll action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction struct {
	// Horizontal scroll amount. Positive values scroll content rightward (reveals
	// content on the left), negative values scroll content leftward (reveals content
	// on the right).
	ScrollX float64 `json:"scrollX,required"`
	// Vertical scroll amount. Positive values scroll content downward (reveals content
	// above), negative values scroll content upward (reveals content below).
	ScrollY float64 `json:"scrollY,required"`
	// X coordinate of the scroll position
	X float64 `json:"x,required"`
	// Y coordinate of the scroll position
	Y float64 `json:"y,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScrollX            respjson.Field
		ScrollY            respjson.Field
		X                  respjson.Field
		Y                  respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed scroll simple action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleAction struct {
	// Direction to scroll. The scroll will be performed from the center of the screen
	// towards this direction. 'up' scrolls content upward (reveals content below),
	// 'down' scrolls content downward (reveals content above), 'left' scrolls content
	// leftward (reveals content on the right), 'right' scrolls content rightward
	// (reveals content on the left).
	//
	// Any of "up", "down", "left", "right".
	Direction string `json:"direction,required"`
	// Distance of the scroll. Can be either a number (in pixels) or a predefined enum
	// value (tiny, short, medium, long). If not provided, the scroll will be performed
	// from the center of the screen to the screen edge
	Distance V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceUnion `json:"distance"`
	// Duration of the scroll
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direction          respjson.Field
		Distance           respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceUnion
// contains all possible properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat
// OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceString]
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceString string `json:",inline"`
	JSON                                                                                                 struct {
		OfFloat                                                                                              respjson.Field
		OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceString respjson.Field
		raw                                                                                                  string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceUnion) AsV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceString string

const (
	V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceStringTiny   V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceString = "tiny"
	V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceStringShort  V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceString = "short"
	V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceStringMedium V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceString = "medium"
	V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceStringLong   V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollSimpleActionDistanceString = "long"
)

// Typed swipe simple action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleAction struct {
	// Direction to swipe. The gesture will be performed from the center of the screen
	// towards this direction.
	//
	// Any of "up", "down", "left", "right", "upLeft", "upRight", "downLeft",
	// "downRight".
	Direction string `json:"direction,required"`
	// Distance of the swipe. Can be either a number (in pixels) or a predefined enum
	// value (tiny, short, medium, long). If not provided, the swipe will be performed
	// from the center of the screen to the screen edge
	Distance V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceUnion `json:"distance"`
	// Duration of the swipe
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Natural language description of the location where the swipe should originate.
	// If not provided, the swipe will be performed from the center of the screen.
	Location string `json:"location"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direction          respjson.Field
		Distance           respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		Location           respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceUnion
// contains all possible properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat
// OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString]
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString string `json:",inline"`
	JSON                                                                                                struct {
		OfFloat                                                                                             respjson.Field
		OfV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString respjson.Field
		raw                                                                                                 string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceUnion) AsV1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString string

const (
	V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceStringTiny   V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString = "tiny"
	V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceStringShort  V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString = "short"
	V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceStringMedium V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString = "medium"
	V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceStringLong   V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleActionDistanceString = "long"
)

// Typed swipe advanced action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedAction struct {
	// End point of the swipe path (coordinates or natural language)
	End V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndUnion `json:"end,required"`
	// Start point of the swipe path (coordinates or natural language)
	Start V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartUnion `json:"start,required"`
	// Duration of the swipe
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End                respjson.Field
		Start              respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndUnion
// contains all possible properties and values from
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath].
	X float64 `json:"x"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath].
	Y    float64 `json:"y"`
	JSON struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndUnion) AsSwipePath() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swipe path
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath struct {
	// Start/end x coordinate of the swipe path
	X float64 `json:"x,required"`
	// Start/end y coordinate of the swipe path
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
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartUnion
// contains all possible properties and values from
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath].
	X float64 `json:"x"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath].
	Y    float64 `json:"y"`
	JSON struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartUnion) AsSwipePath() (v V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swipe path
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath struct {
	// Start/end x coordinate of the swipe path
	X float64 `json:"x,required"`
	// Start/end y coordinate of the swipe path
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
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed press key action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressKeyAction struct {
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
	Keys []string `json:"keys,required"`
	// Whether to press keys as combination (simultaneously) or sequentially. When
	// true, all keys are pressed together as a shortcut (e.g., Ctrl+C). When false,
	// keys are pressed one by one in sequence.
	Combination bool `json:"combination"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Keys               respjson.Field
		Combination        respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressKeyAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressKeyAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed press button action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressButtonAction struct {
	// Button to press
	//
	// Any of "power", "volumeUp", "volumeDown", "volumeMute", "home", "back", "menu",
	// "appSwitch".
	Buttons []string `json:"buttons,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buttons            respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressButtonAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressButtonAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed long press action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedLongPressAction struct {
	// X coordinate of the long press
	X float64 `json:"x,required"`
	// Y coordinate of the long press
	Y float64 `json:"y,required"`
	// Duration to hold the press (e.g. '1s', '500ms')
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 1s
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X                  respjson.Field
		Y                  respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedLongPressAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedLongPressAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed type action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTypeAction struct {
	// Text to type
	Text string `json:"text,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Text input mode: 'append' to add text to existing content, 'replace' to replace
	// all existing text
	//
	// Any of "append", "replace".
	Mode string `json:"mode"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
	// Whether to press Enter after typing the text
	PressEnter bool `json:"pressEnter"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text               respjson.Field
		IncludeScreenshot  respjson.Field
		Mode               respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		PressEnter         respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTypeAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTypeAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed move action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedMoveAction struct {
	// X coordinate to move to
	X float64 `json:"x,required"`
	// Y coordinate to move to
	Y float64 `json:"y,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X                  respjson.Field
		Y                  respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedMoveAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedMoveAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed screen rotation action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenRotationAction struct {
	// Target screen orientation
	//
	// Any of "portrait", "landscapeLeft", "portraitUpsideDown", "landscapeRight".
	Orientation string `json:"orientation,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Orientation        respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenRotationAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenRotationAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed screenshot action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotAction struct {
	// Clipping region for screenshot capture
	Clip V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotActionClip `json:"clip"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
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
	Scale float64 `json:"scale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clip         respjson.Field
		OutputFormat respjson.Field
		Scale        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Clipping region for screenshot capture
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotActionClip struct {
	// Height of the clip
	Height float64 `json:"height,required"`
	// Width of the clip
	Width float64 `json:"width,required"`
	// X coordinate of the clip
	X float64 `json:"x,required"`
	// Y coordinate of the clip
	Y float64 `json:"y,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotActionClip) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotActionClip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed wait action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedWaitAction struct {
	// Duration of the wait (e.g. '3s')
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 3s
	Duration string `json:"duration,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedWaitAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedWaitAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionAIResponseAIActionScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionAIResponseAIActionScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionAIResponseAIActionScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionAIResponseAIActionScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionAIResponseAIActionScreenshotResultScreenshot) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionAIResponseAIActionScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionAIResponseAIActionScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionAIResponseAIActionScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionAIResponseAIActionScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionAIResponseAIActionScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionAIResponseAIActionScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionAIResponseAIActionScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of AI action execution
type V1BoxActionAIResponseAIActionResult struct {
	// Response of AI action execution
	AIResponse V1BoxActionAIResponseAIActionResultAIResponse `json:"aiResponse,required"`
	// output
	Output string `json:"output,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIResponse  respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionAIResponseAIActionResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response of AI action execution
type V1BoxActionAIResponseAIActionResultAIResponse struct {
	// Actions to be executed by the AI with type identifier
	Actions []V1BoxActionAIResponseAIActionResultAIResponseActionUnion `json:"actions,required"`
	// messages returned by the model
	Messages []string `json:"messages,required"`
	// The name of the model that processed this request
	Model string `json:"model,required"`
	// reasoning
	Reasoning string `json:"reasoning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions     respjson.Field
		Messages    respjson.Field
		Model       respjson.Field
		Reasoning   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionAIResponseAIActionResultAIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionResultAIResponseActionUnion contains all possible
// properties and values from
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedClickAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressKeyAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressButtonAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedLongPressAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedTypeAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedMoveAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenRotationAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedWaitAction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionAIResponseAIActionResultAIResponseActionUnion struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedClickAction].
	Button string `json:"button"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedClickAction].
	Double             bool   `json:"double"`
	IncludeScreenshot  bool   `json:"includeScreenshot"`
	OutputFormat       string `json:"outputFormat"`
	PresignedExpiresIn string `json:"presignedExpiresIn"`
	ScreenshotDelay    string `json:"screenshotDelay"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchAction].
	Points []V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPoint `json:"points"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedAction].
	Path     []V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedActionPath `json:"path"`
	Duration string                                                                           `json:"duration"`
	// This field is a union of
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndUnion],
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndUnion]
	End V1BoxActionAIResponseAIActionResultAIResponseActionUnionEnd `json:"end"`
	// This field is a union of
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartUnion],
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartUnion]
	Start V1BoxActionAIResponseAIActionResultAIResponseActionUnionStart `json:"start"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction].
	ScrollX float64 `json:"scrollX"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction].
	ScrollY   float64 `json:"scrollY"`
	Direction string  `json:"direction"`
	// This field is a union of
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceUnion],
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceUnion]
	Distance V1BoxActionAIResponseAIActionResultAIResponseActionUnionDistance `json:"distance"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleAction].
	Location string `json:"location"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressKeyAction].
	Keys []string `json:"keys"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressKeyAction].
	Combination bool `json:"combination"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressButtonAction].
	Buttons []string `json:"buttons"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedTypeAction].
	Text string `json:"text"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedTypeAction].
	Mode string `json:"mode"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedTypeAction].
	PressEnter bool `json:"pressEnter"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenRotationAction].
	Orientation string `json:"orientation"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotAction].
	Clip V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotActionClip `json:"clip"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotAction].
	Scale float64 `json:"scale"`
	JSON  struct {
		X                  respjson.Field
		Y                  respjson.Field
		Button             respjson.Field
		Double             respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		Points             respjson.Field
		Path               respjson.Field
		Duration           respjson.Field
		End                respjson.Field
		Start              respjson.Field
		ScrollX            respjson.Field
		ScrollY            respjson.Field
		Direction          respjson.Field
		Distance           respjson.Field
		Location           respjson.Field
		Keys               respjson.Field
		Combination        respjson.Field
		Buttons            respjson.Field
		Text               respjson.Field
		Mode               respjson.Field
		PressEnter         respjson.Field
		Orientation        respjson.Field
		Clip               respjson.Field
		Scale              respjson.Field
		raw                string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedClickAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedClickAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedTouchAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedDragAdvancedAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedDragSimpleAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedScrollAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedScrollSimpleAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedSwipeSimpleAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedSwipeAdvancedAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedPressKeyAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressKeyAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedPressButtonAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressButtonAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedLongPressAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedLongPressAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedTypeAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedTypeAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedMoveAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedMoveAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedScreenRotationAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenRotationAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedScreenshotAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsV1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsV1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) AsTypedWaitAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedWaitAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionResultAIResponseActionUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionResultAIResponseActionUnionEnd is an implicit
// subunion of [V1BoxActionAIResponseAIActionResultAIResponseActionUnion].
// V1BoxActionAIResponseAIActionResultAIResponseActionUnionEnd provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxActionAIResponseAIActionResultAIResponseActionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionResultAIResponseActionUnionEnd struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string  `json:",inline"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	JSON     struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionUnionEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionResultAIResponseActionUnionStart is an implicit
// subunion of [V1BoxActionAIResponseAIActionResultAIResponseActionUnion].
// V1BoxActionAIResponseAIActionResultAIResponseActionUnionStart provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxActionAIResponseAIActionResultAIResponseActionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionResultAIResponseActionUnionStart struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string  `json:",inline"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	JSON     struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionUnionStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionResultAIResponseActionUnionDistance is an implicit
// subunion of [V1BoxActionAIResponseAIActionResultAIResponseActionUnion].
// V1BoxActionAIResponseAIActionResultAIResponseActionUnionDistance provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxActionAIResponseAIActionResultAIResponseActionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat
// OfV1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString]
type V1BoxActionAIResponseAIActionResultAIResponseActionUnionDistance struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfV1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString string `json:",inline"`
	JSON                                                                                      struct {
		OfFloat                                                                                   respjson.Field
		OfV1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString respjson.Field
		raw                                                                                       string
	} `json:"-"`
}

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionUnionDistance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed click action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedClickAction struct {
	// X coordinate of the click
	X float64 `json:"x,required"`
	// Y coordinate of the click
	Y float64 `json:"y,required"`
	// Mouse button to click
	//
	// Any of "left", "right", "middle".
	Button string `json:"button"`
	// Whether to perform a double click
	Double bool `json:"double"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X                  respjson.Field
		Y                  respjson.Field
		Button             respjson.Field
		Double             respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedClickAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedClickAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed touch action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchAction struct {
	// Array of touch points and their actions
	Points []V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPoint `json:"points,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Points             respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Touch point configuration with start position and actions
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPoint struct {
	// Initial touch point position
	Start V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointStart `json:"start,required"`
	// Sequence of actions to perform after initial touch
	Actions []V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionUnion `json:"actions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Start       respjson.Field
		Actions     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Initial touch point position
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointStart struct {
	// Starting X coordinate
	X float64 `json:"x,required"`
	// Starting Y coordinate
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
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointStart) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionUnion
// contains all possible properties and values from
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionUnion struct {
	Duration string `json:"duration"`
	Type     string `json:"type"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction].
	X float64 `json:"x"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction].
	Y    float64 `json:"y"`
	JSON struct {
		Duration respjson.Field
		Type     respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionUnion) AsTouchPointMoveAction() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionUnion) AsV1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Touch point movement action configuration
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction struct {
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Type        respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointMoveAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto struct {
	// Duration to wait (e.g. "500ms")
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration,required"`
	// Type of the action
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPointActionTouchPointWaitActionDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed drag advanced action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedAction struct {
	// Path of the drag action as a series of coordinates
	Path []V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedActionPath `json:"path,required"`
	// Time interval between points (e.g. "50ms")
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 50ms
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path               respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single point in a drag path
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedActionPath struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
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
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedActionPath) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedActionPath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed drag simple action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleAction struct {
	// End point of the drag path (coordinates or natural language)
	End V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndUnion `json:"end,required"`
	// Start point of the drag path (coordinates or natural language)
	Start V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartUnion `json:"start,required"`
	// Duration to complete the movement from start to end coordinates
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End                respjson.Field
		Start              respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndUnion
// contains all possible properties and values from
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndDragPathPoint],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndDragPathPoint].
	X float64 `json:"x"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndDragPathPoint].
	Y    float64 `json:"y"`
	JSON struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndUnion) AsDragPathPoint() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndDragPathPoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single point in a drag path
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndDragPathPoint struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
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
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndDragPathPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEndDragPathPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartUnion
// contains all possible properties and values from
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartDragPathPoint],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartDragPathPoint].
	X float64 `json:"x"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartDragPathPoint].
	Y    float64 `json:"y"`
	JSON struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartUnion) AsDragPathPoint() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartDragPathPoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single point in a drag path
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartDragPathPoint struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
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
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartDragPathPoint) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStartDragPathPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed scroll action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction struct {
	// Horizontal scroll amount. Positive values scroll content rightward (reveals
	// content on the left), negative values scroll content leftward (reveals content
	// on the right).
	ScrollX float64 `json:"scrollX,required"`
	// Vertical scroll amount. Positive values scroll content downward (reveals content
	// above), negative values scroll content upward (reveals content below).
	ScrollY float64 `json:"scrollY,required"`
	// X coordinate of the scroll position
	X float64 `json:"x,required"`
	// Y coordinate of the scroll position
	Y float64 `json:"y,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScrollX            respjson.Field
		ScrollY            respjson.Field
		X                  respjson.Field
		Y                  respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed scroll simple action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleAction struct {
	// Direction to scroll. The scroll will be performed from the center of the screen
	// towards this direction. 'up' scrolls content upward (reveals content below),
	// 'down' scrolls content downward (reveals content above), 'left' scrolls content
	// leftward (reveals content on the right), 'right' scrolls content rightward
	// (reveals content on the left).
	//
	// Any of "up", "down", "left", "right".
	Direction string `json:"direction,required"`
	// Distance of the scroll. Can be either a number (in pixels) or a predefined enum
	// value (tiny, short, medium, long). If not provided, the scroll will be performed
	// from the center of the screen to the screen edge
	Distance V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceUnion `json:"distance"`
	// Duration of the scroll
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direction          respjson.Field
		Distance           respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceUnion
// contains all possible properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat
// OfV1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceString]
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfV1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceString string `json:",inline"`
	JSON                                                                                       struct {
		OfFloat                                                                                    respjson.Field
		OfV1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceString respjson.Field
		raw                                                                                        string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceUnion) AsV1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceString string

const (
	V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceStringTiny   V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceString = "tiny"
	V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceStringShort  V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceString = "short"
	V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceStringMedium V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceString = "medium"
	V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceStringLong   V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollSimpleActionDistanceString = "long"
)

// Typed swipe simple action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleAction struct {
	// Direction to swipe. The gesture will be performed from the center of the screen
	// towards this direction.
	//
	// Any of "up", "down", "left", "right", "upLeft", "upRight", "downLeft",
	// "downRight".
	Direction string `json:"direction,required"`
	// Distance of the swipe. Can be either a number (in pixels) or a predefined enum
	// value (tiny, short, medium, long). If not provided, the swipe will be performed
	// from the center of the screen to the screen edge
	Distance V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceUnion `json:"distance"`
	// Duration of the swipe
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Natural language description of the location where the swipe should originate.
	// If not provided, the swipe will be performed from the center of the screen.
	Location string `json:"location"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Direction          respjson.Field
		Distance           respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		Location           respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceUnion
// contains all possible properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat
// OfV1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString]
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfV1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString string `json:",inline"`
	JSON                                                                                      struct {
		OfFloat                                                                                   respjson.Field
		OfV1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString respjson.Field
		raw                                                                                       string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceUnion) AsV1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString string

const (
	V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceStringTiny   V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString = "tiny"
	V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceStringShort  V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString = "short"
	V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceStringMedium V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString = "medium"
	V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceStringLong   V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleActionDistanceString = "long"
)

// Typed swipe advanced action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedAction struct {
	// End point of the swipe path (coordinates or natural language)
	End V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndUnion `json:"end,required"`
	// Start point of the swipe path (coordinates or natural language)
	Start V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartUnion `json:"start,required"`
	// Duration of the swipe
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End                respjson.Field
		Start              respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndUnion
// contains all possible properties and values from
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath].
	X float64 `json:"x"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath].
	Y    float64 `json:"y"`
	JSON struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndUnion) AsSwipePath() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swipe path
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath struct {
	// Start/end x coordinate of the swipe path
	X float64 `json:"x,required"`
	// Start/end y coordinate of the swipe path
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
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEndSwipePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartUnion
// contains all possible properties and values from
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath].
	X float64 `json:"x"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath].
	Y    float64 `json:"y"`
	JSON struct {
		OfString respjson.Field
		X        respjson.Field
		Y        respjson.Field
		raw      string
	} `json:"-"`
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartUnion) AsSwipePath() (v V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swipe path
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath struct {
	// Start/end x coordinate of the swipe path
	X float64 `json:"x,required"`
	// Start/end y coordinate of the swipe path
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
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStartSwipePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed press key action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressKeyAction struct {
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
	Keys []string `json:"keys,required"`
	// Whether to press keys as combination (simultaneously) or sequentially. When
	// true, all keys are pressed together as a shortcut (e.g., Ctrl+C). When false,
	// keys are pressed one by one in sequence.
	Combination bool `json:"combination"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Keys               respjson.Field
		Combination        respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressKeyAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressKeyAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed press button action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressButtonAction struct {
	// Button to press
	//
	// Any of "power", "volumeUp", "volumeDown", "volumeMute", "home", "back", "menu",
	// "appSwitch".
	Buttons []string `json:"buttons,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buttons            respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressButtonAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressButtonAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed long press action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedLongPressAction struct {
	// X coordinate of the long press
	X float64 `json:"x,required"`
	// Y coordinate of the long press
	Y float64 `json:"y,required"`
	// Duration to hold the press (e.g. '1s', '500ms')
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 1s
	Duration string `json:"duration"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X                  respjson.Field
		Y                  respjson.Field
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedLongPressAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedLongPressAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed type action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedTypeAction struct {
	// Text to type
	Text string `json:"text,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Text input mode: 'append' to add text to existing content, 'replace' to replace
	// all existing text
	//
	// Any of "append", "replace".
	Mode string `json:"mode"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
	// Whether to press Enter after typing the text
	PressEnter bool `json:"pressEnter"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text               respjson.Field
		IncludeScreenshot  respjson.Field
		Mode               respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		PressEnter         respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedTypeAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedTypeAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed move action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedMoveAction struct {
	// X coordinate to move to
	X float64 `json:"x,required"`
	// Y coordinate to move to
	Y float64 `json:"y,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X                  respjson.Field
		Y                  respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedMoveAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedMoveAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed screen rotation action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenRotationAction struct {
	// Target screen orientation
	//
	// Any of "portrait", "landscapeLeft", "portraitUpsideDown", "landscapeRight".
	Orientation string `json:"orientation,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Orientation        respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenRotationAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenRotationAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed screenshot action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotAction struct {
	// Clipping region for screenshot capture
	Clip V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotActionClip `json:"clip"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
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
	Scale float64 `json:"scale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clip         respjson.Field
		OutputFormat respjson.Field
		Scale        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Clipping region for screenshot capture
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotActionClip struct {
	// Height of the clip
	Height float64 `json:"height,required"`
	// Width of the clip
	Width float64 `json:"width,required"`
	// X coordinate of the clip
	X float64 `json:"x,required"`
	// Y coordinate of the clip
	Y float64 `json:"y,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotActionClip) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotActionClip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed wait action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedWaitAction struct {
	// Duration of the wait (e.g. '3s')
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 3s
	Duration string `json:"duration,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot bool `json:"includeScreenshot"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn string `json:"presignedExpiresIn"`
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
	ScreenshotDelay string `json:"screenshotDelay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration           respjson.Field
		IncludeScreenshot  respjson.Field
		OutputFormat       respjson.Field
		PresignedExpiresIn respjson.Field
		ScreenshotDelay    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedWaitAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedWaitAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionClickResponseUnion contains all possible properties and values from
// [V1BoxActionClickResponseActionIncludeScreenshotResult],
// [V1BoxActionClickResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionClickResponseUnion struct {
	// This field is from variant
	// [V1BoxActionClickResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionClickResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant [V1BoxActionClickResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionClickResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionClickResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionClickResponseUnion) AsActionCommonResult() (v V1BoxActionClickResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionClickResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionClickResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionClickResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionClickResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionClickResponseActionIncludeScreenshotResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionClickResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionClickResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionClickResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionClickResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionClickResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionClickResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionClickResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionClickResponseActionCommonResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionDragResponseUnion contains all possible properties and values from
// [V1BoxActionDragResponseActionIncludeScreenshotResult],
// [V1BoxActionDragResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionDragResponseUnion struct {
	// This field is from variant
	// [V1BoxActionDragResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionDragResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant [V1BoxActionDragResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionDragResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionDragResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionDragResponseUnion) AsActionCommonResult() (v V1BoxActionDragResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionDragResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionDragResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionDragResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionDragResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionDragResponseActionIncludeScreenshotResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionDragResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionDragResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionDragResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionDragResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionDragResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionDragResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionDragResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionDragResponseActionCommonResult) UnmarshalJSON(data []byte) error {
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

// V1BoxActionLongPressResponseUnion contains all possible properties and values
// from [V1BoxActionLongPressResponseActionIncludeScreenshotResult],
// [V1BoxActionLongPressResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionLongPressResponseUnion struct {
	// This field is from variant
	// [V1BoxActionLongPressResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant [V1BoxActionLongPressResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionLongPressResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionLongPressResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionLongPressResponseUnion) AsActionCommonResult() (v V1BoxActionLongPressResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionLongPressResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionLongPressResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionLongPressResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionLongPressResponseActionIncludeScreenshotResult) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionLongPressResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionLongPressResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionLongPressResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionLongPressResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionLongPressResponseActionCommonResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionMoveResponseUnion contains all possible properties and values from
// [V1BoxActionMoveResponseActionIncludeScreenshotResult],
// [V1BoxActionMoveResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionMoveResponseUnion struct {
	// This field is from variant
	// [V1BoxActionMoveResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant [V1BoxActionMoveResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionMoveResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionMoveResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionMoveResponseUnion) AsActionCommonResult() (v V1BoxActionMoveResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionMoveResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionMoveResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionMoveResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionMoveResponseActionIncludeScreenshotResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionMoveResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionMoveResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionMoveResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionMoveResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionMoveResponseActionCommonResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionPressButtonResponseUnion contains all possible properties and values
// from [V1BoxActionPressButtonResponseActionIncludeScreenshotResult],
// [V1BoxActionPressButtonResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionPressButtonResponseUnion struct {
	// This field is from variant
	// [V1BoxActionPressButtonResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant [V1BoxActionPressButtonResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionPressButtonResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionPressButtonResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionPressButtonResponseUnion) AsActionCommonResult() (v V1BoxActionPressButtonResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionPressButtonResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionPressButtonResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionPressButtonResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionPressButtonResponseActionIncludeScreenshotResult) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionPressButtonResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionPressButtonResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionPressButtonResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionPressButtonResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressButtonResponseActionCommonResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionPressKeyResponseUnion contains all possible properties and values
// from [V1BoxActionPressKeyResponseActionIncludeScreenshotResult],
// [V1BoxActionPressKeyResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionPressKeyResponseUnion struct {
	// This field is from variant
	// [V1BoxActionPressKeyResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant [V1BoxActionPressKeyResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionPressKeyResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionPressKeyResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionPressKeyResponseUnion) AsActionCommonResult() (v V1BoxActionPressKeyResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionPressKeyResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionPressKeyResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionPressKeyResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionPressKeyResponseActionIncludeScreenshotResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressKeyResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionPressKeyResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionPressKeyResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionPressKeyResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionPressKeyResponseActionCommonResult) UnmarshalJSON(data []byte) error {
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

// V1BoxActionScreenRotationResponseUnion contains all possible properties and
// values from [V1BoxActionScreenRotationResponseActionIncludeScreenshotResult],
// [V1BoxActionScreenRotationResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionScreenRotationResponseUnion struct {
	// This field is from variant
	// [V1BoxActionScreenRotationResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant
	// [V1BoxActionScreenRotationResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionScreenRotationResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionScreenRotationResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionScreenRotationResponseUnion) AsActionCommonResult() (v V1BoxActionScreenRotationResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionScreenRotationResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionScreenRotationResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionScreenRotationResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionScreenRotationResponseActionIncludeScreenshotResult) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionScreenRotationResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionScreenRotationResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionScreenRotationResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionScreenRotationResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScreenRotationResponseActionCommonResult) UnmarshalJSON(data []byte) error {
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

// V1BoxActionScrollResponseUnion contains all possible properties and values from
// [V1BoxActionScrollResponseActionIncludeScreenshotResult],
// [V1BoxActionScrollResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionScrollResponseUnion struct {
	// This field is from variant
	// [V1BoxActionScrollResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant [V1BoxActionScrollResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionScrollResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionScrollResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionScrollResponseUnion) AsActionCommonResult() (v V1BoxActionScrollResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionScrollResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionScrollResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionScrollResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionScrollResponseActionIncludeScreenshotResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScrollResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionScrollResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionScrollResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionScrollResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionScrollResponseActionCommonResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action setting
type V1BoxActionSettingResponse struct {
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
	Scale float64 `json:"scale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scale       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionSettingResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSettingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action setting
type V1BoxActionSettingResetResponse struct {
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
	Scale float64 `json:"scale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scale       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionSettingResetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSettingResetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action setting
type V1BoxActionSettingUpdateResponse struct {
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
	Scale float64 `json:"scale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scale       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionSettingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSettingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionSwipeResponseUnion contains all possible properties and values from
// [V1BoxActionSwipeResponseActionIncludeScreenshotResult],
// [V1BoxActionSwipeResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionSwipeResponseUnion struct {
	// This field is from variant
	// [V1BoxActionSwipeResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant [V1BoxActionSwipeResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionSwipeResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionSwipeResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionSwipeResponseUnion) AsActionCommonResult() (v V1BoxActionSwipeResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionSwipeResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionSwipeResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionSwipeResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionSwipeResponseActionIncludeScreenshotResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSwipeResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionSwipeResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionSwipeResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionSwipeResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionSwipeResponseActionCommonResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionTapResponseUnion contains all possible properties and values from
// [V1BoxActionTapResponseActionIncludeScreenshotResult],
// [V1BoxActionTapResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionTapResponseUnion struct {
	// This field is from variant
	// [V1BoxActionTapResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionTapResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant [V1BoxActionTapResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionTapResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionTapResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionTapResponseUnion) AsActionCommonResult() (v V1BoxActionTapResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionTapResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionTapResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionTapResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionTapResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionTapResponseActionIncludeScreenshotResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTapResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionTapResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionTapResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTapResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTapResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionTapResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionTapResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTapResponseActionCommonResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionTouchResponseUnion contains all possible properties and values from
// [V1BoxActionTouchResponseActionIncludeScreenshotResult],
// [V1BoxActionTouchResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionTouchResponseUnion struct {
	// This field is from variant
	// [V1BoxActionTouchResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant [V1BoxActionTouchResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionTouchResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionTouchResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionTouchResponseUnion) AsActionCommonResult() (v V1BoxActionTouchResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionTouchResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionTouchResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionTouchResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionTouchResponseActionIncludeScreenshotResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTouchResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTouchResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionTouchResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionTouchResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTouchResponseActionCommonResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxActionTypeResponseUnion contains all possible properties and values from
// [V1BoxActionTypeResponseActionIncludeScreenshotResult],
// [V1BoxActionTypeResponseActionCommonResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxActionTypeResponseUnion struct {
	// This field is from variant
	// [V1BoxActionTypeResponseActionIncludeScreenshotResult].
	Screenshot V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshot `json:"screenshot"`
	// This field is from variant [V1BoxActionTypeResponseActionCommonResult].
	Message string `json:"message"`
	JSON    struct {
		Screenshot respjson.Field
		Message    respjson.Field
		raw        string
	} `json:"-"`
}

func (u V1BoxActionTypeResponseUnion) AsActionIncludeScreenshotResult() (v V1BoxActionTypeResponseActionIncludeScreenshotResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxActionTypeResponseUnion) AsActionCommonResult() (v V1BoxActionTypeResponseActionCommonResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxActionTypeResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxActionTypeResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution with screenshots
type V1BoxActionTypeResponseActionIncludeScreenshotResult struct {
	// Complete screenshot result with operation trace, before and after images
	Screenshot V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshot `json:"screenshot,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Screenshot  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionTypeResponseActionIncludeScreenshotResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTypeResponseActionIncludeScreenshotResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete screenshot result with operation trace, before and after images
type V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshot struct {
	// Screenshot taken after action execution
	After V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotAfter `json:"after,required"`
	// Screenshot taken before action execution
	Before V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotBefore `json:"before,required"`
	// Screenshot with action operation trace
	Trace V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotTrace `json:"trace,required"`
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
func (r V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshot) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken after action execution
type V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotAfter struct {
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
func (r V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotAfter) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot taken before action execution
type V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotBefore struct {
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
func (r V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotBefore) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Screenshot with action operation trace
type V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotTrace struct {
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
func (r V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotTrace) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionTypeResponseActionIncludeScreenshotResultScreenshotTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of an UI action execution
type V1BoxActionTypeResponseActionCommonResult struct {
	// message
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionTypeResponseActionCommonResult) RawJSON() string { return r.JSON.raw }
func (r *V1BoxActionTypeResponseActionCommonResult) UnmarshalJSON(data []byte) error {
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Type of the URI. default is base64.
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

// Type of the URI. default is base64.
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Mouse button to click
	//
	// Any of "left", "right", "middle".
	Button string `json:"button,omitzero"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
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

// Click action configuration with natural language
//
// The property Target is required.
type V1BoxActionClickParamsBodyClickActionWithNaturalLanguage struct {
	// Describe the target to operate using natural language, e.g., 'login button' or
	// 'Chrome'.
	Target string `json:"target,required"`
	// Whether to perform a double click
	Double param.Opt[bool] `json:"double,omitzero"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Mouse button to click
	//
	// Any of "left", "right", "middle".
	Button string `json:"button,omitzero"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
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

type V1BoxActionMoveParams struct {
	// X coordinate to move to
	X float64 `json:"x,required"`
	// Y coordinate to move to
	Y float64 `json:"y,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Whether to press keys as combination (simultaneously) or sequentially. When
	// true, all keys are pressed together as a shortcut (e.g., Ctrl+C). When false,
	// keys are pressed one by one in sequence.
	Combination param.Opt[bool] `json:"combination,omitzero"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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

type V1BoxActionScreenRotationParams struct {
	// Target screen orientation
	//
	// Any of "portrait", "landscapeLeft", "portraitUpsideDown", "landscapeRight".
	Orientation V1BoxActionScreenRotationParamsOrientation `json:"orientation,omitzero,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Type of the URI. default is base64.
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

// Type of the URI. default is base64.
type V1BoxActionScreenRotationParamsOutputFormat string

const (
	V1BoxActionScreenRotationParamsOutputFormatBase64     V1BoxActionScreenRotationParamsOutputFormat = "base64"
	V1BoxActionScreenRotationParamsOutputFormatStorageKey V1BoxActionScreenRotationParamsOutputFormat = "storageKey"
)

type V1BoxActionScreenshotParams struct {
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
	// scroll content downward (reveal content above), negative scrollY to scroll
	// content upward (reveal content below). Use positive scrollX to scroll content
	// rightward (reveal content on the left), negative scrollX to scroll content
	// leftward (reveal content on the right).
	OfScrollAction *V1BoxActionScrollParamsBodyScrollAction `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Simple
	// scroll action configuration. The scroll will be performed from the center of the
	// screen towards the specified direction.
	OfScrollSimple *V1BoxActionScrollParamsBodyScrollSimple `json:",inline"`

	paramObj
}

func (u V1BoxActionScrollParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScrollAction, u.OfScrollSimple)
}
func (r *V1BoxActionScrollParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced scroll action configuration. The scroll will be performed from the
// specified coordinates with the given scroll amounts. Use positive scrollY to
// scroll content downward (reveal content above), negative scrollY to scroll
// content upward (reveal content below). Use positive scrollX to scroll content
// rightward (reveal content on the left), negative scrollX to scroll content
// leftward (reveal content on the right).
//
// The properties ScrollX, ScrollY, X, Y are required.
type V1BoxActionScrollParamsBodyScrollAction struct {
	// Horizontal scroll amount. Positive values scroll content rightward (reveals
	// content on the left), negative values scroll content leftward (reveals content
	// on the right).
	ScrollX float64 `json:"scrollX,required"`
	// Vertical scroll amount. Positive values scroll content downward (reveals content
	// above), negative values scroll content upward (reveals content below).
	ScrollY float64 `json:"scrollY,required"`
	// X coordinate of the scroll position
	X float64 `json:"x,required"`
	// Y coordinate of the scroll position
	Y float64 `json:"y,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
	OutputFormat string `json:"outputFormat,omitzero"`
	paramObj
}

func (r V1BoxActionScrollParamsBodyScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScrollParamsBodyScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScrollParamsBodyScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1BoxActionScrollParamsBodyScrollAction](
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Distance of the scroll. Can be either a number (in pixels) or a predefined enum
	// value (tiny, short, medium, long). If not provided, the scroll will be performed
	// from the center of the screen to the screen edge
	Distance V1BoxActionScrollParamsBodyScrollSimpleDistanceUnion `json:"distance,omitzero"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
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

type V1BoxActionSettingUpdateParams struct {
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
	Scale param.Opt[float64] `json:"scale,omitzero"`
	paramObj
}

func (r V1BoxActionSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSettingUpdateParams) UnmarshalJSON(data []byte) error {
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Natural language description of the location where the swipe should originate.
	// If not provided, the swipe will be performed from the center of the screen.
	Location param.Opt[string] `json:"location,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Distance of the swipe. Can be either a number (in pixels) or a predefined enum
	// value (tiny, short, medium, long). If not provided, the swipe will be performed
	// from the center of the screen to the screen edge
	Distance V1BoxActionSwipeParamsBodySwipeSimpleDistanceUnion `json:"distance,omitzero"`
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
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

// Tap action configuration with natural language
//
// The property Target is required.
type V1BoxActionTapParamsBodyTapActionWithNaturalLanguage struct {
	// Describe the target to operate using natural language, e.g., 'login button' or
	// 'Chrome'.
	Target string `json:"target,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	// Type of the URI. default is base64.
	//
	// Any of "base64", "storageKey".
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

type V1BoxActionTouchParams struct {
	// Array of touch points and their actions
	Points []V1BoxActionTouchParamsPoint `json:"points,omitzero,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
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
	OfTouchPointMoveAction                                *V1BoxActionTouchParamsPointActionTouchPointMoveAction    `json:",omitzero,inline"`
	OfV1BoxActionTouchsPointActionTouchPointWaitActionDto *V1BoxActionTouchParamsPointActionTouchPointWaitActionDto `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxActionTouchParamsPointActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTouchPointMoveAction, u.OfV1BoxActionTouchsPointActionTouchPointWaitActionDto)
}
func (u *V1BoxActionTouchParamsPointActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxActionTouchParamsPointActionUnion) asAny() any {
	if !param.IsOmitted(u.OfTouchPointMoveAction) {
		return u.OfTouchPointMoveAction
	} else if !param.IsOmitted(u.OfV1BoxActionTouchsPointActionTouchPointWaitActionDto) {
		return u.OfV1BoxActionTouchsPointActionTouchPointWaitActionDto
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
	} else if vt := u.OfV1BoxActionTouchsPointActionTouchPointWaitActionDto; vt != nil {
		return (*string)(&vt.Duration)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u V1BoxActionTouchParamsPointActionUnion) GetType() *string {
	if vt := u.OfTouchPointMoveAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfV1BoxActionTouchsPointActionTouchPointWaitActionDto; vt != nil {
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

// The properties Duration, Type are required.
type V1BoxActionTouchParamsPointActionTouchPointWaitActionDto struct {
	// Duration to wait (e.g. "500ms")
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration string `json:"duration,required"`
	// Type of the action
	Type string `json:"type,required"`
	paramObj
}

func (r V1BoxActionTouchParamsPointActionTouchPointWaitActionDto) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionTouchParamsPointActionTouchPointWaitActionDto
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionTouchParamsPointActionTouchPointWaitActionDto) UnmarshalJSON(data []byte) error {
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
	// Presigned url expires in. Only takes effect when outputFormat is storageKey.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30m
	PresignedExpiresIn param.Opt[string] `json:"presignedExpiresIn,omitzero"`
	// Whether to press Enter after typing the text
	PressEnter param.Opt[bool] `json:"pressEnter,omitzero"`
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

// Text input mode: 'append' to add text to existing content, 'replace' to replace
// all existing text
type V1BoxActionTypeParamsMode string

const (
	V1BoxActionTypeParamsModeAppend  V1BoxActionTypeParamsMode = "append"
	V1BoxActionTypeParamsModeReplace V1BoxActionTypeParamsMode = "replace"
)

// Type of the URI. default is base64.
type V1BoxActionTypeParamsOutputFormat string

const (
	V1BoxActionTypeParamsOutputFormatBase64     V1BoxActionTypeParamsOutputFormat = "base64"
	V1BoxActionTypeParamsOutputFormatStorageKey V1BoxActionTypeParamsOutputFormat = "storageKey"
)
