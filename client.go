// Package roku provides a Roku External Control Protocol Client
package roku

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var endpointMap = map[string]string{
	"info":   "/query/device-info",
	"apps":   "/query/apps",
	"active": "/query/active-app",
	"launch": "/launch/%s?contentId=%s&mediaType=%s",
}

type client struct {
	BaseURL    *url.URL
	httpClient *http.Client
}

// ClientOption is a type that can be passed to the NewClient constructor for configuration
type ClientOption func(*client)

// NewClient constructs a client using the address argument
// ClientOptions can be passed in if required
func NewClient(addr string, opts ...ClientOption) (*client, error) {
	if addr == "" {
		return &client{}, errors.New("url is required")
	}

	u, err := url.Parse(addr)
	if err != nil {
		return &client{}, err
	}

	c := &client{
		BaseURL:    u,
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// WithHttpClient is a ClientOption that swaps out the default http.Client
// This is useful for testing/customising the client
func WithHttpClient(hc *http.Client) ClientOption {
	return func(c *client) {
		c.httpClient = hc
	}
}

func (c *client) get(path string) (*http.Response, error) {

	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	resp, err := c.httpClient.Get(u.String())
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeviceInfo returns info on the device
func (c *client) DeviceInfo() (*DeviceInfo, error) {

	resp, err := c.get(endpointMap["info"])
	if err != nil {
		return &DeviceInfo{}, err
	}

	decoder := xml.NewDecoder(resp.Body)
	info := &DeviceInfo{}
	err = decoder.Decode(&info)
	if err != nil {
		return info, err
	}

	return info, err
}

// Apps returns information on the apps installed on the device
func (c *client) Apps() (*Apps, error) {

	resp, err := c.get(endpointMap["apps"])
	if err != nil {
		return &Apps{}, err
	}

	decoder := xml.NewDecoder(resp.Body)
	apps := &Apps{}
	err = decoder.Decode(&apps)
	if err != nil {
		return apps, err
	}

	return apps, err
}

// ActiveApp returns the current app in use on the device
func (c *client) ActiveApp() (*ActiveApp, error) {

	resp, err := c.get(endpointMap["active"])
	if err != nil {
		return &ActiveApp{}, err
	}

	decoder := xml.NewDecoder(resp.Body)
	active := &ActiveApp{}
	err = decoder.Decode(&active)
	if err != nil {
		return active, err
	}

	return active, err
}

// Launch a channel on the device
func (c *client) Launch(channelID string, contentID string, mediaType string) error {
	rel, err := c.BaseURL.Parse(fmt.Sprintf(endpointMap["launch"], channelID, contentID, mediaType))
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Post(rel.String(), "", nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("unable to launch channel %v", resp.StatusCode)
	}

	return nil
}

func printBody(resp *http.Response) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	println(bodyString)
}
