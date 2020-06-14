package entity

//DatasourceAuth El request de autorización
type DatasourceAuth struct {
	ID       		uint 	`gorm:"primary_key;column:id;not null;"`
	DatasourceID	uint 	`gorm:"column:datasource_id;not null;"`
	Bearer  		string  `gorm:"column:bearer"`
	Username  		string  `gorm:"column:username"`
	Password  		string  `gorm:"column:password"`
	APISource  		string  `gorm:"column:apiSource"`
	APIKey  		string  `gorm:"column:apiKey"`
	APIValue  		string  `gorm:"column:apiValue"`
}
