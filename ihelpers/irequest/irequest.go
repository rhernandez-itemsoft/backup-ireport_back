package irequest

import (
	"errors"

	"github.com/kataras/iris/v12"
)

//GetJSON retorna los parametros (request) en JSON
// return error
func GetJSON(ctx iris.Context, params interface{}) error {
	//var response iresponse.ResponseModel = iresponse.IResponse()
	err := ctx.ReadJSON(&params)
	if err != nil {
		return err
	}

	if params == nil {
		return errors.New("No se han recibido parametros")
	}

	return nil
}
