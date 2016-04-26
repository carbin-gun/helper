package database

import (
	"database/sql"

	"github.com/lib/pq"
)

//IsNoMatches charge whether the error is a ErrNoRows error.
func IsNoRows(err error) bool {
	if err == sql.ErrNoRows {
		return true
	}
	return false

}

//IsUniqueError is the error for unique violation(including primary key)
func IsUniqueError(err error) bool {
	switch t := err.(type) {
	case *pq.Error:
		if t.Code == "23505" {
			return true
		}
		return false
	default:
		return false
	}
}
