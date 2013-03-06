package main

import (
	"bufio"
	"bytes"
	"fmt"
	//"github.com/knieriem/markdown"
	"github.com/russross/blackfriday"
	"io"
	"os"
	"strings"
)

func concatene(first string, second string) (retour string) {
	var (
		buffer bytes.Buffer
	)
	buffer.WriteString(first)
	buffer.WriteString(second)
	return buffer.String()
}

func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func writeLines(lines []string, path string) (err error) {
	var (
		file *os.File
	)

	if file, err = os.Create(path); err != nil {
		return
	}
	defer file.Close()

	//writer := bufio.NewWriter(file)
	for _, item := range lines {
		//fmt.Println(item)
		_, err := file.WriteString(strings.TrimSpace(item) + "\n")
		//file.Write([]byte(item)); 
		if err != nil {
			//fmt.Println("debug")
			fmt.Println(err)
			break
		}
	}
	/*content := strings.Join(lines, "\n")
	  _, err = writer.WriteString(content)*/
	return
}

func readfile(path string) (retour string) {
	lines, err := readLines(path)
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}
	var a string
	a = ""
	for _, line := range lines {
		a = concatene(a, "\n"+line)
	}
	return a
}

func writefile(path string, line string) {
	var (
		lines []string
	)

	lines = strings.Split(line, "\n")
	writeLines(lines, path)
}

func main() {

	var lecture string

	lecture = readfile("voyage.md")
	fmt.Println(lecture)
	output := blackfriday.MarkdownBasic([]byte(lecture))
	fmt.Println(string(output))
	writefile("index.html", string(output))
	//p := markdown.NewParser(&markdown.Extensions{Smart: true})

	//fmt.Println(markdown.ToHTML())
	//w := bufio.NewWriter(lecture)
	//p.Markdown(os.Stdin, markdown.ToHTML(&lecture))
	//w.Flush()

}
