//
// Copyright (c) 2020 IPPANEL SMS
// All rights reserved.
//
// Author: ippanel <dev@ippanel.com>

// Package ippanel is an official library for working with ippanel sms api.
// brief documentation for ippanel sms api provided at http://docs.ippanel.com
package ippanel

import (
	"net/http"
	"net/url"
	"time"
)

const (
	// ClientVersion is used in User-Agent request header to provide server with API level.
	ClientVersion = "1.0.1"

	// Endpoint points you to Ippanel REST API.
	Endpoint = "http://rest.ippanel.com/v1"

	// httpClientTimeout is used to limit http.Client waiting time.
	httpClientTimeout = 30 * time.Second
)

// Ippanel ...
type Ippanel struct {
	AccessKey string
	Client    *http.Client
	BaseURL   *url.URL
}

// New create new ippanel sms instance
func New(accesskey string) *Ippanel {
	u, _ := url.Parse(Endpoint)
	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   httpClientTimeout,
	}

	return &Ippanel{
		AccessKey: accesskey,
		Client:    client,
		BaseURL:   u,
	}
}
