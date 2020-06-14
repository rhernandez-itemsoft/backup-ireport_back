package entity

//DatasourceParam Almacena la estructura de cada parametro
type DatasourceParam struct {
	ID           uint   `gorm:"primary_key;column:id;not null;"`
	DatasourceID uint   `gorm:"column:datasource_id;not null;"`
	Key          string `gorm:"column:key"`
	Type         string `gorm:"column:type"`
}
