package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpointStudents(t *testing.T) {
	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()
	setupHandlers()

	for _, tt := range []struct {
		ID          string
		StatusCode  int
		Description string
	}{
		{"", http.StatusBadRequest, "missing ID"},
		{"1", http.StatusBadRequest, "invalid ID"},
		{"1234", http.StatusOK, "valid ID"},
		// IMPORTANT: to get HTMX to display the error response to users we return 200 instead of 404
		{"5678", http.StatusOK, "valid non-existent ID"},
		{"4444", http.StatusOK, "another valid ID"},
	} {
		r, err := http.NewRequest("GET", ts.URL+"/students?id="+tt.ID, nil)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.DefaultClient.Do(r)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != tt.StatusCode {
			t.Errorf(
				"Students endpoint expected status code %d for "+
					"case %s with ID \"%s\" but got %d",
				tt.StatusCode,
				tt.Description, tt.ID, resp.StatusCode,
			)
		}
	}
}
