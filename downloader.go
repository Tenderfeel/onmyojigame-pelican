package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func Downloader() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(2),
		colly.UserAgent(userAgentPC),
	)

	c.SetRequestTimeout(120 * time.Second)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c2 := c.Clone()

	c.OnHTML("div.item", func(e *colly.HTMLElement) {

		e.ForEach(".mask > a", func(index int, el *colly.HTMLElement) {
			// https://yys.res.netease.com/pc/zt/20170731172708/data/picture/20230118/1/1366x768.jpg
			href := el.Attr("href")

			dirName, imagePath, err := parsedImageUrl(href, "")

			checkDir(dirName)

			if err == nil && !exists(imagePath) {
				c2.Visit(e.Request.AbsoluteURL(href))
			}

		})

	})

	c2.OnResponse(func(r *colly.Response) {
		if strings.Contains(r.Headers.Get("Content-Type"), "image") {

			href := r.Request.URL.RequestURI()

			// スラッシュで分割
			arr := strings.Split(href, "/")

			fmt.Printf("🖼 Image URL %s （len: %d） \n", href, len(arr))

			dirName, imagePath, err := parsedImageUrl(href, r.Headers.Get("Content-Type"))

			if err == nil {
				checkDir(dirName)

				err := r.Save(imagePath)

				if err != nil {
					fmt.Println("🔥 Failed image save:", err)
				} else {
					fmt.Println("📥 Saved: " + imagePath) // 年月
				}
			} else {
				fmt.Println("🚀 Skipped: " + imagePath)
			}

			return
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("C1 Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c2.OnError(func(r *colly.Response, err error) {
		fmt.Println("C2 Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)

		file, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err) // ログファイルが開けなかったときエラー出力
		}
		defer file.Close()

		url := r.Request.AbsoluteURL(r.Request.URL.RequestURI()) + "\n"

		file.Write(([]byte)(url))
	})

	c.Visit(pageUrl)
	c.Wait()
	c2.Wait()
}
