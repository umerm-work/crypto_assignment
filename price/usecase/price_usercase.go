package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/umerm-work/crypto_assignment/domain"
	"github.com/umerm-work/crypto_assignment/transform"
	"github.com/umerm-work/crypto_assignment/util"
	"net/http"
	"time"
)

type priceUsecase struct {
	priceRepo      domain.PriceRepository
	contextTimeout time.Duration
}

// NewPriceUsecase will create new an priceUsecase object representation of domain.PriceUsecase interface
func NewPriceUsecase(a domain.PriceRepository, timeout time.Duration) domain.PriceUsecase {
	return &priceUsecase{
		priceRepo:      a,
		contextTimeout: timeout,
	}
}

/*
* In this function below, I'm using errgroup with the pipeline pattern
* Look how this works in this package explanation
 */
func (a *priceUsecase) GetBtcPrice(ctx context.Context, tsyms, fsyms string) (*domain.Price, error) {
	//g, ctx := errgroup.WithContext(ctx)
	res, err := util.SendHTTPRequest(
		http.MethodGet,
		"application/json",
		fmt.Sprintf(`https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%s&tsyms=%s`, fsyms, tsyms),
		nil, nil)
	if err != nil {
		return nil, err
	}
	//logrus.Println(string(res))
	var req domain.Price
	err = json.Unmarshal(res, &req)
	if err != nil {
		resDb, err := a.priceRepo.Fetch(ctx, fsyms, tsyms)
		if err != nil {
			return nil, err
		}
		logrus.Println("unable to marshall error:", err)
		tRes := transform.ModelToApiPrice(resDb)
		return &tRes, nil
	}
	tres := transform.ApiToModelPrice(req)
	for _, v := range tres {
		if err := a.priceRepo.Insert(ctx, v); err != nil {
			logrus.Error(err)
		}
	}
	//logrus.Println(req)
	return &req, nil
}
