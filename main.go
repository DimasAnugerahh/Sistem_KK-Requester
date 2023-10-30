package main

import (
	"kk-requester/config"
	"kk-requester/controller"
	"kk-requester/repository"
	"kk-requester/routes"
	"kk-requester/service"
	"log"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}
	app := echo.New()
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("gagal koneksi database: ", err)
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

	AktaKematianRepository := repository.NewAktaKematianRepository(db)
	AktaKematianService := service.NewAktaKematianService(AktaKematianRepository)
	AktaKematianController := controller.NewAktaKematianController(AktaKematianService)

	KKRequestRepository := repository.NewRequestKKRepository(db)
	KKRequestService := service.NewRequestKKService(KKRequestRepository)
	KKRequestController := controller.NewRequestKKController(KKRequestService)

	KKRepository := repository.NewKkRepository(db)
	KKService := service.NewKkService(KKRepository)
	KKController := controller.NewKKController(KKService)

	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	routes.NewAccountRoutes(app, AccountController)
	routes.NewKTPRoutes(app, KTPController)
	routes.NewAktaKelahiranRoutes(app, AktaKelahiranController)
	routes.NewSuratPindahRoutes(app, SuratPindahController)
	routes.NewAktaKematianRoutes(app, AktaKematianController)
	routes.NewSuratNikahRoutes(app, SuratNikahController)
	routes.NewKKRequestRoutes(app, KKRequestController)
	routes.NewkkRoutes(app, KKController)

	app.Start(":8080")
}
