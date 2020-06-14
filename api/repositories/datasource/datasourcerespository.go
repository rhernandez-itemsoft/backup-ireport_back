package datasourcerepository

import (
	"errors"
	entity "ireport/api/entities"

	"github.com/jinzhu/gorm"
)

//Definition  security controller
type Definition struct {
	Db *gorm.DB
}

func (_app *Definition) Save(_input *entity.Datasource) (uint, error) {
	if !_app.Db.NewRecord(_input) {
		return 0, errors.New("El registro que intenta crear, ya existe")
	}

	err := _app.Db.Create(_input).Error
	return _input.ID, err
}
