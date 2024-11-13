package colly

import (
	"fmt"
	//"os"
	//"strings"
	"time"
	//"github.com/gocolly/colly/v2"
)

type item struct {
	StoryURL  string
	Source    string
	Comments  string
	CrawledAt time.Time
	Title     string
	CommentTexts []string
}

func Scrap(topic string){
	fmt.Printf("STARTED SCRAPING %s TOPIC...", topic)
}