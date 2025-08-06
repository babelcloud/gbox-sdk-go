// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/babelcloud/gbox-sdk-go/internal/apiform"
	"github.com/babelcloud/gbox-sdk-go/internal/apijson"
	"github.com/babelcloud/gbox-sdk-go/internal/requestconfig"
	"github.com/babelcloud/gbox-sdk-go/option"
	"github.com/babelcloud/gbox-sdk-go/packages/respjson"
)

// V1BoxMediaService contains methods and other services that help with interacting
// with the gbox-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1BoxMediaService] method instead.
type V1BoxMediaService struct {
	Options []option.RequestOption
}

// NewV1BoxMediaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1BoxMediaService(opts ...option.RequestOption) (r V1BoxMediaService) {
	r = V1BoxMediaService{}
	r.Options = opts
	return
}

// Create a new album with media files
func (r *V1BoxMediaService) NewAlbum(ctx context.Context, boxID string, body V1BoxMediaNewAlbumParams, opts ...option.RequestOption) (res *V1BoxMediaNewAlbumResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/media/albums", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete an album and all its media files
func (r *V1BoxMediaService) DeleteAlbum(ctx context.Context, albumName string, body V1BoxMediaDeleteAlbumParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if albumName == "" {
		err = errors.New("missing required albumName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/media/albums/%s", body.BoxID, albumName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete a specific media file from an album
func (r *V1BoxMediaService) DeleteMedia(ctx context.Context, mediaName string, body V1BoxMediaDeleteMediaParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if body.AlbumName == "" {
		err = errors.New("missing required albumName parameter")
		return
	}
	if mediaName == "" {
		err = errors.New("missing required mediaName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/media/albums/%s/media/%s", body.BoxID, body.AlbumName, mediaName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Download a specific media file from an album
func (r *V1BoxMediaService) DownloadMedia(ctx context.Context, mediaName string, query V1BoxMediaDownloadMediaParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if query.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if query.AlbumName == "" {
		err = errors.New("missing required albumName parameter")
		return
	}
	if mediaName == "" {
		err = errors.New("missing required mediaName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/media/albums/%s/media/%s", query.BoxID, query.AlbumName, mediaName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get detailed information about a specific album including its media files
func (r *V1BoxMediaService) GetAlbumDetail(ctx context.Context, albumName string, query V1BoxMediaGetAlbumDetailParams, opts ...option.RequestOption) (res *V1BoxMediaGetAlbumDetailResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if albumName == "" {
		err = errors.New("missing required albumName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/media/albums/%s", query.BoxID, albumName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of albums in the box
func (r *V1BoxMediaService) ListAlbums(ctx context.Context, boxID string, opts ...option.RequestOption) (res *V1BoxMediaListAlbumsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if boxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/media/albums", boxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of media files in a specific album
func (r *V1BoxMediaService) ListMedia(ctx context.Context, albumName string, query V1BoxMediaListMediaParams, opts ...option.RequestOption) (res *V1BoxMediaListMediaResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if albumName == "" {
		err = errors.New("missing required albumName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/media/albums/%s/media", query.BoxID, albumName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Add media files to an existing album
func (r *V1BoxMediaService) UpdateAlbum(ctx context.Context, albumName string, params V1BoxMediaUpdateAlbumParams, opts ...option.RequestOption) (res *V1BoxMediaUpdateAlbumResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.BoxID == "" {
		err = errors.New("missing required boxId parameter")
		return
	}
	if albumName == "" {
		err = errors.New("missing required albumName parameter")
		return
	}
	path := fmt.Sprintf("boxes/%s/media/albums/%s", params.BoxID, albumName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Album representation
type V1BoxMediaNewAlbumResponse struct {
	// Last modified time of the album
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// Name of the album
	Name string `json:"name,required"`
	// Full path to the album in the box
	Path string `json:"path,required"`
	// Size of the album
	Size string `json:"size,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxMediaNewAlbumResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxMediaNewAlbumResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Album representation
type V1BoxMediaGetAlbumDetailResponse struct {
	// Last modified time of the album
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// Name of the album
	Name string `json:"name,required"`
	// Full path to the album in the box
	Path string `json:"path,required"`
	// Size of the album
	Size string `json:"size,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxMediaGetAlbumDetailResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxMediaGetAlbumDetailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List albums
type V1BoxMediaListAlbumsResponse struct {
	// List of albums
	Data []V1BoxMediaListAlbumsResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxMediaListAlbumsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxMediaListAlbumsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Album representation
type V1BoxMediaListAlbumsResponseData struct {
	// Last modified time of the album
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// Name of the album
	Name string `json:"name,required"`
	// Full path to the album in the box
	Path string `json:"path,required"`
	// Size of the album
	Size string `json:"size,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxMediaListAlbumsResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1BoxMediaListAlbumsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List album media
type V1BoxMediaListMediaResponse struct {
	// List of media files (photos and videos) in the album
	Data []V1BoxMediaListMediaResponseDataUnion `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxMediaListMediaResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxMediaListMediaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1BoxMediaListMediaResponseDataUnion contains all possible properties and values
// from [V1BoxMediaListMediaResponseDataPhoto],
// [V1BoxMediaListMediaResponseDataVideo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1BoxMediaListMediaResponseDataUnion struct {
	LastModified time.Time `json:"lastModified"`
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	Size         string    `json:"size"`
	Type         string    `json:"type"`
	JSON         struct {
		LastModified respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		Type         respjson.Field
		raw          string
	} `json:"-"`
}

func (u V1BoxMediaListMediaResponseDataUnion) AsPhoto() (v V1BoxMediaListMediaResponseDataPhoto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1BoxMediaListMediaResponseDataUnion) AsVideo() (v V1BoxMediaListMediaResponseDataVideo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1BoxMediaListMediaResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *V1BoxMediaListMediaResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Photo representation
type V1BoxMediaListMediaResponseDataPhoto struct {
	// Last modified time of the photo
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// Name of the photo
	Name string `json:"name,required"`
	// Full path to the photo in the box
	Path string `json:"path,required"`
	// Size of the photo
	Size string `json:"size,required"`
	// Photo type indicator
	//
	// Any of "photo".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxMediaListMediaResponseDataPhoto) RawJSON() string { return r.JSON.raw }
func (r *V1BoxMediaListMediaResponseDataPhoto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Video representation
type V1BoxMediaListMediaResponseDataVideo struct {
	// Last modified time of the video
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// Name of the video
	Name string `json:"name,required"`
	// Full path to the video in the box
	Path string `json:"path,required"`
	// Size of the video
	Size string `json:"size,required"`
	// Video type indicator
	//
	// Any of "video".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxMediaListMediaResponseDataVideo) RawJSON() string { return r.JSON.raw }
func (r *V1BoxMediaListMediaResponseDataVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Album representation
type V1BoxMediaUpdateAlbumResponse struct {
	// Last modified time of the album
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// Name of the album
	Name string `json:"name,required"`
	// Full path to the album in the box
	Path string `json:"path,required"`
	// Size of the album
	Size string `json:"size,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1BoxMediaUpdateAlbumResponse) RawJSON() string { return r.JSON.raw }
func (r *V1BoxMediaUpdateAlbumResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1BoxMediaNewAlbumParams struct {
	// Media files to include in the album (max size: 512MB per file)
	Media []io.Reader `json:"media,omitzero,required" format:"binary"`
	// Name of the album to create
	Name string `json:"name,required"`
	paramObj
}

func (r V1BoxMediaNewAlbumParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type V1BoxMediaDeleteAlbumParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	paramObj
}

type V1BoxMediaDeleteMediaParams struct {
	BoxID     string `path:"boxId,required" json:"-"`
	AlbumName string `path:"albumName,required" json:"-"`
	paramObj
}

type V1BoxMediaDownloadMediaParams struct {
	BoxID     string `path:"boxId,required" json:"-"`
	AlbumName string `path:"albumName,required" json:"-"`
	paramObj
}

type V1BoxMediaGetAlbumDetailParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	paramObj
}

type V1BoxMediaListMediaParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	paramObj
}

type V1BoxMediaUpdateAlbumParams struct {
	BoxID string `path:"boxId,required" json:"-"`
	// Media files to add to the album (max size: 512MB per file)
	Media []io.Reader `json:"media,omitzero,required" format:"binary"`
	paramObj
}

func (r V1BoxMediaUpdateAlbumParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
