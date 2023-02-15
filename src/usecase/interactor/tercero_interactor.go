package interactor

import (
	transaction "app/src/libs/transaction"
	models "app/src/models"
	presenter "app/src/usecase/presenter"
	repository "app/src/usecase/repository"
	"errors"
)

//Builder
type terceroInteractor struct{}

func NewTerceroInteractor() TerceroInteractor {
	return &terceroInteractor{}
}

type TerceroInteractor interface {
	GetAll(params models.TbTerceroQuery) ([]models.TbTerceroResponse, error)
	GetOne(id string) (*models.TbTerceroResponse, error)
	Create(data_insert *models.TbTercero) (*models.TbTercero, error)
	Update(id string, data_update *models.TbTercero) (*models.TbTercero, error)
	Delete(id string) (*models.TbTercero, error)
	SendEmail(id string) (*models.TbTerceroEmailResponse, error)
}

//Methods
func (ctrl terceroInteractor) GetAll(params models.TbTerceroQuery) ([]models.TbTerceroResponse, error) {

	data, err := repository.NewTerceroRepository().FindAll(params)
	if err != nil {
		return nil, err
	}
	return presenter.NewTerceroPresenter().ResponseAll(data), nil
}

func (ctrl terceroInteractor) GetOne(id string) (*models.TbTerceroResponse, error) {

	data, err := repository.NewTerceroRepository().FindOne(id)
	if err != nil {
		return nil, err
	}
	return presenter.NewTerceroPresenter().ResponseOne(data), nil
}

func (ctrl terceroInteractor) Create(data_insert *models.TbTercero) (*models.TbTercero, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewTerceroRepository().Create(data_insert)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbTercero)
	if !ok {
		return nil, errors.New("cast error Create tercero")
	}

	return presenter.NewTerceroPresenter().Response(data_resp), nil
}

func (ctrl terceroInteractor) Update(id string, data_update *models.TbTercero) (*models.TbTercero, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewTerceroRepository().Update(id, data_update)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbTercero)
	if !ok {
		return nil, errors.New("cast error Update tercero")
	}

	return presenter.NewTerceroPresenter().Response(data_resp), nil
}

func (ctrl terceroInteractor) Delete(id string) (*models.TbTercero, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewTerceroRepository().Delete(id)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbTercero)
	if !ok {
		return nil, errors.New("cast error Delete tercero")
	}

	return presenter.NewTerceroPresenter().Response(data_resp), nil
}

func (ctrl terceroInteractor) SendEmail(id string) (*models.TbTerceroEmailResponse, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewTerceroRepository().SendEmail(id)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbTerceroEmailResponse)
	if !ok {
		return nil, errors.New("cast error Delete tercero")
	}

	return presenter.NewTerceroPresenter().ResponseEmail(data_resp), nil
}
