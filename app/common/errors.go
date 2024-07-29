package common

import (
	"github.com/jackc/pgx/v5/pgconn"
)

func IsUniqueKeyViolation(err error) bool {
	return err.(*pgconn.PgError).Code == "23505"
}
