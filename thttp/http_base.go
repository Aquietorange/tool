package thttp

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var defaultClientAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/517.16 (KHTML, like Gecko) Chrome/85.0.1986.12 Safari/517.16"

// Response is the struct for client request response.
type Response struct {
	*http.Response
	request     *http.Request
	requestBody []byte
	cookies     map[string]string
}

// ReadAll retrieves and returns the response content as []byte.
func (r *Response) ReadAll() []byte {
	// Response might be nil.
	if r == nil || r.Response == nil {
		return []byte{}
	}
	body, err := ioutil.ReadAll(r.Response.Body)
	if err != nil {
		return nil
	}
	return body
}

// Close closes the response when it will never be used.
func (r *Response) Close() error {
	if r == nil || r.Response == nil || r.Response.Close {
		return nil
	}
	r.Response.Close = true
	return r.Response.Body.Close()
}

type HandlerFunc = func(c *Client, r *http.Request) (*Response, error)

// Client is the HTTP client for HTTP request management.
type Client struct {
	http.Client                 // Underlying HTTP Client.
	ctx         context.Context // Context for each request.
	//dump              bool              // Mark this request will be dumped.
	parent            *Client           // Parent http client, this is used for chaining operations.
	header            map[string]string // Custom header map.
	cookies           map[string]string // Custom cookie map.
	prefix            string            // Prefix for request.
	authUser          string            // HTTP basic authentication: user.
	authPass          string            // HTTP basic authentication: pass.
	retryCount        int               // Retry count when request fails.
	retryInterval     time.Duration     // Retry interval when request fails.
	middlewareHandler []HandlerFunc     // Interceptor handlers
}

// GetBytes sends a GET request, retrieves and returns the result content as bytes.
func (c *Client) GetBytes(url string, data ...string) []byte {
	if len(data) > 0 && data[0] != "" {
		return c.RequestBytes("GET", url, strings.NewReader(data[0]))
	} else {
		return c.RequestBytes("GET", url, nil)
	}
}

// RequestBytes sends request using given HTTP method and data, retrieves returns the result
// as bytes. It reads and closes the response object internally automatically.
func (c *Client) RequestBytes(method string, url string, body io.Reader) []byte {
	response, err := c.DoRequest(method, url, body)
	if err != nil {
		return nil
	}
	defer response.Close()
	return response.ReadAll()
}

func (c *Client) DoRequest(method, url string, body io.Reader) (resp *Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Custom header.
	if len(c.header) > 0 {
		for k, v := range c.header {
			req.Header.Set(k, v)
		}
	}
	// It's necessary set the req.Host if you want to custom the host value of the request.
	// It uses the "Host" value from header if it's not empty.
	if host := req.Header.Get("Host"); host != "" {
		req.Host = host
	}

	// Custom Cookie.
	if len(c.cookies) > 0 {
		headerCookie := ""
		for k, v := range c.cookies {
			if len(headerCookie) > 0 {
				headerCookie += ";"
			}
			headerCookie += k + "=" + v
		}
		if len(headerCookie) > 0 {
			req.Header.Set("Cookie", headerCookie)
		}
	}

	resp, err = c.callRequest(req)

	return resp, err
}

// callRequest sends request with give http.Request, and returns the responses object.
// Note that the response object MUST be closed if it'll be never used.
func (c *Client) callRequest(req *http.Request) (resp *Response, err error) {
	resp = &Response{
		request: req,
	}

	for {
		if resp.Response, err = c.Do(req); err != nil {
			// The response might not be nil when err != nil.
			if resp.Response != nil {
				if err := resp.Response.Body.Close(); err != nil {
					fmt.Errorf(`%+v`, err)
				}
			}
			if c.retryCount > 0 {
				c.retryCount--
				time.Sleep(c.retryInterval)
			} else {
				//return resp, err
				break
			}
		} else {
			break
		}
	}
	return resp, err
}

// PostContent is a convenience method for sending POST request, which retrieves and returns
// the result content and automatically closes response object.
func (c *Client) PostContent(url string, data ...string) string {
	if len(data) > 0 && data[0] != "" {
		return string(c.RequestBytes("POST", url, strings.NewReader(data[0])))
	} else {
		return string(c.RequestBytes("POST", url, nil))
	}
}

// PostBytes sends a POST request, retrieves and returns the result content as bytes.
func (c *Client) PostBytes(url string, body ...io.Reader) []byte {

	if len(body) > 0 && body[0] != nil {
		return c.RequestBytes("POST", url, body[0])
	} else {
		return c.RequestBytes("POST", url, nil)
	}

}

// New creates and returns a new HTTP client object.
func New() *Client {
	client := &Client{
		Client: http.Client{
			Transport: &http.Transport{
				// No validation for https certification of the server in default.
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				DisableKeepAlives: true,
			},
		},
		header:  make(map[string]string),
		cookies: make(map[string]string),
	}
	client.header["User-Agent"] = defaultClientAgent
	return client
}
