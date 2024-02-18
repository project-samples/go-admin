package currency

import . "github.com/core-go/search"

type CurrencyFilter struct {
	*Filter
    Code *string `mapstructure:"code" json:"code,omitempty" gorm:"column:code;primary_key" bson:"_id,omitempty" dynamodbav:"code,omitempty" firestore:"code,omitempty" avro:"code" validate:"required,max=40" operator:"like" q:"true"`
	Symbol string `mapstructure:"symbol" json:"symbol,omitempty" gorm:"column:symbol" bson:"symbol,omitempty" dynamodbav:"symbol,omitempty" firestore:"symbol,omitempty" avro:"symbol" validate:"required,max=6" operator:"=" q:"true"`
    DecimalDigits *int `mapstructure:"decimal_digits" json:"decimalDigits,omitempty" gorm:"column:decimal_digits" bson:"DecimalDigits,omitempty" dynamodbav:"decimalDigits,omitempty" firestore:"decimalDigits,omitempty" avro:"decimalDigits" validate:"required" operator:"="`
	Status []string `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" match:"equal"`
}
