package vmdl

//Datasource Estructura para almacenar la información del datasource
// esta información permite saber a donde va a conectar el reporte para obtener los datos
type Datasource struct {
	ID uint `json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	//metodo de la petición (GET, POST)
	Method string `json:"method"`

	//url del endpoint
	Endpoint string `json:"endpoint"`

	//tipo de contendo que acepta: application/json, fomultipart/form-data, x-www-form-urlencoded, text/plain
	Accept string `json:"accept"`

	//tipo de autorizacion
	AuthType string `json:"authType"`

	//parametros de autorizacion
	AuthRequest DatasourceAuth `json:"authRequest"`

	//tipo de contendo que envia: application/json, fomultipart/form-data, x-www-form-urlencoded, text/plain
	ContentType string `json:"contentType"`

	Params []DatasourceParam `json:"params"`
}
