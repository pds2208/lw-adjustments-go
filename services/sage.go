package services

import (
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
	client.Header["AuthToken"] = []string{"hwytrwqvf262vsz9"}
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

	_, err := sage.client.R().
		SetResult(&productResponse).
		Get(endpoint)

	if err == nil {
		return productResponse, nil
	}

	return types.ProductResponse{}, err

	//fmt.Println("Response Info:")
	//fmt.Println("  Error      :", err)
	//fmt.Println("  Status Code:", resp.StatusCode())
	//fmt.Println("  Status     :", resp.Status())
	//fmt.Println("  Proto      :", resp.Proto())
	//fmt.Println("  Time       :", resp.Time())
	//fmt.Println("  Received At:", resp.ReceivedAt())
	//fmt.Println("  Body       :\n", resp)
	//fmt.Println()

}
