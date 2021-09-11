package user

import "github.com/core-go/search"

type UserSM struct {
	*search.SearchModel
	UserId      string   `json:"userId,omitempty" gorm:"column:userid;primary_key" bson:"_id,omitempty" validate:"required,max=20,code"`
	Username    string   `json:"username,omitempty" gorm:"column:username" bson:"username,omitempty" dynamodbav:"username,omitempty" firestore:"username,omitempty" validate:"required,max=80"`
	Email       string   `json:"email,omitempty" gorm:"column:email" bson:"email,omitempty" dynamodbav:"email,omitempty" firestore:"email,omitempty" validate:"email,max=100"`
	DisplayName string   `json:"displayName,omitempty" gorm:"column:displayname" bson:"displayName,omitempty" dynamodbav:"displayName,omitempty" firestore:"displayName,omitempty" validate:"max=100"`
	Status      []string `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" match:"equal" validate:"required,max=1,code"`
}
