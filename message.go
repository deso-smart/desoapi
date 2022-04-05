package desoapi

import (
	desoRoutes "github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) CheckPartyMessagingKeys(payload *desoRoutes.CheckPartyMessagingKeysRequest) (*desoRoutes.CheckPartyMessagingKeysResponse, error) {
	data := new(desoRoutes.CheckPartyMessagingKeysResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/check-party-messaging-keys", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetAllMessagingGroupKeys(payload *desoRoutes.GetAllMessagingGroupKeysRequest) (*desoRoutes.GetAllMessagingGroupKeysResponse, error) {
	data := new(desoRoutes.GetAllMessagingGroupKeysResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-all-messaging-group-keys", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetMessagesStateless(payload *desoRoutes.GetMessagesStatelessRequest) (*desoRoutes.GetMessagesResponse, error) {
	data := new(desoRoutes.GetMessagesResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-messages-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) MarkAllMessagesRead(payload *desoRoutes.MarkAllMessagesReadRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/mark-all-messages-read", payload, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) MarkContactMessagesRead(payload *desoRoutes.MarkContactMessagesReadRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/mark-contact-messages-read", payload, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RegisterMessagingGroupKey(payload *desoRoutes.RegisterMessagingGroupKeyRequest) (*desoRoutes.RegisterMessagingGroupKeyResponse, error) {
	data := new(desoRoutes.RegisterMessagingGroupKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/register-messaging-group-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SendMessageStateless(payload *desoRoutes.SendMessageStatelessRequest) (*desoRoutes.SendMessageStatelessResponse, error) {
	data := new(desoRoutes.SendMessageStatelessResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/send-message-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
