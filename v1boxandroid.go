// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/babelcloud/gbox-sdk-go/internal/apiform"
	"github.com/babelcloud/gbox-sdk-go/internal/apijson"
	"github.com/babelcloud/gbox-sdk-go/internal/apiquery"
	"github.com/babelcloud/gbox-sdk-go/internal/requestconfig"
	"github.com/babelcloud/gbox-sdk-go/option"
	"github.com/babelcloud/gbox-sdk-go/packages/param"
	"github.com/babelcloud/gbox-sdk-go/packages/respjson"
)

// V1BoxAndroidService contains methods and other services that help with
// interacting with the gbox-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BoxAndroidService] method instead.
type V1BoxAndroidService struct {
	Options []option.RequestOption
}

// NewV1BoxAndroidService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1BoxAndroidService(opts ...option.RequestOption) (r V1BoxAndroidService) {
	r = V1BoxAndroidService{}
	r.Options = opts
	return
}

// Retrieve detailed information for all installed applications. This endpoint
// provides comprehensive app details
func (r *V1BoxAndroidService) List(ctx context.Context, boxID string, query V1BoxAndroidListParams, opts ...option.RequestOption) (res *V1BoxAndroidListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Backup app
func (r *V1BoxAndroidService) Backup(ctx context.Context, packageName string, body V1BoxAndroidBackupParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if body.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/%s/backup", body.BoxID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Backup all apps
func (r *V1BoxAndroidService) BackupAll(ctx context.Context, boxID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/backup-all", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Close app
func (r *V1BoxAndroidService) Close(ctx context.Context, packageName string, body V1BoxAndroidCloseParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/%s/close", body.BoxID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Close all apps
func (r *V1BoxAndroidService) CloseAll(ctx context.Context, boxID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/close-all", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Get app
func (r *V1BoxAndroidService) Get(ctx context.Context, packageName string, query V1BoxAndroidGetParams, opts ...option.RequestOption) (res *AndroidApp, err error) {
	opts = append(r.Options[:], opts...)
	if query.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/%s", query.BoxID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get connect address
func (r *V1BoxAndroidService) GetConnectAddress(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxAndroidGetConnectAddressResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/connect-address", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Install app
func (r *V1BoxAndroidService) Install(ctx context.Context, boxID string, body V1BoxAndroidInstallParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get app activities
func (r *V1BoxAndroidService) ListActivities(ctx context.Context, packageName string, query V1BoxAndroidListActivitiesParams, opts ...option.RequestOption) (res *V1BoxAndroidListActivitiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/%s/activities", query.BoxID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A faster endpoint to quickly retrieve basic app information. This API provides
// better performance for scenarios where you need to get essential app details
// quickly
func (r *V1BoxAndroidService) ListSimple(ctx context.Context, boxID string, query V1BoxAndroidListSimpleParams, opts ...option.RequestOption) (res *V1BoxAndroidListSimpleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/simple", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Open app
func (r *V1BoxAndroidService) Open(ctx context.Context, packageName string, params V1BoxAndroidOpenParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/%s/open", params.BoxID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Restart app
func (r *V1BoxAndroidService) Restart(ctx context.Context, packageName string, params V1BoxAndroidRestartParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/%s/restart", params.BoxID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Restore app
func (r *V1BoxAndroidService) Restore(ctx context.Context, boxID string, body V1BoxAndroidRestoreParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/restore", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Uninstall app
func (r *V1BoxAndroidService) Uninstall(ctx context.Context, packageName string, params V1BoxAndroidUninstallParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/%s", params.BoxID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Android app information
type AndroidApp struct {
	// Android app apk path
	ApkPath string `json:"apkPath,required"`
	// Application type: system or third-party
	//
	// Any of "system", "third-party".
	AppType AndroidAppAppType `json:"appType,required"`
	// Whether the application is currently running
	IsRunning bool `json:"isRunning,required"`
	// Android app name
	Name string `json:"name,required"`
	// Android app package name
	PackageName string `json:"packageName,required"`
	// Android app version
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApkPath     respjson.Field
		AppType     respjson.Field
		IsRunning   respjson.Field
		Name        respjson.Field
		PackageName respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AndroidApp) RawJSON() string { return r.JSON.raw }
func (r *AndroidApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Application type: system or third-party
type AndroidAppAppType string

const (
	AndroidAppAppTypeSystem     AndroidAppAppType = "system"
	AndroidAppAppTypeThirdParty AndroidAppAppType = "third-party"
)

// Response containing list of Android apps
type V1BoxAndroidListResponse struct {
	// Android app list
	Data []AndroidApp `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Android connection information
type V1BoxAndroidGetConnectAddressResponse struct {
	// Android adb connect address. use `adb connect <adbConnectAddress>` to connect to
	// the Android device
	Adb string `json:"adb,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Adb         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidGetConnectAddressResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidGetConnectAddressResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxAndroidListActivitiesResponse struct {
	// Activity list
	Data []V1BoxAndroidListActivitiesResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidListActivitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidListActivitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Android app activity
type V1BoxAndroidListActivitiesResponseData struct {
	// Activity class name
	ClassName string `json:"className,required"`
	// Activity class name
	IsExported bool `json:"isExported,required"`
	// Whether the activity is the main activity
	IsMain bool `json:"isMain,required"`
	// Activity name
	Name string `json:"name,required"`
	// Activity package name
	PackageName string `json:"packageName,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassName   respjson.Field
		IsExported  respjson.Field
		IsMain      respjson.Field
		Name        respjson.Field
		PackageName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidListActivitiesResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidListActivitiesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing list of Android apps
type V1BoxAndroidListSimpleResponse struct {
	// Android app simple list
	Data []V1BoxAndroidListSimpleResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidListSimpleResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidListSimpleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxAndroidListSimpleResponseData struct {
	// Android app apk path
	ApkPath string `json:"apkPath,required"`
	// Application type: system or third-party
	//
	// Any of "system", "third-party".
	AppType string `json:"appType,required"`
	// Android app package name
	PackageName string `json:"packageName,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApkPath     respjson.Field
		AppType     respjson.Field
		PackageName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidListSimpleResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidListSimpleResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxAndroidListParams struct {
	// Whether to include running apps, default is all
	IsRunning param.Opt[bool] `query:"isRunning,omitzero" json:"-"`
	// Application type: system or third-party, default is third-party
	//
	// Any of "system", "third-party".
	AppType V1BoxAndroidListParamsAppType `query:"appType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BoxAndroidListParams]'s query parameters as `url.Values`.
func (r V1BoxAndroidListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Application type: system or third-party, default is third-party
type V1BoxAndroidListParamsAppType string

const (
	V1BoxAndroidListParamsAppTypeSystem     V1BoxAndroidListParamsAppType = "system"
	V1BoxAndroidListParamsAppTypeThirdParty V1BoxAndroidListParamsAppType = "third-party"
)

type V1BoxAndroidBackupParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	paramObj
}

type V1BoxAndroidCloseParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	paramObj
}

type V1BoxAndroidGetParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	paramObj
}

type V1BoxAndroidInstallParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Request
	// for installing Android app from uploaded APK file
	OfInstallAndroidAppByFile *V1BoxAndroidInstallParamsBodyInstallAndroidAppByFile `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Request
	// for installing Android app from HTTP URL
	OfInstallAndroidAppByURL *V1BoxAndroidInstallParamsBodyInstallAndroidAppByURL `json:",inline"`

	paramObj
}

func (r V1BoxAndroidInstallParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// Request for installing Android app from uploaded APK file
//
// The property Apk is required.
type V1BoxAndroidInstallParamsBodyInstallAndroidAppByFile struct {
	// APK file to install (max file size: 512MB)
	Apk io.Reader `json:"apk,omitzero,required" format:"binary"`
	paramObj
}

func (r V1BoxAndroidInstallParamsBodyInstallAndroidAppByFile) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxAndroidInstallParamsBodyInstallAndroidAppByFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxAndroidInstallParamsBodyInstallAndroidAppByFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request for installing Android app from HTTP URL
//
// The property Apk is required.
type V1BoxAndroidInstallParamsBodyInstallAndroidAppByURL struct {
	// HTTP URL to download APK file (max file size: 512MB)
	Apk string `json:"apk,required"`
	paramObj
}

func (r V1BoxAndroidInstallParamsBodyInstallAndroidAppByURL) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxAndroidInstallParamsBodyInstallAndroidAppByURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxAndroidInstallParamsBodyInstallAndroidAppByURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxAndroidListActivitiesParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	paramObj
}

type V1BoxAndroidListSimpleParams struct {
	// Application type: system or third-party, default is third-party
	//
	// Any of "system", "third-party".
	AppType V1BoxAndroidListSimpleParamsAppType `query:"appType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BoxAndroidListSimpleParams]'s query parameters as
// `url.Values`.
func (r V1BoxAndroidListSimpleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Application type: system or third-party, default is third-party
type V1BoxAndroidListSimpleParamsAppType string

const (
	V1BoxAndroidListSimpleParamsAppTypeSystem     V1BoxAndroidListSimpleParamsAppType = "system"
	V1BoxAndroidListSimpleParamsAppTypeThirdParty V1BoxAndroidListSimpleParamsAppType = "third-party"
)

type V1BoxAndroidOpenParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	// Activity name, default is the main activity.
	ActivityName param.Opt[string] `json:"activityName,omitzero"`
	paramObj
}

func (r V1BoxAndroidOpenParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxAndroidOpenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxAndroidOpenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxAndroidRestartParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	// Activity name, default is the main activity.
	ActivityName param.Opt[string] `json:"activityName,omitzero"`
	paramObj
}

func (r V1BoxAndroidRestartParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxAndroidRestartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxAndroidRestartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxAndroidRestoreParams struct {
	// Backup file to restore (max file size: 100MB)
	Backup io.Reader `json:"backup,omitzero,required" format:"binary"`
	paramObj
}

func (r V1BoxAndroidRestoreParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type V1BoxAndroidUninstallParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	// uninstalls the application while retaining the data/cache
	KeepData param.Opt[bool] `json:"keepData,omitzero"`
	paramObj
}

func (r V1BoxAndroidUninstallParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxAndroidUninstallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxAndroidUninstallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
