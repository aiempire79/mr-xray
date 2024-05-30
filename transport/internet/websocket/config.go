package websocket

import (
	"net/http"
	"time"
	"unicode"
	"math/rand"

	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/transport/internet"
)

const protocolName = "websocket"

func (c *Config) GetNormalizedPath() string {
	path := c.Path
	if path == "" {
		return "/"
	}
	if path[0] != '/' {
		return "/" + path
	}
	return path
}

func (c *Config) GetRequestHeader() http.Header {
	header := http.Header{}
	for k, v := range c.Header {
		header.Add(k, v)
	}
	randomizedHost := randomizeCase(c.Host)
	header.Set("Host", randomizedHost)
	header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 11; SM-G981B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.77 Mobile Safari/537.36")
	return header
}

func randomizeCase(s string) string {
	rand.Seed(time.Now().UnixNano())
	runes := []rune(s)
	for i, r := range runes {
		if rand.Intn(2) == 0 {
			runes[i] = unicode.ToLower(r)
		} else {
			runes[i] = unicode.ToUpper(r)
		}
	}
	return string(runes)
}

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
