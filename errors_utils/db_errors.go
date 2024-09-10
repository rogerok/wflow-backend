package errors_utils

import (
	"database/sql"
	"errors"
	"fmt"
)

func CheckErrorNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func CheckNotFoundError(err error, entity string) error {
	if CheckErrorNoRows(err) {
		return fmt.Errorf("cущность %s не найдена", entity)
	}
	return nil
}
