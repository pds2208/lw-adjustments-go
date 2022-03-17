package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"lw-adjustments/config"
	"lw-adjustments/types"
	"time"
)

type Sage struct {
	client *resty.Client
}

const AdjustmentIn = "adj_in"
const AdjustmentOut = "adj_out"

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

	switch adjustment.AdjustmentType {
	case AdjustmentIn:
		return sage.AddAdjustmentIn(adjustment)
	case AdjustmentOut:
		return sage.AddAdjustmentOut(adjustment)
	}
	return nil
}

func (sage Sage) AddAdjustmentIn(adjustment types.Adjustments) error {
	return nil
}

func (sage Sage) AddAdjustmentOut(adjustment types.Adjustments) error {
	return nil
}

func (sage Sage) GetProductDetail(product string) (types.ProductResponse, error) {

	var productResponse types.ProductResponse
	var endpoint = config.Config.Sage.Endpoint + config.Config.Sage.ProductEndpoint + product

	resp, err := sage.client.R().
		SetResult(&productResponse).
		Get(endpoint)

	if err == nil {
		a, _ := formatJSON(resp.Body())
		log.Debug().
			Int("Status Code", resp.StatusCode()).Msg("")
		log.Debug().
			Str("HTTP Status", resp.Status()).Msg("")
		log.Debug().
			Str("Proto", resp.Proto()).Msg("")
		log.Debug().
			Str("Time", resp.Time().String()).Msg("")
		log.Debug().
			Str("Received At", resp.ReceivedAt().String()).Msg("")
		log.Debug().
			Msg(string(a))

		return productResponse, nil
	}

	return types.ProductResponse{}, err
}

func formatJSON(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err == nil {
		return out.Bytes(), err
	}
	return data, nil
}
