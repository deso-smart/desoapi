package desoapi

import (
	desoRoutes "github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetReferralInfoForReferralHash(payload *desoRoutes.GetReferralInfoForReferralHashRequest) (*desoRoutes.GetReferralInfoForReferralHashResponse, error) {
	data := new(desoRoutes.GetReferralInfoForReferralHashResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-referral-info-for-referral-hash", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetReferralInfoForUser(payload *desoRoutes.GetReferralInfoForUserRequest) (*desoRoutes.GetReferralInfoForUserResponse, error) {
	data := new(desoRoutes.GetReferralInfoForUserResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-referral-info-for-user", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
