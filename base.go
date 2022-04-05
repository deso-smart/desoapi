package desoapi

import (
	"github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetAppState(payload *routes.GetAppStateRequest) (*routes.GetAppStateResponse, error) {
	data := new(routes.GetAppStateResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-app-state", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetExchangeRate() (*routes.GetExchangeRateResponse, error) {
	data := new(routes.GetExchangeRateResponse)

	err := c.executeRequest(fasthttp.MethodGet, "/api/v0/get-exchange-rate", nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) HealthCheck() error {
	err := c.executeRequest(fasthttp.MethodGet, "/api/v0/health-check", nil, nil)
	if err != nil {
		return err
	}

	return nil
}
