package seed

import "spincity.com/spincity/backend/models"

var Records = []models.Vinylrecord{
	{
		Id:1,
		Title: "Break Up",
		Artist: "John Mayer",
		Image:"cloudinary/images/1",
	},
	{
		Id:2,
		Title: "Dancing in the rain",
		Artist: "Kathy Jones",
		Image: "cloudinary/images/2",
	},
}