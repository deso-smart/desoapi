package desoapi

import (
	"github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) GetReferralInfoForReferralHash(payload *routes.GetReferralInfoForReferralHashRequest) (*routes.GetReferralInfoForReferralHashResponse, error) {
	data := new(routes.GetReferralInfoForReferralHashResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-referral-info-for-referral-hash", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetReferralInfoForUser(payload *routes.GetReferralInfoForUserRequest) (*routes.GetReferralInfoForUserResponse, error) {
	data := new(routes.GetReferralInfoForUserResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-referral-info-for-user", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
