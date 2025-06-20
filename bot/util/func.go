package util

import (
	"bot/ent/entgen"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func BindAndValidate(c echo.Context, rq interface{}) error {
	if err := c.Bind(&rq); err != nil {
		return err
	}
	validate := validator.New()
	err := validate.Struct(rq)
	if err != nil {
		return err
	}
	return nil
}

func ExecuteTransaction(ctx context.Context, db *entgen.Client, fn func(tx *entgen.Tx) error) error {
	tx, err := db.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if pErr := recover(); pErr != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				_ = err
			}
			panic(pErr)
		}

		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = rollbackErr
			}
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				err = commitErr
			}
		}
	}()

	err = fn(tx)
	return err
}
