package desoapi

import (
	desoRoutes "github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetBlockTemplate(payload *desoRoutes.GetBlockTemplateRequest) (*desoRoutes.GetBlockTemplateResponse, error) {
	data := new(desoRoutes.GetBlockTemplateResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-block-template", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SubmitBlock(payload *desoRoutes.SubmitBlockRequest) (*desoRoutes.SubmitBlockResponse, error) {
	data := new(desoRoutes.SubmitBlockResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/submit-block", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
