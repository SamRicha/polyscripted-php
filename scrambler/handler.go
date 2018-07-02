package main

import (
	"regexp"
	"io/ioutil"
	"os"
	"encoding/gob"
	"bytes"
)

var ValidWord = regexp.MustCompile("\\w").MatchString

//REGEX used for now for base cases.
var keywordsRegex = regexp.MustCompile( //REGEX found as user @martindilling comment on PHP documentation.
	"((a(bstract|nd|rray|s))|" +
		"(c(a(llable|se|tch)|l(ass|one)|on(st|tinue)))|" +
		"(d(e(clare|fault)|ie|o))|" +
		"(e(cho|lse(if)?|mpty|nd(declare|for(each)?|if|switch|while)|val|x(it|tends)))|" +
		"(f(inal(ly)|or(each)?|unction))|" +
		"(g(lobal|oto))|" +
		"(i(f|mplements|n(clude(_once)?|st(anceof|eadof)|terface)|sset))|" +
		"(n(amespace|ew))|" +
		"(p(r(i(nt|vate)|otected)|ublic))|" +
		"(re(quire(_once)?|turn))|" +
		"(s(tatic|witch))|" +
		"(t(hrow|r(ait|y)))|(u(nset|se))|" +
		"(__halt_compiler|break|list|(x)?or|var|while))")


var lexFlag = []byte("<ST_IN_SCRIPTING>\"")
var yakFlag = []byte("%token T_")


const YAKFILE = "/php/php-src/Zend/zend_language_parser.y"
const LEXFILE = "/php/php-src/Zend/zend_language_scanner.l"

//const YAKFILE = "zend_language_parser.y"
//const LEXFILE = "zend_language_scanner.l"

const (
	YAK = iota
	LEX = iota
)

var polyWords = make(map[string]string)

var b = bytes.Buffer{}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeFile(fileOut string) {
	err := ioutil.WriteFile(fileOut, b.Bytes(), 0644)
	check(err)
}

func writeLineToBuff(s []byte) {
	b.Write([]byte(s))
	b.WriteString("\n")
}

func serializeMap() {
	encodeFile, err := os.Create("/php/scrambled.gob")
	if err != nil {
		panic(err)
	}

	encoder := gob.NewEncoder(encodeFile)

	if err := encoder.Encode(polyWords); err != nil {
		panic(err)
	}
	encodeFile.Close()
}
