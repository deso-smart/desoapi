package desoapi

import (
	"github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetBlockTemplate(payload *routes.GetBlockTemplateRequest) (*routes.GetBlockTemplateResponse, error) {
	data := new(routes.GetBlockTemplateResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-block-template", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SubmitBlock(payload *routes.SubmitBlockRequest) (*routes.SubmitBlockResponse, error) {
	data := new(routes.SubmitBlockResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/submit-block", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
