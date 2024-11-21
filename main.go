package main

import (
	"bufio"
	"fmt"
	"os"
	"scraper/colly"
	"scraper/nlp"
)

func main() {
	fmt.Println("PARALLEL SCRAPER - SCHNEIDER ELECTRIC 2024")
	scanner := bufio.NewScanner(os.Stdin)
	running := true
	var theme string

	fmt.Println()
	fmt.Println("Unesite temu za scraping:")
	scanner.Scan()
	theme = scanner.Text()
	fmt.Println("Odabrali ste \"", theme, "\"")
	for running {
		fmt.Println()
		fmt.Println("1.Pokreni scraper \n2.Pokreni analizu komentara \n3.Izlaz")
		scanner.Scan()
		text := scanner.Text()
		if text == "1" {
			colly.Scrap(theme, "./data/comments.txt")
		} else if text == "2" {
			comments, err := colly.ReadComments("./data/comments.txt")
			if err != nil {
				fmt.Println("Greska pri citanju!")
			}
			comments = nlp.CleanData(comments)
			counts, percentages := nlp.AnalyzeComments(comments)
			fmt.Println("Rezultati sentiment analize:")
			fmt.Printf("Procenat pozitivnih komentara: %d comments (%.2f%%)\n", counts["Positive"], percentages["Positive"])
			fmt.Printf("Procenat negativnih komentara: %d comments (%.2f%%)\n", counts["Negative"], percentages["Negative"])
			nlp.CreatePieChart(percentages["Positive"], percentages["Negative"])
		} else if text == "3" {
			running = false
			fmt.Println("Izlaz...")
		} else {
			fmt.Println("Pogresan unos, pokusajte ponovo!")
		}
	}

}
