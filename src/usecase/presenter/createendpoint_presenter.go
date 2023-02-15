package presenter

import (
	models "app/src/models"
)

//Builder
type createEndpointPresenter struct{}

func NewCreateEndpointPresenter() CreateEndpointPresenter {
	return &createEndpointPresenter{}
}

type CreateEndpointPresenter interface {
	Response(data *models.CreateEndpoint) *models.CreateEndpoint
}

//Methods
func (prs createEndpointPresenter) Response(data *models.CreateEndpoint) *models.CreateEndpoint {
	return data
}
