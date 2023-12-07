package repositories

import (
	"context"
	"fmt"
	"go-webcrawl-api/src/api/models"
	"go-webcrawl-api/src/configs"
)

func SaveNew(ctx context.Context, news models.NewsCreatable) error {
	database, _ := configs.GetGormInstance()
	if err := database.Table(models.NewsCreatable{}.GetTableName()).Create(&news).Error; err != nil {
		fmt.Println("Error while save news to database: ", err)
		return err
	}
	return nil
}
