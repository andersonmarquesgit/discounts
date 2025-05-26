package usecases

// UseCases struct to store the customer use case and other use cases
type UseCases struct {
	CouponGenerateUseCase *CouponGenerateUseCase
}

// NewUseCases struct
func NewUseCases(couponGenerateUseCase *CouponGenerateUseCase) *UseCases {
	return &UseCases{
		CouponGenerateUseCase: couponGenerateUseCase,
	}
}
