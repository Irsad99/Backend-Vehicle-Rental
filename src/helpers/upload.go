package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadImage(image string) (string, error) {

	cld, err := cloudinary.NewFromParams("dsifbeghc", "558379834285585", "uHMSv_bG8DFrWIMwcjtMg5gW-gY")

	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}

	var ctx = context.Background()

	uploadResult, err := cld.Upload.Upload(
		ctx,
		"https://res.cloudinary.com/dsifbeghc/image/upload/go_image/" + image,
		uploader.UploadParams{PublicID: "vehicle"})

	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}

	fmt.Println(uploadResult.SecureURL)

	return uploadResult.SecureURL, nil
}
