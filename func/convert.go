package ascii

import(
	"bufio"
	"fmt"
	"log"
	"os"

)

func Convert (file string) []string{

	f,err := os.Open(file)

	if err != nil {
		log.Fatalf("Error: failed to open file %s",err)
	}

	defer f.Close()

	for scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
}
