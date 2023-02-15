package interactor

import (
	models "app/src/models"
	presenter "app/src/usecase/presenter"
	repository "app/src/usecase/repository"
)

//Builder
type itemWsInteractor struct{}

func NewItemWsInteractor() ItemWsInteractor {
	return &itemWsInteractor{}
}

type ItemWsInteractor interface {
	GetAll(params models.TbItemwsQuery) ([]models.TbItemwsResponse, error)
	GetOne(ws string, id string) (*models.TbItemwsResponse, error)
}

//Methods
func (ctrl itemWsInteractor) GetAll(params models.TbItemwsQuery) ([]models.TbItemwsResponse, error) {

	data, err := repository.NewItemWsRepository().FindAll(params)
	if err != nil {
		return nil, err
	}
	return presenter.NewItemWsPresenter().ResponseAll(data), nil
}

func (ctrl itemWsInteractor) GetOne(ws string, id string) (*models.TbItemwsResponse, error) {

	data, err := repository.NewItemWsRepository().FindOne(ws, id)
	if err != nil {
		return nil, err
	}
	return presenter.NewItemWsPresenter().ResponseOne(data), nil
}
