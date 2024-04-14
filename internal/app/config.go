package app

import (
	"time"

	"github.com/core-go/auth/ldap"
	q "github.com/core-go/auth/sql"
	"github.com/core-go/core"
	"github.com/core-go/core/audit"
	"github.com/core-go/core/builder"
	"github.com/core-go/core/code"
	"github.com/core-go/core/cors"
	redis "github.com/core-go/core/redis/v8"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/log/zap"
	"github.com/core-go/sql"
	sa "github.com/core-go/sql/action"
)

type Config struct {
	Server       core.ServerConf        `mapstructure:"server"`
	Allow        cors.AllowConfig       `mapstructure:"allow"`
	SecuritySkip bool                   `mapstructure:"security_skip"`
	Template     bool                   `mapstructure:"template"`
	Ldap         ldap.LDAPConfig        `mapstructure:"ldap"`
	Auth         q.SqlAuthConfig        `mapstructure:"auth"`
	DB           sql.Config             `mapstructure:"db"`
	Log          log.Config             `mapstructure:"log"`
	MiddleWare   mid.LogConfig          `mapstructure:"middleware"`
	AutoRoleId   *bool                  `mapstructure:"auto_role_id"`
	AutoUserId   *bool                  `mapstructure:"auto_user_id"`
	Role         code.Config            `mapstructure:"role"`
	Code         code.Config            `mapstructure:"code"`
	AuditLog     sa.ActionLogConf       `mapstructure:"audit_log"`
	AuditClient  audit.AuditLogClient   `mapstructure:"audit_client"`
	Action       *core.ActionConfig     `mapstructure:"action"`
	Tracking     builder.TrackingConfig `mapstructure:"tracking"`
	Sql          SqlStatement           `mapstructure:"sql"`
	Session      SessionConfig          `mapstructure:"session"`
	Redis        redis.Config           `mapstructure:"redis"`
}
type SqlStatement struct {
	Privileges        string        `mapstructure:"privileges"`
	PrivilegesByUser  string        `mapstructure:"privileges_by_user"`
	PermissionsByUser string        `mapstructure:"permissions_by_user"`
	Role              RoleStatement `mapstructure:"role"`
}
type RoleStatement struct {
	Check string `mapstructure:"check"`
}
type SessionConfig struct {
	Secret      string        `yaml:"secret" mapstructure:"secret"`
	ExpiredTime time.Duration `mapstructure:"expired_time"`
	Host        string        `mapstructure:"host"`
}
