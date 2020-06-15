package datasource

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	vm "ireport/api/viewmodels"
	"ireport/ihelpers/iobject"
	"ireport/ihelpers/irequest"
	"net/http"
	"strconv"

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
// https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/
func (app *Definition) Connect(ctx iris.Context) interface{} {
	var params *vm.Datasource
	err := irequest.GetJSON(ctx, &params)
	if err != nil {
		return _iresponse.JSON(404, nil, err.Error())
	}

	//urlDumy := "https://jsonplaceholder.typicode.com/todos/1"
	//urlDumy := "http://dummy.restapiexample.com/api/v1/employees"
	url := params.Endpoint
	// jsonData := map[string]string{}
	// jsonValue, _ := json.Marshal(jsonData)
	jsonValue, _ := json.Marshal(params.Params)

	if params.AuthType == "apiKey" && params.AuthRequest.APISource != "header" {
		url = url + "?" + params.AuthRequest.APIKey + "=" + params.AuthRequest.APIValue
	}
	request, _ := http.NewRequest(params.Method, url, bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Encoding", "br")
	request.Header.Set("Accept", "application/json")

	switch params.ContentType {
	case "json":
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		request.Header.Set("Content-Length", strconv.Itoa(len(jsonValue)))
		break
	case "formData":
		//multipart/form-data; boundary=--------------------------000687684202763635373446
		request.Header.Set("Content-Type", "multipart/form-data")
		request.Header.Set("Content-Length", strconv.Itoa(len(jsonValue)))
		break
	case "xwwwFormUrlencoded":
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		request.Header.Set("Content-Length", strconv.Itoa(len(jsonValue)))
		break
	case "raw":
		request.Header.Set("Content-Type", "text/plain")
		request.Header.Set("Content-Length", strconv.Itoa(len(jsonValue)))
		break
	default:
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
	}
	// Establece el tipo de autenticación
	switch params.AuthType {
	case "apiKey":
		if params.AuthRequest.APISource == "header" {
			request.Header.Set(params.AuthRequest.APIKey, params.AuthRequest.APIValue)
		}
		break
	case "bearerToken":
		request.Header.Set("Authorization", "Bearer "+params.AuthRequest.Bearer)
		break
	case "basicAuth":
		userPass := params.AuthRequest.Username + "." + params.AuthRequest.Password
		userPass = base64.StdEncoding.EncodeToString([]byte(userPass))
		request.Header.Set("Authorization", "Basic  "+userPass)
		break

		//default:
	}

	//realiza la petición
	client := &http.Client{}
	responseHTTP, err := client.Do(request)
	if err != nil {
		return _iresponse.JSON(404, nil, err.Error())
	}

	//lee el response
	data, err := ioutil.ReadAll(responseHTTP.Body)
	if err != nil {
		return _iresponse.JSON(404, nil, err.Error())
	}

	//si la respuesta fue un 200
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
