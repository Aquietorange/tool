package thttp

import (
	"context"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

// SetBrowserMode enables browser mode of the client.
// When browser mode is enabled, it automatically saves and sends cookie content
// from and to server.
func (c *Client) SetBrowserMode(enabled bool) *Client {
	if enabled {
		jar, _ := cookiejar.New(nil)
		c.Jar = jar
	}
	return c
}

// SetHeader sets a custom HTTP header pair for the client.
func (c *Client) SetHeader(key, value string) *Client {
	c.header[key] = value
	return c
}

// SetHeaderMap sets custom HTTP headers with map.
func (c *Client) SetHeaderMap(m map[string]string) *Client {
	for k, v := range m {
		c.header[k] = v
	}
	return c
}

// SetAgent sets the User-Agent header for client.
func (c *Client) SetAgent(agent string) *Client {
	c.header["User-Agent"] = agent
	return c
}

// SetContentType sets HTTP content type for the client.
func (c *Client) SetContentType(contentType string) *Client {
	c.header["Content-Type"] = contentType
	return c
}

// SetCookie sets a cookie pair for the client.
func (c *Client) SetCookie(key, value string) *Client {
	c.cookies[key] = value
	return c
}

// SetCookieMap sets cookie items with map.
func (c *Client) SetCookieMap(m map[string]string) *Client {
	for k, v := range m {
		c.cookies[k] = v
	}
	return c
}

// SetPrefix sets the request server URL prefix.
func (c *Client) SetPrefix(prefix string) *Client {
	c.prefix = prefix
	return c
}

// SetTimeout sets the request timeout for the client.
func (c *Client) SetTimeout(t time.Duration) *Client {
	c.Client.Timeout = t
	return c
}

// SetBasicAuth sets HTTP basic authentication information for the client.
func (c *Client) SetBasicAuth(user, pass string) *Client {
	c.authUser = user
	c.authPass = pass
	return c
}

// SetCtx sets context for each request of this client.
func (c *Client) SetCtx(ctx context.Context) *Client {
	c.ctx = ctx
	return c
}

// SetRetry sets retry count and interval.
func (c *Client) SetRetry(retryCount int, retryInterval time.Duration) *Client {
	c.retryCount = retryCount
	c.retryInterval = retryInterval
	return c
}

// SetRedirectLimit limit the number of jumps
func (c *Client) SetRedirectLimit(redirectLimit int) *Client {
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= redirectLimit {
			return http.ErrUseLastResponse
		}
		return nil
	}
	return c
}

// SetProxy set proxy for the client.
// This func will do nothing when the parameter `proxyURL` is empty or in wrong pattern.
// The correct pattern is like `http://USER:PASSWORD@IP:PORT` or `socks5://USER:PASSWORD@IP:PORT`.
// Only `http` and `socks5` proxies are supported currently.
func (c *Client) SetProxy(proxyURL string) {
	if strings.TrimSpace(proxyURL) == "" {
		return
	}
	_proxy, err := url.Parse(proxyURL)
	if err != nil {
		return
	}
	if _proxy.Scheme == "http" {
		if v, ok := c.Transport.(*http.Transport); ok {
			v.Proxy = http.ProxyURL(_proxy)
		}
	} else {
		var auth = &proxy.Auth{}
		user := _proxy.User.Username()

		if user != "" {
			auth.User = user
			password, hasPassword := _proxy.User.Password()
			if hasPassword && password != "" {
				auth.Password = password
			}
		} else {
			auth = nil
		}
		// refer to the source code, error is always nil
		dialer, err := proxy.SOCKS5(
			"tcp",
			_proxy.Host,
			auth,
			&net.Dialer{
				Timeout:   c.Client.Timeout,
				KeepAlive: c.Client.Timeout,
			},
		)
		if err != nil {
			return
		}
		if v, ok := c.Transport.(*http.Transport); ok {
			v.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
				return dialer.Dial(network, addr)
			}
		}
		//c.SetTimeout(10*time.Second)
	}
}
