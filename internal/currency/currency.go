package currency

type Currency struct {
    Code string `mapstructure:"code" json:"code,omitempty" gorm:"column:code;primary_key" bson:"_id,omitempty" dynamodbav:"code,omitempty" firestore:"code,omitempty" avro:"code" validate:"required,max=40"`
    Symbol string `mapstructure:"symbol" json:"symbol,omitempty" gorm:"column:symbol" bson:"symbol,omitempty" dynamodbav:"symbol,omitempty" firestore:"symbol,omitempty" avro:"symbol" validate:"required,max=6"`
    DecimalDigits int16 `mapstructure:"decimal_digits" json:"decimalDigits,omitempty" gorm:"column:decimal_digits" bson:"DecimalDigits,omitempty" dynamodbav:"decimalDigits,omitempty" firestore:"decimalDigits,omitempty" avro:"decimalDigits"`
    Status string `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty"`
}
