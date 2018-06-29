package main

import (
	"bytes"
	"log"
	"io"
)

//TODO: brackets, edge cases.

var bufTok = bytes.Buffer{}
var bufOut = bytes.Buffer{}

func main() {
	r := initReader()
	initMapping()

	for {
		c, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		processState(c)
	}

	bufOut.Write(bufTok.Bytes())
	writeOut(bufOut.Bytes())
}

//Very simple finite state machine, checks base cases.
//Need a use a stack and scanners & parsers to better check syntax and grammar.
func processState(c rune) {
	bufTok.WriteRune(c)
	if skip == true {
		skip = false
		return
	}

	switch state {
	case NonPhp: //Looking for '<?php' -- bug: this will pick up quoted or commented as well.
		if bytes.Contains(bufTok.Bytes(), PhpFlag) {
			RestartScan()
		}
	case Question: //Looks for a > to detect transition to NonPhp
		if c == RBRACKET {
			state = NonPhp
		} else {
			RestartScan()
		}
	case Escaped:
		RestartScan()
	case UserDef:
		if !ValidWord(string(c)) {
			RestartScan()
		}
	case DubQuoted:
		if c == DubQUOTE {
			RestartScan()
			} else if c == BACKSLASH {
			skip = true
		}
	case SingQuoted:
		if c == SingQUOTE {
			RestartScan()
		} else if c == BACKSLASH {
			skip = true
		}
	case Scan:
		if !ValidWord(string(c)) {
			endWord(lookUpBuffer(), c)
		}
	case MultiComment:
		if bytes.Contains(bufTok.Bytes(), endComment) {
			RestartScan()
		}
	case OneLineComment:
		if NewLine(string(c)) {
			RestartScan()
		}
	case FwdSlash:
		if c == ASTRIX {
			state = MultiComment
		} else if c == FwdSLASH {
			state = OneLineComment
		} else {
			RestartScan()
		}
	default:
		err := "Unreachable state."
		log.Fatal(err)
	}
}

func RestartScan() {
	bufOut.Write(bufTok.Bytes())
	bufTok.Reset()
	state = Scan
}

func transitionState(c rune) {
	switch c {
	case QUESTION:
		state = Question
	case HASHTAG:
		state = OneLineComment
	case FwdSLASH:
		state = FwdSlash
	case DubQUOTE:
		state = DubQuoted
	case SingQUOTE:
		state = SingQuoted
	case VARIABLE:
		state = UserDef
	case BACKSLASH:
		state = Escaped
	default:
		//state stays the same- no transition
	}
}

func endWord(str string, c rune) {
	bufOut.WriteString(str)
	bufOut.WriteRune(c)
	bufTok.Reset()
	transitionState(c)
}


