package domain

import (
	"context"
	"time"
)

type DistrictData struct {
	Id        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

// UserUsecase represent the user's usecases
type RegionUsecase interface {
	GetDistrict(ctx context.Context) ([]*DistrictData, error)
}

// UserRepository represent the user's repository contract
type RegionRepository interface {
	GetDistrict(ctx context.Context) ([]*DistrictData, error)
}