package desoapi

import (
	"fmt"
	desoRoutes "github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) AcceptNFTBid(payload *desoRoutes.AcceptNFTBidRequest) (*desoRoutes.AcceptNFTBidResponse, error) {
	data := new(desoRoutes.AcceptNFTBidResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/accept-nft-bid", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) AcceptNFTTransfer(payload *desoRoutes.AcceptNFTTransferRequest) (*desoRoutes.AcceptNFTTransferResponse, error) {
	data := new(desoRoutes.AcceptNFTTransferResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/accept-nft-transfer", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) BurnNFT(payload *desoRoutes.BurnNFTRequest) (*desoRoutes.BurnNFTResponse, error) {
	data := new(desoRoutes.BurnNFTResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/burn-nft", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CreateNFT(payload *desoRoutes.CreateNFTRequest) (*desoRoutes.CreateNFTResponse, error) {
	data := new(desoRoutes.CreateNFTResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/create-nft", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) CreateNFTBid(payload *desoRoutes.CreateNFTBidRequest) (*desoRoutes.CreateNFTBidResponse, error) {
	data := new(desoRoutes.CreateNFTBidResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/create-nft-bid", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetAcceptedBidHistory(postHashHex string) (*desoRoutes.GetAcceptedBidHistoryResponse, error) {
	uri := fmt.Sprintf("/api/v0/accepted-bid-history/%s", postHashHex)
	data := new(desoRoutes.GetAcceptedBidHistoryResponse)

	err := c.executeRequest(fasthttp.MethodGet, uri, nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNextNFTShowcase(payload *desoRoutes.GetNextNFTShowcaseRequest) (*desoRoutes.GetNextNFTShowcaseResponse, error) {
	data := new(desoRoutes.GetNextNFTShowcaseResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-next-nft-showcase", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTBidsForNFTPost(payload *desoRoutes.GetNFTBidsForNFTPostRequest) (*desoRoutes.GetNFTBidsForNFTPostResponse, error) {
	data := new(desoRoutes.GetNFTBidsForNFTPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nft-bids-for-nft-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTBidsForUser(payload *desoRoutes.GetNFTBidsForUserRequest) (*desoRoutes.GetNFTBidsForUserResponse, error) {
	data := new(desoRoutes.GetNFTBidsForUserResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nft-bids-for-user", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTCollectionSummary(payload *desoRoutes.GetNFTCollectionSummaryRequest) (*desoRoutes.GetNFTCollectionSummaryResponse, error) {
	data := new(desoRoutes.GetNFTCollectionSummaryResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nft-collection-summary", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTEntriesForPostHash(payload *desoRoutes.GetNFTEntriesForPostHashRequest) (*desoRoutes.GetNFTEntriesForPostHashResponse, error) {
	data := new(desoRoutes.GetNFTEntriesForPostHashResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nft-entries-for-nft-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTsCreatedByPublicKey(payload *desoRoutes.GetNFTsCreatedByPublicKeyRequest) (*desoRoutes.GetNFTsCreatedByPublicKeyResponse, error) {
	data := new(desoRoutes.GetNFTsCreatedByPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nfts-created-by-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTsForUser(payload *desoRoutes.GetNFTsForUserRequest) (*desoRoutes.GetNFTsForUserResponse, error) {
	data := new(desoRoutes.GetNFTsForUserResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nfts-for-user", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNFTShowcase(payload *desoRoutes.GetNFTShowcaseRequest) (*desoRoutes.GetNFTShowcaseResponse, error) {
	data := new(desoRoutes.GetNFTShowcaseResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-nft-showcase", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) TransferNFT(payload *desoRoutes.TransferNFTRequest) (*desoRoutes.TransferNFTResponse, error) {
	data := new(desoRoutes.TransferNFTResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/transfer-nft", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) UpdateNFT(payload *desoRoutes.UpdateNFTRequest) (*desoRoutes.UpdateNFTResponse, error) {
	data := new(desoRoutes.UpdateNFTResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/update-nft", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
