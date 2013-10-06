package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	initDb()
}

func TestAccountsIndexReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/accounts", nil)
	response := httptest.NewRecorder()

	accountsIndex(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}
}

// func TestHandInsertBookmarkWithStatusOK(t *testing.T) {
// 	bookmark := Bookmark{"wercker", "http://wercker.com"}

// 	b, err := json.Marshal(bookmark)
// 	if err != nil {
// 		t.Fatalf("Unable to marshal Bookmark")
// 	}

// 	request, _ := http.NewRequest("POST", "/new", bytes.NewReader(b))
// 	response := httptest.NewRecorder()

// 	insertBookmark(response, request)

// 	body := response.Body.String()
// 	if !strings.Contains(body, "{'bookmark':'saved'}") {
// 		t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "San Francisco", body)
// 	}
// }

func TestAccountsIndexReturnsJSON(t *testing.T) {
	request, _ := http.NewRequest("GET", "/accounts", nil)
	response := httptest.NewRecorder()

	accountsIndex(response, request)

	ct := response.HeaderMap["Content-Type"][0]
	if !strings.EqualFold(ct, "application/json") {
		t.Fatalf("Content-Type does not equal 'application/json'")
	}
}
