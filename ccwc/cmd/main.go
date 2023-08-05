package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	returnBytes := flag.Bool("c", false, "count number of bytes")
	returnLines := flag.Bool("l", false, "count number of lines")
	returnWords := flag.Bool("w", false, "count number of words")
	returnChars := flag.Bool("m", false, "count number of characters")
	flag.Parse()

	filename := flag.Arg(0)

	if filename == "" {
		fmt.Println("Please give a file path")
		return
	}

	flags := makeFlagMap(returnBytes, returnLines, returnWords, returnChars)
	for _, pair := range flags {
		if pair.Value {
			size, err := pair.Func(filename)
			printMessage(size, filename, err)
			return
		}
	}

	bytes, lines, words, err := getAllThree(filename)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("   %d   %d   %d %s", lines, words, bytes, filename)
}

type Pair struct {
	Value bool
	Func  func(filename string) (int64, error)
}

func makeFlagMap(c, l, w, m *bool) [4]Pair {
	var output [4]Pair
	output[0] = Pair{*c, numberOfBytesInFile}
	output[1] = Pair{*l, numberOfLinesInFile}
	output[2] = Pair{*w, numberOfWordsInFile}
	output[3] = Pair{*m, numberOfCharactersInFile}
	return output
}

func printMessage(size int64, filename string, err error) {
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("   %d %s", size, filename)
}

func getAllThree(filename string) (int64, int64, int64, error) {
	nBytes, err := numberOfBytesInFile(filename)
	if err != nil {
		return -1, -1, -1, err
	}

	nLines, err := numberOfLinesInFile(filename)
	if err != nil {
		return nBytes, -1, -1, err
	}

	nWords, err := numberOfWordsInFile(filename)
	if err != nil {
		return nBytes, nLines, -1, err
	}

	return nBytes, nLines, nWords, nil
}

func numberOfBytesInFile(filename string) (int64, error) {
	return scan(filename, bufio.ScanBytes)
}

func numberOfLinesInFile(filename string) (int64, error) {
	return scan(filename, bufio.ScanLines)
}

func numberOfWordsInFile(filename string) (int64, error) {
	return scan(filename, bufio.ScanWords)
}

func numberOfCharactersInFile(filename string) (int64, error) {
	return scan(filename, bufio.ScanRunes)
}

func scan(filename string, splitFunc bufio.SplitFunc) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(splitFunc)

	var output int64
	for scanner.Scan() {
		output++
	}

	if err := scanner.Err(); err != nil {
		return output, err
	}

	return output, nil
}
