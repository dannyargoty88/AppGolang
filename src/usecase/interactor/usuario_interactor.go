package interactor

import (
	transaction "app/src/libs/transaction"
	models "app/src/models"
	presenter "app/src/usecase/presenter"
	repository "app/src/usecase/repository"
	"errors"
)

//Builder
type usuarioInteractor struct{}

func NewUsuarioInteractor() UsuarioInteractor {
	return &usuarioInteractor{}
}

type UsuarioInteractor interface {
	GetAll(params models.TbUsuarioQuery) ([]models.TbUsuarioResponse, error)
	GetOne(id string) (*models.TbUsuarioResponse, error)
	Create(data_insert *models.TbUsuario) (*models.TbUsuario, error)
	Update(id string, data_update *models.TbUsuario) (*models.TbUsuario, error)
	Delete(id string) (*models.TbUsuario, error)
}

//Methods
func (ctrl usuarioInteractor) GetAll(params models.TbUsuarioQuery) ([]models.TbUsuarioResponse, error) {

	data, err := repository.NewUsuarioRepository().FindAll(params)
	if err != nil {
		return nil, err
	}
	return presenter.NewUsuarioPresenter().ResponseAll(data), nil
}

func (ctrl usuarioInteractor) GetOne(id string) (*models.TbUsuarioResponse, error) {

	data, err := repository.NewUsuarioRepository().FindOne(id)
	if err != nil {
		return nil, err
	}
	return presenter.NewUsuarioPresenter().ResponseOne(data), nil
}

func (ctrl usuarioInteractor) Create(data_insert *models.TbUsuario) (*models.TbUsuario, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewUsuarioRepository().Create(data_insert)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbUsuario)
	if !ok {
		return nil, errors.New("cast error Create usuario")
	}

	return presenter.NewUsuarioPresenter().Response(data_resp), nil
}

func (ctrl usuarioInteractor) Update(id string, data_update *models.TbUsuario) (*models.TbUsuario, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewUsuarioRepository().Update(id, data_update)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbUsuario)
	if !ok {
		return nil, errors.New("cast error Update usuario")
	}

	return presenter.NewUsuarioPresenter().Response(data_resp), nil
}

func (ctrl usuarioInteractor) Delete(id string) (*models.TbUsuario, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewUsuarioRepository().Delete(id)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.TbUsuario)
	if !ok {
		return nil, errors.New("cast error Delete usuario")
	}

	return presenter.NewUsuarioPresenter().Response(data_resp), nil
}
