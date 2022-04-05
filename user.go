package desoapi

import (
	"fmt"
	"github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) BlockPublicKey(payload *routes.BlockPublicKeyRequest) (*routes.BlockPublicKeyResponse, error) {
	data := new(routes.BlockPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/block-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) DeletePII(payload *routes.DeletePIIRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/delete-pii", payload, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetDiamondsForPublicKey(payload *routes.GetDiamondsForPublicKeyRequest) (*routes.GetDiamondsForPublicKeyResponse, error) {
	data := new(routes.GetDiamondsForPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-diamonds-for-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetFollowsStateless(payload *routes.GetFollowsStatelessRequest) (*routes.GetFollowsResponse, error) {
	data := new(routes.GetFollowsResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-follows-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetHodlersForPublicKey(payload *routes.GetHodlersForPublicKeyRequest) (*routes.GetHodlersForPublicKeyResponse, error) {
	data := new(routes.GetHodlersForPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-hodlers-for-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNotifications(payload *routes.GetNotificationsRequest) (*routes.GetNotificationsResponse, error) {
	data := new(routes.GetNotificationsResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-notifications", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetProfiles(payload *routes.GetProfilesRequest) (*routes.GetProfilesResponse, error) {
	data := new(routes.GetProfilesResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-profiles", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetSingleProfile(payload *routes.GetSingleProfileRequest) (*routes.GetSingleProfileResponse, error) {
	data := new(routes.GetSingleProfileResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-single-profile", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetUnreadNotificationsCount(payload *routes.GetNotificationsCountRequest) (*routes.GetNotificationsCountResponse, error) {
	data := new(routes.GetNotificationsCountResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-unread-notifications-count", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetUserDerivedKeys(payload *routes.GetUserDerivedKeysRequest) (*routes.GetUserDerivedKeysResponse, error) {
	data := new(routes.GetUserDerivedKeysResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-user-derived-keys", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetUserGlobalMetadata(payload *routes.GetUserGlobalMetadataRequest) (*routes.GetUserGlobalMetadataResponse, error) {
	data := new(routes.GetUserGlobalMetadataResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-user-global-metadata", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetUserMetadata(publicKeyBase58Check string) (*routes.GetUserMetadataResponse, error) {
	uri := fmt.Sprintf("/api/v0/get-user-metadata/%s", publicKeyBase58Check)
	data := new(routes.GetUserMetadataResponse)

	err := c.executeRequest(fasthttp.MethodGet, uri, nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetUsersStateless(payload *routes.GetUsersStatelessRequest) (*routes.GetUsersResponse, error) {
	data := new(routes.GetUsersResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-users-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) IsFollowingPublicKey(payload *routes.IsFollowingPublicKeyRequest) (*routes.IsFolllowingPublicKeyResponse, error) {
	data := new(routes.IsFolllowingPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/is-following-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) IsHodlingPublicKey(payload *routes.IsHodlingPublicKeyRequest) (*routes.IsHodlingPublicKeyResponse, error) {
	data := new(routes.IsHodlingPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/is-hodling-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SetNotificationMetadata(payload *routes.SetNotificationMetadataRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/set-notification-metadata", payload, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateUserGlobalMetadata(payload *routes.UpdateUserGlobalMetadataRequest) (*routes.UpdateUserGlobalMetadataResponse, error) {
	data := new(routes.UpdateUserGlobalMetadataResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/update-user-global-metadata", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
