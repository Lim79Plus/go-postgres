package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lim79Plus/go-postgres/go/infra"
	"github.com/stretchr/testify/assert"
)

func TestApiSayHello(t *testing.T) {
	router := infra.Router

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	// check the response status code
	if rec.Code != http.StatusOK {
		t.Errorf("rec.Code must be http.StatusOK. :%v", rec.Code)
	}
	// check the response response body
	want := `{"msg":"Hello Workd!"}`
	assert.JSONEq(t, want, rec.Body.String(), "resp body is not expected.")

}
