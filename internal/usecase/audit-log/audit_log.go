package audit

import "time"

type AuditLog struct {
	Id        string     `json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty" validate:"max=40"`
	Resource  string     `json:"resource,omitempty" gorm:"column:resourcetype" bson:"resource,omitempty" dynamodbav:"resource,omitempty" firestore:"resource,omitempty" match:"equal"`
	UserId    string     `json:"userId,omitempty" gorm:"column:userId" bson:"userId,omitempty" dynamodbav:"userId,omitempty" firestore:"userId,omitempty"`
	Ip        string     `json:"ip,omitempty" gorm:"column:ip" bson:"ip,omitempty" dynamodbav:"ip,omitempty" firestore:"ip,omitempty" match:"equal"`
	Action    string     `json:"action,omitempty" gorm:"column:action" bson:"action,omitempty" dynamodbav:"action,omitempty" firestore:"action,omitempty" match:"equal"`
	Timestamp *time.Time `json:"timestamp,omitempty" gorm:"column:timestamp" bson:"timestamp,omitempty" dynamodbav:"timestamp,omitempty" firestore:"timestamp,omitempty"`
	Status    string     `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" match:"equal"`
	Remark    string     `json:"remark,omitempty" gorm:"column:remark" bson:"remark,omitempty" dynamodbav:"remark,omitempty" firestore:"remark,omitempty" validate:"max=255"`
}
