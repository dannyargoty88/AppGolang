package interactor

import (
	transaction "app/src/libs/transaction"
	models "app/src/models"
	presenter "app/src/usecase/presenter"
	repository "app/src/usecase/repository"
	"errors"
)

//Builder
type ciudadInteractor struct{}

func NewCiudadInteractor() CiudadInteractor {
	return &ciudadInteractor{}
}

type CiudadInteractor interface {
	GetAll(params models.TbCiudadQuery) ([]models.TbCiudadResponse, error)
	GetOne(id string) (*models.TbCiudadResponse, error)
	Create(data_insert *models.TbCiudad) (*models.TbCiudad, error)
	Update(id string, data_update *models.TbCiudad) (*models.TbCiudad, error)
	Delete(id string) (*models.TbCiudad, error)
}

//Methods
func (ctrl ciudadInteractor) GetAll(params models.TbCiudadQuery) ([]models.TbCiudadResponse, error) {

	data, err := repository.NewCiudadRepository().FindAll(params)
	if err != nil {
		return nil, err
	}
	return presenter.NewCiudadPresenter().ResponseAll(data), nil
}

func (ctrl ciudadInteractor) GetOne(id string) (*models.TbCiudadResponse, error) {

	data, err := repository.NewCiudadRepository().FindOne(id)
	if err != nil {
		return nil, err
	}
	return presenter.NewCiudadPresenter().ResponseOne(data), nil
}

func (ctrl ciudadInteractor) Create(data_insert *models.TbCiudad) (*models.TbCiudad, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewCiudadRepository().Create(data_insert)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbCiudad)
	if !ok {
		return nil, errors.New("cast error Create ciudad")
	}

	return presenter.NewCiudadPresenter().Response(data_resp), nil
}

func (ctrl ciudadInteractor) Update(id string, data_update *models.TbCiudad) (*models.TbCiudad, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewCiudadRepository().Update(id, data_update)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbCiudad)
	if !ok {
		return nil, errors.New("cast error Update ciudad")
	}

	return presenter.NewCiudadPresenter().Response(data_resp), nil
}

func (ctrl ciudadInteractor) Delete(id string) (*models.TbCiudad, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewCiudadRepository().Delete(id)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbCiudad)
	if !ok {
		return nil, errors.New("cast error Delete ciudad")
	}

	return presenter.NewCiudadPresenter().Response(data_resp), nil
}
