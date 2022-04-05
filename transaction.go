package desoapi

import (
	"github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) AppendExtraData(payload *routes.AppendExtraDataRequest) (*routes.AppendExtraDataResponse, error) {
	data := new(routes.AppendExtraDataResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/append-extra-data", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) AuthorizeDerivedKey(payload *routes.AuthorizeDerivedKeyRequest) (*routes.AuthorizeDerivedKeyResponse, error) {
	data := new(routes.AuthorizeDerivedKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/authorize-derived-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) BuyOrSellCreatorCoin(payload *routes.BuyOrSellCreatorCoinRequest) (*routes.BuyOrSellCreatorCoinResponse, error) {
	data := new(routes.BuyOrSellCreatorCoinResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/buy-or-sell-creator-coin", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CreateFollowTxnStateless(payload *routes.CreateFollowTxnStatelessRequest) (*routes.CreateFollowTxnStatelessResponse, error) {
	data := new(routes.CreateFollowTxnStatelessResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/create-follow-txn-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CreateLikeStateless(payload *routes.CreateLikeStatelessRequest) (*routes.CreateLikeStatelessResponse, error) {
	data := new(routes.CreateLikeStatelessResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/create-like-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) DAOCoin(payload *routes.DAOCoinRequest) (*routes.DAOCoinResponse, error) {
	data := new(routes.DAOCoinResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/dao-coin", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) ExchangeBitcoin(payload *routes.ExchangeBitcoinRequest) (*routes.ExchangeBitcoinResponse, error) {
	data := new(routes.ExchangeBitcoinResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/exchange-bitcoin", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetTransactionSpending(payload *routes.GetTransactionSpendingRequest) (*routes.GetTransactionSpendingResponse, error) {
	data := new(routes.GetTransactionSpendingResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-transaction-spending", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetTxn(payload *routes.GetTxnRequest) (*routes.GetTxnResponse, error) {
	data := new(routes.GetTxnResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-txn", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SendDeSo(payload *routes.SendDeSoRequest) (*routes.SendDeSoResponse, error) {
	data := new(routes.SendDeSoResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/send-deso", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SendDiamonds(payload *routes.SendDiamondsRequest) (*routes.SendDiamondsResponse, error) {
	data := new(routes.SendDiamondsResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/send-diamonds", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SubmitPost(payload *routes.SubmitPostRequest) (*routes.SubmitPostResponse, error) {
	data := new(routes.SubmitPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/submit-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SubmitTransaction(payload *routes.SubmitTransactionRequest) (*routes.SubmitTransactionResponse, error) {
	data := new(routes.SubmitTransactionResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/submit-transaction", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) TransferCreatorCoin(payload *routes.TransferCreatorCoinRequest) (*routes.TransferCreatorCoinResponse, error) {
	data := new(routes.TransferCreatorCoinResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/transfer-creator-coin", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) TransferDAOCoin(payload *routes.TransferDAOCoinRequest) (*routes.TransferDAOCoinResponse, error) {
	data := new(routes.TransferDAOCoinResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/transfer-dao-coin", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) UpdateProfile(payload *routes.UpdateProfileRequest) (*routes.UpdateProfileResponse, error) {
	data := new(routes.UpdateProfileResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/update-profile", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
