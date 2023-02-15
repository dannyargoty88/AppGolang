package interactor

import (
	transaction "app/src/libs/transaction"
	models "app/src/models"
	presenter "app/src/usecase/presenter"
	repository "app/src/usecase/repository"
	"errors"
)

//Builder
type departamentoInteractor struct{}

func NewDepartamentoInteractor() DepartamentoInteractor {
	return &departamentoInteractor{}
}

type DepartamentoInteractor interface {
	GetAll(params models.TbDepartamentoQuery) ([]models.TbDepartamentoResponse, error)
	GetOne(id string) (*models.TbDepartamentoResponse, error)
	Create(data_insert *models.TbDepartamento) (*models.TbDepartamento, error)
	Update(id string, data_update *models.TbDepartamento) (*models.TbDepartamento, error)
	Delete(id string) (*models.TbDepartamento, error)
	GetAllDepartamentoCiudad(params models.TbDepartamentoQuery) ([]models.TbDepartamentoCiudadResponse, error)
}

//Methods
func (ctrl departamentoInteractor) GetAll(params models.TbDepartamentoQuery) ([]models.TbDepartamentoResponse, error) {

	data, err := repository.NewDepartamentoRepository().FindAll(params)
	if err != nil {
		return nil, err
	}
	return presenter.NewDepartamentoPresenter().ResponseAll(data), nil
}

func (ctrl departamentoInteractor) GetOne(id string) (*models.TbDepartamentoResponse, error) {

	data, err := repository.NewDepartamentoRepository().FindOne(id)
	if err != nil {
		return nil, err
	}
	return presenter.NewDepartamentoPresenter().ResponseOne(data), nil
}

func (ctrl departamentoInteractor) Create(data_insert *models.TbDepartamento) (*models.TbDepartamento, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewDepartamentoRepository().Create(data_insert)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbDepartamento)
	if !ok {
		return nil, errors.New("cast error Create departamento")
	}

	return presenter.NewDepartamentoPresenter().Response(data_resp), nil
}

func (ctrl departamentoInteractor) Update(id string, data_update *models.TbDepartamento) (*models.TbDepartamento, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewDepartamentoRepository().Update(id, data_update)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbDepartamento)
	if !ok {
		return nil, errors.New("cast error Update departamento")
	}

	return presenter.NewDepartamentoPresenter().Response(data_resp), nil
}

func (ctrl departamentoInteractor) Delete(id string) (*models.TbDepartamento, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewDepartamentoRepository().Delete(id)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbDepartamento)
	if !ok {
		return nil, errors.New("cast error Delete departamento")
	}

	return presenter.NewDepartamentoPresenter().Response(data_resp), nil
}

func (ctrl departamentoInteractor) GetAllDepartamentoCiudad(params models.TbDepartamentoQuery) ([]models.TbDepartamentoCiudadResponse, error) {

	data, err := repository.NewDepartamentoRepository().FindAllDepartamentoCiudad(params)
	if err != nil {
		return nil, err
	}
	return presenter.NewDepartamentoPresenter().ResponseAllDepartamentoCiudad(data), nil
}
