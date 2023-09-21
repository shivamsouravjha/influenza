package cloudinary

import (
	passwordConfig "github.com/shivamsouravjha/influenza/config"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/config"
)

var (
	Cloudinary *cloudinary.Cloudinary
)

func Init() error {
	Cloudinary = &cloudinary.Cloudinary{}
	Cloudinary.Config = config.Configuration{
		Cloud: config.Cloud{
			CloudName: passwordConfig.Get().CloudName,
			APIKey:    passwordConfig.Get().CloudPublic,
			APISecret: passwordConfig.Get().CloudSecret,
		},
	}
	Cloudinary, _ = cloudinary.NewFromParams(Cloudinary.Config.Cloud.CloudName, Cloudinary.Config.Cloud.APIKey, Cloudinary.Config.Cloud.APISecret)

	return nil
}

func Client() *cloudinary.Cloudinary {
	return Cloudinary
}
