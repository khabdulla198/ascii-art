package main

import (
	ascii "ascii/func"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	wordTofind, file, error := ascii.Validation(args)
	if error != nil {
		fmt.Println(error)
		return
	}
	//turn the file into an array
	fileArray := ascii.Convert(file)

	//convert the wordToFind into an array for each \n
	splitstring := strings.Split(wordTofind, "\\n")
	result := ""

	// For loop to generate for each split
	for _, word := range splitstring {
		if word == ""{
			fmt.Print("\n")
			break
		} else{
		//fmt.Println(string('\n'))
		asciiArt, err := ascii.GenerateAscii(word, fileArray)
		if err != nil {
			fmt.Println(err)
			return
		}
		result += asciiArt //+ "\n"
		}
	}
	//print the word as an ascii art into the terminal
	fmt.Print(result)
	}