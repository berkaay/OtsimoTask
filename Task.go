package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Bu kodu çalıştırmak için command line'dan `go run task.go` çalıştırman yeterli.
var (
	paragraphCount int
	chapterCount   int
	chapters       [][]string
	paragraphs     []string
)

//readBook reads the book at filePath. Keep the at a glabal variable at access it at 'count' and 'query' functions
func readBook(filePath string) {
	//	YOUR CODE HERE. Read the book and save it to a global variable, something like `var Book [][]string`
	//open file
	paragraphCount = 0
	chapterCount = 0
	tempPar := ""
	emptyCount := 0
	endOfParagraph := false
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	//close file after you are done with it
	defer file.Close()
	//--regexes
	isChapter, _ := regexp.Compile("Chapter [0-9][0-9]?")
	isEmptyLine, _ := regexp.Compile("^$")
	//--
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if isChapter.MatchString(line) {
			//fmt.Println("chapter")
			chapters = append(chapters, paragraphs)
			chapterCount++
			paragraphs = []string{}
		}
		if isEmptyLine.MatchString(line) {
			//fmt.Println("empty")
			emptyCount++
			endOfParagraph = true
		} else { //para
			if endOfParagraph {
				//	fmt.Println("Paragraph")
				paragraphs = append(paragraphs, tempPar)
				paragraphCount++
				endOfParagraph = false
				tempPar = line
			} else {
				tempPar += line
			}

		}

	}
	//fmt.Println(paragraphCount)
	//fmt.Println(chapterCount)
	//fmt.Println(len(paragraphs))
	//fmt.Println(len(chapters))
	//fmt.Println(chapters[4][0])
	//fmt.Println(len(chapters[chapterCount]))
}
func query(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	c := q.Get("c") //Get Chapter from url as string. you have to convert it to int by using strconv.Atoi
	p := q.Get("p") //Get Paragraph from url as string, you have to convert it to int
	result := ""
	fmt.Println(p)
	fmt.Println(c)
	cint, _ := strconv.Atoi(c)
	pint, _ := strconv.Atoi(p)
	if p == "" {
		fmt.Fprint(w, "no p")
		fmt.Fprint(w, chapters[cint])

	} else {
		fmt.Fprint(w, chapters[cint][pint])

	}
	fmt.Fprint(w, result)
}
func count(w http.ResponseWriter, r *http.Request) {
	chapCount := 0
	paraCount := 0
	chapCount = chapterCount
	paraCount = paragraphCount

	fmt.Fprintf(w, "chapter: %d\nparagraph: %d\n", chapCount, paraCount)
}
func otherwise(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
func main() {
	readBook("book.txt")
	http.HandleFunc("/count", count)
	http.HandleFunc("/query", query)
	http.HandleFunc("/", otherwise)
	http.ListenAndServe(":8080", nil)
}
