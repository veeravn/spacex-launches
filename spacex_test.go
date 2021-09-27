package main

import (
	"net/url"
)

func newSpaceXQueryForm() *url.Values {
	form := url.Values{}
	form.Set("limit", "10")
	form.Set("offset", "0")
	return &form
}

// func TestRouter(t *testing.T) {
//
// 	// form := newSpaceXQueryForm()
// 	req, err := http.NewRequest("GET", "/launchesPast", nil)
// 	q := req.URL.Query()
// 	q.Add("limit", "10")
// 	q.Add("offset", "0")
// 	req.URL.RawQuery = q.Encode()
// 	log.Println(req)
// 	recorder := httptest.NewRecorder()
// 	hf := http.HandlerFunc(handler)
//
// 	hf.ServeHTTP(recorder, req)
// 	status := recorder.Code
// 	if status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	// In the next few lines, the response body is read, and converted to a string
// 	// read the body into a bunch of bytes (b)
// 	bytes, err := ioutil.ReadAll(recorder.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// Check the response body is what we expect.
// 	expected := `[{"MissionName":"Starlink-15 (v1.0)","SiteNameLong":"Cape Canaveral Air Force Station Space Launch Complex 40","RocketName":"Falcon 9"},{"MissionName":"Sentinel-6 Michael Freilich","SiteNameLong":"Vandenberg Air Force Base Space Launch Complex 4E","RocketName":"Falcon 9"},{"MissionName":"Crew-1","SiteNameLong":"Kennedy Space Center Historic Launch Complex 39A","RocketName":"Falcon 9"},{"MissionName":"GPS III SV04 (Sacagawea)","SiteNameLong":"Cape Canaveral Air Force Station Space Launch Complex 40","RocketName":"Falcon 9"},{"MissionName":"Starlink-14 (v1.0)","SiteNameLong":"Cape Canaveral Air Force Station Space Launch Complex 40","RocketName":"Falcon 9"},{"MissionName":"Starlink-13 (v1.0)","SiteNameLong":"Kennedy Space Center Historic Launch Complex 39A","RocketName":"Falcon 9"},{"MissionName":"Starlink-12 (v1.0)","SiteNameLong":"Kennedy Space Center Historic Launch Complex 39A","RocketName":"Falcon 9"},{"MissionName":"Starlink-11 (v1.0)","SiteNameLong":"Kennedy Space Center Historic Launch Complex 39A","RocketName":"Falcon 9"},{"MissionName":"SAOCOM 1B, GNOMES-1, Tyvak-0172","SiteNameLong":"Cape Canaveral Air Force Station Space Launch Complex 40","RocketName":"Falcon 9"},{"MissionName":"Starlink-10 (v1.0) \u0026 SkySat 19-21","SiteNameLong":"Cape Canaveral Air Force Station Space Launch Complex 40","RocketName":"Falcon 9"}]`
//
// 	actual := string(bytes)
// 	if actual != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
// 	}
// }

// func TestStaticFileServer(t *testing.T) {
// 	r := newRouter()
// 	mockServer := httptest.NewServer(r)
//
// 	// We want to hit the `GET /assets/` route to get the index.html file response
// 	resp, err := http.Get(mockServer.URL + "/assets/")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// We want our status to be 200 (ok)
// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("Status should be 200, got %d", resp.StatusCode)
// 	}
//
// 	// It isn't wise to test the entire content of the HTML file.
// 	// Instead, we test that the content-type header is "text/html; charset=utf-8"
// 	// so that we know that an html file has been served
// 	contentType := resp.Header.Get("Content-Type")
// 	expectedContentType := "text/html; charset=utf-8"
//
// 	if expectedContentType != contentType {
// 		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
// 	}
//
// }
