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

// Backup
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
	path := fmt.Sprintf("boxes/%s/android/packages/%s/backup", body.BoxID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Backup all
func (r *V1BoxAndroidService) BackupAll(ctx context.Context, boxID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/packages/backup-all", boxID)
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
	path := fmt.Sprintf("boxes/%s/android/packages/%s/close", body.BoxID, packageName)
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
	path := fmt.Sprintf("boxes/%s/android/packages/close-all", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Get app
func (r *V1BoxAndroidService) Get(ctx context.Context, packageName string, query V1BoxAndroidGetParams, opts ...option.RequestOption) (res *V1BoxAndroidGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/packages/%s", query.BoxID, packageName)
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
func (r *V1BoxAndroidService) Install(ctx context.Context, boxID string, body V1BoxAndroidInstallParams, opts ...option.RequestOption) (res *V1BoxAndroidInstallResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/packages", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get pkg activities
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
	path := fmt.Sprintf("boxes/%s/android/packages/%s/activities", query.BoxID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List apps
func (r *V1BoxAndroidService) ListApp(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxAndroidListAppResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve detailed information for all installed pkgs. This endpoint provides
// comprehensive pkg details
func (r *V1BoxAndroidService) ListPkg(ctx context.Context, boxID string, query V1BoxAndroidListPkgParams, opts ...option.RequestOption) (res *V1BoxAndroidListPkgResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/packages", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A faster endpoint to quickly retrieve basic pkg information. This API provides
// better performance for scenarios where you need to get essential pkg details
// quickly
func (r *V1BoxAndroidService) ListPkgSimple(ctx context.Context, boxID string, query V1BoxAndroidListPkgSimpleParams, opts ...option.RequestOption) (res *V1BoxAndroidListPkgSimpleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/packages/simple", boxID)
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
	path := fmt.Sprintf("boxes/%s/android/packages/%s/open", params.BoxID, packageName)
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
	path := fmt.Sprintf("boxes/%s/android/packages/%s/restart", params.BoxID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Restore
func (r *V1BoxAndroidService) Restore(ctx context.Context, boxID string, body V1BoxAndroidRestoreParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/packages/restore", boxID)
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
	path := fmt.Sprintf("boxes/%s/android/packages/%s", params.BoxID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Android pkg information
type V1BoxAndroidGetResponse struct {
	// Android apk path
	ApkPath string `json:"apkPath,required"`
	// Whether the pkg is currently running
	IsRunning bool `json:"isRunning,required"`
	// Android pkg name
	Name string `json:"name,required"`
	// Android package name
	PackageName string `json:"packageName,required"`
	// Package type: system or thirdParty
	//
	// Any of "system", "thirdParty".
	PkgType V1BoxAndroidGetResponsePkgType `json:"pkgType,required"`
	// Android pkg version
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApkPath     respjson.Field
		IsRunning   respjson.Field
		Name        respjson.Field
		PackageName respjson.Field
		PkgType     respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Package type: system or thirdParty
type V1BoxAndroidGetResponsePkgType string

const (
	V1BoxAndroidGetResponsePkgTypeSystem     V1BoxAndroidGetResponsePkgType = "system"
	V1BoxAndroidGetResponsePkgTypeThirdParty V1BoxAndroidGetResponsePkgType = "thirdParty"
)

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

// Response containing the result of installing an Android pkg
type V1BoxAndroidInstallResponse struct {
	// Activity list
	Activities []V1BoxAndroidInstallResponseActivity `json:"activities,required"`
	// Android apk path
	ApkPath string `json:"apkPath,required"`
	// Android pkg package name
	PackageName string `json:"packageName,required"`
	// Package type: system or thirdParty
	//
	// Any of "system", "thirdParty".
	PkgType V1BoxAndroidInstallResponsePkgType `json:"pkgType,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Activities  respjson.Field
		ApkPath     respjson.Field
		PackageName respjson.Field
		PkgType     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidInstallResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidInstallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Android pkg activity
type V1BoxAndroidInstallResponseActivity struct {
	// Activity class name
	ClassName string `json:"className,required"`
	// Activity class name
	IsExported bool `json:"isExported,required"`
	// Whether the activity is a launcher activity. Launcher activities appear in the
	// device's pkg launcher/home screen and can be directly launched by the user.
	IsLauncher bool `json:"isLauncher,required"`
	// Whether the activity is the main activity. Main activity is the entry point of
	// the pkg and is typically launched when the pkg is started.
	IsMain bool `json:"isMain,required"`
	// Activity name
	Name string `json:"name,required"`
	// Activity package name
	PackageName string `json:"packageName,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassName   respjson.Field
		IsExported  respjson.Field
		IsLauncher  respjson.Field
		IsMain      respjson.Field
		Name        respjson.Field
		PackageName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidInstallResponseActivity) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidInstallResponseActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Package type: system or thirdParty
type V1BoxAndroidInstallResponsePkgType string

const (
	V1BoxAndroidInstallResponsePkgTypeSystem     V1BoxAndroidInstallResponsePkgType = "system"
	V1BoxAndroidInstallResponsePkgTypeThirdParty V1BoxAndroidInstallResponsePkgType = "thirdParty"
)

// Android pkg activity list
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

// Android pkg activity
type V1BoxAndroidListActivitiesResponseData struct {
	// Activity class name
	ClassName string `json:"className,required"`
	// Activity class name
	IsExported bool `json:"isExported,required"`
	// Whether the activity is a launcher activity. Launcher activities appear in the
	// device's pkg launcher/home screen and can be directly launched by the user.
	IsLauncher bool `json:"isLauncher,required"`
	// Whether the activity is the main activity. Main activity is the entry point of
	// the pkg and is typically launched when the pkg is started.
	IsMain bool `json:"isMain,required"`
	// Activity name
	Name string `json:"name,required"`
	// Activity package name
	PackageName string `json:"packageName,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassName   respjson.Field
		IsExported  respjson.Field
		IsLauncher  respjson.Field
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

// Android app list
type V1BoxAndroidListAppResponse struct {
	// App list
	Data []V1BoxAndroidListAppResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidListAppResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidListAppResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Android pkg information
type V1BoxAndroidListAppResponseData struct {
	// Android apk path
	ApkPath string `json:"apkPath,required"`
	// Whether the pkg is currently running
	IsRunning bool `json:"isRunning,required"`
	// Android pkg name
	Name string `json:"name,required"`
	// Android package name
	PackageName string `json:"packageName,required"`
	// Package type: system or thirdParty
	//
	// Any of "system", "thirdParty".
	PkgType string `json:"pkgType,required"`
	// Android pkg version
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApkPath     respjson.Field
		IsRunning   respjson.Field
		Name        respjson.Field
		PackageName respjson.Field
		PkgType     respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidListAppResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidListAppResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing list of Android pkgs
type V1BoxAndroidListPkgResponse struct {
	// Android pkg list
	Data []V1BoxAndroidListPkgResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidListPkgResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidListPkgResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Android pkg information
type V1BoxAndroidListPkgResponseData struct {
	// Android apk path
	ApkPath string `json:"apkPath,required"`
	// Whether the pkg is currently running
	IsRunning bool `json:"isRunning,required"`
	// Android pkg name
	Name string `json:"name,required"`
	// Android package name
	PackageName string `json:"packageName,required"`
	// Package type: system or thirdParty
	//
	// Any of "system", "thirdParty".
	PkgType string `json:"pkgType,required"`
	// Android pkg version
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApkPath     respjson.Field
		IsRunning   respjson.Field
		Name        respjson.Field
		PackageName respjson.Field
		PkgType     respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidListPkgResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidListPkgResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing list of Android pkgs
type V1BoxAndroidListPkgSimpleResponse struct {
	// Android pkg simple list
	Data []V1BoxAndroidListPkgSimpleResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidListPkgSimpleResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidListPkgSimpleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Android pkg simple
type V1BoxAndroidListPkgSimpleResponseData struct {
	// Android apk path
	ApkPath string `json:"apkPath,required"`
	// Android pkg package name
	PackageName string `json:"packageName,required"`
	// Package type: system or thirdParty
	//
	// Any of "system", "thirdParty".
	PkgType string `json:"pkgType,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApkPath     respjson.Field
		PackageName respjson.Field
		PkgType     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxAndroidListPkgSimpleResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BoxAndroidListPkgSimpleResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// for installing Android pkg from uploaded APK file
	OfInstallAndroidPkgByFile *V1BoxAndroidInstallParamsBodyInstallAndroidPkgByFile `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Request
	// for installing Android pkg from HTTP URL
	OfInstallAndroidPkgByURL *V1BoxAndroidInstallParamsBodyInstallAndroidPkgByURL `json:",inline"`

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

// Request for installing Android pkg from uploaded APK file
//
// The property Apk is required.
type V1BoxAndroidInstallParamsBodyInstallAndroidPkgByFile struct {
	// APK file to install (max file size: 512MB)
	Apk io.Reader `json:"apk,omitzero,required" format:"binary"`
	paramObj
}

func (r V1BoxAndroidInstallParamsBodyInstallAndroidPkgByFile) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxAndroidInstallParamsBodyInstallAndroidPkgByFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxAndroidInstallParamsBodyInstallAndroidPkgByFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request for installing Android pkg from HTTP URL
//
// The property Apk is required.
type V1BoxAndroidInstallParamsBodyInstallAndroidPkgByURL struct {
	// HTTP URL to download APK file (max file size: 512MB)
	Apk string `json:"apk,required"`
	paramObj
}

func (r V1BoxAndroidInstallParamsBodyInstallAndroidPkgByURL) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxAndroidInstallParamsBodyInstallAndroidPkgByURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxAndroidInstallParamsBodyInstallAndroidPkgByURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxAndroidListActivitiesParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	paramObj
}

type V1BoxAndroidListPkgParams struct {
	// Package type: system or thirdParty, default is thirdParty
	//
	// Any of "system", "thirdParty".
	PkgType []string `query:"pkgType,omitzero" json:"-"`
	// Filter pkgs by running status: running (show only running pkgs), notRunning
	// (show only non-running pkgs). Default is all
	//
	// Any of "running", "notRunning".
	RunningFilter []string `query:"runningFilter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BoxAndroidListPkgParams]'s query parameters as
// `url.Values`.
func (r V1BoxAndroidListPkgParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1BoxAndroidListPkgSimpleParams struct {
	// Package type: system or thirdParty, default is thirdParty
	//
	// Any of "system", "thirdParty".
	PkgType []string `query:"pkgType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1BoxAndroidListPkgSimpleParams]'s query parameters as
// `url.Values`.
func (r V1BoxAndroidListPkgSimpleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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
	// uninstalls the pkg while retaining the data/cache
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
