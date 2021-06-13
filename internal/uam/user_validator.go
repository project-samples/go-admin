package uam

import (
	"context"
	"database/sql"

	sv "github.com/core-go/service"
	s "github.com/core-go/sql"
)

type UserValidator struct {
	db       *sql.DB
	query    string
	validate func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error)
}

func NewUserValidator(db *sql.DB, query string, validate func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error)) *UserValidator {
	driver := s.GetDriver(db)
	q := s.ReplaceQueryArgs(driver, query)
	return &UserValidator{db, q, validate}
}

func (v *UserValidator) Validate(ctx context.Context, req interface{}) (errors []sv.ErrorMessage, err error) {
	errors, err = v.validate(ctx, req)
	if err != nil {
		return errors, err
	}
	user, ok := req.(*User)
	if !ok {
		return errors, err
	}
	i := 0
	rows, err := v.db.Query(v.query, user.Username, user.UserId)
	if err != nil {
		return errors, err
	}
	for rows.Next() {
		err := rows.Scan(&i)
		if err != nil {
			return errors, err
		}
		if i > 0 {
			er := sv.ErrorMessage{Field: "username", Code: "duplicate"}
			return append(errors, er), nil
		}
	}
	return errors, err
}
