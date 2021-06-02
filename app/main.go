package main

import (
	context2 "context"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	config2 "github.com/umerm-work/crypto_assignment/config"
	"github.com/umerm-work/crypto_assignment/price/delivery/http"
	"github.com/umerm-work/crypto_assignment/price/repository/postgres"
	"github.com/umerm-work/crypto_assignment/price/usecase"
	"time"
)

func main() {
	router := gin.Default()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	config, err := config2.Init()
	if err != nil {
		panic("config error")
	}
	dbRepo := postgres.New(config)
	au := usecase.NewPriceUsecase(dbRepo, timeoutContext)
	http.NewPriceHandler(router, au)

	// Scheduler every one hour
	gocron.Every(1).Hour().Do(au.GetBtcPrice(context2.Background(), config.CurrencyPair.Tsyms, config.CurrencyPair.Fsyms))

	log.Fatal(router.Run(":" + config.DefaultAppPort))
}
