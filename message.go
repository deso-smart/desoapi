package desoapi

import (
	"github.com/deso-smart/deso-backend/v2/routes"
	"github.com/valyala/fasthttp"
)

func (c *Client) CheckPartyMessagingKeys(payload *routes.CheckPartyMessagingKeysRequest) (*routes.CheckPartyMessagingKeysResponse, error) {
	data := new(routes.CheckPartyMessagingKeysResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/check-party-messaging-keys", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetAllMessagingGroupKeys(payload *routes.GetAllMessagingGroupKeysRequest) (*routes.GetAllMessagingGroupKeysResponse, error) {
	data := new(routes.GetAllMessagingGroupKeysResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-all-messaging-group-keys", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetMessagesStateless(payload *routes.GetMessagesStatelessRequest) (*routes.GetMessagesResponse, error) {
	data := new(routes.GetMessagesResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/get-messages-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) MarkAllMessagesRead(payload *routes.MarkAllMessagesReadRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/mark-all-messages-read", payload, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) MarkContactMessagesRead(payload *routes.MarkContactMessagesReadRequest) error {
	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/mark-contact-messages-read", payload, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RegisterMessagingGroupKey(payload *routes.RegisterMessagingGroupKeyRequest) (*routes.RegisterMessagingGroupKeyResponse, error) {
	data := new(routes.RegisterMessagingGroupKeyResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/register-messaging-group-key", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) SendMessageStateless(payload *routes.SendMessageStatelessRequest) (*routes.SendMessageStatelessResponse, error) {
	data := new(routes.SendMessageStatelessResponse)

	err := c.executeRequest(fasthttp.MethodPost, "/api/v0/send-message-stateless", payload, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
