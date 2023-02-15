package interactor

import (
	transaction "app/src/libs/transaction"
	models "app/src/models"
	presenter "app/src/usecase/presenter"
	repository "app/src/usecase/repository"
	"errors"
)

//Builder
type wsInteractor struct{}

func NewWsInteractor() WsInteractor {
	return &wsInteractor{}
}

type WsInteractor interface {
	GetAll(params models.TbWsQuery) ([]models.TbWsResponse, error)
	GetOne(id string) (*models.TbWsResponse, error)
	Create(data_insert *models.TbWs) (*models.TbWs, error)
	Update(id string, data_update *models.TbWs) (*models.TbWs, error)
	Delete(id string) (*models.TbWs, error)
}

//Methods
func (ctrl wsInteractor) GetAll(params models.TbWsQuery) ([]models.TbWsResponse, error) {

	data, err := repository.NewWsRepository().FindAll(params)
	if err != nil {
		return nil, err
	}
	return presenter.NewWsPresenter().ResponseAll(data), nil
}

func (ctrl wsInteractor) GetOne(id string) (*models.TbWsResponse, error) {

	data, err := repository.NewWsRepository().FindOne(id)
	if err != nil {
		return nil, err
	}
	return presenter.NewWsPresenter().ResponseOne(data), nil
}

func (ctrl wsInteractor) Create(data_insert *models.TbWs) (*models.TbWs, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewWsRepository().Create(data_insert)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbWs)
	if !ok {
		return nil, errors.New("cast error Create ws")
	}

	return presenter.NewWsPresenter().Response(data_resp), nil
}

func (ctrl wsInteractor) Update(id string, data_update *models.TbWs) (*models.TbWs, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewWsRepository().Update(id, data_update)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbWs)
	if !ok {
		return nil, errors.New("cast error Update ws")
	}

	return presenter.NewWsPresenter().Response(data_resp), nil
}

func (ctrl wsInteractor) Delete(id string) (*models.TbWs, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewWsRepository().Delete(id)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbWs)
	if !ok {
		return nil, errors.New("cast error Delete ws")
	}

	return presenter.NewWsPresenter().Response(data_resp), nil
}
