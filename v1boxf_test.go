// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gbox-sdk-go"
	"github.com/stainless-sdks/gbox-sdk-go/internal/testutil"
	"github.com/stainless-sdks/gbox-sdk-go/option"
)

func TestV1BoxFListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gboxsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Boxes.Fs.List(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxFListParams{
			Path:  "/home/user/documents",
			Depth: gboxsdk.Float(2),
		},
	)
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BoxFRead(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gboxsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Boxes.Fs.Read(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxFReadParams{
			Path: "/home/user/documents/config.json",
		},
	)
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BoxFRemove(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gboxsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Boxes.Fs.Remove(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxFRemoveParams{
			Path: "/home/user/documents/output.txt",
		},
	)
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BoxFRename(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gboxsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Boxes.Fs.Rename(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxFRenameParams{
			NewPath: "/home/user/documents/new-name.txt",
			OldPath: "/home/user/documents/output.txt",
		},
	)
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BoxFWrite(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gboxsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Boxes.Fs.Write(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxFWriteParams{
			Content: "Hello, World!\nThis is file content.",
			Path:    "/home/user/documents/output.txt",
		},
	)
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
