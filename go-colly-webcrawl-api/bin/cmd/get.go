/*
Copyright © 2023 dsoFresherXuanHoa <dso.intern.xuanhoa@gmail.com>
*/
package cmd

import (
	"context"
	"encoding/xml"
	"fmt"
	"go-webcrawl-api/src/api/models"
	"go-webcrawl-api/src/api/repositories"
	"io"
	"os"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get content from URL",
	Long:  `Get content from default URL and save it into database!`,
	Run: func(cmd *cobra.Command, args []string) {
		var filePath = "./sitemap.xml"
		if xmlFile, err := os.Open(filePath); err != nil {
			fmt.Println("Error while open sitemap.xml file: ", err)
			return
		} else {
			defer xmlFile.Close()
			var urlSet models.UrlSet

			byteValue, _ := io.ReadAll(xmlFile)
			xml.Unmarshal(byteValue, &urlSet)

			// Crawling
			for i := 0; i < len(urlSet.Urls); i++ {
				c := colly.NewCollector()
				c.OnHTML("#dark_theme > section.section.page-detail.top-detail > div > div.sidebar-1", func(h *colly.HTMLElement) {
					title := h.ChildText("#dark_theme > section.section.page-detail.top-detail > div > div.sidebar-1 > h1")
					content := ""
					h.ForEach("#dark_theme > section.section.page-detail.top-detail > div > div.sidebar-1 > article > .Normal", func(_ int, h *colly.HTMLElement) {
						contentOrigin := regexp.MustCompile(`\n`)
						contentConverted := contentOrigin.ReplaceAllString(h.Text, "<br/>")
						content += "<p>" + contentConverted + "</p>"
					})
					if title != content {
						news := models.NewsCreatable{Title: &title, Content: &content}
						// Save
						if err := repositories.SaveNew(context.Background(), news); err != nil {
							fmt.Println("Error while save news to database: ", err)
							return
						}
						fmt.Println("Save a record into database: ", news.Title)
					}
				})
				c.OnRequest(func(r *colly.Request) {
					fmt.Println("Crawling: ", r.URL)
				})
				c.Visit(urlSet.Urls[i].Loc)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
