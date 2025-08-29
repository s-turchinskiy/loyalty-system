package common

import (
	"fmt"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"net/http"
	"runtime"
)

func WrapError(err error) error {

	if err == nil {
		return nil
	}

	_, filename, line, _ := runtime.Caller(1)
	return fmt.Errorf("[error] %s %d: %w", filename, line, err)
}

func CheckResponseStatus(statusCode int, body []byte, url string) error {

	if statusCode != http.StatusOK {

		logger.Log.Infow("error. status code <> 200",
			"status code", statusCode,
			"url", url,
			"body", string(body))
		err := fmt.Errorf("status code <> 200, = %d, url : %s", statusCode, url)
		return err
	}

	return nil
}
