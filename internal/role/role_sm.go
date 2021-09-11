package role

import "github.com/core-go/search"

type RoleSM struct {
	*search.SearchModel
	RoleId   string   `json:"roleId,omitempty" gorm:"column:roleid;primary_key" bson:"_id,omitempty" dynamodbav:"roleId,omitempty" firestore:"roleId,omitempty" validate:"max=40"`
	RoleName string   `json:"roleName,omitempty" gorm:"column:rolename" bson:"roleName,omitempty" dynamodbav:"roleName,omitempty" firestore:"roleName,omitempty" validate:"required,max=255"`
	Status   []string `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" match:"equal" validate:"required,max=1,code"`
}
