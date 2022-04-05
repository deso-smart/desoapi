package desoapi

import (
	"github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) Balance(payload *routes.APIBalanceRequest) (*routes.APIBalanceResponse, error) {
	data := new(routes.APIBalanceResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v1/balance", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) Base() (*routes.APIBaseResponse, error) {
	data := new(routes.APIBaseResponse)

	err := c.executeRequest(fasthttp.MethodGet, "/api/v1", nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) Block(payload *routes.APIBlockRequest) (*routes.APIBlockResponse, error) {
	data := new(routes.APIBlockResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v1/block", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) KeyPair(payload *routes.APIKeyPairRequest) (*routes.APIKeyPairResponse, error) {
	data := new(routes.APIKeyPairResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v1/key-pair", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) TransactionInfo(payload *routes.APITransactionInfoRequest) (*routes.APITransactionInfoResponse, error) {
	data := new(routes.APITransactionInfoResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v1/transaction-info", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) TransferDeSo(payload *routes.APITransferDeSoRequest) (*routes.APITransferDeSoResponse, error) {
	data := new(routes.APITransferDeSoResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v1/transfer-deso", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
