package desoapi

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	desoCore "github.com/deso-smart/deso-core/v2/lib"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"net/url"
)

const (
	version   = "0.0.0"
	userAgent = "desoapi/" + version
)

type Client struct {
	baseUri   *url.URL
	http      *fasthttp.Client
	json      jsoniter.API
	userAgent string
}

func NewClient(baseUri string) (*Client, error) {
	parsedBaseUri, err := url.Parse(baseUri)
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseUri: parsedBaseUri,
		http: &fasthttp.Client{
			MaxConnsPerHost: 65536,
		},
		json:      jsoniter.ConfigCompatibleWithStandardLibrary,
		userAgent: userAgent,
	}

	return c, nil
}

type errorResponse struct {
	Error string
}

func (c *Client) executeRequest(method string, uri string, body interface{}, data interface{}) error {
	resolvedUri, err := c.baseUri.Parse(uri)
	if err != nil {
		return err
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(resolvedUri.String())
	req.Header.SetMethod(method)
	req.Header.SetUserAgent(userAgent)
	req.Header.Set("Accept", "application/json")

	if body != nil {
		buf, err := c.json.Marshal(body)
		if err != nil {
			return err
		}
		req.Header.SetContentType("application/json")
		req.SetBody(buf)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := c.http.Do(req, resp); err != nil {
		return err
	}

	// TODO: or status >= 200 && status <= 204?
	if resp.StatusCode() < fasthttp.StatusOK || resp.StatusCode() >= fasthttp.StatusBadRequest {
		errResp := new(errorResponse)
		if err := c.json.Unmarshal(resp.Body(), errResp); err != nil {
			return err
		}

		if errResp.Error != "" {
			return errors.New(errResp.Error)
		}

		return errors.New(string(resp.Body()))
	}

	if data != nil {
		if err := c.json.Unmarshal(resp.Body(), data); err != nil {
			return err
		}
	}

	return nil
}

func SignTransactionWithDerivedKey(transactionHex string, derivedKeySeedHex string) (string, error) {
	txnBytes, err := hex.DecodeString(transactionHex)
	if err != nil {
		return "", fmt.Errorf("problem decoding transaction hex: %w", err)
	}

	derivedKeyBytes, err := hex.DecodeString(derivedKeySeedHex)
	if err != nil {
		return "", fmt.Errorf("problem decoding derived key seed hex: %w", err)
	}

	privateKeyBytes, _ := btcec.PrivKeyFromBytes(btcec.S256(), derivedKeyBytes)
	newTxnBytes, txnSignatureBytes, err := desoCore.SignTransactionBytes(txnBytes, privateKeyBytes, true)
	if err != nil {
		return "", fmt.Errorf("problem signing transaction: %w", err)
	}

	var signedTxnHex []byte
	signedTxnHex = newTxnBytes[0 : len(newTxnBytes)-1]
	signedTxnHex = append(signedTxnHex, desoCore.UintToBuf(uint64(len(txnSignatureBytes)))...)
	signedTxnHex = append(signedTxnHex, txnSignatureBytes...)

	return hex.EncodeToString(signedTxnHex), nil
}
