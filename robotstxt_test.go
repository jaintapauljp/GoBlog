package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_robotsTXT(t *testing.T) {

	app := &goBlog{
		cfg: &config{
			Server: &configServer{
				PublicAddress: "https://example.com",
			},
		},
	}

	h := http.HandlerFunc(app.serveRobotsTXT)
	assert.HTTPStatusCode(t, h, http.MethodGet, "", nil, 200)
	txt := assert.HTTPBody(h, http.MethodGet, "", nil)
	assert.Equal(t, "User-agent: *\nSitemap: https://example.com/sitemap.xml", txt)

	app.cfg.PrivateMode = &configPrivateMode{
		Enabled: true,
	}

	h = http.HandlerFunc(app.serveRobotsTXT)
	assert.HTTPStatusCode(t, h, http.MethodGet, "", nil, 200)
	txt = assert.HTTPBody(h, http.MethodGet, "", nil)
	assert.Equal(t, "User-agent: *\nDisallow: /", txt)

}
