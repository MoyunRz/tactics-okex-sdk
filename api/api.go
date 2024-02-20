package api

import (
	"chain-tactics/okex-sdk"
	"chain-tactics/okex-sdk/api/rest"
	"chain-tactics/okex-sdk/api/ws"
	"context"
)

// Client is the main retrieval wrapper of okex-sdk
type Client struct {
	Rest *rest.ClientRest
	Ws   *ws.ClientWs
	ctx  context.Context
}

// NewClient returns a pointer to a fresh Client
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, destination okex_sdk.Destination) (*Client, error) {
	restURL := okex_sdk.RestURL
	wsPubURL := okex_sdk.PublicWsURL
	wsPriURL := okex_sdk.PrivateWsURL
	switch destination {
	case okex_sdk.AwsServer:
		restURL = okex_sdk.AwsRestURL
		wsPubURL = okex_sdk.AwsPublicWsURL
		wsPriURL = okex_sdk.AwsPrivateWsURL
	case okex_sdk.DemoServer:
		restURL = okex_sdk.DemoRestURL
		wsPubURL = okex_sdk.DemoPublicWsURL
		wsPriURL = okex_sdk.DemoPrivateWsURL
	}

	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination, "")
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[bool]okex_sdk.BaseURL{true: wsPriURL, false: wsPubURL})

	return &Client{r, c, ctx}, nil
}

// NewClientProxy returns a pointer to a fresh Client proxy
func NewClientProxy(ctx context.Context, apiKey, secretKey, passphrase string, destination okex_sdk.Destination, proxyUrl string) (*Client, error) {
	restURL := okex_sdk.RestURL
	wsPubURL := okex_sdk.PublicWsURL
	wsPriURL := okex_sdk.PrivateWsURL
	switch destination {
	case okex_sdk.AwsServer:
		restURL = okex_sdk.AwsRestURL
		wsPubURL = okex_sdk.AwsPublicWsURL
		wsPriURL = okex_sdk.AwsPrivateWsURL
	case okex_sdk.DemoServer:
		restURL = okex_sdk.DemoRestURL
		wsPubURL = okex_sdk.DemoPublicWsURL
		wsPriURL = okex_sdk.DemoPrivateWsURL
	}

	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination, proxyUrl)
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[bool]okex_sdk.BaseURL{true: wsPriURL, false: wsPubURL})

	return &Client{r, c, ctx}, nil
}
