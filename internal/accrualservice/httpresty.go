package accrualservice

import (
	"github.com/go-resty/resty/v2"
	"github.com/mailru/easyjson"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"go.uber.org/zap"
	"strings"
)

type HTTPResty struct {
	client *resty.Client
	url    string
}

type AccrualGetter interface {
	GetAccrual(url string) (result *models.AccrualData, err error)
}

func NewHTTPResty(url string) *HTTPResty {

	return &HTTPResty{
		client: resty.New(),
		url:    url,
	}
}

func (r *HTTPResty) GetAccrual(numberOrder string) (result *models.AccrualData, err error) {

	request := r.client.R()

	url := strings.Replace(r.url, "{number}", numberOrder, 1)
	resp, err := request.Get(url)

	if err != nil {
		return nil, common.WrapError(err)
	}

	if err := common.CheckResponseStatus(
		resp.StatusCode(),
		resp.Body(),
		url,
	); err != nil {
		return nil, err
	}

	logger.Log.Debugln(zap.String("body", string(resp.Body())))
	
	err = easyjson.Unmarshal(resp.Body(), result)
	if err != nil {
		logger.Log.Info("error encoding response", zap.Error(err))
		return nil, err
	}

	return result, nil

}
