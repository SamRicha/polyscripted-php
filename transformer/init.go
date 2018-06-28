package main

import (
	"regexp"
	"io"
	"io/ioutil"
	"fmt"
	"os"
	"bytes"
	"flag"
	"errors"
	"path/filepath"
)

var FILEIN = ""
var FILEOUT = ""


var state = NonPhp
var ValidWord = regexp.MustCompile("\\w").MatchString
var NewLine = regexp.MustCompile("\\r\\n|\\r|\\n|;").MatchString

var PhpFlag = []byte("<?php")
var endComment = []byte("*/")

const (
	UserDef        = iota
	DubQuoted      = iota
	Scan           = iota
	Escaped        = iota
	NonPhp         = iota
	FwdSlash       = iota
	MultiComment   = iota
	OneLineComment = iota
	Question 	   = iota
	SingQuoted 	   = iota
)

const (
	DubQUOTE  = rune('"')
	VARIABLE  = rune('$')
	BACKSLASH = rune('\\')
	RBRACKET  = rune('>')
	HASHTAG   = rune('#')
	ASTRIX    = rune('*')
	FwdSLASH  = rune('/')
	QUESTION  = rune('?')
	SingQUOTE  = rune('\'')
)


func initReader() io.RuneReader {

	parseCmdLn()

	original, err := ioutil.ReadFile(FILEIN)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return bytes.NewReader(original)
}

func writeOut(b []byte) {
	if err := ioutil.WriteFile(FILEOUT, b, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Generated polyscripted file- %s.\n", FILEOUT)
}

func parseCmdLn() { //TODO: This should take multiple files eventually.
	flag.StringVar(&FILEIN, "f", "", "File to transform.")
	var replace = flag.Bool("replace", false, "Set to true to replace original file.")

	flag.Parse()
	if *replace == true {
	 	FILEOUT = FILEIN
	} else {
		dir, file := filepath.Split(FILEIN)
		FILEOUT = dir + "ps-" + file
	}


	if FILEIN == "" || FILEOUT == "" {
		err := errors.New("required field '-f' missing. Please input filename" )
		fmt.Println(err)
		os.Exit(1)
	}

}
