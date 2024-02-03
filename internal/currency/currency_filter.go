package currency

import . "github.com/core-go/search"

type CurrencyFilter struct {
	*Filter
    Code *string `mapstructure:"code" json:"code,omitempty" gorm:"column:code;primary_key" bson:"_id,omitempty" dynamodbav:"code,omitempty" firestore:"code,omitempty" avro:"code" validate:"required,max=40"`
    CurrencySymbol *string `mapstructure:"currency_symbol" json:"currencySymbol,omitempty" gorm:"column:currency_symbol" bson:"CurrencySymbol,omitempty" dynamodbav:"currencySymbol,omitempty" firestore:"currencySymbol,omitempty" avro:"currencySymbol" validate:"required,max=6"`
    DecimalDigits *int16 `mapstructure:"decimal_digits" json:"decimalDigits,omitempty" gorm:"column:decimal_digits" bson:"DecimalDigits,omitempty" dynamodbav:"decimalDigits,omitempty" firestore:"decimalDigits,omitempty" avro:"decimalDigits" validate:"required"`
	Status []string `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" match:"equal"`
}
