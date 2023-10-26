package main

import (
	"kk-requester/config"
	"kk-requester/controller"
	"kk-requester/repository"
	"kk-requester/routes"
	"kk-requester/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	config.InitialMigration(db)

	AccountRepository := repository.NewAccountRepository(db)
	AccountService := service.NewAccountService(AccountRepository)
	AccountController := controller.NewAccountController(AccountService)

	KTPRepository := repository.NewKTPRepository(db)
	KTPService := service.NewKTPService(KTPRepository)
	KTPController := controller.NewKTPController(KTPService)

	AktaKelahiranRepository := repository.NewAktaKelahiranRepository(db)
	AktaKelahiranService := service.NewAktaKelahiranService(AktaKelahiranRepository)
	AktaKelahiranController := controller.NewAktaKelahiranController(AktaKelahiranService)

	SuratPindahRepository := repository.NewSuratPindahRepository(db)
	SuratPindahService := service.NewSuratPindahService(SuratPindahRepository)
	SuratPindahController := controller.NewSuratPindahController(SuratPindahService)

	SuratNikahRepository := repository.NewSuratNikahRepository(db)
	SuratNikahService := service.NewSuratNikahService(SuratNikahRepository)
	SuratNikahController := controller.NewSuratNikahController(SuratNikahService)

	SuratKematianRepository := repository.NewSuratKematianRepository(db)
	SuratKematianService := service.NewSuratKematianService(SuratKematianRepository)
	SuratKematianController := controller.NewSuratKematianController(SuratKematianService)

	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	routes.NewAccountRoutes(app, AccountController)
	routes.NewKTPRoutes(app, KTPController)
	routes.NewAktaKelahiranRoutes(app, AktaKelahiranController)
	routes.NewSuratPindahRoutes(app, SuratPindahController)
	routes.NewSuratKematianRoutes(app, SuratKematianController)
	routes.NewSuratNikahRoutes(app, SuratNikahController)

	app.Start(":8080")
}
