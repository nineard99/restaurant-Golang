package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cld *cloudinary.Cloudinary

func InitCloudinary() error {
	var err error
	cld, err = cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	return err
}

func UploadImageToCloudinary(file multipart.File, folder string, publicID string) (string, error) {
	ctx := context.Background()

	uploadParams := uploader.UploadParams{
		PublicID: fmt.Sprintf("%s/%s", folder, publicID),
	}

	uploadResult, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}
	return uploadResult.SecureURL, nil
}
