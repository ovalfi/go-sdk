package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/go-sdk/helpers"
	"github.com/ovalfi/go-sdk/model"
)

const yieldAPIVersion = "v1/configuration/yield-offering"

// GetBusinessPortfolios makes an API request using Call to get business portfolios
func (c *Call) GetBusinessPortfolios(ctx context.Context) ([]model.Portfolio, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, "v1/configuration/portfolio")

	fL := c.logger.With().Str("func", "GetBusinessPortfolios").Str("endpoint", endpoint).Logger()
	fL.Info().Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data []model.Portfolio `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return []model.Portfolio{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return []model.Portfolio{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// CreateYieldOfferingProfile makes an API request using Call to create a yield offering profile
func (c *Call) CreateYieldOfferingProfile(ctx context.Context, request model.CreateYieldOfferingProfilesRequest) (model.YieldOfferingProfile, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, yieldAPIVersion)

	fL := c.logger.With().Str("func", "CreateYieldOfferingProfile").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Str("name", request.Name).Str("reference", request.Reference).
		Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	signature := helpers.GetSignatureFromReferenceAndPubKey(request.Reference, c.publicKey)
	response := struct {
		Data model.YieldOfferingProfile `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetHeader("Signature", signature).
		SetContext(ctx).
		Post(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.YieldOfferingProfile{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.YieldOfferingProfile{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// UpdateYieldOfferingProfile makes an API request using Call to update a yield offering profile
func (c *Call) UpdateYieldOfferingProfile(ctx context.Context, request model.UpdateYieldOfferingProfilesRequest) (model.UpdatedYieldOfferingProfile, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, yieldAPIVersion)

	fL := c.logger.With().Str("func", "UpdateYieldOfferingProfile").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Str("name", request.Name).Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.UpdatedYieldOfferingProfile `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetBody(request).
		SetResult(&response).
		SetContext(ctx).
		Put(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.UpdatedYieldOfferingProfile{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.UpdatedYieldOfferingProfile{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetAllYieldProfiles makes an API request using Call to get all yield offering profiles
func (c *Call) GetAllYieldProfiles(ctx context.Context) ([]model.YieldOfferingProfile, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, yieldAPIVersion)

	fL := c.logger.With().Str("func", "GetAllYieldProfiles").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data []model.YieldOfferingProfile `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return []model.YieldOfferingProfile{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return []model.YieldOfferingProfile{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}

// GetYieldProfileByID makes an API request using Call to get a yield offering profile by ID
func (c *Call) GetYieldProfileByID(ctx context.Context, request model.GetYieldProfileByIDRequest) (model.YieldOfferingProfile, error) {
	endpoint := fmt.Sprintf("%s%s/%s", c.baseURL, yieldAPIVersion, request.YieldProfileID)

	fL := c.logger.With().Str("func", "GetYieldProfileByID").Str("endpoint", endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, "empty").Msg("request")
	defer fL.Info().Msg("done...")

	response := struct {
		Data model.YieldOfferingProfile `json:"data"`
	}{}
	res, err := c.client.R().
		SetAuthToken(c.bearerToken).
		SetResult(&response).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("error occurred")
		return model.YieldOfferingProfile{}, err
	}

	if res.StatusCode() != http.StatusOK {
		fL.Info().Str(model.LogErrorCode, fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return model.YieldOfferingProfile{}, model.ErrNetworkError
	}

	fL.Info().Interface(model.LogStrResponse, response.Data).Msg("response")
	return response.Data, nil
}
