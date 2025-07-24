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

// Use natural language instructions to perform UI operations on the box. You can
// describe what you want to do in plain language (e.g., 'click the login button',
// 'scroll down to find settings', 'input my email address'), and the AI will
// automatically convert your instruction into the appropriate UI action and
// execute it on the box.
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

// Press button on the device. like power button, volume up button, volume down
// button, etc.
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

// Rotate screen
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

// Scroll
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
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressKeyAction],
// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressButtonAction],
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
	Double            bool   `json:"double"`
	IncludeScreenshot bool   `json:"includeScreenshot"`
	OutputFormat      string `json:"outputFormat"`
	ScreenshotDelay   string `json:"screenshotDelay"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchAction].
	Points []V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedTouchActionPoint `json:"points"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedAction].
	Path     []V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragAdvancedActionPath `json:"path"`
	Duration string                                                                                     `json:"duration"`
	// This field is a union of
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEnd],
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEnd]
	End V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionEnd `json:"end"`
	// This field is a union of
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStart],
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStart]
	Start V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionStart `json:"start"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction].
	ScrollX float64 `json:"scrollX"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction].
	ScrollY   float64 `json:"scrollY"`
	Direction string  `json:"direction"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleAction].
	Distance float64 `json:"distance"`
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
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenRotationAction].
	Angle float64 `json:"angle"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotAction].
	Clip V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScreenshotActionClip `json:"clip"`
	JSON struct {
		X                 respjson.Field
		Y                 respjson.Field
		Button            respjson.Field
		Double            respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		Points            respjson.Field
		Path              respjson.Field
		Duration          respjson.Field
		End               respjson.Field
		Start             respjson.Field
		ScrollX           respjson.Field
		ScrollY           respjson.Field
		Direction         respjson.Field
		Distance          respjson.Field
		Keys              respjson.Field
		Combination       respjson.Field
		Buttons           respjson.Field
		Text              respjson.Field
		Mode              respjson.Field
		Angle             respjson.Field
		Clip              respjson.Field
		raw               string
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
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionEnd struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	JSON struct {
		X   respjson.Field
		Y   respjson.Field
		raw string
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
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionStart struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	JSON struct {
		X   respjson.Field
		Y   respjson.Field
		raw string
	} `json:"-"`
}

func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionUnionStart) UnmarshalJSON(data []byte) error {
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
		X                 respjson.Field
		Y                 respjson.Field
		Button            respjson.Field
		Double            respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
		Points            respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
		Path              respjson.Field
		Duration          respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
	// Single point in a drag path
	End V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEnd `json:"end,required"`
	// Single point in a drag path
	Start V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStart `json:"start,required"`
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
		End               respjson.Field
		Start             respjson.Field
		Duration          respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single point in a drag path
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEnd struct {
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
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEnd) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single point in a drag path
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStart struct {
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
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStart) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedDragSimpleActionStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed scroll action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction struct {
	// Horizontal scroll amount
	ScrollX float64 `json:"scrollX,required"`
	// Vertical scroll amount
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
		ScrollX           respjson.Field
		ScrollY           respjson.Field
		X                 respjson.Field
		Y                 respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed swipe simple action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleAction struct {
	// Direction to swipe. The gesture will be performed from the center of the screen
	// towards this direction.
	//
	// Any of "up", "down", "left", "right", "upLeft", "upRight", "downLeft",
	// "downRight".
	Direction string `json:"direction,required"`
	// Distance of the swipe in pixels. If not provided, the swipe will be performed
	// from the center of the screen to the screen edge
	Distance float64 `json:"distance"`
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
		Direction         respjson.Field
		Distance          respjson.Field
		Duration          respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeSimpleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed swipe advanced action
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedAction struct {
	// Swipe path
	End V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEnd `json:"end,required"`
	// Swipe path
	Start V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStart `json:"start,required"`
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
		End               respjson.Field
		Start             respjson.Field
		Duration          respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swipe path
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEnd struct {
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
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEnd) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swipe path
type V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStart struct {
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
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStart) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedSwipeAdvancedActionStart) UnmarshalJSON(data []byte) error {
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
		Keys              respjson.Field
		Combination       respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
		Buttons           respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressButtonAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionScreenshotResultAIResponseActionTypedPressButtonAction) UnmarshalJSON(data []byte) error {
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
		Text              respjson.Field
		IncludeScreenshot respjson.Field
		Mode              respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
		X                 respjson.Field
		Y                 respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
	// Rotation angle in degrees
	//
	// Any of 90, 180, 270.
	Angle float64 `json:"angle,required"`
	// Rotation direction
	//
	// Any of "clockwise", "counter-clockwise".
	Direction string `json:"direction,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Angle       respjson.Field
		Direction   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clip         respjson.Field
		OutputFormat respjson.Field
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
		Duration          respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressKeyAction],
// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressButtonAction],
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
	Double            bool   `json:"double"`
	IncludeScreenshot bool   `json:"includeScreenshot"`
	OutputFormat      string `json:"outputFormat"`
	ScreenshotDelay   string `json:"screenshotDelay"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchAction].
	Points []V1BoxActionAIResponseAIActionResultAIResponseActionTypedTouchActionPoint `json:"points"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedAction].
	Path     []V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragAdvancedActionPath `json:"path"`
	Duration string                                                                           `json:"duration"`
	// This field is a union of
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEnd],
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEnd]
	End V1BoxActionAIResponseAIActionResultAIResponseActionUnionEnd `json:"end"`
	// This field is a union of
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStart],
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStart]
	Start V1BoxActionAIResponseAIActionResultAIResponseActionUnionStart `json:"start"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction].
	ScrollX float64 `json:"scrollX"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction].
	ScrollY   float64 `json:"scrollY"`
	Direction string  `json:"direction"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleAction].
	Distance float64 `json:"distance"`
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
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenRotationAction].
	Angle float64 `json:"angle"`
	// This field is from variant
	// [V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotAction].
	Clip V1BoxActionAIResponseAIActionResultAIResponseActionTypedScreenshotActionClip `json:"clip"`
	JSON struct {
		X                 respjson.Field
		Y                 respjson.Field
		Button            respjson.Field
		Double            respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		Points            respjson.Field
		Path              respjson.Field
		Duration          respjson.Field
		End               respjson.Field
		Start             respjson.Field
		ScrollX           respjson.Field
		ScrollY           respjson.Field
		Direction         respjson.Field
		Distance          respjson.Field
		Keys              respjson.Field
		Combination       respjson.Field
		Buttons           respjson.Field
		Text              respjson.Field
		Mode              respjson.Field
		Angle             respjson.Field
		Clip              respjson.Field
		raw               string
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
type V1BoxActionAIResponseAIActionResultAIResponseActionUnionEnd struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	JSON struct {
		X   respjson.Field
		Y   respjson.Field
		raw string
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
type V1BoxActionAIResponseAIActionResultAIResponseActionUnionStart struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	JSON struct {
		X   respjson.Field
		Y   respjson.Field
		raw string
	} `json:"-"`
}

func (r *V1BoxActionAIResponseAIActionResultAIResponseActionUnionStart) UnmarshalJSON(data []byte) error {
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
		X                 respjson.Field
		Y                 respjson.Field
		Button            respjson.Field
		Double            respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
		Points            respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
		Path              respjson.Field
		Duration          respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
	// Single point in a drag path
	End V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEnd `json:"end,required"`
	// Single point in a drag path
	Start V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStart `json:"start,required"`
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
		End               respjson.Field
		Start             respjson.Field
		Duration          respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single point in a drag path
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEnd struct {
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
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEnd) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single point in a drag path
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStart struct {
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
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStart) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedDragSimpleActionStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed scroll action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction struct {
	// Horizontal scroll amount
	ScrollX float64 `json:"scrollX,required"`
	// Vertical scroll amount
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
		ScrollX           respjson.Field
		ScrollY           respjson.Field
		X                 respjson.Field
		Y                 respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed swipe simple action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleAction struct {
	// Direction to swipe. The gesture will be performed from the center of the screen
	// towards this direction.
	//
	// Any of "up", "down", "left", "right", "upLeft", "upRight", "downLeft",
	// "downRight".
	Direction string `json:"direction,required"`
	// Distance of the swipe in pixels. If not provided, the swipe will be performed
	// from the center of the screen to the screen edge
	Distance float64 `json:"distance"`
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
		Direction         respjson.Field
		Distance          respjson.Field
		Duration          respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeSimpleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed swipe advanced action
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedAction struct {
	// Swipe path
	End V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEnd `json:"end,required"`
	// Swipe path
	Start V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStart `json:"start,required"`
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
		End               respjson.Field
		Start             respjson.Field
		Duration          respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swipe path
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEnd struct {
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
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEnd) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swipe path
type V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStart struct {
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
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStart) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedSwipeAdvancedActionStart) UnmarshalJSON(data []byte) error {
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
		Keys              respjson.Field
		Combination       respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
		Buttons           respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressButtonAction) RawJSON() string {
	return r.JSON.raw
}
func (r *V1BoxActionAIResponseAIActionResultAIResponseActionTypedPressButtonAction) UnmarshalJSON(data []byte) error {
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
		Text              respjson.Field
		IncludeScreenshot respjson.Field
		Mode              respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
		X                 respjson.Field
		Y                 respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
	// Rotation angle in degrees
	//
	// Any of 90, 180, 270.
	Angle float64 `json:"angle,required"`
	// Rotation direction
	//
	// Any of "clockwise", "counter-clockwise".
	Direction string `json:"direction,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Angle       respjson.Field
		Direction   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clip         respjson.Field
		OutputFormat respjson.Field
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
		Duration          respjson.Field
		IncludeScreenshot respjson.Field
		OutputFormat      respjson.Field
		ScreenshotDelay   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// X coordinate of the click
	X float64 `json:"x,required"`
	// Y coordinate of the click
	Y float64 `json:"y,required"`
	// Whether to perform a double click
	Double param.Opt[bool] `json:"double,omitzero"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
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
	// Single point in a drag path
	End V1BoxActionDragParamsBodyDragSimpleEnd `json:"end,omitzero,required"`
	// Single point in a drag path
	Start V1BoxActionDragParamsBodyDragSimpleStart `json:"start,omitzero,required"`
	// Duration to complete the movement from start to end coordinates
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration param.Opt[string] `json:"duration,omitzero"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
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

// Single point in a drag path
//
// The properties X, Y are required.
type V1BoxActionDragParamsBodyDragSimpleEnd struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionDragParamsBodyDragSimpleEnd) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsBodyDragSimpleEnd
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsBodyDragSimpleEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single point in a drag path
//
// The properties X, Y are required.
type V1BoxActionDragParamsBodyDragSimpleStart struct {
	// X coordinate of a point in the drag path
	X float64 `json:"x,required"`
	// Y coordinate of a point in the drag path
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionDragParamsBodyDragSimpleStart) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionDragParamsBodyDragSimpleStart
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionDragParamsBodyDragSimpleStart) UnmarshalJSON(data []byte) error {
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

type V1BoxActionMoveParams struct {
	// X coordinate to move to
	X float64 `json:"x,required"`
	// Y coordinate to move to
	Y float64 `json:"y,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
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

type V1BoxActionScreenRotationParams struct {
	// Rotation angle in degrees
	//
	// Any of 90, 180, 270.
	Angle float64 `json:"angle,omitzero,required"`
	// Rotation direction
	//
	// Any of "clockwise", "counter-clockwise".
	Direction V1BoxActionScreenRotationParamsDirection `json:"direction,omitzero,required"`
	paramObj
}

func (r V1BoxActionScreenRotationParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionScreenRotationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionScreenRotationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rotation direction
type V1BoxActionScreenRotationParamsDirection string

const (
	V1BoxActionScreenRotationParamsDirectionClockwise        V1BoxActionScreenRotationParamsDirection = "clockwise"
	V1BoxActionScreenRotationParamsDirectionCounterClockwise V1BoxActionScreenRotationParamsDirection = "counter-clockwise"
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
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
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
	// Distance of the swipe in pixels. If not provided, the swipe will be performed
	// from the center of the screen to the screen edge
	Distance param.Opt[float64] `json:"distance,omitzero"`
	// Duration of the swipe
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration param.Opt[string] `json:"duration,omitzero"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
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

// Swipe action configuration. The gesture will start from the specified start
// point and move towards the end point.
//
// The properties End, Start are required.
type V1BoxActionSwipeParamsBodySwipeAdvanced struct {
	// Swipe path
	End V1BoxActionSwipeParamsBodySwipeAdvancedEnd `json:"end,omitzero,required"`
	// Swipe path
	Start V1BoxActionSwipeParamsBodySwipeAdvancedStart `json:"start,omitzero,required"`
	// Duration of the swipe
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 500ms
	Duration param.Opt[string] `json:"duration,omitzero"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
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

// Swipe path
//
// The properties X, Y are required.
type V1BoxActionSwipeParamsBodySwipeAdvancedEnd struct {
	// Start/end x coordinate of the swipe path
	X float64 `json:"x,required"`
	// Start/end y coordinate of the swipe path
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionSwipeParamsBodySwipeAdvancedEnd) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSwipeParamsBodySwipeAdvancedEnd
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSwipeParamsBodySwipeAdvancedEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Swipe path
//
// The properties X, Y are required.
type V1BoxActionSwipeParamsBodySwipeAdvancedStart struct {
	// Start/end x coordinate of the swipe path
	X float64 `json:"x,required"`
	// Start/end y coordinate of the swipe path
	Y float64 `json:"y,required"`
	paramObj
}

func (r V1BoxActionSwipeParamsBodySwipeAdvancedStart) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxActionSwipeParamsBodySwipeAdvancedStart
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxActionSwipeParamsBodySwipeAdvancedStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxActionTouchParams struct {
	// Array of touch points and their actions
	Points []V1BoxActionTouchParamsPoint `json:"points,omitzero,required"`
	// Whether to include screenshots in the action response. If false, the screenshot
	// object will still be returned but with empty URIs. Default is false.
	IncludeScreenshot param.Opt[bool] `json:"includeScreenshot,omitzero"`
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
