package interactor

import (
	transaction "app/src/libs/transaction"
	models "app/src/models"
	presenter "app/src/usecase/presenter"
	repository "app/src/usecase/repository"
	"errors"
)

//Builder
type errorInteractor struct{}

func NewErrorInteractor() ErrorInteractor {
	return &errorInteractor{}
}

type ErrorInteractor interface {
	GetAll(params models.TbErrorQuery) ([]models.TbError, error)
	GetOne(id string) (*models.TbError, error)
	Delete(id string) (*models.TbError, error)
}

//Methods
func (ctrl errorInteractor) GetAll(params models.TbErrorQuery) ([]models.TbError, error) {

	data, err := repository.NewErrorRepository().FindAll(params)
	if err != nil {
		return nil, err
	}
	return presenter.NewErrorPresenter().ResponseAll(data), nil
}

func (ctrl errorInteractor) GetOne(id string) (*models.TbError, error) {

	data, err := repository.NewErrorRepository().FindOne(id)
	if err != nil {
		return nil, err
	}
	return presenter.NewErrorPresenter().ResponseOne(data), nil
}

func (ctrl errorInteractor) Delete(id string) (*models.TbError, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewErrorRepository().Delete(id)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbError)
	if !ok {
		return nil, errors.New("cast error Delete error")
	}

	return presenter.NewErrorPresenter().Response(data_resp), nil
}
