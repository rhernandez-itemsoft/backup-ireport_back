package entity

//Datasource Estructura para almacenar la información del datasource
// esta información permite saber a donde va a conectar el reporte para obtener los datos
type Datasource struct {
	ID          uint   `gorm:"primary_key;column:id;not null;"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Method      string `gorm:"column:method"`
	Endpoint    string `gorm:"column:endpoint"`
	AuthType    string `gorm:"column:authType"`
	//DataType    string `gorm:"column:dataType"`
	ContentType string `gorm:"column:contentType"`
	Accept      string `gorm:"column:accept"`
}
