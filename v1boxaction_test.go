// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/babelcloud/gbox-sdk-go"
	"github.com/babelcloud/gbox-sdk-go/internal/testutil"
	"github.com/babelcloud/gbox-sdk-go/option"
)

func TestV1BoxActionAIWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.AI(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionAIParams{
			Instruction:        "click the login button",
			Background:         gboxsdk.String("The user is on the login page"),
			IncludeScreenshot:  gboxsdk.Bool(false),
			OutputFormat:       gboxsdk.V1BoxActionAIParamsOutputFormatBase64,
			PresignedExpiresIn: gboxsdk.String("30m"),
			ScreenshotDelay:    gboxsdk.String("500ms"),
			Settings: gboxsdk.V1BoxActionAIParamsSettings{
				DisableActions: []string{"swipe"},
				SystemPrompt:   gboxsdk.String("You are a helpful assistant specialized in UI automation. When given a screenshot and instruction, analyze the visual elements carefully and execute the most appropriate action. Always prioritize user safety and avoid destructive actions unless explicitly requested."),
			},
			Stream: gboxsdk.Bool(false),
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

func TestV1BoxActionClickWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.Click(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionClickParams{
			X:                  100,
			Y:                  100,
			Button:             gboxsdk.V1BoxActionClickParamsButtonLeft,
			Double:             gboxsdk.Bool(false),
			IncludeScreenshot:  gboxsdk.Bool(false),
			OutputFormat:       gboxsdk.V1BoxActionClickParamsOutputFormatBase64,
			PresignedExpiresIn: gboxsdk.String("30m"),
			ScreenshotDelay:    gboxsdk.String("500ms"),
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

func TestV1BoxActionDragWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.Drag(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionDragParams{
			OfDragSimple: &gboxsdk.V1BoxActionDragParamsBodyDragSimple{
				End: gboxsdk.V1BoxActionDragParamsBodyDragSimpleEnd{
					X: 200,
					Y: 200,
				},
				Start: gboxsdk.V1BoxActionDragParamsBodyDragSimpleStart{
					X: 100,
					Y: 100,
				},
				Duration:           gboxsdk.String("500ms"),
				IncludeScreenshot:  gboxsdk.Bool(false),
				OutputFormat:       "base64",
				PresignedExpiresIn: gboxsdk.String("30m"),
				ScreenshotDelay:    gboxsdk.String("500ms"),
			},
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

func TestV1BoxActionExtractWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.Extract(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionExtractParams{
			Instruction: "Extract the email address from the UI interface",
			Schema:      map[string]interface{}{},
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

func TestV1BoxActionMoveWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.Move(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionMoveParams{
			X:                  200,
			Y:                  300,
			IncludeScreenshot:  gboxsdk.Bool(false),
			OutputFormat:       gboxsdk.V1BoxActionMoveParamsOutputFormatBase64,
			PresignedExpiresIn: gboxsdk.String("30m"),
			ScreenshotDelay:    gboxsdk.String("500ms"),
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

func TestV1BoxActionPressButtonWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.PressButton(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionPressButtonParams{
			Buttons:            []string{"power"},
			IncludeScreenshot:  gboxsdk.Bool(false),
			OutputFormat:       gboxsdk.V1BoxActionPressButtonParamsOutputFormatBase64,
			PresignedExpiresIn: gboxsdk.String("30m"),
			ScreenshotDelay:    gboxsdk.String("500ms"),
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

func TestV1BoxActionPressKeyWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.PressKey(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionPressKeyParams{
			Keys:               []string{"enter"},
			Combination:        gboxsdk.Bool(true),
			IncludeScreenshot:  gboxsdk.Bool(false),
			OutputFormat:       gboxsdk.V1BoxActionPressKeyParamsOutputFormatBase64,
			PresignedExpiresIn: gboxsdk.String("30m"),
			ScreenshotDelay:    gboxsdk.String("500ms"),
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

func TestV1BoxActionRecordingStart(t *testing.T) {
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
	err := client.V1.Boxes.Actions.RecordingStart(context.TODO(), "c9bdc193-b54b-4ddb-a035-5ac0c598d32d")
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BoxActionRecordingStop(t *testing.T) {
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
	err := client.V1.Boxes.Actions.RecordingStop(context.TODO(), "c9bdc193-b54b-4ddb-a035-5ac0c598d32d")
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BoxActionScreenLayout(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.ScreenLayout(context.TODO(), "c9bdc193-b54b-4ddb-a035-5ac0c598d32d")
	if err != nil {
		var apierr *gboxsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1BoxActionScreenRotationWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.ScreenRotation(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionScreenRotationParams{
			Orientation:        gboxsdk.V1BoxActionScreenRotationParamsOrientationLandscapeLeft,
			IncludeScreenshot:  gboxsdk.Bool(false),
			OutputFormat:       gboxsdk.V1BoxActionScreenRotationParamsOutputFormatBase64,
			PresignedExpiresIn: gboxsdk.String("30m"),
			ScreenshotDelay:    gboxsdk.String("500ms"),
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

func TestV1BoxActionScreenshotWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.Screenshot(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionScreenshotParams{
			Clip: gboxsdk.V1BoxActionScreenshotParamsClip{
				Height: 600,
				Width:  800,
				X:      100,
				Y:      50,
			},
			OutputFormat: gboxsdk.V1BoxActionScreenshotParamsOutputFormatBase64,
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

func TestV1BoxActionScrollWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.Scroll(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionScrollParams{
			ScrollX:            0,
			ScrollY:            100,
			X:                  100,
			Y:                  100,
			IncludeScreenshot:  gboxsdk.Bool(false),
			OutputFormat:       gboxsdk.V1BoxActionScrollParamsOutputFormatBase64,
			PresignedExpiresIn: gboxsdk.String("30m"),
			ScreenshotDelay:    gboxsdk.String("500ms"),
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

func TestV1BoxActionSwipeWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.Swipe(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionSwipeParams{
			OfSwipeSimple: &gboxsdk.V1BoxActionSwipeParamsBodySwipeSimple{
				Direction:          "up",
				Distance:           gboxsdk.Float(300),
				Duration:           gboxsdk.String("500ms"),
				IncludeScreenshot:  gboxsdk.Bool(false),
				OutputFormat:       "base64",
				PresignedExpiresIn: gboxsdk.String("30m"),
				ScreenshotDelay:    gboxsdk.String("500ms"),
			},
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

func TestV1BoxActionTouchWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.Touch(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionTouchParams{
			Points: []gboxsdk.V1BoxActionTouchParamsPoint{{
				Start: gboxsdk.V1BoxActionTouchParamsPointStart{
					X: 100,
					Y: 150,
				},
				Actions: []gboxsdk.V1BoxActionTouchParamsPointActionUnion{{
					OfTouchPointMoveAction: &gboxsdk.V1BoxActionTouchParamsPointActionTouchPointMoveAction{
						Duration: "200ms",
						Type:     "move",
						X:        400,
						Y:        300,
					},
				}},
			}},
			IncludeScreenshot:  gboxsdk.Bool(false),
			OutputFormat:       gboxsdk.V1BoxActionTouchParamsOutputFormatBase64,
			PresignedExpiresIn: gboxsdk.String("30m"),
			ScreenshotDelay:    gboxsdk.String("500ms"),
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

func TestV1BoxActionTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Boxes.Actions.Type(
		context.TODO(),
		"c9bdc193-b54b-4ddb-a035-5ac0c598d32d",
		gboxsdk.V1BoxActionTypeParams{
			Text:               "Hello World",
			IncludeScreenshot:  gboxsdk.Bool(false),
			Mode:               gboxsdk.V1BoxActionTypeParamsModeAppend,
			OutputFormat:       gboxsdk.V1BoxActionTypeParamsOutputFormatBase64,
			PresignedExpiresIn: gboxsdk.String("30m"),
			ScreenshotDelay:    gboxsdk.String("500ms"),
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
