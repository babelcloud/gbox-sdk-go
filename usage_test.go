// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gboxsdk_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/babelcloud/gbox-sdk-go"
	"github.com/babelcloud/gbox-sdk-go/option"
)

func TestUsage(t *testing.T) {
	// 使用 httptest.Server mock /boxes 接口
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/boxes" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json") // 设置返回类型为 JSON
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":"mocked-box-id"}`)) // 返回一个简单的 JSON
			return
		}
		http.NotFound(w, r)
	}))
	defer mockServer.Close()

	client := gboxsdk.NewClient(
		option.WithBaseURL(mockServer.URL),
		option.WithAPIKey("My API Key"),
	)
	box, err := client.V1.Boxes.New(context.TODO(), gboxsdk.V1BoxNewParams{
		OfCreateLinuxBox: &gboxsdk.CreateLinuxBoxParam{
			Type: gboxsdk.CreateLinuxBoxTypeLinux,
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", box)
}
