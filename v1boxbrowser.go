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

// V1BoxBrowserService contains methods and other services that help with
// interacting with the gbox-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BoxBrowserService] method instead.
type V1BoxBrowserService struct {
	Options []option.RequestOption
}

// NewV1BoxBrowserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1BoxBrowserService(opts ...option.RequestOption) (r V1BoxBrowserService) {
	r = V1BoxBrowserService{}
	r.Options = opts
	return
}

// This endpoint allows you to generate a pre-signed URL for accessing the Chrome
// DevTools Protocol (CDP) of a running box. The URL is valid for a limited time
// and can be used to interact with the box's browser environment
func (r *V1BoxBrowserService) CdpURL(ctx context.Context, boxID string, body V1BoxBrowserCdpURLParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/browser/connect-url/cdp", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Close a specific browser tab identified by its index. This endpoint will
// permanently close the tab and free up the associated resources. The tab index
// corresponds to the index returned when listing tabs or opening new tabs. After
// closing a tab, the indices of subsequent tabs may shift down to fill the gap.
// It's important to refresh the tab list after closing tabs to get the current
// indices. You cannot close the last remaining tab - at least one tab must remain
// open in the browser context.
func (r *V1BoxBrowserService) CloseTab(ctx context.Context, tabIndex string, body V1BoxBrowserCloseTabParams, opts ...option.RequestOption) (res *V1BoxBrowserCloseTabResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if tabIndex == "" {
		err = errors.New("missing required tabIndex parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/browser/tabs/%s", body.BoxID, tabIndex)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve a comprehensive list of all currently open browser tabs in the
// specified box. This endpoint returns detailed information about each tab
// including its index, title, current URL, and favicon. The tab index can be used
// for subsequent operations like navigation, closing, or updating tabs. This is
// essential for managing multiple browser sessions and understanding the current
// state of the browser environment.
func (r *V1BoxBrowserService) GetTabs(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxBrowserGetTabsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/browser/tabs", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create and open a new browser tab with the specified URL. This endpoint will
// navigate to the provided URL and return the new tab's information including its
// assigned index, loaded title, final URL (after any redirects), and favicon. The
// returned tab index can be used for future operations on this specific tab. The
// browser will attempt to load the page and will wait for the DOM content to be
// loaded before returning the response. If the URL is invalid or unreachable, an
// error will be returned.
func (r *V1BoxBrowserService) OpenTab(ctx context.Context, boxID string, body V1BoxBrowserOpenTabParams, opts ...option.RequestOption) (res *V1BoxBrowserOpenTabResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/browser/tabs", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Navigate an existing browser tab to a new URL. This endpoint updates the
// specified tab by navigating it to the provided URL and returns the updated tab
// information. The browser will wait for the DOM content to be loaded before
// returning the response. This operation preserves the tab's position and index
// while updating its content. If the navigation fails due to an invalid URL or
// network issues, an error will be returned. The updated tab information will
// include the new title, final URL (after any redirects), and favicon from the new
// page.
func (r *V1BoxBrowserService) UpdateTab(ctx context.Context, tabIndex string, params V1BoxBrowserUpdateTabParams, opts ...option.RequestOption) (res *V1BoxBrowserUpdateTabResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if tabIndex == "" {
		err = errors.New("missing required tabIndex parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/browser/tabs/%s", params.BoxID, tabIndex)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type V1BoxBrowserCloseTabResponse struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxBrowserCloseTabResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxBrowserCloseTabResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List tabs
type V1BoxBrowserGetTabsResponse struct {
	// The tabs
	Tabs []V1BoxBrowserGetTabsResponseTab `json:"tabs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tabs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxBrowserGetTabsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxBrowserGetTabsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser tab
type V1BoxBrowserGetTabsResponseTab struct {
	// The tab favicon
	Favicon string `json:"favicon,required"`
	// The tab index, starting from 0
	Index float64 `json:"index,required"`
	// The tab title
	Title string `json:"title,required"`
	// The tab url
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Favicon     respjson.Field
		Index       respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxBrowserGetTabsResponseTab) RawJSON() string { return r.JSON.raw }
func (r *V1BoxBrowserGetTabsResponseTab) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser tab
type V1BoxBrowserOpenTabResponse struct {
	// The tab favicon
	Favicon string `json:"favicon,required"`
	// The tab index, starting from 0
	Index float64 `json:"index,required"`
	// The tab title
	Title string `json:"title,required"`
	// The tab url
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Favicon     respjson.Field
		Index       respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxBrowserOpenTabResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxBrowserOpenTabResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser tab
type V1BoxBrowserUpdateTabResponse struct {
	// The tab favicon
	Favicon string `json:"favicon,required"`
	// The tab index, starting from 0
	Index float64 `json:"index,required"`
	// The tab title
	Title string `json:"title,required"`
	// The tab url
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Favicon     respjson.Field
		Index       respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxBrowserUpdateTabResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxBrowserUpdateTabResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxBrowserCdpURLParams struct {
	// The CDP url will be alive for the given duration
	//
	// Supported time units: ms (milliseconds), s (seconds), m (minutes), h (hours)
	// Example formats: "500ms", "30s", "5m", "1h" Default: 120m
	ExpiresIn param.Opt[string] `json:"expiresIn,omitzero"`
	paramObj
}

func (r V1BoxBrowserCdpURLParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxBrowserCdpURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxBrowserCdpURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxBrowserCloseTabParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	paramObj
}

type V1BoxBrowserOpenTabParams struct {
	// The tab url
	URL string `json:"url,required"`
	paramObj
}

func (r V1BoxBrowserOpenTabParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxBrowserOpenTabParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxBrowserOpenTabParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxBrowserUpdateTabParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	// The tab new url
	URL string `json:"url,required"`
	paramObj
}

func (r V1BoxBrowserUpdateTabParams) MarshalJSON() (data []byte, err error) {
	type shadow V1BoxBrowserUpdateTabParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1BoxBrowserUpdateTabParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
