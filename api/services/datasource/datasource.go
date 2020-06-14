package datasource

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	vm "ireport/api/viewmodels"
	"ireport/ihelpers/iobject"
	"ireport/ihelpers/irequest"
	"net/http"

	"github.com/kataras/iris/v12"
)

//Save guarda los datos del diseñador de reportes
func (app *Definition) Save(ctx iris.Context) interface{} {
	var params *vm.Datasource
	err := irequest.GetJSON(ctx, &params)
	if err != nil {
		return _iresponse.JSON(404, nil, err.Error())
	}

	err = app.saveManager(params)
	if err != nil {
		return _iresponse.JSON(404, nil, err.Error())
	}

	return _iresponse.JSON(200, params, "")

}

//Connect connecta al endpoint y retorna su response
func (app *Definition) xConnect(ctx iris.Context) interface{} {
	var params *vm.Datasource
	err := irequest.GetJSON(ctx, &params)
	if err != nil {
		return _iresponse.JSON(404, nil, err.Error())
	}

	return nil
}

// Connect Obtiene el tipo de dato de la estructura de datos que se intenta procesar
// @param URL el endpoint a donde vamos a conectar
// @param Method puede ser GET, POST
func (app *Definition) Connect() interface{} {

	//https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/

	//urlDumy := "https://jsonplaceholder.typicode.com/todos/1"
	urlDumy := "http://dummy.restapiexample.com/api/v1/employees"
	url := urlDumy
	jsonData := map[string]string{}
	jsonValue, _ := json.Marshal(jsonData)
	request, _ := http.NewRequest("GET", url, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Content-Encoding", "br")
	request.Header.Set("Accept", "application/json")

	client := &http.Client{}
	responseHTTP, err := client.Do(request)
	if err != nil {
		// pc, fn, line, _ := runtime.Caller(1)
		// errorX := fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err.Error())
		return _iresponse.JSON(404, nil, err.Error())
	}
	data, err := ioutil.ReadAll(responseHTTP.Body)
	if err != nil {
		return _iresponse.JSON(404, nil, err.Error())
	}

	if responseHTTP.StatusCode == 200 {
		//obtiene el json del response
		jsonMap, err := iobject.ByteToMapStringInterface(data)
		if err != nil {
			return _iresponse.JSON(404, nil, err.Error())
		}

		//obtiene el tipo de dato del json
		var dataType map[string]interface{}
		dataType = iobject.GetDataType(jsonMap)
		return _iresponse.JSON(responseHTTP.StatusCode, dataType)
	}

	return _iresponse.JSON(responseHTTP.StatusCode, string(data))
}
