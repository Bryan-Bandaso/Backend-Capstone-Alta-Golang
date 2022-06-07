package upload_csv

import (
	upload_csv2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/controller/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	admin2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/repository/admin"
	admin3 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/service/admin"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)
	repo := admin2.NewUploadCsvRepository(db)
	svc := admin3.NewUploadCsvService(repo, conf)
	controller := upload_csv2.UploadCsvController{
		Svc: svc,
	}

	echo.POST("/admin/upload_csv", controller.UploadCsvController)
}
