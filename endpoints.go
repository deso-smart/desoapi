package desoapi

import (
	"github.com/valyala/fasthttp"
)

func (c *Client) AuthorizeDerivedKey(payload *AuthorizeDerivedKeyRequest) (*AuthorizeDerivedKeyResponse, error) {
	data := new(AuthorizeDerivedKeyResponse)

	err := c.Request(fasthttp.MethodPost, "/api/v0/authorize-derived-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) Base() (*BaseResponse, error) {
	data := new(BaseResponse)

	err := c.Request(fasthttp.MethodGet, "/api/v1", nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetExchangeRate() (*GetExchangeRateResponse, error) {
	data := new(GetExchangeRateResponse)

	err := c.Request(fasthttp.MethodGet, "/api/v0/get-exchange-rate", nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetSingleProfile(payload *GetSingleProfileRequest) (*GetSingleProfileResponse, error) {
	data := new(GetSingleProfileResponse)

	err := c.Request(fasthttp.MethodPost, "/api/v0/get-single-profile", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SubmitTransaction(payload *SubmitTransactionRequest) (*SubmitTransactionResponse, error) {
	data := new(SubmitTransactionResponse)

	err := c.Request(fasthttp.MethodPost, "/api/v0/submit-transaction", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) TransactionInfo(payload *TransactionInfoRequest) (*TransactionInfoResponse, error) {
	data := new(TransactionInfoResponse)

	err := c.Request(fasthttp.MethodPost, "/api/v1/transaction-info", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
