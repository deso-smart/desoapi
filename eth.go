package desoapi

import (
	desoRoutes "github.com/deso-smart/deso-backend/v3/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) QueryETHRPC(payload *desoRoutes.QueryETHRPCRequest) (*desoRoutes.InfuraResponse, error) {
	data := new(desoRoutes.InfuraResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/query-eth-rpc", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SubmitETHTx(payload *desoRoutes.SubmitETHTxRequest) (*desoRoutes.SubmitETHTxResponse, error) {
	data := new(desoRoutes.SubmitETHTxResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/submit-eth-tx", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) MetamaskSignIn(payload *desoRoutes.MetamaskSignInRequest) (*desoRoutes.MetamaskSignInResponse, error) {
	data := new(desoRoutes.MetamaskSignInResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/send-starter-deso-for-metamask-account", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
