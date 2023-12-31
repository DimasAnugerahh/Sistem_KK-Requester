package helper

import (
	"context"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func CloudinaryUpdload(c echo.Context, fileheader string) string {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}
	fileHeader, _ := c.FormFile(fileheader)

	if fileHeader != nil {
		file, _ := fileHeader.Open()
		ctx := context.Background()
		urlCloudinary := os.Getenv("CLOUDINARY_URL")
		cldService, _ := cloudinary.NewFromURL(urlCloudinary)
		response, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "kk-requester"})
		return response.SecureURL
	}
	return ""

}
