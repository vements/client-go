/*
Copyright 2023 Monster Street Systems LLC

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// NB: This is a generated file; any changes will be lost.

package vements

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Achievement AchievementEndpoint
	Participant ParticipantEndpoint
	Scoreboard  ScoreboardEndpoint
	BaseUrl     string
	Config      *Config
}

func NewClient(apiKey string, tags []string, config *Config) (*Client, bool) {
	if url, ok := config.ServerUrl(tags); ok {
		return NewClientFromConfig(apiKey, url, config), true
	}
	return nil, false
}

func NewClientFromConfig(apiKey string, baseUrl string, config *Config) *Client {
	return &Client{
		Achievement: AchievementEndpoint{
			ApiKey:  apiKey,
			BaseUrl: baseUrl,
			Rc:      resty.New(),
		},
		Participant: ParticipantEndpoint{
			ApiKey:  apiKey,
			BaseUrl: baseUrl,
			Rc:      resty.New(),
		},
		Scoreboard: ScoreboardEndpoint{
			ApiKey:  apiKey,
			BaseUrl: baseUrl,
			Rc:      resty.New(),
		},
		BaseUrl: baseUrl,
		Config:  config,
	}
}

