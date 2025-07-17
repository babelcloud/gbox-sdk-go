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

	"github.com/babelcloud/gbox-sdk-go/internal/apijson"
	"github.com/babelcloud/gbox-sdk-go/internal/apiquery"
	"github.com/babelcloud/gbox-sdk-go/internal/requestconfig"
	"github.com/babelcloud/gbox-sdk-go/option"
	"github.com/babelcloud/gbox-sdk-go/packages/param"
	"github.com/babelcloud/gbox-sdk-go/packages/respjson"
)

// V1BoxService contains methods and other services that help with interacting with
// the gbox-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BoxService] method instead.
type V1BoxService struct {
	Options []option.RequestOption
	Actions V1BoxActionService
	Fs      V1BoxFService
	Browser V1BoxBrowserService
	Android V1BoxAndroidService
}

// NewV1BoxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1BoxService(opts ...option.RequestOption) (r V1BoxService) {
	r = V1BoxService{}
	r.Options = opts
	r.Actions = NewV1BoxActionService(opts...)
	r.Fs = NewV1BoxFService(opts...)
	r.Browser = NewV1BoxBrowserService(opts...)
	r.Android = NewV1BoxAndroidService(opts...)
	return
}

// Get box
func (r *V1BoxService) Get(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List box
func (r *V1BoxService) List(ctx context.Context, query V1BoxListParams, opts ...option.RequestOption) (res *V1BoxListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "boxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create android box
func (r *V1BoxService) NewAndroid(ctx context.Context, body V1BoxNewAndroidParams, opts ...option.RequestOption) (res *AndroidBox, err error) {
	opts = append(r.Options[:], opts...)
	path := "boxes/android"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create linux box
func (r *V1BoxService) NewLinux(ctx context.Context, body V1BoxNewLinuxParams, opts ...option.RequestOption) (res *LinuxBox, err error) {
	opts = append(r.Options[:], opts...)
	path := "boxes/linux"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the current display properties for a running box. This endpoint
// provides details about the box's screen resolution, orientation, and other
// visual properties
func (r *V1BoxService) Display(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxDisplayResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/display", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Execute a command on a running box. This endpoint allows you to send commands to
// the box and receive the output
func (r *V1BoxService) ExecuteCommands(ctx context.Context, boxID string, body V1BoxExecuteCommandsParams, opts ...option.RequestOption) (res *V1BoxExecuteCommandsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/commands", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows you to generate a pre-signed URL for accessing the live
// view of a running box. The URL is valid for a limited time and can be used to
// view the box's live stream
func (r *V1BoxService) LiveViewURL(ctx context.Context, boxID string, body V1BoxLiveViewURLParams, opts ...option.RequestOption) (res *V1BoxLiveViewURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/live-view-url", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Run code on the box
func (r *V1BoxService) RunCode(ctx context.Context, boxID string, body V1BoxRunCodeParams, opts ...option.RequestOption) (res *V1BoxRunCodeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/run-code", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Start box
func (r *V1BoxService) Start(ctx context.Context, boxID string, body V1BoxStartParams, opts ...option.RequestOption) (res *V1BoxStartResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/start", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop box
func (r *V1BoxService) Stop(ctx context.Context, boxID string, body V1BoxStopParams, opts ...option.RequestOption) (res *V1BoxStopResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/stop", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Terminate a running box. This action will stop the box and release its resources
func (r *V1BoxService) Terminate(ctx context.Context, boxID string, body V1BoxTerminateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/terminate", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// This endpoint allows you to generate a pre-signed URL for accessing the web
// terminal of a running box. The URL is valid for a limited time and can be used
// to access the box's terminal interface
func (r *V1BoxService) WebTerminalURL(ctx context.Context, boxID string, body V1BoxWebTerminalURLParams, opts ...option.RequestOption) (res *V1BoxWebTerminalURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/web-terminal-url", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Android box instance with full configuration and status
type AndroidBox struct {
	// Unique identifier for the box
	ID string `json:"id,required"`
	// Complete configuration for Android box instance
	Config AndroidBoxConfig `json:"config,required"`
	// Creation timestamp of the box
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Expiration timestamp of the box
	ExpiresAt time.Time `json:"expiresAt,required" format:"date-time"`
	// The current status of a box instance
	//
	// Any of "pending", "running", "error", "terminated".
	Status AndroidBoxStatus `json:"status,required"`
	// Box type is Android
	//
	// Any of "android".
	Type AndroidBoxType `json:"type,required"`
	// Last update timestamp of the box
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The reason for the current status, if any
	Reason string `json:"reason,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Config      respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AndroidBox) RawJSON() string { return r.JSON.raw }
func (r *AndroidBox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete configuration for Android box instance
type AndroidBoxConfig struct {
	// CPU cores allocated to the box
	CPU float64 `json:"cpu,required"`
	// Environment variables for the box. These variables will be available in all
	// operations including command execution, code running, and other box behaviors
	Envs map[string]string `json:"envs,required"`
	// Key-value pairs of labels for the box. Labels are used to add custom metadata to
	// help identify, categorize, and manage boxes. Common use cases include project
	// names, environments, teams, applications, or any other organizational tags that
	// help you organize and filter your boxes.
	Labels map[string]string `json:"labels,required"`
	// Memory allocated to the box in MiB
	Memory float64 `json:"memory,required"`
	// Android operating system configuration
	Os AndroidBoxConfigOs `json:"os,required"`
	// Storage allocated to the box in GiB
	Storage float64 `json:"storage,required"`
	// Device type - virtual or physical Android device
	//
	// Any of "virtual", "physical".
	DeviceType string `json:"deviceType"`
	// Working directory path for the box. This directory serves as the default
	// starting point for all operations including command execution, code running, and
	// file system operations. When you execute commands or run code, they will start
	// from this directory unless explicitly specified otherwise.
	WorkingDir string `json:"workingDir"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Envs        respjson.Field
		Labels      respjson.Field
		Memory      respjson.Field
		Os          respjson.Field
		Storage     respjson.Field
		DeviceType  respjson.Field
		WorkingDir  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AndroidBoxConfig) RawJSON() string { return r.JSON.raw }
func (r *AndroidBoxConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Android operating system configuration
type AndroidBoxConfigOs struct {
	// Supported Android versions
	//
	// Any of "12", "13", "15".
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AndroidBoxConfigOs) RawJSON() string { return r.JSON.raw }
func (r *AndroidBoxConfigOs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of a box instance
type AndroidBoxStatus string

const (
	AndroidBoxStatusPending    AndroidBoxStatus = "pending"
	AndroidBoxStatusRunning    AndroidBoxStatus = "running"
	AndroidBoxStatusError      AndroidBoxStatus = "error"
	AndroidBoxStatusTerminated AndroidBoxStatus = "terminated"
)

// Box type is Android
type AndroidBoxType string

const (
	AndroidBoxTypeAndroid AndroidBoxType = "android"
)

// Request body for creating a new Android box instance
type CreateAndroidBoxParam struct {
	// Timeout for waiting the box to transition from pending to running state, default
	// is 30s. If the box doesn't reach running state within this timeout, the API will
	// return HTTP status code 408. The timed-out box will be automatically deleted and
	// will not count towards your quota.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30s Maximum allowed: 5m
	Timeout param.Opt[string] `json:"timeout,omitzero"`
	// Wait for the box operation to be completed, default is true
	Wait param.Opt[bool] `json:"wait,omitzero"`
	// Configuration for a Android box instance
	Config CreateAndroidBoxConfigParam `json:"config,omitzero"`
	paramObj
}

func (r CreateAndroidBoxParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAndroidBoxParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAndroidBoxParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a Android box instance
type CreateAndroidBoxConfigParam struct {
	// The box will be alive for the given duration
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 60m
	ExpiresIn param.Opt[string] `json:"expiresIn,omitzero"`
	// Device type - virtual or physical Android device
	//
	// Any of "virtual", "physical".
	DeviceType string `json:"deviceType,omitzero"`
	// Environment variables for the box. These variables will be available in all
	// operations including command execution, code running, and other box behaviors
	Envs map[string]string `json:"envs,omitzero"`
	// Key-value pairs of labels for the box. Labels are used to add custom metadata to
	// help identify, categorize, and manage boxes. Common use cases include project
	// names, environments, teams, applications, or any other organizational tags that
	// help you organize and filter your boxes.
	Labels map[string]string `json:"labels,omitzero"`
	paramObj
}

func (r CreateAndroidBoxConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAndroidBoxConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAndroidBoxConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CreateAndroidBoxConfigParam](
		"deviceType", "virtual", "physical",
	)
}

// Request body for creating a new Linux box instance
type CreateLinuxBoxParam struct {
	// Timeout for waiting the box to transition from pending to running state, default
	// is 30s. If the box doesn't reach running state within this timeout, the API will
	// return HTTP status code 408. The timed-out box will be automatically deleted and
	// will not count towards your quota.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30s Maximum allowed: 5m
	Timeout param.Opt[string] `json:"timeout,omitzero"`
	// Wait for the box operation to be completed, default is true
	Wait param.Opt[bool] `json:"wait,omitzero"`
	// Configuration for a Linux box instance
	Config CreateLinuxBoxConfigParam `json:"config,omitzero"`
	paramObj
}

func (r CreateLinuxBoxParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateLinuxBoxParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateLinuxBoxParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a Linux box instance
type CreateLinuxBoxConfigParam struct {
	// The box will be alive for the given duration
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 60m
	ExpiresIn param.Opt[string] `json:"expiresIn,omitzero"`
	// Environment variables for the box. These variables will be available in all
	// operations including command execution, code running, and other box behaviors
	Envs map[string]string `json:"envs,omitzero"`
	// Key-value pairs of labels for the box. Labels are used to add custom metadata to
	// help identify, categorize, and manage boxes. Common use cases include project
	// names, environments, teams, applications, or any other organizational tags that
	// help you organize and filter your boxes.
	Labels map[string]string `json:"labels,omitzero"`
	paramObj
}

func (r CreateLinuxBoxConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateLinuxBoxConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateLinuxBoxConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Linux box instance with full configuration and status
type LinuxBox struct {
	// Unique identifier for the box
	ID string `json:"id,required"`
	// Complete configuration for Linux box instance
	Config LinuxBoxConfig `json:"config,required"`
	// Creation timestamp of the box
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Expiration timestamp of the box
	ExpiresAt time.Time `json:"expiresAt,required" format:"date-time"`
	// The current status of a box instance
	//
	// Any of "pending", "running", "error", "terminated".
	Status LinuxBoxStatus `json:"status,required"`
	// Box type is Linux
	//
	// Any of "linux".
	Type LinuxBoxType `json:"type,required"`
	// Last update timestamp of the box
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The reason for the current status, if any
	Reason string `json:"reason,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Config      respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinuxBox) RawJSON() string { return r.JSON.raw }
func (r *LinuxBox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete configuration for Linux box instance
type LinuxBoxConfig struct {
	// CPU cores allocated to the box
	CPU float64 `json:"cpu,required"`
	// Environment variables for the box. These variables will be available in all
	// operations including command execution, code running, and other box behaviors
	Envs map[string]string `json:"envs,required"`
	// Key-value pairs of labels for the box. Labels are used to add custom metadata to
	// help identify, categorize, and manage boxes. Common use cases include project
	// names, environments, teams, applications, or any other organizational tags that
	// help you organize and filter your boxes.
	Labels map[string]string `json:"labels,required"`
	// Memory allocated to the box in MiB
	Memory float64 `json:"memory,required"`
	// Linux operating system configuration
	Os LinuxBoxConfigOs `json:"os,required"`
	// Storage allocated to the box in GiB.
	Storage float64 `json:"storage,required"`
	// Working directory path for the box. This directory serves as the default
	// starting point for all operations including command execution, code running, and
	// file system operations. When you execute commands or run code, they will start
	// from this directory unless explicitly specified otherwise.
	WorkingDir string `json:"workingDir,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Envs        respjson.Field
		Labels      respjson.Field
		Memory      respjson.Field
		Os          respjson.Field
		Storage     respjson.Field
		WorkingDir  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinuxBoxConfig) RawJSON() string { return r.JSON.raw }
func (r *LinuxBoxConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Linux operating system configuration
type LinuxBoxConfigOs struct {
	// OS version string (e.g. 'ubuntu-20.04')
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinuxBoxConfigOs) RawJSON() string { return r.JSON.raw }
func (r *LinuxBoxConfigOs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of a box instance
type LinuxBoxStatus string

const (
	LinuxBoxStatusPending    LinuxBoxStatus = "pending"
	LinuxBoxStatusRunning    LinuxBoxStatus = "running"
	LinuxBoxStatusError      LinuxBoxStatus = "error"
	LinuxBoxStatusTerminated LinuxBoxStatus = "terminated"
)

// Box type is Linux
type LinuxBoxType string

const (
	LinuxBoxTypeLinux LinuxBoxType = "linux"
)

// V1BoxGetResponseUnion contains all possible properties and values from
// [LinuxBox], [AndroidBox].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxGetResponseUnion struct {
	ID string `json:"id"`
	// This field is a union of [LinuxBoxConfig], [AndroidBoxConfig]
	Config    V1BoxGetResponseUnionConfig `json:"config"`
	CreatedAt time.Time                   `json:"createdAt"`
	ExpiresAt time.Time                   `json:"expiresAt"`
	Status    string                      `json:"status"`
	Type      string                      `json:"type"`
	UpdatedAt time.Time                   `json:"updatedAt"`
	Reason    string                      `json:"reason"`
	JSON      struct {
		ID        respjson.Field
		Config    respjson.Field
		CreatedAt respjson.Field
		ExpiresAt respjson.Field
		Status    respjson.Field
		Type      respjson.Field
		UpdatedAt respjson.Field
		Reason    respjson.Field
		raw       string
	} `json:"-"`
}

func (u V1BoxGetResponseUnion) AsLinuxBox() (v LinuxBox) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxGetResponseUnion) AsAndroidBox() (v AndroidBox) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxGetResponseUnionConfig is an implicit subunion of [V1BoxGetResponseUnion].
// V1BoxGetResponseUnionConfig provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxGetResponseUnion].
type V1BoxGetResponseUnionConfig struct {
	CPU    float64 `json:"cpu"`
	Envs   string  `json:"envs"`
	Labels string  `json:"labels"`
	Memory float64 `json:"memory"`
	// This field is a union of [LinuxBoxConfigOs], [AndroidBoxConfigOs]
	Os         V1BoxGetResponseUnionConfigOs `json:"os"`
	Storage    float64                       `json:"storage"`
	WorkingDir string                        `json:"workingDir"`
	// This field is from variant [AndroidBoxConfig].
	DeviceType string `json:"deviceType"`
	JSON       struct {
		CPU        respjson.Field
		Envs       respjson.Field
		Labels     respjson.Field
		Memory     respjson.Field
		Os         respjson.Field
		Storage    respjson.Field
		WorkingDir respjson.Field
		DeviceType respjson.Field
		raw        string
	} `json:"-"`
}

func (r *V1BoxGetResponseUnionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxGetResponseUnionConfigOs is an implicit subunion of
// [V1BoxGetResponseUnion]. V1BoxGetResponseUnionConfigOs provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxGetResponseUnion].
type V1BoxGetResponseUnionConfigOs struct {
	Version string `json:"version"`
	JSON    struct {
		Version respjson.Field
		raw     string
	} `json:"-"`
}

func (r *V1BoxGetResponseUnionConfigOs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing paginated list of box instances
type V1BoxListResponse struct {
	// A box instance that can be either Linux or Android type
	Data []V1BoxListResponseDataUnion `json:"data,required"`
	// Page number
	Page int64 `json:"page,required"`
	// Page size
	PageSize int64 `json:"pageSize,required"`
	// Total number of items
	Total int64 `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Page        respjson.Field
		PageSize    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxListResponseDataUnion contains all possible properties and values from
// [LinuxBox], [AndroidBox].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxListResponseDataUnion struct {
	ID string `json:"id"`
	// This field is a union of [LinuxBoxConfig], [AndroidBoxConfig]
	Config    V1BoxListResponseDataUnionConfig `json:"config"`
	CreatedAt time.Time                        `json:"createdAt"`
	ExpiresAt time.Time                        `json:"expiresAt"`
	Status    string                           `json:"status"`
	Type      string                           `json:"type"`
	UpdatedAt time.Time                        `json:"updatedAt"`
	Reason    string                           `json:"reason"`
	JSON      struct {
		ID        respjson.Field
		Config    respjson.Field
		CreatedAt respjson.Field
		ExpiresAt respjson.Field
		Status    respjson.Field
		Type      respjson.Field
		UpdatedAt respjson.Field
		Reason    respjson.Field
		raw       string
	} `json:"-"`
}

func (u V1BoxListResponseDataUnion) AsLinuxBox() (v LinuxBox) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxListResponseDataUnion) AsAndroidBox() (v AndroidBox) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxListResponseDataUnionConfig is an implicit subunion of
// [V1BoxListResponseDataUnion]. V1BoxListResponseDataUnionConfig provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxListResponseDataUnion].
type V1BoxListResponseDataUnionConfig struct {
	CPU    float64 `json:"cpu"`
	Envs   string  `json:"envs"`
	Labels string  `json:"labels"`
	Memory float64 `json:"memory"`
	// This field is a union of [LinuxBoxConfigOs], [AndroidBoxConfigOs]
	Os         V1BoxListResponseDataUnionConfigOs `json:"os"`
	Storage    float64                            `json:"storage"`
	WorkingDir string                             `json:"workingDir"`
	// This field is from variant [AndroidBoxConfig].
	DeviceType string `json:"deviceType"`
	JSON       struct {
		CPU        respjson.Field
		Envs       respjson.Field
		Labels     respjson.Field
		Memory     respjson.Field
		Os         respjson.Field
		Storage    respjson.Field
		WorkingDir respjson.Field
		DeviceType respjson.Field
		raw        string
	} `json:"-"`
}

func (r *V1BoxListResponseDataUnionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxListResponseDataUnionConfigOs is an implicit subunion of
// [V1BoxListResponseDataUnion]. V1BoxListResponseDataUnionConfigOs provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxListResponseDataUnion].
type V1BoxListResponseDataUnionConfigOs struct {
	Version string `json:"version"`
	JSON    struct {
		Version respjson.Field
		raw     string
	} `json:"-"`
}

func (r *V1BoxListResponseDataUnionConfigOs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Box display
type V1BoxDisplayResponse struct {
	// Orientation of the box
	//
	// Any of "portrait", "landscape", "landscape-reverse", "portrait-reverse".
	Orientation V1BoxDisplayResponseOrientation `json:"orientation,required"`
	// Resolution configuration
	Resolution V1BoxDisplayResponseResolution `json:"resolution,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Orientation respjson.Field
		Resolution  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxDisplayResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxDisplayResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Orientation of the box
type V1BoxDisplayResponseOrientation string

const (
	V1BoxDisplayResponseOrientationPortrait         V1BoxDisplayResponseOrientation = "portrait"
	V1BoxDisplayResponseOrientationLandscape        V1BoxDisplayResponseOrientation = "landscape"
	V1BoxDisplayResponseOrientationLandscapeReverse V1BoxDisplayResponseOrientation = "landscape-reverse"
	V1BoxDisplayResponseOrientationPortraitReverse  V1BoxDisplayResponseOrientation = "portrait-reverse"
)

// Resolution configuration
type V1BoxDisplayResponseResolution struct {
	// Height of the screen
	Height float64 `json:"height,required"`
	// Width of the screen
	Width float64 `json:"width,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxDisplayResponseResolution) RawJSON() string { return r.JSON.raw }
func (r *V1BoxDisplayResponseResolution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of command execution
type V1BoxExecuteCommandsResponse struct {
	// The exit code of the command
	ExitCode float64 `json:"exitCode,required"`
	// The standard error output of the command
	Stderr string `json:"stderr,required"`
	// The standard output of the command
	Stdout string `json:"stdout,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExitCode    respjson.Field
		Stderr      respjson.Field
		Stdout      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxExecuteCommandsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxExecuteCommandsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Live view result
type V1BoxLiveViewURLResponse struct {
	// Raw live view url without additional layout content, typically used for
	// embedding into your own application
	RawURL string `json:"rawUrl,required"`
	// Live view url with Gbox interface and basic information, typically used for
	// real-time observation of box usage status
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RawURL      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxLiveViewURLResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxLiveViewURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of code execution
type V1BoxRunCodeResponse struct {
	// The exit code of the code
	ExitCode float64 `json:"exitCode,required"`
	// The stderr of the code
	Stderr string `json:"stderr,required"`
	// The stdout of the code
	Stdout string `json:"stdout,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExitCode    respjson.Field
		Stderr      respjson.Field
		Stdout      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxRunCodeResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxRunCodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxStartResponseUnion contains all possible properties and values from
// [LinuxBox], [AndroidBox].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxStartResponseUnion struct {
	ID string `json:"id"`
	// This field is a union of [LinuxBoxConfig], [AndroidBoxConfig]
	Config    V1BoxStartResponseUnionConfig `json:"config"`
	CreatedAt time.Time                     `json:"createdAt"`
	ExpiresAt time.Time                     `json:"expiresAt"`
	Status    string                        `json:"status"`
	Type      string                        `json:"type"`
	UpdatedAt time.Time                     `json:"updatedAt"`
	Reason    string                        `json:"reason"`
	JSON      struct {
		ID        respjson.Field
		Config    respjson.Field
		CreatedAt respjson.Field
		ExpiresAt respjson.Field
		Status    respjson.Field
		Type      respjson.Field
		UpdatedAt respjson.Field
		Reason    respjson.Field
		raw       string
	} `json:"-"`
}

func (u V1BoxStartResponseUnion) AsLinuxBox() (v LinuxBox) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxStartResponseUnion) AsAndroidBox() (v AndroidBox) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxStartResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxStartResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxStartResponseUnionConfig is an implicit subunion of
// [V1BoxStartResponseUnion]. V1BoxStartResponseUnionConfig provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxStartResponseUnion].
type V1BoxStartResponseUnionConfig struct {
	CPU    float64 `json:"cpu"`
	Envs   string  `json:"envs"`
	Labels string  `json:"labels"`
	Memory float64 `json:"memory"`
	// This field is a union of [LinuxBoxConfigOs], [AndroidBoxConfigOs]
	Os         V1BoxStartResponseUnionConfigOs `json:"os"`
	Storage    float64                         `json:"storage"`
	WorkingDir string                          `json:"workingDir"`
	// This field is from variant [AndroidBoxConfig].
	DeviceType string `json:"deviceType"`
	JSON       struct {
		CPU        respjson.Field
		Envs       respjson.Field
		Labels     respjson.Field
		Memory     respjson.Field
		Os         respjson.Field
		Storage    respjson.Field
		WorkingDir respjson.Field
		DeviceType respjson.Field
		raw        string
	} `json:"-"`
}

func (r *V1BoxStartResponseUnionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxStartResponseUnionConfigOs is an implicit subunion of
// [V1BoxStartResponseUnion]. V1BoxStartResponseUnionConfigOs provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxStartResponseUnion].
type V1BoxStartResponseUnionConfigOs struct {
	Version string `json:"version"`
	JSON    struct {
		Version respjson.Field
		raw     string
	} `json:"-"`
}

func (r *V1BoxStartResponseUnionConfigOs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxStopResponseUnion contains all possible properties and values from
// [LinuxBox], [AndroidBox].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxStopResponseUnion struct {
	ID string `json:"id"`
	// This field is a union of [LinuxBoxConfig], [AndroidBoxConfig]
	Config    V1BoxStopResponseUnionConfig `json:"config"`
	CreatedAt time.Time                    `json:"createdAt"`
	ExpiresAt time.Time                    `json:"expiresAt"`
	Status    string                       `json:"status"`
	Type      string                       `json:"type"`
	UpdatedAt time.Time                    `json:"updatedAt"`
	Reason    string                       `json:"reason"`
	JSON      struct {
		ID        respjson.Field
		Config    respjson.Field
		CreatedAt respjson.Field
		ExpiresAt respjson.Field
		Status    respjson.Field
		Type      respjson.Field
		UpdatedAt respjson.Field
		Reason    respjson.Field
		raw       string
	} `json:"-"`
}

func (u V1BoxStopResponseUnion) AsLinuxBox() (v LinuxBox) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxStopResponseUnion) AsAndroidBox() (v AndroidBox) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxStopResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxStopResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxStopResponseUnionConfig is an implicit subunion of
// [V1BoxStopResponseUnion]. V1BoxStopResponseUnionConfig provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxStopResponseUnion].
type V1BoxStopResponseUnionConfig struct {
	CPU    float64 `json:"cpu"`
	Envs   string  `json:"envs"`
	Labels string  `json:"labels"`
	Memory float64 `json:"memory"`
	// This field is a union of [LinuxBoxConfigOs], [AndroidBoxConfigOs]
	Os         V1BoxStopResponseUnionConfigOs `json:"os"`
	Storage    float64                        `json:"storage"`
	WorkingDir string                         `json:"workingDir"`
	// This field is from variant [AndroidBoxConfig].
	DeviceType string `json:"deviceType"`
	JSON       struct {
		CPU        respjson.Field
		Envs       respjson.Field
		Labels     respjson.Field
		Memory     respjson.Field
		Os         respjson.Field
		Storage    respjson.Field
		WorkingDir respjson.Field
		DeviceType respjson.Field
		raw        string
	} `json:"-"`
}

func (r *V1BoxStopResponseUnionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxStopResponseUnionConfigOs is an implicit subunion of
// [V1BoxStopResponseUnion]. V1BoxStopResponseUnionConfigOs provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1BoxStopResponseUnion].
type V1BoxStopResponseUnionConfigOs struct {
	Version string `json:"version"`
	JSON    struct {
		Version respjson.Field
		raw     string
	} `json:"-"`
}

func (r *V1BoxStopResponseUnionConfigOs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Web terminal result
type V1BoxWebTerminalURLResponse struct {
	// Web terminal url
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxWebTerminalURLResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxWebTerminalURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxListParams struct {
	// Filter boxes by their device type (virtual, physical)
	DeviceType param.Opt[string] `query:"deviceType,omitzero" json:"-"`
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Page size
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Filter boxes by their labels. Labels are key-value pairs that help identify and
	// categorize boxes. Use this to filter boxes that match specific label criteria.
	// For example, you can filter by project, environment, team, or any custom labels
	// you've added to your boxes.
	Labels any `query:"labels,omitzero" json:"-"`
	// Filter boxes by their current status (pending, running, stopped, error,
	// terminated, all). Must be an array of statuses. Use 'all' to get boxes with any
	// status.
	//
	// Any of "all", "pending", "running", "error", "terminated".
	Status []string `query:"status,omitzero" json:"-"`
	// Filter boxes by their type (linux, android, all). Must be an array of types. Use
	// 'all' to get boxes of any type.
	//
	// Any of "all", "linux", "android".
	Type []string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BoxListParams]'s query parameters as `url.Values`.
func (r V1BoxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1BoxNewAndroidParams struct {
	// Request body for creating a new Android box instance
	CreateAndroidBox CreateAndroidBoxParam
	paramObj
}

func (r V1BoxNewAndroidParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.CreateAndroidBox)
}
func (r *V1BoxNewAndroidParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateAndroidBox)
}

type V1BoxNewLinuxParams struct {
	// Request body for creating a new Linux box instance
	CreateLinuxBox CreateLinuxBoxParam
	paramObj
}

func (r V1BoxNewLinuxParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.CreateLinuxBox)
}
func (r *V1BoxNewLinuxParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateLinuxBox)
}

type V1BoxExecuteCommandsParams struct {
	// The command to run. Can be a single string or an array of strings
	Commands V1BoxExecuteCommandsParamsCommandsUnion `json:"commands,omitzero,required"`
	// The timeout of the command. If the command times out, the exit code will be 124.
	// For example: 'timeout 5s sleep 10s' will result in exit code 124.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30s
	Timeout param.Opt[string] `json:"timeout,omitzero"`
	// The working directory of the command. It not provided, the command will be run
	// in the `box.config.workingDir` directory.
	WorkingDir param.Opt[string] `json:"workingDir,omitzero"`
	// The environment variables to run the command
	Envs any `json:"envs,omitzero"`
	paramObj
}

func (r V1BoxExecuteCommandsParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxExecuteCommandsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxExecuteCommandsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1BoxExecuteCommandsParamsCommandsUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u V1BoxExecuteCommandsParamsCommandsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *V1BoxExecuteCommandsParamsCommandsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V1BoxExecuteCommandsParamsCommandsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type V1BoxLiveViewURLParams struct {
	// The live view will be alive for the given duration
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 180m
	ExpiresIn param.Opt[string] `json:"expiresIn,omitzero"`
	paramObj
}

func (r V1BoxLiveViewURLParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxLiveViewURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxLiveViewURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxRunCodeParams struct {
	// The code to run
	Code string `json:"code,required"`
	// The timeout of the code execution. If the code execution times out, the exit
	// code will be 124.
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 30s
	Timeout param.Opt[string] `json:"timeout,omitzero"`
	// The working directory of the code. It not provided, the code will be run in the
	// `box.config.workingDir` directory.
	WorkingDir param.Opt[string] `json:"workingDir,omitzero"`
	// The arguments to run the code. For example, if you want to run "python index.py
	// --help", you should pass ["--help"] as arguments.
	Argv []string `json:"argv,omitzero"`
	// The environment variables to run the code
	Envs any `json:"envs,omitzero"`
	// The language of the code.
	//
	// Any of "bash", "python", "typescript".
	Language V1BoxRunCodeParamsLanguage `json:"language,omitzero"`
	paramObj
}

func (r V1BoxRunCodeParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxRunCodeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxRunCodeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The language of the code.
type V1BoxRunCodeParamsLanguage string

const (
	V1BoxRunCodeParamsLanguageBash       V1BoxRunCodeParamsLanguage = "bash"
	V1BoxRunCodeParamsLanguagePython     V1BoxRunCodeParamsLanguage = "python"
	V1BoxRunCodeParamsLanguageTypescript V1BoxRunCodeParamsLanguage = "typescript"
)

type V1BoxStartParams struct {
	// Wait for the box operation to be completed, default is true
	Wait param.Opt[bool] `json:"wait,omitzero"`
	paramObj
}

func (r V1BoxStartParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxStopParams struct {
	// Wait for the box operation to be completed, default is true
	Wait param.Opt[bool] `json:"wait,omitzero"`
	paramObj
}

func (r V1BoxStopParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxStopParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxStopParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxTerminateParams struct {
	// Wait for the box operation to be completed, default is true
	Wait param.Opt[bool] `json:"wait,omitzero"`
	paramObj
}

func (r V1BoxTerminateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxTerminateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxTerminateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxWebTerminalURLParams struct {
	// The web terminal will be alive for the given duration
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 180m
	ExpiresIn param.Opt[string] `json:"expiresIn,omitzero"`
	paramObj
}

func (r V1BoxWebTerminalURLParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxWebTerminalURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxWebTerminalURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
