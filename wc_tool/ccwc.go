package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func options(opt string, file os.File, scanner *bufio.Scanner) {

	data, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	switch opt {
	case "-c", "--bytes":
		fmt.Println(data.Size())
		break
	case "-l", "--lines":
		lineCount := 0
		for scanner.Scan() {
			lineCount++
		}
		fmt.Println(lineCount)
		break
	case "-w", "--words":
		scanner.Split(bufio.ScanWords)
		wordCount := 0
		for scanner.Scan() {
			wordCount++
		}
		fmt.Println(wordCount)
		break

	case "-m", "--chars":
		charCount := 0
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			word := scanner.Text()
			charCount += utf8.RuneCountInString(word)
		}
		fmt.Println(charCount)
		break
	case "-h", "--help":
		fmt.Println("Usage: ccwc <filename> <option>")
		fmt.Println("Options:")
		fmt.Println("  -c, --bytes: print the byte counts")
		fmt.Println("  -l, --lines: print the newline counts")
		fmt.Println("  -w, --words: print the word counts")
		fmt.Println("  -m, --chars: print the character counts")
		fmt.Println("  -h, --help: print this help message")
		break
	default:
		lineCount := 0
		for scanner.Scan() {
			lineCount++
		}
		fmt.Print(lineCount, " ")

		file.Seek(0, 0)
		scanner = bufio.NewScanner(&file)
		scanner.Split(bufio.ScanWords)
		wordCount := 0
		for scanner.Scan() {
			wordCount++
		}
		fmt.Print(wordCount, " ")

		fmt.Print(data.Size(), " ")
	}
}

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: ccwc <filename> <option>")
		return
	}

	if len(args) == 1 {
		if args[0] == "-h" || args[0] == "--help" {
			options(args[0], *os.Stdin, bufio.NewScanner(os.Stdin))
			return
		}

		args = append(args, "-a")
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// if no file is provided read from stdin
	if args[0] == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		wordCount := 0
		for scanner.Scan() {
			wordCount++
		}
		fmt.Println(wordCount)
		return
	}

	scanner := bufio.NewScanner(file)
	options(args[1], *file, scanner)
}
