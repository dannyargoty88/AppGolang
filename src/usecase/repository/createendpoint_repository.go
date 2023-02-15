package repository

import (
	models "app/src/models"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

//Builder
type createEndpointRepository struct{}

func NewCreateEndpointRepository() CreateEndpointRepository {
	return &createEndpointRepository{}
}

type CreateEndpointRepository interface {
	Create(data *models.CreateEndpoint) (*models.CreateEndpoint, error)
}

//Methods
func (repo createEndpointRepository) Create(data *models.CreateEndpoint) (*models.CreateEndpoint, error) {

	//Routes
	err := createFile(data,
		"./infrastructure/storage/template/file/route/route_template.txt",
		"./infrastructure/routes/"+data.Name_folderroute+"/"+data.Name_file+"/"+data.Name_file+"_routes.go", "SI")
	if err != nil {
		return nil, err
	}

	//Model
	err = createFile(data,
		"./infrastructure/storage/template/file/src/model/model_template.txt",
		"./src/models/"+data.Name_file+"_models.go", "NO")
	if err != nil {
		return nil, err
	}

	//Controller
	err = createFile(data,
		"./infrastructure/storage/template/file/src/controller/controller_template.txt",
		"./src/controllers/"+data.Name_file+"_controller.go", "NO")
	if err != nil {
		return nil, err
	}

	//Usecase Interactor
	err = createFile(data,
		"./infrastructure/storage/template/file/src/usecase/interactor/interactor_template.txt",
		"./src/usecase/interactor/"+data.Name_file+"_interactor.go", "NO")
	if err != nil {
		return nil, err
	}

	//Usecase Presenter
	err = createFile(data,
		"./infrastructure/storage/template/file/src/usecase/presenter/presenter_template.txt",
		"./src/usecase/presenter/"+data.Name_file+"_presenter.go", "NO")
	if err != nil {
		return nil, err
	}

	//Usecase Repository
	err = createFile(data,
		"./infrastructure/storage/template/file/src/usecase/repository/repository_template.txt",
		"./src/usecase/repository/"+data.Name_file+"_repository.go", "NO")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func createFile(contentNewFile *models.CreateEndpoint, templateTxt string, newFile string, createFolder string) error {

	if dir := filepath.Dir(newFile); dir != "" && createFolder == "SI" {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	t, err := template.ParseFiles(templateTxt)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, contentNewFile)
	if err != nil {
		return err
	}
	content := buf.String()

	file, errTx := os.OpenFile(newFile, os.O_CREATE|os.O_WRONLY, 0777)
	if errTx != nil {
		return errTx
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s", content))
	if err != nil {
		return err
	}

	return nil
}
