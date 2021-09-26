package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {

	r := newRouter()
	// In case there is an error in forming the request, we fail and stop the test
	mockServer := httptest.NewServer(r)
	// recorder := httptest.NewRecorder()
	resp, err := http.Get(mockServer.URL + "/")

	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// In the next few lines, the response body is read, and converted to a string
	defer resp.Body.Close()
	// read the body into a bunch of bytes (b)
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response body is what we expect.
	expected := `0 Cape Canaveral Air Force Station Space Launch Complex 401 Vandenberg Air Force Base Space Launch Complex 4E2 Kennedy Space Center Historic Launch Complex 39A3 Cape Canaveral Air Force Station Space Launch Complex 404 Cape Canaveral Air Force Station Space Launch Complex 405 Kennedy Space Center Historic Launch Complex 39A6 Kennedy Space Center Historic Launch Complex 39A7 Kennedy Space Center Historic Launch Complex 39A8 Cape Canaveral Air Force Station Space Launch Complex 409 Cape Canaveral Air Force Station Space Launch Complex 40`

	actual := string(bytes)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	// Most of the code is similar. The only difference is that now we make a
	//request to a route we know we didn't define, like the `POST /hello` route.
	resp, err := http.Post(mockServer.URL+"/", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	// The code to test the body is also mostly the same, except this time, we
	// expect an empty body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}
