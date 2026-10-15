package updater

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// releaseServer serves both the REST API endpoint and the releases/latest
// web redirect, so fetchLatest can be exercised without the network.
func releaseServer(t *testing.T, api http.HandlerFunc, web http.HandlerFunc) {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/repos/o/r/releases/latest", api)
	mux.HandleFunc("/o/r/releases/latest", web)
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)

	oldAPI, oldWeb := releasesAPIFmt, releasesWebFmt
	releasesAPIFmt = srv.URL + "/api/repos/%s/releases/latest"
	releasesWebFmt = srv.URL + "/%s/releases/latest"
	t.Cleanup(func() { releasesAPIFmt, releasesWebFmt = oldAPI, oldWeb })
	t.Setenv(envRepo, "o/r")
}

func rateLimited(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("X-RateLimit-Remaining", "0")
	w.Header().Set("X-RateLimit-Reset", "1792062213")
	w.WriteHeader(http.StatusForbidden)
}

func redirectTo(tag string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/o/r/releases/tag/"+tag, http.StatusFound)
	}
}

func notCalled(t *testing.T) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		t.Error("fallback was called although the API succeeded")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func TestFetchLatestUsesAPIWhenAvailable(t *testing.T) {
	releaseServer(t, func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `{"tag_name":"v1.2.3","html_url":"https://example.com/v1.2.3"}`)
	}, notCalled(t))

	r, err := fetchLatest("1.0.0")
	if err != nil {
		t.Fatal(err)
	}
	if r.Latest != "1.2.3" || r.HTMLURL != "https://example.com/v1.2.3" || !r.HasUpdate() {
		t.Fatalf("unexpected result: %+v", r)
	}
}

func TestFetchLatestFallsBackOnRateLimit(t *testing.T) {
	releaseServer(t, rateLimited, redirectTo("v1.0.2"))

	r, err := fetchLatest("1.0.1")
	if err != nil {
		t.Fatalf("expected fallback to succeed, got %v", err)
	}
	if r.Latest != "1.0.2" || !r.HasUpdate() {
		t.Fatalf("unexpected result: %+v", r)
	}
	if !strings.HasSuffix(r.HTMLURL, "/o/r/releases/tag/v1.0.2") {
		t.Fatalf("HTMLURL = %q, want the tag page", r.HTMLURL)
	}
}

func TestFetchLatestFallsBackOnOtherAPIErrors(t *testing.T) {
	releaseServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	}, redirectTo("v2.0.0"))

	r, err := fetchLatest("1.0.0")
	if err != nil {
		t.Fatalf("expected fallback to succeed, got %v", err)
	}
	if r.Latest != "2.0.0" {
		t.Fatalf("Latest = %q, want 2.0.0", r.Latest)
	}
}

func TestFetchLatestBothFailNamesRateLimit(t *testing.T) {
	releaseServer(t, rateLimited, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	})

	_, err := fetchLatest("1.0.0")
	if err == nil {
		t.Fatal("expected an error when both lookups fail")
	}
	msg := err.Error()
	for _, want := range []string{"rate limit exceeded", "HTTP 403", "resets at", "fallback via releases page also failed", "HTTP 503"} {
		if !strings.Contains(msg, want) {
			t.Errorf("error %q missing %q", msg, want)
		}
	}
}

func TestFetchLatestWebRejectsNonTagRedirect(t *testing.T) {
	// A repo with no releases redirects releases/latest to the releases list.
	releaseServer(t, rateLimited, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/o/r/releases", http.StatusFound)
	})

	if _, err := fetchLatest("1.0.0"); err == nil || !strings.Contains(err.Error(), "not a release tag") {
		t.Fatalf("expected a not-a-release-tag error, got %v", err)
	}
}
