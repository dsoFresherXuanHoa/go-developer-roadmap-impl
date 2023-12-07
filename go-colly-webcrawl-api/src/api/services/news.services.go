package services

import (
	"context"
	"fmt"
	"go-webcrawl-api/src/api/models"
	"go-webcrawl-api/src/api/repositories"
	"go-webcrawl-api/src/api/utils"
	"reflect"
	"regexp"
	"sync"

	"github.com/gocolly/colly/v2"
	"github.com/ttacon/chalk"
)

func SyncLink2News(urlSet models.UrlSet) models.NewsCreatableList {
	var newsList models.NewsCreatableList
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
				newsList = append(newsList, news)
			}
			fmt.Println("Title: ", title)
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println(chalk.Green, "Crawling: ", r.URL)
		})
		c.Visit(urlSet.Urls[i].Loc)
	}
	return newsList
}

func SyncSaveNews(ctx context.Context, urlSet models.UrlSet) error {
	newsList := SyncLink2News(urlSet)
	for _, news := range newsList {
		if err := repositories.SaveNew(context.Background(), news); err != nil {
			fmt.Println(chalk.Red, "Error while save news to database: ", err)
			return err
		} else if err := utils.New2CSV(news); err != nil {
			fmt.Println(chalk.Red, "Error while save news to csv file: ", err)
			return err
		}
	}
	return nil
}

// Push Url to Channel Queue
func Urls2Queue(urlSet models.UrlSet) chan models.Url {
	urlsChan := make(chan models.Url, 512)
	defer close(urlsChan)
	for _, url := range urlSet.Urls {
		urlsChan <- url
		fmt.Println(chalk.Yellow.Color("Url to Queue: "), url)
	}
	return urlsChan
}

// Crawling New
func News2Queue(urlChan chan models.Url, numberOfWorker int) chan models.NewsCreatable {
	newChan := make(chan models.NewsCreatable, len(urlChan))
	defer close(newChan)

	wg := new(sync.WaitGroup)
	for i := 0; i < numberOfWorker; i++ {
		wg.Add(1)
		go func(worker int) {
			for url := range urlChan {
				var news models.NewsCreatable
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
						news = models.NewsCreatable{Title: &title, Content: &content}
						newChan <- news
					}
				})
				c.OnRequest(func(r *colly.Request) {
					fmt.Println(chalk.Green.Color("[Worker"), worker, chalk.Green.Color("] Crawling:"), r.URL)
				})
				c.Visit(url.Loc)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return newChan
}

// Save to Database and CSV file
func AsyncSaveNews(urlSet models.UrlSet, numberOfWorker int, wg *sync.WaitGroup) {
	urlChan := Urls2Queue(urlSet)
	newChan := News2Queue(urlChan, numberOfWorker)
	for i := 0; i < numberOfWorker; i++ {
		wg.Add(1)
		go func(worker int) {
			for news := range newChan {
				if !reflect.DeepEqual(news, models.NewsCreatable{}) {
					if err := repositories.SaveNew(context.Background(), news); err != nil {
						fmt.Println(chalk.Red.Color("Error while save record to database: "), err)
					} else if err := utils.New2CSV(news); err != nil {
						fmt.Println(chalk.Red.Color("Error while save record to csv file: "), err)
					} else {
						fmt.Println(chalk.Red.Color("Save record successes by worker "), worker)
					}
				}
			}
			wg.Done()
		}(i)
	}
}
