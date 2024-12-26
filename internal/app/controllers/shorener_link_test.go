package controllers

import (
	"go-shorener-links/config"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRequest(t *testing.T, method,
	path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, path, body)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Если будет редирект, ничего не делать, просто не следовать за ним
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	require.NoError(t, err)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestShortenerLink(t *testing.T) {
	appConfig := config.NewConfig()
	ts := httptest.NewServer(AppRouter(appConfig))
	defer ts.Close()

	type args struct {
		link string
	}
	type want struct {
		status int
	}
	tests := []struct {
		name string
		url  string
		args args
		want want
	}{
		{
			name: "Should generate short link",
			url:  "/",
			args: args{
				link: "https://kagi.com/",
			},
			want: want{
				status: http.StatusCreated,
			},
		},
		{
			name: "Should return 403 for non-existing short link",
			url:  "/nonexisting",
			args: args{
				link: "",
			},
			want: want{
				status: http.StatusForbidden,
			},
		},
		{
			name: "Should return 403 for invalid method",
			url:  "/",
			args: args{
				link: "",
			},
			want: want{
				status: http.StatusForbidden,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Should generate short link" {
				res, resBody := testRequest(t, http.MethodPost, ts.URL+tt.url, strings.NewReader(tt.args.link))
				defer res.Body.Close()
				assert.Equal(t, res.StatusCode, tt.want.status)

				redirectResponse, _ := testRequest(t, http.MethodGet, string(resBody), nil)
				defer redirectResponse.Body.Close()
				assert.Equal(t, http.StatusTemporaryRedirect, redirectResponse.StatusCode)

				location := redirectResponse.Header.Get("location")
				assert.Equal(t, tt.args.link, location)
			} else if tt.name == "Should return 403 for non-existing short link" {
				res, _ := testRequest(t, http.MethodGet, ts.URL+tt.url, nil)
				defer res.Body.Close()
				assert.Equal(t, res.StatusCode, tt.want.status)
			} else if tt.name == "Should return 403 for invalid method" {
				res, _ := testRequest(t, http.MethodPut, ts.URL+tt.url, nil)
				defer res.Body.Close()
				assert.Equal(t, res.StatusCode, tt.want.status)
			}
		})
	}
}
