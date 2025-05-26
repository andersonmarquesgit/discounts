package usecases

import (
	"coupon/internal/domain/models"
	"coupon/internal/infrastructure/logging"
	"coupon/internal/interfaces/dto"
)

type CouponGenerateUseCase struct {
}

func NewCouponGenerateUseCase() *CouponGenerateUseCase {
	return &CouponGenerateUseCase{}
}

func (uc *CouponGenerateUseCase) CouponGenerate(req dto.CouponGenerateRequest, traceID string) (*models.Coupon, error) {
	logging.Info("CouponGenerate called")

	return nil, nil
}
