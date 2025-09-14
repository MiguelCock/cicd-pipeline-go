package go_test

import (
	"cicd-pipeline-go/endpoints"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestIndexGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	endpoints.IndexHandler(w, req)

	res := w.Result()
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	body := w.Body.String()
	if !strings.Contains(body, "<!DOCTYPE html>") {
		t.Errorf("expected <!DOCTYPE html> in response, got: %s", body)
	}
}

func testTemplate(t *testing.T, num1, num2, operacion, expected string) {
	form := url.Values{}
	form.Add("num1", num1)
	form.Add("num2", num2)
	form.Add("operacion", operacion)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	endpoints.IndexHandler(w, req)

	body := w.Body.String()
	if !strings.Contains(body, expected) {
		t.Errorf("expected result %s in response, got: %s", expected, body)
	}
}

func TestPostCalculator(t *testing.T) {
	test := []struct {
		num1, num2 string
		operacion  string
		expect     string
	}{
		{"2", "3", "sumar", "5"},
		{"5", "3", "restar", "2"},
		{"2", "3", "multiplicar", "6"},
		{"6", "3", "dividir", "2"},
		{"6", "0", "dividir", "error: Division by zero"},
		{"6", "3", "invalid", "error: invalid operation"},
		{"a", "b", "sumar", "error: enter valid numbers"},
	}

	for _, tt := range test {
		testTemplate(t, tt.num1, tt.num2, tt.operacion, tt.expect)
	}
}
