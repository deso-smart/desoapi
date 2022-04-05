package desoapi

import (
	desoRoutes "github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) Balance(payload *desoRoutes.APIBalanceRequest) (*desoRoutes.APIBalanceResponse, error) {
	data := new(desoRoutes.APIBalanceResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v1/balance", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) Base() (*desoRoutes.APIBaseResponse, error) {
	data := new(desoRoutes.APIBaseResponse)

	err := c.executeRequest(fasthttp.MethodGet, "/api/v1", nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) Block(payload *desoRoutes.APIBlockRequest) (*desoRoutes.APIBlockResponse, error) {
	data := new(desoRoutes.APIBlockResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v1/block", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) KeyPair(payload *desoRoutes.APIKeyPairRequest) (*desoRoutes.APIKeyPairResponse, error) {
	data := new(desoRoutes.APIKeyPairResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v1/key-pair", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) TransactionInfo(payload *desoRoutes.APITransactionInfoRequest) (*desoRoutes.APITransactionInfoResponse, error) {
	data := new(desoRoutes.APITransactionInfoResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v1/transaction-info", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) TransferDeSo(payload *desoRoutes.APITransferDeSoRequest) (*desoRoutes.APITransferDeSoResponse, error) {
	data := new(desoRoutes.APITransferDeSoResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v1/transfer-deso", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
