package colly

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WriteComments(filename string, arr []string) error {

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
	return err
}

func ReadComments(filename string) ([]string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModeExclusive)
	if err != nil {
		fmt.Println("Greska pri otvaranju fajla!")
	}
	defer file.Close()

	var comments []string
	var currentComment string
	var newLine bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "COMMENT###") {
			if newLine {
				comments = append(comments, currentComment)
				currentComment = line
			} else {
				currentComment = line
				newLine = true
			}
		} else {
			currentComment = fmt.Sprintf("%s\n%s", currentComment, line)
		}
	}

	if currentComment != "" {
		comments = append(comments, currentComment)
	}
	return comments, err
}
