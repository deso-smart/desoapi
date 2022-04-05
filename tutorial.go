package desoapi

import (
	"github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetTutorialCreators(payload *routes.GetTutorialCreatorsRequest) (*routes.GetTutorialCreatorResponse, error) {
	data := new(routes.GetTutorialCreatorResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-tutorial-creators", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) StartOrSkipTutorial(payload *routes.StartOrSkipTutorialRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/start-or-skip-tutorial", payload, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateTutorialStatus(payload *routes.UpdateTutorialStatusRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/update-tutorial-status", payload, nil)
	if err != nil {
		return err
	}

	return nil
}
