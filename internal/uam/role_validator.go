package uam

import (
	"context"
	"database/sql"

	sv "github.com/core-go/service"
	s "github.com/core-go/sql"
)

type RoleValidator struct {
	db       *sql.DB
	query    string
	validate func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error)
}

func NewRoleValidator(db *sql.DB, query string, validate func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error)) *RoleValidator {
	driver := s.GetDriver(db)
	q := s.ReplaceQueryArgs(driver, query)
	return &RoleValidator{db: db, query: q, validate: validate}
}

func (v *RoleValidator) Validate(ctx context.Context, req interface{}) (errors []sv.ErrorMessage, err error) {
	errors, err = v.validate(ctx, req)
	if err != nil {
		return errors, err
	}
	i := 0
	role, ok := req.(*Role)
	if !ok {
		return errors, err
	}
	rows, err := v.db.Query(v.query, role.RoleName, role.RoleId)
	if err != nil {
		return errors, err
	}
	for rows.Next() {
		err := rows.Scan(&i)
		if err != nil {
			return errors, err
		}
		if i > 0 {
			er := sv.ErrorMessage{Field: "roleName", Code: "duplicate"}
			return append(errors, er), nil
		}
	}
	return errors, err
}
