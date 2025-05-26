package models

import (
	"github.com/google/uuid"
	"time"
)

type Coupon struct {
	ID          uuid.UUID  `db:"id"`
	Code        string     `db:"code"`
	CampaignID  int32      `db:"campaign_id"`
	Status      string     `db:"status"` // ACTIVE, USED, EXPIRED
	GeneratedAt time.Time  `db:"generated_at"`
	UsedAt      *time.Time `db:"used_at"`
	ExpiresAt   *time.Time `db:"expires_at"`
}

func NewCoupon() *Coupon {
	return &Coupon{}
}
