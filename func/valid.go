package ascii

import (
	"fmt"
	"strings"
)

func Validation(args []string) (string, string, error) {
	file := "standard.txt" //making default as standard
	wordTofind := args[1]  // declaring word/phrase
	if len(args) > 3 || len(args) == 1 {
		//error message for incorrect amount of input
		return "", "", fmt.Errorf("Error:There should be one or two string argument\n[STRING] [OPTIONAL BANNER]")
	}

	if len(args) == 3 {
		switch strings.ToLower(args[2]) {
		// if standard is input, standard.txt will be the file
		case "standard":
			file = "standard.txt"
		// if shadow is input, shadow.txt will become the file
		case "shadow":
			file = "shadow.txt"
		//if thinkertoy is input, thinkertoy will become the file
		case "thinkertoy":
			file = "thinkertoy.txt"
		// error message in case of invalid input
		default:
			return "", "", fmt.Errorf("error: only two arguments are allowed [STRING] [OPTIONAL BANNER]\n Banners that are available are: standard, thinkertoy, and shadow")
		}
	}

	return wordTofind, file, nil
}
