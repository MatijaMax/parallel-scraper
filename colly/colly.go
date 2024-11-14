package colly

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type item struct {
	StoryURL     string
	Source       string
	Comments     string
	CrawledAt    time.Time
	Title        string
	CommentTexts []string
}

func Scrap(topic string, path string) {
	fmt.Printf("STARTED SCRAPING %s TOPIC...", topic)
	stories := []item{}

	c := colly.NewCollector(
		colly.AllowedDomains("old.reddit.com"),
		colly.Async(true),
	)

	commentCollector := colly.NewCollector(
		colly.AllowedDomains("old.reddit.com"),
		colly.Async(true),
	)

	c.OnHTML(".top-matter", func(e *colly.HTMLElement) {
		temp := item{
			StoryURL:  e.ChildAttr("a[data-event-action=title]", "href"),
			Source:    "https://old.reddit.com/r/politics/",
			Title:     e.ChildText("a[data-event-action=title]"),
			Comments:  e.ChildAttr("a[data-event-action=comments]", "href"),
			CrawledAt: time.Now(),
		}

		if strings.Contains(strings.ToLower(temp.Title), strings.ToLower(topic)) {
			stories = append(stories, temp)
			commentCollector.Visit(e.Request.AbsoluteURL(temp.Comments))
		}
	})

	c.OnHTML("span.next-button", func(h *colly.HTMLElement) {
		nextPage := h.ChildAttr("a", "href")
		c.Visit(nextPage)
	})

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 2 * time.Second,
		//Parallelism: 4,
		//RandomDelay: 2 * time.Second,
	})

	commentCollector.OnHTML(".comment .md", func(e *colly.HTMLElement) {
		commentText := e.Text
		for i := range stories {
			if strings.Contains(e.Request.URL.String(), stories[i].Comments) {
				stories[i].CommentTexts = append(stories[i].CommentTexts, commentText)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	commentCollector.OnRequest(func(r *colly.Request) {
		//fmt.Println("Scraping comments from", r.URL.String())
	})

	c.Visit("https://old.reddit.com/r/politics")
	c.Wait()
	commentCollector.Wait()

	fmt.Println("Scraped Stories:")
	var toFile []string
	if /*len(stories) <= 3 &&*/ len(stories) > 0 {
		story := stories[0]
		fmt.Printf("Title: %s\nURL: %s\n", story.Title, story.StoryURL)

		for i, comment := range story.CommentTexts {
			fmt.Printf("Comment %d: %s\n", i+1, comment)
			c := fmt.Sprintf("COMMENT### %s\n", comment)
			toFile = append(toFile, c)
		}
		toFile = toFile[1:]
	}
	/*
		if len(stories) > 3 {
			for _, story := range stories[:4] {
				fmt.Printf("Title: %s\nURL: %s\n", story.Title, story.StoryURL)
				for i, comment := range story.CommentTexts {
					fmt.Printf("Comment %d: %s\n", i+1, comment)
					toFile = append(toFile, story.CommentTexts...)
				}
			}
		}
	*/
	if len(stories) > 1 {
		WriteComments(path, toFile)
	}
}
