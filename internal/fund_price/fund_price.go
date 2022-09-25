package fund

import (
	"time"
)

type User struct {
	FundID             string     `json:"FundID,omitempty" gorm:"column:FUNDID;primary_key" `
	EffectiveDate      *time.Time `json:"EffectiveDate,omitempty" gorm:"column:EFFECTIVEDATE"  match:"equal"`
	AuthStat           *time.Time `json:"AuthStat,omitempty" gorm:"column:AUTH_STAT"  match:"equal"`
	RealNAVPerUnit     float64    `json:"RealNAVPerUnit,omitempty" gorm:"column:REALNAVPERUNIT"  match:"equal"`
	DeclaredNAV        float64    `json:"DeclaredNAV,omitempty" gorm:"column:DECLAREDNAV"  match:"equal"`
	TotalNETAsset      *time.Time `json:"TotalNETAsset,omitempty" gorm:"column:TOTALNETASSET"  match:"equal"`
	OutStandingUnits   float64    `json:"OutStandingUnits,omitempty" gorm:"column:OUTSTANDINGUNITS"  match:"equal"`
	FIOutStanding      float64    `json:"FIOutStanding,omitempty" gorm:"column:FIOUTSTANDINGUNITS"  match:"equal"`
	NumberOfUnitHolder string     `json:"NumberOfUnitHolder,omitempty" gorm:"column:NOOFUNITHOLDERS"  match:"equal"`
	LastedPrice        string     `json:"LastedPrice,omitempty" gorm:"column:LATESTPRICE"  match:"equal"`
	FloorPrice         float64    `json:"FloorPrice,omitempty" gorm:"column:FLOORPRICE"  match:"equal"`
	CeilingPrice       float64    `json:"CeilingPrice,omitempty" gorm:"column:CEILINGPRICE"  match:"equal"`
}
