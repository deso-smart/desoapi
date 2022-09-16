package desoapi

import (
	"fmt"
	desoRoutes "github.com/deso-smart/deso-backend/v3/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) BlockPublicKey(payload *desoRoutes.BlockPublicKeyRequest) (*desoRoutes.BlockPublicKeyResponse, error) {
	data := new(desoRoutes.BlockPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/block-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) DeletePII(payload *desoRoutes.DeletePIIRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/delete-pii", payload, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetDiamondsForPublicKey(payload *desoRoutes.GetDiamondsForPublicKeyRequest) (*desoRoutes.GetDiamondsForPublicKeyResponse, error) {
	data := new(desoRoutes.GetDiamondsForPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-diamonds-for-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetFollowsStateless(payload *desoRoutes.GetFollowsStatelessRequest) (*desoRoutes.GetFollowsResponse, error) {
	data := new(desoRoutes.GetFollowsResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-follows-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetHodlersForPublicKey(payload *desoRoutes.GetHodlersForPublicKeyRequest) (*desoRoutes.GetHodlersForPublicKeyResponse, error) {
	data := new(desoRoutes.GetHodlersForPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-hodlers-for-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetNotifications(payload *desoRoutes.GetNotificationsRequest) (*desoRoutes.GetNotificationsResponse, error) {
	data := new(desoRoutes.GetNotificationsResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-notifications", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetProfiles(payload *desoRoutes.GetProfilesRequest) (*desoRoutes.GetProfilesResponse, error) {
	data := new(desoRoutes.GetProfilesResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-profiles", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetSingleProfile(payload *desoRoutes.GetSingleProfileRequest) (*desoRoutes.GetSingleProfileResponse, error) {
	data := new(desoRoutes.GetSingleProfileResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-single-profile", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetUnreadNotificationsCount(payload *desoRoutes.GetNotificationsCountRequest) (*desoRoutes.GetNotificationsCountResponse, error) {
	data := new(desoRoutes.GetNotificationsCountResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-unread-notifications-count", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetUserDerivedKeys(payload *desoRoutes.GetUserDerivedKeysRequest) (*desoRoutes.GetUserDerivedKeysResponse, error) {
	data := new(desoRoutes.GetUserDerivedKeysResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-user-derived-keys", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetUserGlobalMetadata(payload *desoRoutes.GetUserGlobalMetadataRequest) (*desoRoutes.GetUserGlobalMetadataResponse, error) {
	data := new(desoRoutes.GetUserGlobalMetadataResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-user-global-metadata", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetUserMetadata(publicKeyBase58Check string) (*desoRoutes.GetUserMetadataResponse, error) {
	uri := fmt.Sprintf("/api/v0/get-user-metadata/%s", publicKeyBase58Check)
	data := new(desoRoutes.GetUserMetadataResponse)

	err := c.executeRequest(fasthttp.MethodGet, uri, nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetUsersStateless(payload *desoRoutes.GetUsersStatelessRequest) (*desoRoutes.GetUsersResponse, error) {
	data := new(desoRoutes.GetUsersResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-users-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) IsFollowingPublicKey(payload *desoRoutes.IsFollowingPublicKeyRequest) (*desoRoutes.IsFolllowingPublicKeyResponse, error) {
	data := new(desoRoutes.IsFolllowingPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/is-following-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) IsHodlingPublicKey(payload *desoRoutes.IsHodlingPublicKeyRequest) (*desoRoutes.IsHodlingPublicKeyResponse, error) {
	data := new(desoRoutes.IsHodlingPublicKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/is-hodling-public-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SetNotificationMetadata(payload *desoRoutes.SetNotificationMetadataRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/set-notification-metadata", payload, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateUserGlobalMetadata(payload *desoRoutes.UpdateUserGlobalMetadataRequest) (*desoRoutes.UpdateUserGlobalMetadataResponse, error) {
	data := new(desoRoutes.UpdateUserGlobalMetadataResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/update-user-global-metadata", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetTransactionSpendingLimitHexString(payload *desoRoutes.GetTransactionSpendingLimitHexStringRequest) (*desoRoutes.GetTransactionSpendingLimitHexStringResponse, error) {
	data := new(desoRoutes.GetTransactionSpendingLimitHexStringResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-transaction-spending-limit-hex-string", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetTransactionSpendingLimitResponseFromHex(transactionSpendingLimitHex string) (*desoRoutes.TransactionSpendingLimitResponse, error) {
	uri := fmt.Sprintf("/api/v0/get-transaction-spending-limit-response-from-hex/%s", transactionSpendingLimitHex)
	data := new(desoRoutes.TransactionSpendingLimitResponse)

	err := c.executeRequest(fasthttp.MethodGet, uri, nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetHodlersCountForPublicKeys(payload *desoRoutes.GetHolderCountForPublicKeysRequest) (map[string]int, error) {
	var data map[string]int

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-hodlers-count-for-public-keys", payload, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetSingleDerivedKey(ownerPublicKeyBase58Check string, derivedPublicKeyBase58Check string) (*desoRoutes.GetSingleDerivedKeyResponse, error) {
	uri := fmt.Sprintf("/api/v0/get-single-derived-key/%s/%s", ownerPublicKeyBase58Check, derivedPublicKeyBase58Check)
	data := new(desoRoutes.GetSingleDerivedKeyResponse)

	err := c.executeRequest(fasthttp.MethodGet, uri, nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetAccessBytes(payload *desoRoutes.GetAccessBytesRequest) (*desoRoutes.GetAccessBytesResponse, error) {
	data := new(desoRoutes.GetAccessBytesResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-access-bytes", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
