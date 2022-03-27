package desoapi

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"net/url"
)

const (
	version   = "0.0.0"
	userAgent = "desoapi/" + version
	BaseUri   = "https://node.desosmart.com"
)

type Client struct {
	baseUri   *url.URL
	http      *fasthttp.Client
	json      jsoniter.API
	userAgent string
}

func NewClient(baseUri string) (*Client, error) {
	parsedBaseUri, err := url.Parse(baseUri)
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseUri:   parsedBaseUri,
		http:      &fasthttp.Client{},
		json:      jsoniter.ConfigCompatibleWithStandardLibrary,
		userAgent: userAgent,
	}

	return c, nil
}

type errorResponse struct {
	Error string
}

func (c *Client) Request(method string, uri string, body interface{}, data interface{}) error {
	resolvedUri, err := c.baseUri.Parse(uri)
	if err != nil {
		return err
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(resolvedUri.String())
	req.Header.SetMethod(method)
	req.Header.SetUserAgent(userAgent)
	req.Header.Set("Accept", "application/json")

	if body != nil {
		buf, err := c.json.Marshal(body)
		if err != nil {
			return err
		}
		req.Header.SetContentType("application/json")
		req.SetBody(buf)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := c.http.Do(req, resp); err != nil {
		return err
	}

	// or status >= 200 && status <= 204
	if resp.StatusCode() < fasthttp.StatusOK || resp.StatusCode() >= fasthttp.StatusBadRequest {
		errResp := new(errorResponse)
		if err := c.json.Unmarshal(resp.Body(), errResp); err != nil {
			return err
		}

		if errResp.Error != "" {
			return errors.New(errResp.Error)
		}

		return errors.New(string(resp.Body()))
	}

	if data != nil {
		if err := c.json.Unmarshal(resp.Body(), data); err != nil {
			return err
		}
	}

	return nil
}
