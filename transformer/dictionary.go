package main

import (
	"os"
	"encoding/gob"
	"strings"
	"fmt"
)

var lookup = make(map[string]string)

//TODO process dictionary

func lookUpBuffer() string {
	s := bufTok.String()
	if s == "<nil>" {
		return "<nil>"
	}

	sOut := string([]rune(s)[:(len(s) - 1)])

	if val, ok := lookup[strings.ToLower(sOut)]; ok {
		return val
	} else {
		return sOut
	}
}


//Grab dictionary -- FromFile
func initMapping() {
	decodeFile, err := os.Open("/php/scrambled.gob")
	if err != nil {
		fmt.Println("Error-- transformation not preformed.")
		fmt.Println(err)
		os.Exit(1)

	}
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)
	decoder.Decode(&lookup)
}


