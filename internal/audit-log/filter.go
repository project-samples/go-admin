package audit

import "github.com/core-go/search"

type AuditLogFilter struct {
	*search.Filter
	Resource  string            `json:"resource,omitempty" gorm:"column:resource" bson:"resource,omitempty" dynamodbav:"resource,omitempty" firestore:"resource,omitempty" match:"equal"`
	Resources []string          `json:"resources,omitempty" gorm:"column:resources" bson:"resources,omitempty" dynamodbav:"resources,omitempty" firestore:"resources,omitempty"`
	Ip        string            `json:"ip,omitempty" gorm:"column:ip" bson:"ip,omitempty" dynamodbav:"ip,omitempty" firestore:"ip,omitempty" match:"equal"`
	UserId    string            `json:"userId,omitempty" gorm:"column:userid;primary_key" bson:"userId,omitempty" dynamodbav:"userId,omitempty" firestore:"userId,omitempty"`
	Users     []string          `json:"users,omitempty" gorm:"column:users" bson:"users,omitempty" dynamodbav:"users,omitempty" firestore:"users,omitempty"`
	Action    string            `json:"action,omitempty" gorm:"column:action" bson:"action,omitempty" dynamodbav:"action,omitempty" firestore:"action,omitempty" match:"equal"`
	Actions   []string          `json:"actions,omitempty" gorm:"column:actions" bson:"actions,omitempty" dynamodbav:"actions,omitempty" firestore:"actions,omitempty"`
	Time      *search.TimeRange `json:"time,omitempty" gorm:"column:time" bson:"time,omitempty" dynamodbav:"time,omitempty" firestore:"time,omitempty"`
	Status    []string          `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" match:"equal"`
}
