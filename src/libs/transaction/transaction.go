package repository

import (
	db "app/infrastructure/datastore"
	"errors"
)

func Transaction(txFunc func(interface{}) (interface{}, error)) (data interface{}, err error) {
	tx := db.Conn().Begin()
	if !errors.Is(tx.Error, nil) {
		return nil, tx.Error
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else if !errors.Is(err, nil) {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	data, err = txFunc(tx)
	return data, err
}
