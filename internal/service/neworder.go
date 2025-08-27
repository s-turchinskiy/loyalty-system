package service

import (
	"context"
	"github.com/ShiraazMoollatjie/goluhn"
	"github.com/s-turchinskiy/loyalty-system/internal/servicecommon"
	"time"
)

func (s *Service) NewOrder(ctx context.Context, login, orderID string) (err error) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	err = goluhn.Validate(orderID)
	if err != nil {
		return servicecommon.ErrNoLuhnValidate
	}

	var uploaded bool
	var who string
	for _, delay := range s.retryStrategy {
		time.Sleep(delay)
		uploaded, who, err = s.Repository.GetOrderAlreadyUploaded(ctx, orderID)
		if err == nil {
			break
		} else if !IsConnectionError(err) {
			return err
		}
	}

	if err != nil {
		return err
	}

	if uploaded {
		if who == login {
			return servicecommon.ErrOrderNumberAlreadyUploadedByThisUser
		} else {
			return servicecommon.ErrOrderNumberAlreadyUploadedByAnotherUser
		}
	}

	for _, delay := range s.retryStrategy {
		time.Sleep(delay)
		err = s.Repository.NewOrder(ctx, login, orderID)
		if err == nil {
			break
		} else if !IsConnectionError(err) {
			return err
		}
	}

	return err

}
