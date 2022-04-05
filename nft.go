package desoapi

import (
	"fmt"
	"github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) AcceptNFTBid(payload *routes.AcceptNFTBidRequest) (*routes.AcceptNFTBidResponse, error) {
	data := new(routes.AcceptNFTBidResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/accept-nft-bid", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) AcceptNFTTransfer(payload *routes.AcceptNFTTransferRequest) (*routes.AcceptNFTTransferResponse, error) {
	data := new(routes.AcceptNFTTransferResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/accept-nft-transfer", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) BurnNFT(payload *routes.BurnNFTRequest) (*routes.BurnNFTResponse, error) {
	data := new(routes.BurnNFTResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/burn-nft", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CreateNFT(payload *routes.CreateNFTRequest) (*routes.CreateNFTResponse, error) {
	data := new(routes.CreateNFTResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/create-nft", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CreateNFTBid(payload *routes.CreateNFTBidRequest) (*routes.CreateNFTBidResponse, error) {
	data := new(routes.CreateNFTBidResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/create-nft-bid", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetAcceptedBidHistory(postHashHex string) (*routes.GetAcceptedBidHistoryResponse, error) {
	uri := fmt.Sprintf("/api/v0/accepted-bid-history/%s", postHashHex)
	data := new(routes.GetAcceptedBidHistoryResponse)

	err := c.executeRequest(fasthttp.MethodGet, uri, nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNextNFTShowcase(payload *routes.GetNextNFTShowcaseRequest) (*routes.GetNextNFTShowcaseResponse, error) {
	data := new(routes.GetNextNFTShowcaseResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-next-nft-showcase", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTBidsForNFTPost(payload *routes.GetNFTBidsForNFTPostRequest) (*routes.GetNFTBidsForNFTPostResponse, error) {
	data := new(routes.GetNFTBidsForNFTPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nft-bids-for-nft-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTBidsForUser(payload *routes.GetNFTBidsForUserRequest) (*routes.GetNFTBidsForUserResponse, error) {
	data := new(routes.GetNFTBidsForUserResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nft-bids-for-user", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTCollectionSummary(payload *routes.GetNFTCollectionSummaryRequest) (*routes.GetNFTCollectionSummaryResponse, error) {
	data := new(routes.GetNFTCollectionSummaryResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nft-collection-summary", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTEntriesForPostHash(payload *routes.GetNFTEntriesForPostHashRequest) (*routes.GetNFTEntriesForPostHashResponse, error) {
	data := new(routes.GetNFTEntriesForPostHashResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nft-entries-for-nft-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTsCreatedByPublicKey(payload *routes.GetNFTsCreatedByPublicKeyRequest) (*routes.GetNFTsCreatedByPublicKeyResponse, error) {
	data := new(routes.GetNFTsCreatedByPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nfts-created-by-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTsForUser(payload *routes.GetNFTsForUserRequest) (*routes.GetNFTsForUserResponse, error) {
	data := new(routes.GetNFTsForUserResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nfts-for-user", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTShowcase(payload *routes.GetNFTShowcaseRequest) (*routes.GetNFTShowcaseResponse, error) {
	data := new(routes.GetNFTShowcaseResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nft-showcase", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) TransferNFT(payload *routes.TransferNFTRequest) (*routes.TransferNFTResponse, error) {
	data := new(routes.TransferNFTResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/transfer-nft", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) UpdateNFT(payload *routes.UpdateNFTRequest) (*routes.UpdateNFTResponse, error) {
	data := new(routes.UpdateNFTResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/update-nft", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
