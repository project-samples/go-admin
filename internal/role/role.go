package role

import "time"

type Role struct {
	RoleId     string     `json:"roleId,omitempty" gorm:"column:roleId;primary_key" bson:"_id,omitempty" dynamodbav:"roleId,omitempty" firestore:"roleId,omitempty" validate:"max=40"`
	RoleName   string     `json:"roleName,omitempty" gorm:"column:roleName" bson:"roleName,omitempty" dynamodbav:"roleName,omitempty" firestore:"roleName,omitempty" validate:"required,max=255"`
	Status     string     `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" match:"equal" validate:"required,max=1,code"`
	Remark     *string    `json:"remark,omitempty" gorm:"column:remark" bson:"remark,omitempty" dynamodbav:"remark,omitempty" firestore:"remark,omitempty" validate:"max=255"`
	CreatedBy  *string    `json:"createdBy,omitempty" gorm:"column:createdBy" bson:"createdBy,omitempty" dynamodbav:"createdBy,omitempty" firestore:"createdBy,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty" gorm:"column:createdAt" bson:"createdAt,omitempty" dynamodbav:"createdAt,omitempty" firestore:"createdAt,omitempty"`
	UpdatedBy  *string    `json:"updatedBy,omitempty" gorm:"column:updatedBy" bson:"updatedBy,omitempty" dynamodbav:"updatedBy,omitempty" firestore:"updatedBy,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt" bson:"updatedAt,omitempty" dynamodbav:"updatedAt,omitempty" firestore:"updatedAt,omitempty"`
	Privileges []string   `json:"privileges,omitempty" bson:"privileges,omitempty" dynamodbav:"privileges,omitempty" firestore:"privileges,omitempty"`
}

type RoleModule struct {
	RoleId      string `json:"roleId,omitempty" gorm:"column:roleId" bson:"roleId,omitempty" dynamodbav:"roleId,omitempty" firestore:"roleId,omitempty" validate:"required"`
	ModuleId    string `json:"moduleId,omitempty" gorm:"column:moduleId" bson:"moduleId,omitempty" dynamodbav:"moduleId,omitempty" firestore:"moduleId,omitempty" validate:"required"`
	Permissions int32  `json:"permissions,omitempty" gorm:"column:permissions" bson:"permissions,omitempty" dynamodbav:"permissions,omitempty" firestore:"permissions,omitempty" validate:"required"`
}
