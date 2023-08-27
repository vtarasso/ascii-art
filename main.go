package main

import (
	"ascii-art/datafile"
	"fmt"
	"os"
	"strings"
)

const standardHash = "ac85e83127e49ec42487f272d9b9db8b"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR: wrong number of arguments")
		return
	}
	if args[0] == "" {
		return
	}
	if args[0] == "\\n" {
		return
	}

	for _, alter := range args[0] {
		if (rune(alter) < rune(32) || rune(alter) > rune(127)) && alter != rune(10) {
			fmt.Println("ERROR: non printable character")
			return
		}
	}
	filename := "standard.txt"
	// Chaek for hash
	if standardHash != datafile.GetHash(filename) {
		fmt.Println("Error: Wrong hash")
		return
	}

	asciiLines, err := datafile.GetStrings("standard.txt")
	if err != nil {
		fmt.Println("ERROR: can't read file")
		return
	}
	asciiMap := make(map[rune][]string)
	x := 1
	y := 9
	for key := 32; key < 127; key++ {
		asciiMap[rune(key)] = asciiLines[x:y]
		x = x + 9
		y = y + 9

	}

	res := ""
	text := strings.ReplaceAll(os.Args[1], "\n", "\\n")
	arg := strings.Split(text, "\\n")
	for i, v := range arg {
		if v == "" {
			arg[i] = ""
		}
	}
	// fmt.Println(arg)
	newline := forNewLines(arg)
	for w := 0; w < len(arg); w++ {
		if newline && w == len(arg)-1 {
			break
		}
		if arg[w] != "" {
			for i := 0; i < 8; i++ {
				for _, ch := range arg[w] {
					res = res + asciiMap[ch][i]
				}
				res = res + string(rune(10))
			}
		} else if arg[w] == "" {
			res = res + string(rune(10))
		}
	}
	fmt.Print(res)
}

func forNewLines(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			return false
		}
	}
	return true
}