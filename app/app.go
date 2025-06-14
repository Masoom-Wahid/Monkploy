package app

import (
	"platform/app/v1/controllers"
	"platform/app/v1/repositories"
	"platform/config"
	"platform/database"
	"platform/database/migrations"
	"platform/database/seeder"
	"platform/middlewares"
	"platform/pkg/env"
	"platform/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppEntity interface {
	Run()
}

type App struct {
	RepoSupplier repositories.RepoSupplier
	Controllers  controllers.ControllerSupplier
	Middlewares  middlewares.MiddlewareSupplier
	Database     *gorm.DB
	// routes 	routes.Routes
	Config config.Config
}

func NewApp() AppEntity {
	env.SetupEnvFile()

	database := database.Connect()

	config := config.NewConfig()

	repoSupplier := repositories.NewRepoSupplier(config)
	controllers := controllers.NewControllerSupplier(&repoSupplier)
	middlewares := middlewares.NewMiddlewareSupplier()

	app := new(App)
	app.RepoSupplier = repoSupplier
	app.Database = database
	app.Controllers = controllers
	app.Middlewares = middlewares
	app.Config = config

	return app
}

func (app *App) Run() {
	r := gin.Default()

	// r.Static("/assets/", "./views/dist/assets")
	// r.LoadHTMLGlob("./views/dist/*.html")

	r.Use(middlewares.CORSMiddleware())
	if app.Config.AppConfig().Env == "development" {
		migrations.Migrate(app.Database)

		seederSupplier := seeder.NewSeederSupplier()
		seederSupplier.Run()
	}

	// r.NoRoute(func(c *gin.Context) {
	// 	c.HTML(200, "index.html", nil)
	// })

	routes.SetUpRoutes(&r.RouterGroup, app.Controllers, app.Middlewares)

	r.Run()
}
