package desoapi

import (
	desoRoutes "github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetHotFeed(payload *desoRoutes.HotFeedPageRequest) (*desoRoutes.HotFeedPageResponse, error) {
	data := new(desoRoutes.HotFeedPageResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-hot-feed", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
