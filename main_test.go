package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreateURL(t *testing.T) {
	tt := []struct {
		name           string
		baseURL        string
		params         []string
		expectedResult string
		expectedToFail bool
	}{
		{
			name:           "with params",
			baseURL:        "https://test.com/",
			params:         []string{"a", "b", "c"},
			expectedResult: "https://test.com/a,b,c",
			expectedToFail: false,
		},
		{
			name:           "without params",
			baseURL:        "https://test.com/",
			params:         []string{},
			expectedResult: "",
			expectedToFail: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url, err := createURL(tc.baseURL, tc.params)
			if err != nil {
				if tc.expectedToFail {
					t.Skipf("test failed as expected: %v", err)
				}

				t.Errorf("while creating URL: %v", err)
			}

			if url != tc.expectedResult {
				t.Errorf("expected URL to be %q. got=%q", tc.expectedResult, url)
			}
		})
	}
}

func TestFetch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))
	defer ts.Close()

	tt := []struct {
		name           string
		url            string
		expectedToFail bool
	}{
		{
			name:           "with valid URL",
			url:            ts.URL,
			expectedToFail: false,
		},
		{
			name:           "without valid URL",
			url:            "",
			expectedToFail: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := fetch(tc.url)
			if err != nil {
				if tc.expectedToFail {
					t.Skipf("test failed as expected: %v", err)
				}

				t.Errorf("while fetching URL %q: %v", tc.url, err)
			}
		})
	}
}

func TestCreatingFile(t *testing.T) {
	filename := "./tests/testfile.txt"

	// tests directory should not contain testfile.txt
	if checkIfExists(filename) {
		t.Fatalf("test file %q already exists", filename)
	}

	_, err := createFile(filename)
	if err != nil {
		t.Fatalf("while creating test file %q: %v", filename, err)
	}

	// tests directory should contain now testfile.txt
	if !checkIfExists(filename) {
		t.Fatalf("test file %q should exists", filename)
	}

	// Delete test file
	if err := os.Remove(filename); err != nil {
		t.Fatalf("while removing test file %q: %v", filename, err)
	}

	// Check if test file has already been deleted
	if checkIfExists(filename) {
		t.Fatalf("test file %q should not exists", filename)
	}
}

func TestWriteTo(t *testing.T) {
	var buf bytes.Buffer
	if err := writeTo(&buf, []byte("Hello, World!")); err != nil {
		t.Fatalf("while writing contents: %v", err)
	}

	if buf.String() != "Hello, World!" {
		t.Fatalf("expected written content to be %q. got=%q", "Hello, World!", buf.String())
	}
}
