package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"strings"
	"usersApi/utils/errors"
)

const (
	errNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errNoRows) {
			return errors.NewNotFoundError("no records matching")
		}
		return errors.NewInternalServerError("Error parsing data")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("Error processing request")
}
