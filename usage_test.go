// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk_test

import (
	"context"
	"os"
	"testing"

	gboxsdk "github.com/babelcloud/gbox-sdk-go"
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
	box, err := client.V1.Boxes.NewLinux(context.TODO(), gboxsdk.V1BoxNewLinuxParams{
		CreateLinuxBox: gboxsdk.CreateLinuxBoxParam{
			Type: gboxsdk.CreateLinuxBoxTypeLinux,
			Config: gboxsdk.CreateBoxConfigParam{
				Envs:      map[string]interface{}{"FOO": "bar"},
				Labels:    map[string]interface{}{"env": "test"},
				ExpiresIn: gboxsdk.String("10m"),
			},
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", box)
}
