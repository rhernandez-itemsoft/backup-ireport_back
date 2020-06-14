package vmdl

//DatasourceParam Almacena la estructura de cada parametro
type DatasourceParam struct {
	ID           uint   `json:"id"`
	DatasourceID uint   `json:"datasourceId"`
	Key          string `json:"key"`
	Type         string `json:"type"`
}
