package app

import (
	. "github.com/core-go/auth"
	. "github.com/core-go/auth/ldap"
	as "github.com/core-go/auth/sql"
	"github.com/core-go/code"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/log/zap"
	sv "github.com/core-go/service"
	"github.com/core-go/service/audit"
	. "github.com/core-go/service/model-builder"
	"github.com/core-go/sql"
)

type Root struct {
	Server              sv.ServerConfig       `mapstructure:"server"`
	SecuritySkip        bool                  `mapstructure:"security_skip"`
	Ldap                LDAPConfig            `mapstructure:"ldap"`
	Status              *StatusConfig         `mapstructure:"status"`
	Auth                as.SqlConfig          `mapstructure:"auth"`
	Token               TokenConfig           `mapstructure:"token"`
	Payload             PayloadConfig         `mapstructure:"payload"`
	DB                  sql.Config            `mapstructure:"db"`
	Log                 log.Config            `mapstructure:"log"`
	MiddleWare          mid.LogConfig         `mapstructure:"middleware"`
	AutoRoleId          *bool                 `mapstructure:"auto_role_id"`
	AutoUserId          *bool                 `mapstructure:"auto_user_id"`
	Role                code.Config           `mapstructure:"role"`
	Code                code.Config           `mapstructure:"code"`
	AuditLog            sql.ActionLogConf     `mapstructure:"audit_log"`
	AuditClient         audit.AuditLogClient  `mapstructure:"audit_client"`
	Writer              sv.WriterConfig       `mapstructure:"writer"`
	Tracking            TrackingConfig        `mapstructure:"tracking"`
	Sql                 SqlStatement          `mapstructure:"sql"`
}
type SqlStatement struct {
	Privileges        string        `mapstructure:"privileges"`
	PrivilegesByUser  string        `mapstructure:"privileges_by_user"`
	PermissionsByUser string        `mapstructure:"permissions_by_user"`
	User              string        `mapstructure:"user"`
	Role              RoleStatement `mapstructure:"role"`
}
type RoleStatement struct {
	Duplicate string `mapstructure:"duplicate"`
	Check     string `mapstructure:"check"`
}
