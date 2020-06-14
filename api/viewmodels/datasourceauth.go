package vmdl

//DatasourceAuth El request de autorización
type DatasourceAuth struct {
	ID           uint   `json:"id"`
	DatasourceID uint   `json:"datasourceId"`
	Bearer       string `json:"bearer"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	APISource    string `json:"apiSource"`
	APIKey       string `json:"apiKey"`
	APIValue     string `json:"apiValue"`
}
