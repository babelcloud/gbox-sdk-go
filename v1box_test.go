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

func TestV1BoxGet(t *testing.T) {
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
	_, err := client.V1.Boxes.Get(context.TODO(), "c9bdc193-b54b-4ddb-a035-5ac0c598d32d")
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BoxListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.List(context.TODO(), gboxsdk.V1BoxListParams{
		Page:     gboxsdk.Int(1),
		PageSize: gboxsdk.Int(10),
		Status:   gboxsdk.String("running"),
		Type:     gboxsdk.String("linux"),
	})
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BoxDeleteWithOptionalParams(t *testing.T) {
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
	err := client.V1.Boxes.Delete(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxDeleteParams{
			Timeout: gboxsdk.String("30s"),
			Wait:    gboxsdk.Bool(true),
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

func TestV1BoxNewAndroidWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.NewAndroid(context.TODO(), gboxsdk.V1BoxNewAndroidParams{
		CreateAndroidBox: gboxsdk.CreateAndroidBoxParam{
			Config: gboxsdk.CreateBoxConfigParam{
				Envs: map[string]interface{}{
					"DEBUG":   "true",
					"API_URL": "https://api.example.com",
				},
				ExpiresIn: gboxsdk.String("10m"),
				Labels: map[string]interface{}{
					"project":     "web-automation",
					"environment": "testing",
				},
			},
			Timeout: gboxsdk.String("30s"),
			Wait:    gboxsdk.Bool(true),
		},
	})
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BoxNewLinuxWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.NewLinux(context.TODO(), gboxsdk.V1BoxNewLinuxParams{
		CreateLinuxBox: gboxsdk.CreateLinuxBoxParam{
			Config: gboxsdk.CreateBoxConfigParam{
				Envs: map[string]interface{}{
					"DEBUG":   "true",
					"API_URL": "https://api.example.com",
				},
				ExpiresIn: gboxsdk.String("10m"),
				Labels: map[string]interface{}{
					"project":     "web-automation",
					"environment": "testing",
				},
			},
			Timeout: gboxsdk.String("30s"),
			Wait:    gboxsdk.Bool(true),
		},
	})
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BoxExecuteCommandsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.ExecuteCommands(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxExecuteCommandsParams{
			Commands: gboxsdk.V1BoxExecuteCommandsParamsCommandsUnion{
				OfStringArray: []string{"ls", "-l"},
			},
			Envs: map[string]interface{}{
				"PATH":     "/usr/bin:/bin",
				"NODE_ENV": "production",
			},
			Timeout:    gboxsdk.String("30s"),
			WorkingDir: gboxsdk.String("/home/user/projects"),
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

func TestV1BoxRunCodeWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.RunCode(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxRunCodeParams{
			Code: `print("Hello, World!")`,
			Argv: []string{"-v", "--help"},
			Envs: map[string]interface{}{
				"PYTHONPATH": "/usr/lib/python3",
				"DEBUG":      "true",
			},
			Language:   gboxsdk.V1BoxRunCodeParamsLanguagePython3,
			Timeout:    gboxsdk.String("timeout"),
			WorkingDir: gboxsdk.String("/home/user/scripts"),
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

func TestV1BoxStartWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Start(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxStartParams{
			Timeout: gboxsdk.String("30s"),
			Wait:    gboxsdk.Bool(true),
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

func TestV1BoxStopWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Stop(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxStopParams{
			Timeout: gboxsdk.String("30s"),
			Wait:    gboxsdk.Bool(true),
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
