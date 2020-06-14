package datasource

import (
	entity "ireport/api/entities"
	vmdl "ireport/api/viewmodels"
)

func (app *Definition) saveManager(_input *vmdl.Datasource) error {
	var err error

	dataSource := entity.Datasource{
		ID:          0,
		Name:        _input.Name,
		Description: _input.Description,
		Method:      _input.Method,
		Endpoint:    _input.Endpoint,
		Accept:      _input.Accept,
		ContentType: _input.ContentType,
		AuthType:    _input.AuthType,
	}

	datasourceAuth := entity.DatasourceAuth{
		ID:           0,
		DatasourceID: _input.AuthRequest.DatasourceID,
		Bearer:       _input.AuthRequest.Bearer,
		Username:     _input.AuthRequest.Username,
		Password:     _input.AuthRequest.Password,
		APISource:    _input.AuthRequest.APISource,
		APIKey:       _input.AuthRequest.APIKey,
		APIValue:     _input.AuthRequest.APIValue,
	}

	_input.ID, err = _datasourcerepository.Save(&dataSource)
	if err != nil {
		return err
	}

	_input.AuthRequest.ID, err = _datasourceauthrespository.Save(&datasourceAuth)
	if err != nil {
		return err
	}

	for k, row := range _input.Params {
		var param = entity.DatasourceParam{
			ID:           0,
			DatasourceID: row.DatasourceID,
			Key:          row.Key,
			Type:         row.Type,
		}

		_input.Params[k].ID, err = _datasourceparam.Save(&param)
		if err != nil {
			return err
		}
	}

	return err
}
