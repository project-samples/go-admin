package fund

import (
	"time"
)

type Fund struct {
	FundID            string     `json:"FundID,omitempty" gorm:"column:FUNDID;primary_key" validate:"required,max=20,code"`
	EffectiveDate     *time.Time `json:"EffectiveDate,omitempty" gorm:"column:RULEEFFECTIVEDATE" `
	FundRuleStartDate *time.Time `json:"FundRuleStartDate,omitempty" gorm:"column:FUNDSTARTDATE"  `
	FundName          string     `json:"FundName,omitempty" gorm:"column:FUNDNAME"    `
	FundNameShort     string     `json:"FundNameShort,omitempty" gorm:"column:FUNDNAMESHORT"    match:"equal"`
	FundClass         string     `json:"FundClass,omitempty" gorm:"column:FUNDCLASS"   match:"equal" `
	FundType          string     `json:"FundType,omitempty" gorm:"column:FUNDTYPE"    match:"equal" `
	FundBaseCurrency  string     `json:"FundBaseCurrency,omitempty" gorm:"column:FUNDBASECURRENCY"    match:"equal" `
	FundContry        string     `json:"FundContry,omitempty" gorm:"column:FUNDCOUNTRY"   match:"equal" `
	FiscalStartYear   *time.Time `json:"FiscalStartYear,omitempty" gorm:"column:FISCALSTARTYEAR"   match:"equal" `
	FiscalEndYear     *time.Time `json:"FiscalEndYear,omitempty" gorm:"column:FISCALENDYEAR"   match:"equal" `
	AMCID             string     `json:"AMCID,omitempty" gorm:"column:AMCID"   match:"equal" `
	FundEnabled       int8       `json:"FundEnabled,omitempty" gorm:"column:FUNDENABLED"   match:"equal" `
	LastestRule       int8       `json:"LastestRule,omitempty" gorm:"column:LATESTRULE"   match:"equal" `
}
