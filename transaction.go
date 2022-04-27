package desoapi

import (
	desoRoutes "github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) AppendExtraData(payload *desoRoutes.AppendExtraDataRequest) (*desoRoutes.AppendExtraDataResponse, error) {
	data := new(desoRoutes.AppendExtraDataResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/append-extra-data", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) AuthorizeDerivedKey(payload *desoRoutes.AuthorizeDerivedKeyRequest) (*desoRoutes.AuthorizeDerivedKeyResponse, error) {
	data := new(desoRoutes.AuthorizeDerivedKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/authorize-derived-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) BuyOrSellCreatorCoin(payload *desoRoutes.BuyOrSellCreatorCoinRequest) (*desoRoutes.BuyOrSellCreatorCoinResponse, error) {
	data := new(desoRoutes.BuyOrSellCreatorCoinResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/buy-or-sell-creator-coin", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CreateFollowTxnStateless(payload *desoRoutes.CreateFollowTxnStatelessRequest) (*desoRoutes.CreateFollowTxnStatelessResponse, error) {
	data := new(desoRoutes.CreateFollowTxnStatelessResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/create-follow-txn-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CreateLikeStateless(payload *desoRoutes.CreateLikeStatelessRequest) (*desoRoutes.CreateLikeStatelessResponse, error) {
	data := new(desoRoutes.CreateLikeStatelessResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/create-like-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) DAOCoin(payload *desoRoutes.DAOCoinRequest) (*desoRoutes.DAOCoinResponse, error) {
	data := new(desoRoutes.DAOCoinResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/dao-coin", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) ExchangeBitcoin(payload *desoRoutes.ExchangeBitcoinRequest) (*desoRoutes.ExchangeBitcoinResponse, error) {
	data := new(desoRoutes.ExchangeBitcoinResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/exchange-bitcoin", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetTransactionSpending(payload *desoRoutes.GetTransactionSpendingRequest) (*desoRoutes.GetTransactionSpendingResponse, error) {
	data := new(desoRoutes.GetTransactionSpendingResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-transaction-spending", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetTxn(payload *desoRoutes.GetTxnRequest) (*desoRoutes.GetTxnResponse, error) {
	data := new(desoRoutes.GetTxnResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-txn", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SendDeSo(payload *desoRoutes.SendDeSoRequest) (*desoRoutes.SendDeSoResponse, error) {
	data := new(desoRoutes.SendDeSoResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/send-deso", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SendDiamonds(payload *desoRoutes.SendDiamondsRequest) (*desoRoutes.SendDiamondsResponse, error) {
	data := new(desoRoutes.SendDiamondsResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/send-diamonds", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SubmitPost(payload *desoRoutes.SubmitPostRequest) (*desoRoutes.SubmitPostResponse, error) {
	data := new(desoRoutes.SubmitPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/submit-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SubmitTransaction(payload *desoRoutes.SubmitTransactionRequest) (*desoRoutes.SubmitTransactionResponse, error) {
	data := new(desoRoutes.SubmitTransactionResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/submit-transaction", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) TransferCreatorCoin(payload *desoRoutes.TransferCreatorCoinRequest) (*desoRoutes.TransferCreatorCoinResponse, error) {
	data := new(desoRoutes.TransferCreatorCoinResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/transfer-creator-coin", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) TransferDAOCoin(payload *desoRoutes.TransferDAOCoinRequest) (*desoRoutes.TransferDAOCoinResponse, error) {
	data := new(desoRoutes.TransferDAOCoinResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/transfer-dao-coin", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) UpdateProfile(payload *desoRoutes.UpdateProfileRequest) (*desoRoutes.UpdateProfileResponse, error) {
	data := new(desoRoutes.UpdateProfileResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/update-profile", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CreateDAOCoinLimitOrder(payload *desoRoutes.DAOCoinLimitOrderWithExchangeRateAndQuantityRequest) (*desoRoutes.DAOCoinLimitOrderResponse, error) {
	data := new(desoRoutes.DAOCoinLimitOrderResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/create-dao-coin-limit-order", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CancelDAOCoinLimitOrder(payload *desoRoutes.DAOCoinLimitOrderWithCancelOrderIDRequest) (*desoRoutes.DAOCoinLimitOrderResponse, error) {
	data := new(desoRoutes.DAOCoinLimitOrderResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/cancel-dao-coin-limit-order", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CreateDAOCoinMarketOrder(payload *desoRoutes.DAOCoinMarketOrderWithQuantityRequest) (*desoRoutes.DAOCoinLimitOrderResponse, error) {
	data := new(desoRoutes.DAOCoinLimitOrderResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/create-dao-coin-market-order", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
