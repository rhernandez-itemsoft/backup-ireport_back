package routes

import (
	"ireport/api/services/datasource"

	"github.com/kataras/iris/v12/mvc"

	conf "ireport/config"

	"ireport/ihelpers/iresponse"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"

	"github.com/iris-contrib/swagger/v12"              // swagger middleware for Iris
	"github.com/iris-contrib/swagger/v12/swaggerFiles" // swagger embed files
)

//LanguageHandler handler para manejar internacionalización
//var LanguageHandler context.Handler = conf.LanguageConf

//incializa el ruteo para mostrar un mensaje de error en las rutas que no son válidas
func initRouteErrors(app *iris.Application) {
	app.OnAnyErrorCode(func(ctx iris.Context) {
		_response := iresponse.New(ctx)
		_response.JSON(iris.StatusUnauthorized, nil, "404 Not found")
		return
	})
}

//inicializamos el ruteo para mostrar la documentacion
func initSwaggerDocumentation(app *iris.Application) {
	config := &swagger.Config{
		URL: "http://localhost/swagger/doc.json", //The url pointing to API definition
	}

	// use swagger middleware to
	app.Get("/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))

	app.Handle("GET", "/", func(ctx iris.Context) {
		//ctx.HTML("<h1>Welcome</h1>")
		//ctx.View("index.html")
		ctx.Redirect("/index.html")
	})

	//http.Handle("/", http.StripPrefix(strings.TrimRight(path, "/"), http.FileServer(http.Dir(directory))))

}

//InitRoutes esta funciona es la que creará el ruteo
func InitRoutes(app *iris.Application, db *gorm.DB) {
	initSwaggerDocumentation(app)

	initRouteErrors(app)

	mvc.New(app.Party("/report/datasource")).Handle(&datasource.Definition{
		Db:         db,
		TimeFormat: conf.TimeFormat,
	})

}
