package datasource

import (
	"fmt"
	datasourcerepository "ireport/api/repositories/datasource"
	datasourceauthrespository "ireport/api/repositories/datasourceauth"
	datasourceparam "ireport/api/repositories/datasourceparam"
	"ireport/ihelpers/iresponse"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

var _iresponse *iresponse.Definition
var _datasourcerepository *datasourcerepository.Definition
var _datasourceauthrespository *datasourceauthrespository.Definition
var _datasourceparam *datasourceparam.Definition

//Definition controlador "reportdesignctrl"
type Definition struct {
	Db         *gorm.DB
	TimeFormat string
}

//BeforeActivation antes de activar el controller
func (app *Definition) BeforeActivation(b mvc.BeforeActivation) {
	app.initRespositories()

	b.Handle("POST", "/save", "Save", serviceHandler)
	b.Handle("POST", "/connect", "Connect", serviceHandler)
}

func (app *Definition) initRespositories() {
	_datasourcerepository = &datasourcerepository.Definition{
		Db: app.Db,
	}

	_datasourceauthrespository = &datasourceauthrespository.Definition{
		Db: app.Db,
	}

	_datasourceparam = &datasourceparam.Definition{
		Db: app.Db,
	}
}

//esto es para poder pasar el iresponse inicializado en el context
func serviceHandler(ctx iris.Context) {
	_iresponse = iresponse.New(ctx)

	if ctx == nil {
		fmt.Println("ctx es null")
	}

	ctx.Next()
	return
}
