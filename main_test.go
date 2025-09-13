package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestIndexGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	indexHandler(w, req)

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

func TestIndexPostSumar(t *testing.T) {
	form := url.Values{}
	form.Add("num1", "2")
	form.Add("num2", "3")
	form.Add("operacion", "sumar")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	indexHandler(w, req)

	body := w.Body.String()
	if !strings.Contains(body, "5") {
		t.Errorf("expected result 5 in response, got: %s", body)
	}
}

func TestIndexPostRestar(t *testing.T) {
	form := url.Values{}
	form.Add("num1", "5")
	form.Add("num2", "3")
	form.Add("operacion", "restar")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	indexHandler(w, req)

	body := w.Body.String()
	if !strings.Contains(body, "2") {
		t.Errorf("expected result 2 in response, got: %s", body)
	}
}

func TestIndexPostMultiplicar(t *testing.T) {
	form := url.Values{}
	form.Add("num1", "2")
	form.Add("num2", "3")
	form.Add("operacion", "multiplicar")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	indexHandler(w, req)

	body := w.Body.String()
	if !strings.Contains(body, "6") {
		t.Errorf("expected result 6 in response, got: %s", body)
	}
}

func TestIndexPostDividir(t *testing.T) {
	form := url.Values{}
	form.Add("num1", "6")
	form.Add("num2", "3")
	form.Add("operacion", "dividir")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	indexHandler(w, req)

	body := w.Body.String()
	if !strings.Contains(body, "2") {
		t.Errorf("expected result 2 in response, got: %s", body)
	}
}

func TestIndexPostDividirByZero(t *testing.T) {
	form := url.Values{}
	form.Add("num1", "6")
	form.Add("num2", "0")
	form.Add("operacion", "dividir")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	indexHandler(w, req)

	body := w.Body.String()
	if !strings.Contains(body, "error: Division by zero") {
		t.Errorf("expected divide by zero error, got: %s", body)
	}
}

func TestIndexPostInvalidOperation(t *testing.T) {
	form := url.Values{}
	form.Add("num1", "6")
	form.Add("num2", "3")
	form.Add("operacion", "invalid")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	indexHandler(w, req)

	body := w.Body.String()
	if !strings.Contains(body, "error: invalid operation") {
		t.Errorf("expected invalid operation message, got: %s", body)
	}
}

func TestIndexPostInvalidNumbers(t *testing.T) {
	form := url.Values{}
	form.Add("num1", "a")
	form.Add("num2", "b")
	form.Add("operacion", "sumar")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	indexHandler(w, req)

	body := w.Body.String()
	if !strings.Contains(body, "error: enter valid numbers") {
		t.Errorf("expected invalid numbers message, got: %s", body)
	}
}
