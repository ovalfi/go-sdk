package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/model"
)

const walletAPIVersion = "v1/customer/wallet"

// GetWallet makes an API request using Call to get customer wallet
func (c Call) GetWallet(ctx context.Context, request model.WalletRequest) (model.Wallet, error) {
	endpoint := fmt.Sprintf("%s%s?customer_id=%s&network=%s&asset=%s", c.baseURL, walletAPIVersion, request.CustomerID, request.Network, request.Asset)

	fL := c.logger.With().Str("func", "GetWallet").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.Wallet `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.Wallet{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.Wallet{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetWallets makes an API request using Call to get all customers wallets
func (c Call) GetWallets(ctx context.Context, customerID string) ([]*model.Wallet, error) {
	endpoint := fmt.Sprintf("%s%s/%s", c.baseURL, walletAPIVersion, customerID)

	fL := c.logger.With().Str("func", "GetWallets").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, customerID).Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data []*model.Wallet `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return nil, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}
