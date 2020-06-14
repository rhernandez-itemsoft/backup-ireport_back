package main

import (
	goCtx "context"
	"fmt"
	conf "ireport/config"
	"ireport/ihelpers/errors"
	"ireport/routes"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/iris-contrib/middleware/cors"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//realiza la migración de la BD
	migrateDataBase()

	//Inicializa la conección con la Base de datos
	db, err := initDataBase()
	defer db.Close()
	if err != nil {
		errors.Catch(err, true)
	}

	//Programamos que si ocurre alguna falla, entonces la Base de datos se cerrará
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	//Inicializa la configuración de iris para levantar el api
	app := initIrisApp()

	//inicializa el ruteo a los servicios
	routes.InitRoutes(app, db)

	//ejecuta el server
	app.Run(iris.Addr("localhost:8080"), iris.WithConfiguration(conf.IrisConfig))
}

func migrateDataBase() error {

	db, err := gorm.Open("mysql", conf.DBConf.User+":"+conf.DBConf.Password+"@("+conf.DBConf.Server+":"+conf.DBConf.Port+")/"+conf.DBConf.DataBase+"?multiStatements=true")
	defer db.Close()
	if err != nil {
		return err
	}
	driver, _ := mysql.WithInstance(db.DB(), &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///Personal\\ItemsoftMX\\Projects\\Golang\\src\\ireport-back\\DB",

		"ireport",
		driver,
	)
	if err != nil {
		return err
	}
	m.Steps(2)
	return nil
}

//Conecta a la base de datos, y obtenemos la version de la BD para mostrarlo en consola
func initDataBase() (*gorm.DB, error) {
	// Create the database handle, confirm driver is present
	//db, err := gorm.Open("mysql", "root:root@(localhost:3306)/isystem")
	db, err := gorm.Open("mysql", conf.DBConf.User+":"+conf.DBConf.Password+"@("+conf.DBConf.Server+":"+conf.DBConf.Port+")/"+conf.DBConf.DataBase+"?multiStatements=true")
	if err != nil {
		return nil, err
	}

	var version string
	rows, err := db.Raw("SELECT VERSION()").Rows()
	defer rows.Close()
	if err != nil {
		fmt.Println("No se pudo conectar")
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&version)
		fmt.Println("Conectado a: " + version)
	}

	return db, nil
}

//Establecemos toda la configuración necesario para inicializar IRIS
func initIrisApp() *iris.Application {
	app := iris.New()

	// Prevenimos que se cierre el server, cuando un error fatal ocurre
	iris.RegisterOnInterrupt(func() {
		timeout := 10000 * time.Second
		ctx, cancel := goCtx.WithTimeout(goCtx.Background(), timeout)
		defer cancel()

		// close all hosts
		app.Shutdown(ctx)
	})

	//configuramos el CORS
	/*
		opts := cors.Options{
			AllowCredentials: false,
			AllowedOrigins: []string{"http://localhost:4200","http://localhost:4200/*","localhost:4200", "*"},
			AllowedHeaders: []string{
				"Accept",
				"Accept-Encoding",
				"Access-Control-Allow-Origin",
				"Authorization",
				"Content-Type",
				"Content-Length",
				"Origin",
				"X-Auth-Token",
				"X-Requested-With",
				"X-Force-Content-Type",
			},
			AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "OPTION", "HEAD", "XHR"},
			ExposedHeaders:     []string{"X-Header"},
			MaxAge:             int((24 * time.Hour).Seconds()),
			OptionsPassthrough: true,
			// Debug:          true,
		}

		app.Use(cors.New(opts))*/
	app.Use(cors.AllowAll())
	app.AllowMethods(iris.MethodOptions)

	//i 18n Lenguage handlers
	//Establece los handlers para el manejo de errores y para el manejo de traducciones
	// app.Use(LanguageHandler)

	app.I18n.Load("./resources/languages/*/*", "en-US", "es-MX")
	app.I18n.SetDefault("en-US")

	return app
}
