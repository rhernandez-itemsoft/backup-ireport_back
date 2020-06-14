package vmdl

//DatasourceAuth El request de autorización
type DatasourceAuth struct {
	ID           uint `json:"id"`
	DatasourceID uint `json:"datasourceId"`

	//bearer
	Bearer string `json:"bearer"`

	//basicAuth
	Username string `json:"username"`
	Password string `json:"password"`

	// apiKey
	APISource string `json:"apiSource"` //indica que debemos agregarlo al header o a queryParams
	APIKey    string `json:"apiKey"`
	APIValue  string `json:"apiValue"`
}
