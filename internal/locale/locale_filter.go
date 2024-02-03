package locale

import . "github.com/core-go/search"

type LocaleFilter struct {
	*Filter
    Code *string `mapstructure:"code" json:"code,omitempty" gorm:"column:code;primary_key" bson:"_id,omitempty" dynamodbav:"code,omitempty" firestore:"code,omitempty" avro:"code" validate:"required,max=40"`
    CountryCode *string `mapstructure:"country_code" json:"countryCode,omitempty" gorm:"column:country_code" bson:"CountryCode,omitempty" dynamodbav:"countryCode,omitempty" firestore:"countryCode,omitempty" avro:"countryCode" validate:"required,max=5"`
    CountryName *string `mapstructure:"country_name" json:"countryName,omitempty" gorm:"column:country_name" bson:"CountryName,omitempty" dynamodbav:"countryName,omitempty" firestore:"countryName,omitempty" avro:"countryName" validate:"required,max=255"`
    NativeCountryName *string `mapstructure:"native_country_name" json:"nativeCountryName,omitempty" gorm:"column:native_country_name" bson:"NativeCountryName,omitempty" dynamodbav:"nativeCountryName,omitempty" firestore:"nativeCountryName,omitempty" avro:"nativeCountryName" validate:"required,max=255"`
    Name *string `mapstructure:"name" json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty" avro:"name" validate:"required,max=255"`
    NativeName *string `mapstructure:"native_name" json:"nativeName,omitempty" gorm:"column:native_name" bson:"NativeName,omitempty" dynamodbav:"nativeName,omitempty" firestore:"nativeName,omitempty" avro:"nativeName" validate:"required,max=255"`
    DateFormat *string `mapstructure:"date_format" json:"dateFormat,omitempty" gorm:"column:date_format" bson:"DateFormat,omitempty" dynamodbav:"dateFormat,omitempty" firestore:"dateFormat,omitempty" avro:"dateFormat" validate:"required,max=14"`
    FirstDayOfWeek *int16 `mapstructure:"first_day_of_week" json:"firstDayOfWeek,omitempty" gorm:"column:first_day_of_week" bson:"FirstDayOfWeek,omitempty" dynamodbav:"firstDayOfWeek,omitempty" firestore:"firstDayOfWeek,omitempty" avro:"firstDayOfWeek" validate:"required"`
    DecimalSeparator *string `mapstructure:"decimal_separator" json:"decimalSeparator,omitempty" gorm:"column:decimal_separator" bson:"DecimalSeparator,omitempty" dynamodbav:"decimalSeparator,omitempty" firestore:"decimalSeparator,omitempty" avro:"decimalSeparator" validate:"required,max=3"`
    GroupSeparator *string `mapstructure:"group_separator" json:"groupSeparator,omitempty" gorm:"column:group_separator" bson:"GroupSeparator,omitempty" dynamodbav:"groupSeparator,omitempty" firestore:"groupSeparator,omitempty" avro:"groupSeparator" validate:"required,max=3"`
    CurrencyCode *string `mapstructure:"currency_code" json:"currencyCode,omitempty" gorm:"column:currency_code" bson:"CurrencyCode,omitempty" dynamodbav:"currencyCode,omitempty" firestore:"currencyCode,omitempty" avro:"currencyCode" validate:"required,max=3"`
    CurrencySymbol *string `mapstructure:"currency_symbol" json:"currencySymbol,omitempty" gorm:"column:currency_symbol" bson:"CurrencySymbol,omitempty" dynamodbav:"currencySymbol,omitempty" firestore:"currencySymbol,omitempty" avro:"currencySymbol" validate:"required,max=6"`
    CurrencyDecimalDigits *int16 `mapstructure:"currency_decimal_digits" json:"currencyDecimalDigits,omitempty" gorm:"column:currency_decimal_digits" bson:"CurrencyDecimalDigits,omitempty" dynamodbav:"currencyDecimalDigits,omitempty" firestore:"currencyDecimalDigits,omitempty" avro:"currencyDecimalDigits" validate:"required"`
    CurrencyPattern *int16 `mapstructure:"currency_pattern" json:"currencyPattern,omitempty" gorm:"column:currency_pattern" bson:"CurrencyPattern,omitempty" dynamodbav:"currencyPattern,omitempty" firestore:"currencyPattern,omitempty" avro:"currencyPattern" validate:"required"`
    CurrencySample *string `mapstructure:"currency_sample" json:"currencySample,omitempty" gorm:"column:currency_sample" bson:"CurrencySample,omitempty" dynamodbav:"currencySample,omitempty" firestore:"currencySample,omitempty" avro:"currencySample" validate:"required,max=40"`
}
