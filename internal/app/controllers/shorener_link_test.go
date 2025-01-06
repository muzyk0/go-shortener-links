package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/muzyk0/go-shortener-links/internal/app/logger"
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
	path string, body io.Reader) (*http.Response, []byte) {
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

	return resp, respBody
}

type (
	args struct {
		link string
	}
	want struct {
		status int
	}

	testCase struct {
		name     string
		url      string
		args     args
		want     want
		testFunc func(*testing.T, testCase)
	}
)

func TestShortenerLink(t *testing.T) {
	appConfig := config.NewConfig()
	appLogger, err := logger.NewLogger(appConfig.FlagLogLevel)
	require.NoError(t, err)
	ts := httptest.NewServer(AppRouter(appConfig, *appLogger))
	defer ts.Close()

	tests := []testCase{
		{
			name: "Should generate short link",
			url:  "/",
			args: args{
				link: "https://kagi.com/",
			},
			want: want{
				status: http.StatusCreated,
			},
			testFunc: func(t *testing.T, tc testCase) {
				res, resBody := testRequest(t, http.MethodPost, ts.URL+tc.url, strings.NewReader(tc.args.link))
				defer res.Body.Close()
				assert.Equal(t, res.StatusCode, tc.want.status)

				redirectResponse, _ := testRequest(t, http.MethodGet, string(resBody), nil)
				defer redirectResponse.Body.Close()
				assert.Equal(t, http.StatusTemporaryRedirect, redirectResponse.StatusCode)

				location := redirectResponse.Header.Get("location")
				assert.Equal(t, tc.args.link, location)
			},
		},
		{
			name: "Should generate short link with JSON",
			url:  "/api/shorten",
			args: args{
				link: "https://kagi.com/",
			},
			want: want{
				status: http.StatusCreated,
			},
			testFunc: func(t *testing.T, tc testCase) {
				type RequestBody struct {
					Url string `json:"url"`
				}

				type ResponseBody struct {
					Result string `json:"result"`
				}

				jsonBody, err := json.Marshal(RequestBody{Url: tc.args.link})
				require.NoError(t, err)

				// Convert jsonBody to an io.Reader
				jsonBodyReader := bytes.NewReader(jsonBody)

				res, resBody := testRequest(t, http.MethodPost, ts.URL+tc.url, jsonBodyReader)
				defer res.Body.Close()
				assert.Equal(t, tc.want.status, res.StatusCode)

				var responseBody ResponseBody
				err = json.Unmarshal(resBody, &responseBody)
				require.NoError(t, err)

				redirectResponse, _ := testRequest(t, http.MethodGet, responseBody.Result, nil)
				defer redirectResponse.Body.Close()
				assert.Equal(t, http.StatusTemporaryRedirect, redirectResponse.StatusCode)

				location := redirectResponse.Header.Get("location")
				assert.Equal(t, tc.args.link, location)
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
			testFunc: func(t *testing.T, tc testCase) {
				res, _ := testRequest(t, http.MethodGet, ts.URL+tc.url, nil)
				defer res.Body.Close()
				assert.Equal(t, res.StatusCode, tc.want.status)
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
			testFunc: func(t *testing.T, tc testCase) {
				res, _ := testRequest(t, http.MethodPut, ts.URL+tc.url, nil)
				defer res.Body.Close()
				assert.Equal(t, res.StatusCode, tc.want.status)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.testFunc(t, tc)

			//if tc.name == "Should generate short link" {
			//
			//} else if tc.name == "Should return 403 for non-existing short link" {
			//
			//} else if tc.name == "Should return 403 for invalid method" {
			//
			//}
		})
	}
}
