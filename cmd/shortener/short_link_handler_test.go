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
		// shortLink string
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
				// shortLink: "http://localhost:8080/123",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, tt.url, strings.NewReader(tt.args.link))
			w := httptest.NewRecorder()
			r.Header.Set("Content-Type", "text/plain")

			ShortLinkHandler(w, r)

			res := w.Result()

			assert.Equal(t, res.StatusCode, tt.want.status)

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)

			rn := httptest.NewRequest(http.MethodGet, string(resBody), nil)
			wn := httptest.NewRecorder()
			rn.Header.Set("Content-Type", "text/plain")

			ShortLinkHandler(wn, rn)

			res2 := wn.Result()

			assert.Equal(t, res2.StatusCode, http.StatusTemporaryRedirect)

			location := res2.Header.Get("location")

			assert.Equal(t, tt.args.link, location)
		})
	}
}
