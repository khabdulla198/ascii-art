package main

import (
	ascii "ascii/func"
	"os"
)

func main() {
	args := os.Args

	wordTofind, file := ascii.Validation(args)

	//turn the file into an array
	fileArray := ascii.FileToArray(file)
	//find the location of each letter of the word

	//print the word as an ascii art into the terminal

}
