package errors_utils

import (
	"database/sql"
	"errors"
	"fmt"
)

func CheckNotFoundError(err error, entity string) error {
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("cущность %s не найдена", entity)
	}
	return nil
}
