package accrualservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"github.com/s-turchinskiy/loyalty-system/internal/repository"
	"github.com/s-turchinskiy/loyalty-system/internal/service"
	"time"
)

type AccrualService struct {
	ctx           context.Context
	repository    repository.Repository
	retryStrategy []time.Duration
	interval      time.Duration
	accrualGetter AccrualGetter
	workerCount   int
	errorsCh      chan error
	doneCh        chan struct{}

	jobs    chan models.OrdersForAccrualCalculation
	numJobs int
	results chan result
}

type result struct {
	order   models.OrdersForAccrualCalculation
	accrual *models.AccrualData
	err     error
}

func New(
	ctx context.Context,
	rep repository.Repository,
	retryStrategy []time.Duration,
	interval time.Duration,
	accrualGetter AccrualGetter,
	workerCount int,
	errorsCh chan error,
	doneCh chan struct{}) *AccrualService {

	return &AccrualService{
		ctx:           ctx,
		repository:    rep,
		retryStrategy: retryStrategy,
		interval:      interval,
		accrualGetter: accrualGetter,
		workerCount:   workerCount,
		errorsCh:      errorsCh,
		doneCh:        doneCh,
	}

}

func (a *AccrualService) RunPeriodically() {

	ticker := time.NewTicker(a.interval)
	for range ticker.C {
		a.Run()
	}
}

func (a *AccrualService) Run() {

	var orders []models.OrdersForAccrualCalculation
	var err error

	for _, delay := range a.retryStrategy {
		time.Sleep(delay)
		orders, err = a.repository.GetOrdersForAccrualCalculation(a.ctx)
		if err == nil {
			break
		} else {
			logger.Log.Infoln("failed GetOrdersForAccrualCalculation", common.WrapError(err).Error())
		}
	}

	a.numJobs = len(orders)
	a.jobs = a.generator(a.doneCh, orders)
	a.results = make(chan result, cap(a.jobs))

	for w := 1; w <= a.workerCount; w++ {
		go a.workerSender()
	}

	a.resultHandling()

}

func (a *AccrualService) generator(doneCh chan struct{}, input []models.OrdersForAccrualCalculation) chan models.OrdersForAccrualCalculation {

	jobs := make(chan models.OrdersForAccrualCalculation, len(input))

	go func() {
		defer close(jobs)

		for _, data := range input {
			select {
			case <-doneCh:
				return
			case jobs <- data:
			}
		}
	}()

	return jobs

}

func (a *AccrualService) workerSender() {

	for order := range a.jobs {

		select {
		case <-a.doneCh:
			return
		default:

			a.sendResult(order)
		}
	}
}

func (a *AccrualService) sendResult(order models.OrdersForAccrualCalculation) {

	var accrual *models.AccrualData
	var err error

	for _, delay := range a.retryStrategy {
		time.Sleep(delay)
		accrual, err = a.accrualGetter.GetAccrual(order.OrderID)
		if err == nil {
			a.results <- result{order, accrual, nil}
			return
		} else if !service.IsConnectionError(err) {
			a.results <- result{order, nil, err}
			return
		}
	}

	a.results <- result{order, nil, err}

}

func (a *AccrualService) resultHandling() {

	var result result
	var errs []error
	for i := 1; i <= a.numJobs; i++ {
		select {
		case <-a.doneCh:
			return
		case result = <-a.results:
			if result.err != nil {
				errs = append(errs, fmt.Errorf("%v %w", result, result.err))
				continue
			}

			oldStatus := result.order.CurrentStatus
			newStatus := result.accrual.Status.AsString()
			if oldStatus == newStatus {
				continue
			}

			if result.accrual.Status == models.PROCESSED {
				err := a.repository.UpdateOrderStatusAndNewRefill(
					a.ctx,
					result.order.OrderID,
					newStatus,
					result.accrual.Accrual,
					result.order.UserID,
				)

				if err != nil {
					errs = append(errs, fmt.Errorf("%v %w", result, err))
				}
				continue
			}

			err := a.repository.UpdateOrderStatus(a.ctx, result.order.OrderID, newStatus)
			if err != nil {
				errs = append(errs, fmt.Errorf("%v %w", result, err))
				continue
			}
		}
	}

	close(a.results)

	if len(errs) != 0 {
		logger.Log.Info(fmt.Sprintf("Success got %d accruals, unsuccess %d", a.numJobs-len(errs), len(errs)))
		logger.Log.Info("errors got accruals", errors.Join(errs...))
	} else {
		logger.Log.Info(fmt.Sprintf("Success got %d accruals", a.numJobs))
	}

}
