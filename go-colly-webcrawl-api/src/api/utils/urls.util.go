package utils

import (
	"encoding/xml"
	"fmt"
	"go-webcrawl-api/src/api/models"
	"io"
	"os"
)

func UrlsFromSiteMap(filePath string) (*models.UrlSet, error) {
	if xmlFile, err := os.Open(filePath); err != nil {
		fmt.Println("Error while open sitemap.xml file: ", err)
		return nil, err
	} else {
		defer xmlFile.Close()
		var urlSet models.UrlSet
		fmt.Println("Successfully open sitemap.xml")
		byteValue, _ := io.ReadAll(xmlFile)
		xml.Unmarshal(byteValue, &urlSet)
		return &urlSet, nil
	}
}
