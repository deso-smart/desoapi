package desoapi

import (
	"github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) QueryETHRPC(payload *routes.QueryETHRPCRequest) (*routes.InfuraResponse, error) {
	data := new(routes.InfuraResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/query-eth-rpc", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SubmitETHTx(payload *routes.SubmitETHTxRequest) (*routes.SubmitETHTxResponse, error) {
	data := new(routes.SubmitETHTxResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/submit-eth-tx", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
