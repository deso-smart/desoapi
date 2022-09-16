package desoapi

import (
	desoRoutes "github.com/deso-smart/deso-backend/v3/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetTutorialCreators(payload *desoRoutes.GetTutorialCreatorsRequest) (*desoRoutes.GetTutorialCreatorResponse, error) {
	data := new(desoRoutes.GetTutorialCreatorResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-tutorial-creators", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) StartOrSkipTutorial(payload *desoRoutes.StartOrSkipTutorialRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/start-or-skip-tutorial", payload, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateTutorialStatus(payload *desoRoutes.UpdateTutorialStatusRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/update-tutorial-status", payload, nil)
	if err != nil {
		return err
	}

	return nil
}
