package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortLinkHandler(t *testing.T) {
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
				r := httptest.NewRequest(http.MethodPost, tt.url, strings.NewReader(tt.args.link))
				w := httptest.NewRecorder()
				r.Header.Set("Content-Type", "text/plain")

				ShortLinkHandler(w, r)

				res := w.Result()
				defer res.Body.Close()

				assert.Equal(t, res.StatusCode, tt.want.status)

				resBody, err := io.ReadAll(res.Body)
				require.NoError(t, err)

				rn := httptest.NewRequest(http.MethodGet, string(resBody), nil)
				wn := httptest.NewRecorder()

				ShortLinkHandler(wn, rn)

				res2 := wn.Result()
				defer res2.Body.Close()

				assert.Equal(t, res2.StatusCode, http.StatusTemporaryRedirect)

				location := res2.Header.Get("location")
				assert.Equal(t, tt.args.link, location)
			} else if tt.name == "Should return 403 for non-existing short link" {
				r := httptest.NewRequest(http.MethodGet, tt.url, nil)
				w := httptest.NewRecorder()

				ShortLinkHandler(w, r)

				res := w.Result()
				defer res.Body.Close()

				assert.Equal(t, res.StatusCode, tt.want.status)
			} else if tt.name == "Should return 403 for invalid method" {
				r := httptest.NewRequest(http.MethodPut, tt.url, nil)
				w := httptest.NewRecorder()

				ShortLinkHandler(w, r)

				res := w.Result()
				defer res.Body.Close()

				assert.Equal(t, res.StatusCode, tt.want.status)
			}
		})
	}
}
