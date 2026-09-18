package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPlaygroundRoutes(t *testing.T) {
	for _, path := range []string{"/", "/?url=https://example.com", "/playground.html", "/playground.js", "/url-compressor.js", "/huffman.js"} {
		t.Run(path, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, path, nil)
			response := httptest.NewRecorder()
			serveRoot(response, request)
			if response.Code != http.StatusOK {
				t.Fatalf("status = %d: %s", response.Code, response.Body.String())
			}
			name := strings.TrimPrefix(request.URL.Path, "/")
			if name == "" {
				name = "playground.html"
			}
			want, err := playgroundFiles.ReadFile(name)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(response.Body.Bytes(), want) {
				t.Fatal("response differs from embedded asset")
			}
			contentType := response.Header().Get("Content-Type")
			if strings.HasSuffix(name, ".html") {
				if !strings.HasPrefix(contentType, "text/html") {
					t.Fatalf("HTML Content-Type = %q", contentType)
				}
			} else if !strings.Contains(contentType, "javascript") {
				t.Fatalf("module Content-Type = %q", contentType)
			}
			head := httptest.NewRecorder()
			serveRoot(head, httptest.NewRequest(http.MethodHead, path, nil))
			if head.Code != http.StatusOK || head.Body.Len() != 0 {
				t.Fatalf("HEAD status = %d, body length = %d", head.Code, head.Body.Len())
			}
			post := httptest.NewRecorder()
			serveRoot(post, httptest.NewRequest(http.MethodPost, path, nil))
			if post.Code != http.StatusMethodNotAllowed || post.Header().Get("Allow") != "GET, HEAD" {
				t.Fatalf("POST response = %v", post.Result())
			}
		})
	}
}

func TestCompressedLinkRoute(t *testing.T) {
	response := httptest.NewRecorder()
	serveRoot(response, httptest.NewRequest(http.MethodGet, "/"+encode([]byte("example.com/path")), nil))
	if response.Code != http.StatusOK || !strings.Contains(response.Body.String(), "https://example.com/path") {
		t.Fatalf("compressed link response = %d %s", response.Code, response.Body.String())
	}
}

func TestNonAssetRoutes(t *testing.T) {
	for _, path := range []string{"/main.go", "/README.md", "/.git/config", "/unknown.js"} {
		response := httptest.NewRecorder()
		serveRoot(response, httptest.NewRequest(http.MethodGet, path, nil))
		if response.Code != http.StatusBadRequest {
			t.Fatalf("%s: status = %d", path, response.Code)
		}
	}
}
