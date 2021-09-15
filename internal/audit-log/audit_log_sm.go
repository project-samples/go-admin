package audit

import "github.com/core-go/search"

type AuditLogSM struct {
	*search.SearchModel
	Resource  string            `json:"resource,omitempty" gorm:"column:resourcetype" bson:"resource,omitempty" dynamodbav:"resource,omitempty" firestore:"resource,omitempty" match:"equal"`
	UserId    string            `json:"userId,omitempty" gorm:"column:userid;primary_key" bson:"userId,omitempty" dynamodbav:"userId,omitempty" firestore:"userId,omitempty"`
	Ip        string            `json:"ip,omitempty" gorm:"column:ip" bson:"ip,omitempty" dynamodbav:"ip,omitempty" firestore:"ip,omitempty" match:"equal"`
	Action    string            `json:"action,omitempty" gorm:"column:action" bson:"action,omitempty" dynamodbav:"action,omitempty" firestore:"action,omitempty" match:"equal"`
	Timestamp *search.TimeRange `json:"timestamp,omitempty" gorm:"column:timestamp" bson:"timestamp,omitempty" dynamodbav:"timestamp,omitempty" firestore:"timestamp,omitempty"`
	Status    string            `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" match:"equal"`
}
