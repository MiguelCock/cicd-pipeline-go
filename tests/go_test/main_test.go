package go_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"cicd-pipeline-go/endpoints"
)

func TestIndexGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	endpoints.IndexHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

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

func TestIndexPostSumar(t *testing.T) {
	testTemplate(t, "2", "3", "sumar", "5")
}

func TestIndexPostRestar(t *testing.T) {
	testTemplate(t, "5", "3", "restar", "2")
}

func TestIndexPostMultiplicar(t *testing.T) {
	testTemplate(t, "2", "3", "multiplicar", "6")
}

func TestIndexPostDividir(t *testing.T) {
	testTemplate(t, "6", "3", "dividir", "2")
}

func TestIndexPostDividirByZero(t *testing.T) {
	testTemplate(t, "6", "0", "dividir", "error: Division by zero")
}

func TestIndexPostInvalidOperation(t *testing.T) {
	testTemplate(t, "6", "3", "invalid", "error: invalid operation")
}

func TestIndexPostInvalidNumbers(t *testing.T) {
	testTemplate(t, "a", "b", "sumar", "error: enter valid numbers")
}
