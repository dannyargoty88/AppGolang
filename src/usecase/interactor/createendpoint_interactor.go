package interactor

import (
	transaction "app/src/libs/transaction"
	models "app/src/models"
	presenter "app/src/usecase/presenter"
	repository "app/src/usecase/repository"
	"errors"
)

//Builder
type createEndpointInteractor struct{}

func NewCreateEndpointInteractor() CreateEndpointInteractor {
	return &createEndpointInteractor{}
}

type CreateEndpointInteractor interface {
	Create(data_insert *models.CreateEndpoint) (*models.CreateEndpoint, error)
}

//Methods
func (ctrl createEndpointInteractor) Create(data_insert *models.CreateEndpoint) (*models.CreateEndpoint, error) {

	data_transaction, err := transaction.Transaction(func(i interface{}) (interface{}, error) {
		res, err := repository.NewCreateEndpointRepository().Create(data_insert)
		return res, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	data_resp, ok := data_transaction.(*models.CreateEndpoint)
	if !ok {
		return nil, errors.New("cast error Create createEndpoint")
	}

	return presenter.NewCreateEndpointPresenter().Response(data_resp), nil
}
