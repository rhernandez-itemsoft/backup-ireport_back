package datasource

import (
	"encoding/json"
	entity "ireport/api/entities"
	vmdl "ireport/api/viewmodels"
)

func (app *Definition) saveManager(_input *vmdl.Datasource) error {
	var err error
	_response, _ := json.Marshal(_input.Response)

	dataSource := entity.Datasource{
		ID:          0,
		Name:        _input.Name,
		Description: _input.Description,
		Method:      _input.Method,
		Endpoint:    _input.Endpoint,
		Accept:      _input.Accept,
		ContentType: _input.ContentType,
		AuthType:    _input.AuthType,
		Response:    string(_response),
	}

	_input.ID, err = _datasourcerepository.Save(&dataSource)
	if err != nil {
		return err
	}

	_input.AuthParams.DatasourceID = _input.ID
	datasourceAuth := entity.DatasourceAuth{
		ID:           0,
		DatasourceID: _input.AuthParams.DatasourceID,
		Bearer:       _input.AuthParams.Bearer,
		Username:     _input.AuthParams.Username,
		Password:     _input.AuthParams.Password,
		APISource:    _input.AuthParams.APISource,
		APIKey:       _input.AuthParams.APIKey,
		APIValue:     _input.AuthParams.APIValue,
	}
	_input.AuthParams.ID, err = _datasourceauthrespository.Save(&datasourceAuth)
	if err != nil {
		return err
	}

	for k, row := range _input.RequestParams {
		_input.RequestParams[k].DatasourceID = _input.ID
		var param = entity.DatasourceParam{
			ID:           0,
			DatasourceID: row.DatasourceID,
			Key:          row.Key,
			Type:         row.Type,
		}

		_input.RequestParams[k].ID, err = _datasourceparam.Save(&param)
		if err != nil {
			return err
		}
	}

	return err
}
