package main

import (
	"os"
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	scrambleFile(LEX)
	fmt.Println("Mapping Built. \nLex Scrambled.")
	b.Reset()
	scrambleFile(YAK)
	fmt.Println("Yak Scrambled.")
	serializeMap()
	fmt.Println("Map Serialized.")
}

func scrambleFile(file int) {
	switch file {
	case LEX:
		scanLines(LEXFILE, lexFlag, LEX)
	case YAK:
		scanLines(YAKFILE, yakFlag, YAK)
	}
}

func scanLines(fileIn string, flag []byte, state int) {
	file, err := os.Open(fileIn)
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Bytes()

		if bytes.HasPrefix(line, flag) && keywordsRegex.Match(line) {
			getWords(line, state)
		} else {
			writeLineToBuff(line)
		}
	}
	writeFile(fileIn)
}

func getWords(s []byte, state int) {
	keyWord := keywordsRegex.Find(s)
	index := keywordsRegex.FindIndex(s)
	suffix := string(s[index[1]])
	prefix := string(s[index[0]-1])

	if ValidWord(suffix) || ValidWord(prefix) { //word found was part of larger word, return
		writeLineToBuff(s)
		return
	}

	key := string(keyWord)

	if _, ok := polyWords[key]; !ok {
		if state != YAK { //For now, only words scrambled within lex will be scrambled in Yak.
			polyWords[key] = RandomStringGen() // Add to map, generate random string (need checks here?)
			key = polyWords[key]
		}
	} else {
		key = polyWords[key]
	}

	out := keywordsRegex.ReplaceAll([]byte(s), []byte(key)) //Replace word with random string
	writeLineToBuff(out)
}