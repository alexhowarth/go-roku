// package roku_test tests client functions
package roku_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/alexhowarth/go-roku"
)

// Provide a Transport for testing purposes (pass into the constructor to override the default http.Client)
type roundTripFunc func(r *http.Request) (*http.Response, error)

func (s roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return s(r)
}

func TestDeviceInfo(t *testing.T) {

	// RoundTrip test data
	tc := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(helperLoadBytes(t, "deviceinfo.xml"))),
		}, nil
	})}

	c, err := roku.NewClient("http://foo:8060", roku.WithHttpClient(tc))
	if err != nil {
		t.Error("Unable to create client")
	}

	got, err := c.DeviceInfo()
	if err != nil {
		t.Error("Unable to get DeviceInfo")
	}

	expected := `Alex's Roku Streaming Stick+`
	if got.FriendlyDeviceName != expected {
		t.Errorf("Expected %v got %v", expected, got.FriendlyDeviceName)
	}
}

func TestApps(t *testing.T) {

	// RoundTrip test data
	tc := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(helperLoadBytes(t, "apps.xml"))),
		}, nil
	})}

	c, err := roku.NewClient("http://foo:8060", roku.WithHttpClient(tc))
	if err != nil {
		t.Errorf("Unable to create client: %v", err)
	}

	a, err := c.Apps()
	if err != nil {
		t.Errorf("Unable to get apps: %v", err)
	}

	// table driven tests
	var nameTests = []struct {
		text  string
		expID int
		expOK bool
	}{
		{"FloSports", 112048, true},
		{"HBO Max", 61322, true},
		{"Netflix", 12, true},
		{"BBC America", 168243, true},
		{"", 0, false},
		{"Foo", 0, false},
	}

	for _, tst := range nameTests {

		res, ok := a.FindByName(tst.text)
		if res.ID != tst.expID {
			t.Errorf(`FindByText("%v") = %v, expected %v`,
				tst.text, res.ID, tst.expID)
		}
		if ok != tst.expOK {
			t.Errorf(`FindByText("%v") = %v, expected %v`,
				tst.text, ok, tst.expOK)
		}
	}

	var idTests = []struct {
		id      int
		expText string
		expOK   bool
	}{
		{19977, "Spotify Music", true},
		{13, "Prime Video", true},
		{999, "", false},
	}

	for _, tst := range idTests {

		res, ok := a.FindByID(tst.id)
		if res.Text != tst.expText {
			t.Errorf(`FindByID("%v") = %v, expected %v`,
				tst.id, res.Text, tst.expText)
		}
		if ok != tst.expOK {
			t.Errorf(`FindByID("%v") = %v, expected %v`,
				tst.id, ok, tst.expOK)
		}
	}

}

func TestActiveApp(t *testing.T) {

	// RoundTrip test data
	tc := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(helperLoadBytes(t, "apps-active.xml"))),
		}, nil
	})}

	c, err := roku.NewClient("http://foo:8060", roku.WithHttpClient(tc))
	if err != nil {
		t.Errorf("Unable to create client: %v", err)
	}

	active, err := c.ActiveApp()
	if err != nil {
		t.Errorf("Unable to get apps: %v", err)
	}

	expected := "Roku"
	got := active.App
	if got != expected {
		t.Errorf("Expected %v got %v", expected, got)
	}

	expected = "City Stroll: Movie Magic"
	got = active.Screensaver.Text
	if got != expected {
		t.Errorf("Expected %v got %v", expected, got)
	}

	expId := 55545
	gotID := active.Screensaver.ID
	if gotID != expId {
		t.Errorf("Expected %v got %v", expId, gotID)
	}
}

func TestLaunch(t *testing.T) {

	// RoundTrip test data
	tc200 := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})}

	c, err := roku.NewClient("http://foo:8060", roku.WithHttpClient(tc200))
	if err != nil {
		t.Error("Unable to create client")
	}

	// good request
	got := c.Launch("837", "yk8yvt5lWVc", "")
	if got != nil {
		t.Error("Unable to launch channel")
	}

	// bad request
	tc404 := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
		}, nil
	})}

	c, err = roku.NewClient("http://foo:8060", roku.WithHttpClient(tc404))
	if err != nil {
		t.Error("Unable to create client")
	}

	got = c.Launch("999", "yk8yvt5lWVc", "")
	if got == nil {
		t.Error("Expected 404")
	}
}

func helperLoadBytes(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}
