package desoapi

import (
	desoRoutes "github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetAppState(payload *desoRoutes.GetAppStateRequest) (*desoRoutes.GetAppStateResponse, error) {
	data := new(desoRoutes.GetAppStateResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-app-state", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetExchangeRate() (*desoRoutes.GetExchangeRateResponse, error) {
	data := new(desoRoutes.GetExchangeRateResponse)

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
