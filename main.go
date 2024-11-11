package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	fmt.Println("PARALLEL SCRAPER - SCHNEIDER ELECTRIC 2024")
	scanner := bufio.NewScanner(os.Stdin)
	running := true
	var theme string

	fmt.Println()
	fmt.Println("Unesite temu za scraping:")
	scanner.Scan()
	theme = scanner.Text()
	fmt.Println("Odabrali ste \"",theme,"\"")
	for running {
		fmt.Println()
		fmt.Println("1.Pokreni scraper \n2.Pokreni analizu komentara \n3.Izlaz \n")
		scanner.Scan()
		text := scanner.Text()
		if text == "1"{
			fmt.Println("//TODO")
		}else if text == "2"{
			fmt.Println("//TODO")
		}else if text == "3"{
			running = false
			fmt.Println("Izlaz...")
		}else {
			fmt.Println("Pogresan unos, pokusajte ponovo!")			
		}
	}

}