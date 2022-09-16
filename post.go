package desoapi

import (
	desoRoutes "github.com/deso-smart/deso-backend/v3/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetDiamondedPosts(payload *desoRoutes.GetPostsDiamondedBySenderForReceiverRequest) (*desoRoutes.GetPostsDiamondedBySenderForReceiverResponse, error) {
	data := new(desoRoutes.GetPostsDiamondedBySenderForReceiverResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-diamonded-posts", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetDiamondsForPost(payload *desoRoutes.GetDiamondsForPostRequest) (*desoRoutes.GetDiamondsForPostResponse, error) {
	data := new(desoRoutes.GetDiamondsForPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-diamonds-for-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetLikesForPost(payload *desoRoutes.GetLikesForPostRequest) (*desoRoutes.GetLikesForPostResponse, error) {
	data := new(desoRoutes.GetLikesForPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-likes-for-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetPostsForPublicKey(payload *desoRoutes.GetPostsForPublicKeyRequest) (*desoRoutes.GetPostsForPublicKeyResponse, error) {
	data := new(desoRoutes.GetPostsForPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-posts-for-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetPostsStateless(payload *desoRoutes.GetPostsStatelessRequest) (*desoRoutes.GetPostsStatelessResponse, error) {
	data := new(desoRoutes.GetPostsStatelessResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-posts-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetQuoteRepostsForPost(payload *desoRoutes.GetQuoteRepostsForPostRequest) (*desoRoutes.GetQuoteRepostsForPostResponse, error) {
	data := new(desoRoutes.GetQuoteRepostsForPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-quote-reposts-for-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetRepostsForPost(payload *desoRoutes.GetRepostsForPostRequest) (*desoRoutes.GetRepostsForPostResponse, error) {
	data := new(desoRoutes.GetRepostsForPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-reposts-for-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetSinglePost(payload *desoRoutes.GetSinglePostRequest) (*desoRoutes.GetSinglePostResponse, error) {
	data := new(desoRoutes.GetSinglePostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-single-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
