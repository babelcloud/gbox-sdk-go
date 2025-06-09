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

	"github.com/stainless-sdks/gbox-sdk-go/internal/apiform"
	"github.com/stainless-sdks/gbox-sdk-go/internal/apijson"
	"github.com/stainless-sdks/gbox-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/gbox-sdk-go/option"
	"github.com/stainless-sdks/gbox-sdk-go/packages/param"
	"github.com/stainless-sdks/gbox-sdk-go/packages/respjson"
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

// List Android app
func (r *V1BoxAndroidService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *V1BoxAndroidListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Android app
func (r *V1BoxAndroidService) Get(ctx context.Context, packageName string, query V1BoxAndroidGetParams, opts ...option.RequestOption) (res *AndroidApp, err error) {
	opts = append(r.Options[:], opts...)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/%s", query.ID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Install Android app
func (r *V1BoxAndroidService) Install(ctx context.Context, id string, body V1BoxAndroidInstallParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Uninstall Android app
func (r *V1BoxAndroidService) Uninstall(ctx context.Context, packageName string, params V1BoxAndroidUninstallParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if packageName == "" {
		err = errors.New("missing required packageName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/android/apps/%s", params.ID, packageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Android app information
type AndroidApp struct {
	// Android app apk path
	ApkPath string `json:"apkPath,required"`
	// Android app name
	Name string `json:"name,required"`
	// Android app package name
	PackageName string `json:"packageName,required"`
	// Android app version
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApkPath     respjson.Field
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

type V1BoxAndroidGetParams struct {
	ID string `path:"id,required" json:"-"`
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
	// APK file to install (max file size: 200MB)
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
	// HTTP URL to download APK file (max file size: 200MB)
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

type V1BoxAndroidUninstallParams struct {
	ID string `path:"id,required" json:"-"`
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
