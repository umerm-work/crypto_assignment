package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/umerm-work/crypto_assignment/domain"
	"net/http"
	"time"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// PriceHandler  represent the httphandler for article
type PriceHandler struct {
	PriceUsecase domain.PriceUsecase
}

// NewPriceHandler will initialize the articles/ resources endpoint
func NewPriceHandler(e *gin.Engine, us domain.PriceUsecase) {
	handler := &PriceHandler{
		PriceUsecase: us,
	}
	e.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	v1 := e.Group("/service")
	{
		v1.GET("/price", handler.FetchPrice)
	}
}

// FetchPrice will fetch the btc price based on given params
func (a *PriceHandler) FetchPrice(c *gin.Context) {
	tsyms := c.Query("tsyms")
	fsyms := c.Query("fsyms")
	price, err := a.PriceUsecase.GetBtcPrice(context.Background(), tsyms, fsyms)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	c.JSON(http.StatusOK, price)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
