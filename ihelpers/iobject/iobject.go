package iobject

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

//GetDataType Retorna el tipo de objeto de una interface del tipo ma[string]interface{}
//https://stackoverflow.com/questions/20170275/how-to-find-a-type-of-an-object-in-go
//mapDataType mapeo la estructura y regresa el tipo de datos de cada campo
// Realiza un mapeo recursivo de la estructura de datos
func GetDataType(dataInterface map[string]interface{}) map[string]interface{} {
	mapD := map[string]interface{}{}
	for field, data := range dataInterface { // .(map[string]interface{}) {
		//verificamos el tipo de dato para saber cual es la accion a tomar
		dataType := typeof(data)
		switch dataType {
		case "[]interface {}":
			{
				arrInterface := data.([]interface{})
				//puede ser un arreglo de objetos o un arreglo de un tipo simple
				if typeof(arrInterface[0]) == "map[string]interface {}" {
					mapD[field] = GetDataType(arrInterface[0].(map[string]interface{}))
				} else {
					mapD[field] = "[]" + typeof(arrInterface[0])
				}
			}
		case "map[string]interface {}":
			{
				mapD[field] = GetDataType(data.(map[string]interface{}))
			}
		default:
			{
				mapD[field] = dataType
			}
		}
	}
	return mapD
}

//typeof obtiene el tipo de datos
func typeof(data interface{}) string {
	//return fmt.Sprintf("%T", data)

	if data != nil {
		return reflect.TypeOf(data).String()
	}

	return "nil"
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// ByteToMapStringInterface Convierte un arreglo de bytes a un map[string]interface{}
func ByteToMapStringInterface(data []byte) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	var jsonArray []map[string]interface{}

	isJSONArray := false
	if err := json.Unmarshal(data, &jsonMap); err != nil {
		isJSONArray = strings.Index(err.Error(), "unmarshal array") > -1
	}

	if isJSONArray {
		if err := json.Unmarshal(data, &jsonArray); err != nil {
			return nil, err
		}

		jsonMap = jsonArray[0]
	}

	return jsonMap, nil
}
