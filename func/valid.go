package ascii

import "fmt"

func Validation(args []string) bool{
	flag:= false
	
	if len(args) > 3 || len(args) == 1 {
		fmt.Println("Error:There should be one or two string argument\n[STRING] [OPTIONAL BANNER]")
	} else if len(args) == 3 && args[2] != "standard" && args[2] != "thinkertoy" && args[2] != "shadow" {
		fmt.Println("Error:only two arguments are allowed [STRING] [OPTIONAL BANNER]\n Banners that are available are: standard, thinkertoy, and shadow")
	}
}
