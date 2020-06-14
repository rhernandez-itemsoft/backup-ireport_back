package vmdl

//Datasource Estructura para almacenar la información del datasource
// esta información permite saber a donde va a conectar el reporte para obtener los datos
type Datasource struct {
	ID       		uint `json:"id"`
	
	Name  			string  `json:"name"`
	Description  	string  `json:"description"`
	Method  		string  `json:"method"`
	Endpoint  		string  `json:"endpoint"`
	AuthType  		string  `json:"authType"`
	DataType  		string  `json:"dataType"`
	ContentType  	string  `json:"contentType"`

	AuthRequest  	DatasourceAuth  `json:"authRequest"`
	Params  		[]DatasourceParam  `json:"params"`
}