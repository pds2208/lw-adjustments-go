package services

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"lw-adjustments/config"
	"lw-adjustments/types"
	"time"
)

type Sage struct {
	client *resty.Client
}

func NewSage() Sage {
	client := resty.New()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))
	client.SetRetryCount(3)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.Header["AuthToken"] = []string{config.Config.Sage.AuthToken}
	client.SetTimeout(15 * time.Second)
	return Sage{client: client}
}

func (sage Sage) addAdjustment(adjustment types.Adjustments) error {
	log.Info().
		Str("stock code", adjustment.StockCode).
		Str("type", adjustment.AdjustmentType).
		Float64("amount", adjustment.Amount).
		Msg("updating Sage...")

	return nil
}

func (sage Sage) GetProductDetail(product string) (types.ProductResponse, error) {

	var productResponse types.ProductResponse
	var endpoint = config.Config.Sage.Endpoint + config.Config.Sage.ProductEndpoint + product

	resp, err := sage.client.R().
		SetResult(&productResponse).
		Get(endpoint)

	if err == nil {
		log.Debug().
			Str("  Error      :", err.Error()).
			Int("  Status Code:", resp.StatusCode()).
			Str("  Status     :", resp.Status()).
			Str("  Proto      :", resp.Proto()).
			Str("  Time      :", resp.Time().String()).
			Str("  Received At:", resp.ReceivedAt().String()).
			Str("  Body:", string(resp.Body())).
			Msg("Response Info:")

		return productResponse, nil
	}

	return types.ProductResponse{}, err

}
