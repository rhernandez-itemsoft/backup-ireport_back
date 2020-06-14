package config

import (
	// i18n "github.com/iris-contrib/middleware/go-i18n"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

const (
	//TimeFormat formato  de tiempo por default
	TimeFormat = "2006-01-02 15:04:05.9999999"
)

//AppCtrl definición generalizada de los controladores
type AppCtrl struct {
	Db         *gorm.DB
	TimeFormat string
}

type _configDataBase struct {
	Server   string
	Port     string
	User     string
	Password string
	DataBase string
}

//DBConf contiene la configuración a la base de datos
var DBConf = _configDataBase{
	Server:   "localhost",
	Port:     "3306",
	User:     "root",
	Password: "root",
	DataBase: "ireport", //"isystem",
}

//LanguageConf configuración para la internacionalización
// var LanguageConf = i18n.New(
// 	i18n.Config{
// 		Default:      "es-MX",
// 		URLParameter: "language",
// 		Languages: map[string]string{
// 			"en-US": "./resources/languages/en-US.ini",
// 			"es-MX": "./resources/languages/es-MX.ini",
// 		},
// 	},
// )

// //IrisConfig Establece la configuración de IRIS framework
var IrisConfig = iris.Configuration{
	DisableStartupLog:                 false,
	DisableInterruptHandler:           false,
	DisablePathCorrection:             false,
	EnablePathEscape:                  false,
	FireMethodNotAllowed:              true,
	DisableBodyConsumptionOnUnmarshal: false,
	DisableAutoFireStatusCode:         false,
	TimeFormat:                        "Mon, Jan 02 2006 15:04:05 GMT",
	Charset:                           "UTF-8",

	// PostMaxMemory is for post body max memory.
	//
	// The request body the size limit
	// can be set by the middleware `LimitRequestBodySize`
	// or `context#SetMaxRequestBodySize`.
	PostMaxMemory:               32 << 20, // 32MB
	//TranslateFunctionContextKey: "iris.translate",
	//TranslateLanguageContextKey: "Accept-Language",
	ViewLayoutContextKey:        "iris.viewLayout",
	ViewDataContextKey:          "iris.viewData",
	RemoteAddrHeaders: map[string]bool{
		"X-Real-Ip":        true,
		"X-Forwarded-For":  true,
		"CF-Connecting-IP": true,
	}, // make(map[string]bool) ,
	EnableOptimizations: false,
	Other:               make(map[string]interface{}),
}
