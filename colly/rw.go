package colly

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WriteComments(filename string, arr []string) {

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		fmt.Println("Greska pri otvaranju fajla")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, c := range arr {
		if c == "" {
			continue
		}
		//fmt.Println(line)
		_, err = writer.WriteString(c)
		if err != nil {
			fmt.Println("Neuspesan upis komentara u fajl!")
		} else {
			//fmt.Println("Komentar uspesno unet u fajl!")
		}
	}
	writer.Flush()
}

func ReadComments(filename string) []string {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModeExclusive)
	if err != nil {
		fmt.Println("Greska pri otvaranju fajla!")
	}
	defer file.Close()

	var comments []string
	var currentComment string
	var inCommentBlock bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "COMMENT###") {
			if inCommentBlock && currentComment != "" {
				comments = append(comments, currentComment)
				currentComment = ""
			}

			inCommentBlock = true
		}

		if inCommentBlock {
			currentComment += line + "\n"
		}

		if line == "" && inCommentBlock {
			inCommentBlock = false
		}
	}

	if currentComment != "" {
		comments = append(comments, currentComment)
	}

	return comments
}
