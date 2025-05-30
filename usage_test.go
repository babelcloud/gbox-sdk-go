// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/babelcloud/gbox-sdk-go"
	"github.com/babelcloud/gbox-sdk-go/internal/testutil"
	"github.com/babelcloud/gbox-sdk-go/option"
)

func TestUsage(t *testing.T) {
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
	box, err := client.V1.Boxes.New(context.TODO(), gboxsdk.V1BoxNewParamsCreateLinuxBox{
		CreateLinuxBox: gboxsdk.CreateLinuxBoxParam{
			Type: gboxsdk.CreateLinuxBoxTypeLinux,
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", box)
}
