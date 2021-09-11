package user

import "time"

type User struct {
	UserId      string     `json:"userId,omitempty" gorm:"column:userId;primary_key" bson:"_id,omitempty" validate:"required,max=20,code"`
	Username    string     `json:"username,omitempty" gorm:"column:username" bson:"username,omitempty" dynamodbav:"username,omitempty" firestore:"username,omitempty" validate:"required,max=80"`
	Email       string     `json:"email,omitempty" gorm:"column:email" bson:"email,omitempty" dynamodbav:"email,omitempty" firestore:"email,omitempty" validate:"email,max=100"`
	DisplayName string     `json:"displayName,omitempty" gorm:"column:displayName" bson:"displayName,omitempty" dynamodbav:"displayName,omitempty" firestore:"displayName,omitempty" validate:"max=100"`
	ImageURL    string     `json:"imageURL,omitempty" gorm:"column:imageurl" bson:"imageURL,omitempty" dynamodbav:"imageURL,omitempty" firestore:"imageURL,omitempty" match:"equal"`
	Status      string     `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" match:"equal" validate:"required,max=1,code"`
	Gender      string     `json:"gender,omitempty" gorm:"column:gender" bson:"gender,omitempty" dynamodbav:"gender,omitempty" firestore:"gender,omitempty" match:"equal" validate:"required,oneof=M F"`
	Phone       string     `json:"phone,omitempty" gorm:"column:phone" bson:"phone,omitempty" dynamodbav:"phone,omitempty" firestore:"phone,omitempty" match:"equal" validate:"phone,max=16"`
	Title       string     `json:"title,omitempty" gorm:"column:title" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty" match:"equal" validate:"max=5,code"`
	Position    string     `json:"position,omitempty" gorm:"column:position" bson:"position,omitempty" dynamodbav:"position,omitempty" firestore:"position,omitempty" match:"equal" validate:"max=1,code"`
	CreatedBy   *string    `json:"createdBy,omitempty" gorm:"column:createdBy" bson:"createdBy,omitempty" dynamodbav:"createdBy,omitempty" firestore:"createdBy,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty" gorm:"column:createdAt" bson:"createdAt,omitempty" dynamodbav:"createdAt,omitempty" firestore:"createdAt,omitempty"`
	UpdatedBy   *string    `json:"updatedBy,omitempty" gorm:"column:updatedBy" bson:"updatedBy,omitempty" dynamodbav:"updatedBy,omitempty" firestore:"updatedBy,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt" bson:"updatedAt,omitempty" dynamodbav:"updatedAt,omitempty" firestore:"updatedAt,omitempty"`
	LastLogin   *time.Time `json:"lastLogin,omitempty" gorm:"lastLogin" bson:"lastLogin,omitempty" dynamodbav:"lastLogin,omitempty" firestore:"lastLogin,omitempty"`
	Roles       []string   `json:"roles,omitempty" bson:"roles,omitempty" dynamodbav:"roles,omitempty" firestore:"roles,omitempty"`
}

type UserRole struct {
	UserId string `json:"userId,omitempty" gorm:"column:userId;primary_key" bson:"_id,omitempty" validate:"required,max=20,code"`
	RoleId string `json:"roleId,omitempty" gorm:"column:roleId;primary_key" bson:"_id,omitempty" dynamodbav:"roleId,omitempty" firestore:"roleId,omitempty" validate:"max=40"`
}
