package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"go-webcrawl-api/src/api/constants"
	"go-webcrawl-api/src/api/models"
	"os"
)

func New2CSV(news models.NewsCreatable) error {
	if _, err := os.Stat(constants.CSV_FILE_PATH); errors.Is(err, os.ErrNotExist) {
		if file, err := os.Create(constants.CSV_FILE_PATH); err != nil {
			fmt.Println("Error while create CSV file: ", err)
			return err
		} else {
			defer file.Close()
			writer := csv.NewWriter(file)
			writer.Write([]string{*news.Title, *news.Content})
			writer.Flush()
		}
	} else if file, err := os.OpenFile(constants.CSV_FILE_PATH, os.O_RDWR|os.O_APPEND, 0644); err != nil {
		fmt.Println("Error while open CSV file: ", err)
		return err
	} else {
		defer file.Close()
		writer := csv.NewWriter(file)
		writer.Write([]string{*news.Title, *news.Content})
		writer.Flush()
	}
	return nil
}
