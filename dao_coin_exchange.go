package desoapi

import (
	desoRoutes "github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetDAOCoinLimitOrders(payload *desoRoutes.GetDAOCoinLimitOrdersRequest) (*desoRoutes.GetDAOCoinLimitOrdersResponse, error) {
	data := new(desoRoutes.GetDAOCoinLimitOrdersResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-dao-coin-limit-orders", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetTransactorDAOCoinLimitOrders(payload *desoRoutes.GetTransactorDAOCoinLimitOrdersRequest) (*desoRoutes.GetDAOCoinLimitOrdersResponse, error) {
	data := new(desoRoutes.GetDAOCoinLimitOrdersResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-transactor-dao-coin-limit-orders", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
