package client

import (
	"net/http"
	"time"
)

type Client struct {
	APIURL     string
	APIVersion string
	Username   string
	Password   string
	Tenant     string
	PageSize   int

	Client *http.Client
}

func defaultOrString(d, v string) string {
	if v != "" {
		return v
	}
	return d
}

func defaultOrInt(d, v int) int {
	if v != 0 {
		return v
	}
	return d
}

func NewClient(c Config) *Client {
	return &Client{
		Username:   c.Username,
		Password:   c.Password,
		Tenant:     c.Tenant,
		APIURL:     defaultOrString(DefaultAPIURL, c.APIURL),
		APIVersion: defaultOrString(DefaultAPIVersion, c.APIVersion),
		PageSize:   defaultOrInt(DefaultPageSize, c.PageSize),
		Client:     &http.Client{Timeout: time.Duration(defaultOrInt(DefaultTimeout, c.Timeout)) * time.Second},
	}
}