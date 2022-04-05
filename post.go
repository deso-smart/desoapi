package desoapi

import (
	"github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetDiamondedPosts(payload *routes.GetPostsDiamondedBySenderForReceiverRequest) (*routes.GetPostsDiamondedBySenderForReceiverResponse, error) {
	data := new(routes.GetPostsDiamondedBySenderForReceiverResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-diamonded-posts", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetDiamondsForPost(payload *routes.GetDiamondsForPostRequest) (*routes.GetDiamondsForPostResponse, error) {
	data := new(routes.GetDiamondsForPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-diamonds-for-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetLikesForPost(payload *routes.GetLikesForPostRequest) (*routes.GetLikesForPostResponse, error) {
	data := new(routes.GetLikesForPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-likes-for-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetPostsForPublicKey(payload *routes.GetPostsForPublicKeyRequest) (*routes.GetPostsForPublicKeyResponse, error) {
	data := new(routes.GetPostsForPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-posts-for-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetPostsStateless(payload *routes.GetPostsStatelessRequest) (*routes.GetPostsStatelessResponse, error) {
	data := new(routes.GetPostsStatelessResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-posts-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetQuoteRepostsForPost(payload *routes.GetQuoteRepostsForPostRequest) (*routes.GetQuoteRepostsForPostResponse, error) {
	data := new(routes.GetQuoteRepostsForPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-quote-reposts-for-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetRepostsForPost(payload *routes.GetRepostsForPostRequest) (*routes.GetRepostsForPostResponse, error) {
	data := new(routes.GetRepostsForPostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-reposts-for-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetSinglePost(payload *routes.GetSinglePostRequest) (*routes.GetSinglePostResponse, error) {
	data := new(routes.GetSinglePostResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-single-post", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
